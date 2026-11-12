// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRate struct {
	// The minimum number of things that will be notified of a pending job, per minute, at the start of the job rollout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#base_rate_per_minute IotOtaUpdate#base_rate_per_minute}
	BaseRatePerMinute *float64 `field:"optional" json:"baseRatePerMinute" yaml:"baseRatePerMinute"`
	// The rate of increase for a job rollout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#increment_factor IotOtaUpdate#increment_factor}
	IncrementFactor *float64 `field:"optional" json:"incrementFactor" yaml:"incrementFactor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#rate_increase_criteria IotOtaUpdate#rate_increase_criteria}.
	RateIncreaseCriteria *IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteria `field:"optional" json:"rateIncreaseCriteria" yaml:"rateIncreaseCriteria"`
}

