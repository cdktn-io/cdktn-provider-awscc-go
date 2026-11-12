// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateAwsJobExecutionsRolloutConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#exponential_rate IotOtaUpdate#exponential_rate}.
	ExponentialRate *IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRate `field:"optional" json:"exponentialRate" yaml:"exponentialRate"`
	// The maximum number of OTA update job executions started per minute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#maximum_per_minute IotOtaUpdate#maximum_per_minute}
	MaximumPerMinute *float64 `field:"optional" json:"maximumPerMinute" yaml:"maximumPerMinute"`
}

