// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjobtemplate


type IotJobTemplateJobExecutionsRolloutConfigExponentialRolloutRate struct {
	// The minimum number of things that will be notified of a pending job, per minute at the start of job rollout.
	//
	// This parameter allows you to define the initial rate of rollout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_job_template#base_rate_per_minute IotJobTemplate#base_rate_per_minute}
	BaseRatePerMinute *float64 `field:"optional" json:"baseRatePerMinute" yaml:"baseRatePerMinute"`
	// The exponential factor to increase the rate of rollout for a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_job_template#increment_factor IotJobTemplate#increment_factor}
	IncrementFactor *float64 `field:"optional" json:"incrementFactor" yaml:"incrementFactor"`
	// The criteria to initiate the increase in rate of rollout for a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_job_template#rate_increase_criteria IotJobTemplate#rate_increase_criteria}
	RateIncreaseCriteria *IotJobTemplateJobExecutionsRolloutConfigExponentialRolloutRateRateIncreaseCriteria `field:"optional" json:"rateIncreaseCriteria" yaml:"rateIncreaseCriteria"`
}

