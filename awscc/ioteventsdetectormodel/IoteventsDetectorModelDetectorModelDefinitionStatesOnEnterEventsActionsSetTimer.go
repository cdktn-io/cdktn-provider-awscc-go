// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetTimer struct {
	// The duration of the timer, in seconds.
	//
	// You can use a string expression that includes numbers, variables (``$variable.<variable-name>``), and input values (``$input.<input-name>.<path-to-datum>``) as the duration. The range of the duration is 1-31622400 seconds. To ensure accuracy, the minimum duration is 60 seconds. The evaluated result of the duration is rounded down to the nearest whole number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_detector_model#duration_expression IoteventsDetectorModel#duration_expression}
	DurationExpression *string `field:"optional" json:"durationExpression" yaml:"durationExpression"`
	// The number of seconds until the timer expires.
	//
	// The minimum value is 60 seconds to ensure accuracy. The maximum value is 31622400 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_detector_model#seconds IoteventsDetectorModel#seconds}
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
	// The name of the timer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_detector_model#timer_name IoteventsDetectorModel#timer_name}
	TimerName *string `field:"optional" json:"timerName" yaml:"timerName"`
}

