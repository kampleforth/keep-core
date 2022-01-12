import { helpers, ethers } from "hardhat"
import type { BigNumber } from "ethers"

import { constants } from "../fixtures"
import type { Operator } from "./operators"
import type { SortitionPool } from "../../typechain"

const { mineBlocks } = helpers.time

export async function selectGroup(
  sortitionPool: SortitionPool,
  seed: BigNumber
): Promise<Operator[]> {
  const identifiers = await sortitionPool.selectGroup(
    constants.groupSize,
    seed.toHexString()
  )
  const addresses = await sortitionPool.getIDOperators(identifiers)

  return identifiers.map((identifier, i) => ({
    id: identifier,
    address: addresses[i],
  }))
}
