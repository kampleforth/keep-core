import React from "react"
import centered from "@storybook/addon-centered/react"
import StepNav from "../components/StepNav"

const mockedSteps = ["step1", "step2", "step3", "step4"]

export default {
  title: "StepNav",
  component: StepNav,
  decorators: [centered],
}

const Template = (args) => <StepNav {...args} />

export const Default = Template.bind({})
Default.args = {
  steps: mockedSteps,
  activeStep: 1,
}
