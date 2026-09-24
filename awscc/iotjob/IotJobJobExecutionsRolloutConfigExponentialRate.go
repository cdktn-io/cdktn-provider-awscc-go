// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjob


type IotJobJobExecutionsRolloutConfigExponentialRate struct {
	// The minimum number of things that will be notified of a pending job, per minute at the start of job rollout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_job#base_rate_per_minute IotJob#base_rate_per_minute}
	BaseRatePerMinute *float64 `field:"optional" json:"baseRatePerMinute" yaml:"baseRatePerMinute"`
	// The exponential factor to increase the rate of rollout for a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_job#increment_factor IotJob#increment_factor}
	IncrementFactor *float64 `field:"optional" json:"incrementFactor" yaml:"incrementFactor"`
	// Allows you to define a criteria to initiate the increase in rate of rollout for a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_job#rate_increase_criteria IotJob#rate_increase_criteria}
	RateIncreaseCriteria *IotJobJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteria `field:"optional" json:"rateIncreaseCriteria" yaml:"rateIncreaseCriteria"`
}

