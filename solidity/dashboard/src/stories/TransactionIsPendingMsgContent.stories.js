import React from "react"
import TransactionIsPendingMsgContent from "../components/ViewTxMsgContent"
import centered from "@storybook/addon-centered/react"

const mockedTransactionHash =
  "6146ccf6a66d994f7c363db875e31ca35581450a4bf6d3be6cc9ac79233a69d0"

export default {
  title: "TransactionIsPendingMsgContent",
  component: TransactionIsPendingMsgContent,
  decorators: [centered],
}

const Template = (args) => <TransactionIsPendingMsgContent {...args} />

export const Default = Template.bind({})
Default.args = {
  txHash: mockedTransactionHash,
}