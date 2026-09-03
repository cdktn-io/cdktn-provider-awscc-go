// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjob


type IotJobJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteria struct {
	// The threshold for number of notified things that will initiate the increase in rate of rollout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job#number_of_notified_things IotJob#number_of_notified_things}
	NumberOfNotifiedThings *float64 `field:"optional" json:"numberOfNotifiedThings" yaml:"numberOfNotifiedThings"`
	// The threshold for number of succeeded things that will initiate the increase in rate of rollout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job#number_of_succeeded_things IotJob#number_of_succeeded_things}
	NumberOfSucceededThings *float64 `field:"optional" json:"numberOfSucceededThings" yaml:"numberOfSucceededThings"`
}

