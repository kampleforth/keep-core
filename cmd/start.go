package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/keep-network/keep-core/build"
	"github.com/keep-network/keep-core/config"
	"github.com/keep-network/keep-core/pkg/beacon"
	"github.com/keep-network/keep-core/pkg/chain"
	"github.com/keep-network/keep-core/pkg/chain/ethereum"
	"github.com/keep-network/keep-core/pkg/diagnostics"
	"github.com/keep-network/keep-core/pkg/firewall"
	"github.com/keep-network/keep-core/pkg/generator"
	"github.com/keep-network/keep-core/pkg/metrics"
	"github.com/keep-network/keep-core/pkg/net"
	"github.com/keep-network/keep-core/pkg/net/libp2p"
	"github.com/keep-network/keep-core/pkg/net/retransmission"
	"github.com/keep-network/keep-core/pkg/storage"
	"github.com/keep-network/keep-core/pkg/tbtc"
)

// StartCommand contains the definition of the start command-line subcommand.
var StartCommand = &cobra.Command{
	Use:   "start",
	Short: "Starts the Keep Client",
	Long:  "Starts the Keep Client in the foreground",
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := clientConfig.ReadConfig(configFilePath, cmd.Flags(), config.AllCategories...); err != nil {
			logger.Fatalf("error reading config: %v", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := start(cmd); err != nil {
			logger.Fatal(err)
		}
	},
}

func init() {
	initFlags(StartCommand, &configFilePath, clientConfig, config.AllCategories...)

	StartCommand.SetUsageTemplate(
		fmt.Sprintf(`%s
Environment variables:
    %s    Password for Keep operator account keyfile decryption.
    %s                 Space-delimited set of log level directives; set to "help" for help.
`,
			StartCommand.UsageString(),
			config.EthereumPasswordEnvVariable,
			config.LogLevelEnvVariable,
		),
	)
}

// start starts a node
func start(cmd *cobra.Command) error {
	ctx := context.Background()

	logger.Infof(
		"Starting the client against [%s] ethereum network...",
		clientConfig.Ethereum.Network,
	)

	beaconChain, tbtcChain, blockCounter, signing, operatorPrivateKey, err :=
		ethereum.Connect(ctx, clientConfig.Ethereum)
	if err != nil {
		return fmt.Errorf("error connecting to Ethereum node: [%v]", err)
	}

	firewall := firewall.AnyApplicationPolicy(
		[]firewall.Application{beaconChain, tbtcChain},
	)

	netProvider, err := libp2p.Connect(
		ctx,
		clientConfig.LibP2P,
		operatorPrivateKey,
		firewall,
		retransmission.NewTicker(blockCounter.WatchBlocks(ctx)),
	)
	if err != nil {
		return fmt.Errorf("failed while creating the network provider: [%v]", err)
	}

	nodeHeader(
		netProvider.ConnectionManager().AddrStrings(),
		beaconChain.Signing().Address().String(),
		clientConfig.LibP2P.Port,
		clientConfig.Ethereum,
	)

	storage, err := storage.Initialize(
		clientConfig.Storage,
		clientConfig.Ethereum.KeyFilePassword,
	)
	if err != nil {
		return fmt.Errorf("cannot initialize storage: [%w]", err)
	}

	beaconKeyStorePersistence, err := storage.InitializeKeyStorePersistence("beacon")
	if err != nil {
		return fmt.Errorf("cannot initialize beacon keystore persistence: [%w]", err)
	}

	tbtcKeyStorePersistence, err := storage.InitializeKeyStorePersistence("tbtc")
	if err != nil {
		return fmt.Errorf("cannot initialize tbtc keystore persistence: [%w]", err)
	}

	tbtcDataPersistence, err := storage.InitializeWorkPersistence("tbtc")
	if err != nil {
		return fmt.Errorf("cannot initialize tbtc data persistence: [%w]", err)
	}

	scheduler := generator.StartScheduler()

	err = beacon.Initialize(
		ctx,
		beaconChain,
		netProvider,
		beaconKeyStorePersistence,
		scheduler,
	)
	if err != nil {
		return fmt.Errorf("error initializing beacon: [%v]", err)
	}

	initializeMetrics(
		ctx,
		clientConfig,
		netProvider,
		blockCounter,
	)

	diagnosticsRegistry, diagnosticsConfigured := initializeDiagnostics(clientConfig)
	if diagnosticsConfigured {
		diagnosticsRegistry.RegisterConnectedPeersSource(netProvider, signing)
		diagnosticsRegistry.RegisterClientInfoSource(
			netProvider,
			signing,
			build.Version,
			build.Revision,
		)
	}

	err = tbtc.Initialize(
		ctx,
		tbtcChain,
		netProvider,
		tbtcKeyStorePersistence,
		tbtcDataPersistence,
		scheduler,
		clientConfig.Tbtc,
		diagnosticsRegistry,
	)
	if err != nil {
		return fmt.Errorf("error initializing TBTC: [%v]", err)
	}

	select {
	case <-ctx.Done():
		if err != nil {
			return err
		}

		return fmt.Errorf("uh-oh, we went boom boom for no reason")
	}
}

func initializeMetrics(
	ctx context.Context,
	config *config.Config,
	netProvider net.Provider,
	blockCounter chain.BlockCounter,
) (*metrics.Registry, bool) {
	registry, isConfigured := metrics.Initialize(
		ctx, config.Metrics.Port,
	)
	if !isConfigured {
		logger.Infof("metrics are not configured")
		return nil, false
	}

	logger.Infof(
		"enabled metrics on port [%v]",
		config.Metrics.Port,
	)

	registry.ObserveConnectedPeersCount(
		netProvider,
		config.Metrics.NetworkMetricsTick,
	)

	registry.ObserveConnectedBootstrapCount(
		netProvider,
		config.LibP2P.Peers,
		config.Metrics.NetworkMetricsTick,
	)

	registry.ObserveEthConnectivity(
		blockCounter,
		config.Metrics.EthereumMetricsTick,
	)

	return registry, true
}

func initializeDiagnostics(
	config *config.Config,
) (*diagnostics.Registry, bool) {
	registry, isConfigured := diagnostics.Initialize(
		config.Diagnostics.Port,
	)
	if !isConfigured {
		logger.Infof("diagnostics are not configured")
		return nil, false
	}

	logger.Infof(
		"enabled diagnostics on port [%v]",
		config.Diagnostics.Port,
	)

	return registry, true
}
