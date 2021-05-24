import React from "react"
import centered from "@storybook/addon-centered/react"
import MaxAmountAddon from "../components/MaxAmountAddon"

export default {
  title: "MaxAmountAddon",
  component: MaxAmountAddon,
  argTypes: {
    onClick: {
      action: "onClick clicked",
    },
  },
  decorators: [centered],
}

const Template = (args) => <MaxAmountAddon {...args} />

export const MaxStake = Template.bind({})
MaxStake.args = { text: "Max Stake" }

export const MaxKEEP = Template.bind({})
MaxKEEP.args = { text: "Max KEEP" }
