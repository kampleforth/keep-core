const {web3} = require("@openzeppelin/test-environment")

async function grantTokens(
    grantContract,
    token, amount,
    from, grantee,
    unlockingDuration, start, cliff,
    revocable,
    stakingPolicy) {
  let grantData = web3.eth.abi.encodeParameters(
    ['address', 'uint256', 'uint256', 'uint256', 'bool', 'address'],
    [grantee, unlockingDuration.toNumber(), start.toNumber(), cliff.toNumber(), revocable, stakingPolicy]
  );

  await token.approveAndCall(grantContract.address, amount, grantData, {from: from})
  return (await grantContract.getPastEvents())[0].args[0].toNumber()
}
module.exports = grantTokens
