package sortition

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ipfs/go-log"
	"github.com/keep-network/keep-core/pkg/internal/testutils"
	"github.com/keep-network/keep-core/pkg/sortition/internal/local"
)

const (
	testOperatorAddress        = "0x3c5eBAcFe5aE12D82d43602a12b8bBb76b893CfA"
	testStakingProviderAddress = "0x80C63B577DC79B2432357BECC5b431dfb8E181DD"

	statusCheckTick = 10 * time.Millisecond
)

// If environment variable `PRINT_LOGS_IN_TEST` is set to `true`, logger in
// the code called by unit tests prints to the console.
func TestMain(m *testing.M) {
	if os.Getenv("PRINT_LOGS_IN_TEST") == "true" {
		err := log.SetLogLevel("*", "DEBUG")
		if err != nil {
			fmt.Fprintf(os.Stderr, "logger initialization failed: [%v]\n", err)
			os.Exit(-1)
		}
	}

	os.Exit(m.Run())
}

func TestMonitorPool_NotRegisteredOperator(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	localChain := local.Connect(testOperatorAddress)

	err := MonitorPool(ctx, localChain, statusCheckTick)
	testutils.AssertErrorsEqual(t, errOperatorUnknown, err)
}

func TestMonitorPool_NoStake(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	localChain := local.Connect(testOperatorAddress)
	localChain.RegisterOperator(testStakingProviderAddress, testOperatorAddress)

	err := MonitorPool(ctx, localChain, statusCheckTick)
	if err != nil {
		t.Fatal(err)
	}

	isOperatorInPool, err := localChain.IsOperatorInPool()
	if err != nil {
		t.Fatal(err)
	}

	if isOperatorInPool {
		t.Fatal("expected the operator not to be in the pool")
	}
}

func TestMonitor_JoinPool(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	localChain := local.Connect(testOperatorAddress)
	localChain.RegisterOperator(testStakingProviderAddress, testOperatorAddress)
	localChain.SetEligibleStake(testStakingProviderAddress, big.NewInt(100))

	err := MonitorPool(ctx, localChain, statusCheckTick)
	if err != nil {
		t.Fatal(err)
	}

	isOperatorInPool, err := localChain.IsOperatorInPool()
	if err != nil {
		t.Fatal(err)
	}

	if !isOperatorInPool {
		t.Fatal("expected the operator to join the pool")
	}
}

func TestMonitor_UpdatePool(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	localChain := local.Connect(testOperatorAddress)
	localChain.RegisterOperator(testStakingProviderAddress, testOperatorAddress)
	localChain.SetEligibleStake(testStakingProviderAddress, big.NewInt(100))
	localChain.JoinSortitionPool()

	localChain.SetEligibleStake(testStakingProviderAddress, big.NewInt(101))

	err := MonitorPool(ctx, localChain, statusCheckTick)
	if err != nil {
		t.Fatal(err)
	}

	isOperatorUpToDate, err := localChain.IsOperatorUpToDate()
	if err != nil {
		t.Fatal(err)
	}
	if !isOperatorUpToDate {
		t.Fatal("expected the operator to be up to date")
	}
}

func TestMonitor_JoinPool_WithDelay(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	localChain := local.Connect(testOperatorAddress)
	localChain.RegisterOperator(testStakingProviderAddress, testOperatorAddress)

	err := MonitorPool(ctx, localChain, statusCheckTick)
	if err != nil {
		t.Fatal(err)
	}

	isOperatorInPool, err := localChain.IsOperatorInPool()
	if err != nil {
		t.Fatal(err)
	}
	if isOperatorInPool {
		t.Fatal("expected the operator not to be in the pool")
	}

	localChain.SetEligibleStake(testStakingProviderAddress, big.NewInt(100))

	// Let's give some time for the monitoring loop to react...
	time.Sleep(50 * time.Millisecond)

	isOperatorInPool, err = localChain.IsOperatorInPool()
	if err != nil {
		t.Fatal(err)
	}
	if !isOperatorInPool {
		t.Fatal("expected the operator to join the pool")
	}
}

func TestMonitor_UpdatePool_WithDelay(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	localChain := local.Connect(testOperatorAddress)
	localChain.RegisterOperator(testStakingProviderAddress, testOperatorAddress)
	localChain.SetEligibleStake(testStakingProviderAddress, big.NewInt(100))
	localChain.JoinSortitionPool()

	err := MonitorPool(ctx, localChain, statusCheckTick)
	if err != nil {
		t.Fatal(err)
	}

	localChain.SetEligibleStake(testStakingProviderAddress, big.NewInt(101))
	isOperatorUpToDate, err := localChain.IsOperatorUpToDate()
	if err != nil {
		t.Fatal(err)
	}
	if isOperatorUpToDate {
		t.Fatal("expected the operator not to be up to date")
	}

	// Let's give some time for the monitoring loop to react...
	time.Sleep(50 * time.Millisecond)

	isOperatorUpToDate, err = localChain.IsOperatorUpToDate()
	if err != nil {
		t.Fatal(err)
	}
	if !isOperatorUpToDate {
		t.Fatal("expected the operator to be up to date")
	}
}

func TestMonitor_EligibleForRewards(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	localChain := local.Connect(testOperatorAddress)
	localChain.RegisterOperator(testStakingProviderAddress, testOperatorAddress)
	localChain.SetEligibleStake(testStakingProviderAddress, big.NewInt(100))
	localChain.JoinSortitionPool()

	err := MonitorPool(ctx, localChain, statusCheckTick)
	if err != nil {
		t.Fatal(err)
	}

	isEligibleForRewards, err := localChain.IsEligibleForRewards()
	if err != nil {
		t.Fatal(err)
	}
	if !isEligibleForRewards {
		t.Fatal("expected the operator to be eligible for rewards")
	}
}

func TestMonitor_CannotRestoreRewardsEligibility(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	localChain := local.Connect(testOperatorAddress)
	localChain.RegisterOperator(testStakingProviderAddress, testOperatorAddress)
	localChain.SetEligibleStake(testStakingProviderAddress, big.NewInt(100))
	localChain.JoinSortitionPool()

	localChain.SetRewardIneligibility(big.NewInt(1))
	localChain.SetCurrentTimestamp(big.NewInt(0))
	err := MonitorPool(ctx, localChain, statusCheckTick)
	if err != nil {
		t.Fatal(err)
	}

	isEligibleForRewards, err := localChain.IsEligibleForRewards()
	if err != nil {
		t.Fatal(err)
	}
	if isEligibleForRewards {
		t.Fatal("expected the operator not to be eligible for rewards")
	}

	canRestoreRewardsEligibility, err := localChain.CanRestoreRewardEligibility()
	if err != nil {
		t.Fatal(err)
	}
	if canRestoreRewardsEligibility {
		t.Fatal("expected the operator cannot restore rewards eligibility")
	}
}

func TestMonitor_CanRestoreRewardsEligibility(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	localChain := local.Connect(testOperatorAddress)
	localChain.RegisterOperator(testStakingProviderAddress, testOperatorAddress)
	localChain.SetEligibleStake(testStakingProviderAddress, big.NewInt(100))
	localChain.JoinSortitionPool()

	localChain.SetRewardIneligibility(big.NewInt(1))
	localChain.SetCurrentTimestamp(big.NewInt(2))

	isEligibleForRewards, err := localChain.IsEligibleForRewards()
	if err != nil {
		t.Fatal(err)
	}
	if isEligibleForRewards {
		t.Fatal("expected the operator not to be eligible for rewards yet")
	}

	canRestoreRewardsEligibility, err := localChain.CanRestoreRewardEligibility()
	if err != nil {
		t.Fatal(err)
	}
	if !canRestoreRewardsEligibility {
		t.Fatal("expected the operator to be able to restore rewards eligibility")
	}

	err = MonitorPool(ctx, localChain, statusCheckTick)
	if err != nil {
		t.Fatal(err)
	}

	isEligibleForRewards, err = localChain.IsEligibleForRewards()
	if err != nil {
		t.Fatal(err)
	}
	if !isEligibleForRewards {
		t.Fatal("expected the operator to be restored for rewards")
	}
}
