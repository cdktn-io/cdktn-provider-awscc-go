// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateAwsJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteria struct {
	// When this number of things have been notified, it will initiate an increase in the rollout rate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#number_of_notified_things IotOtaUpdate#number_of_notified_things}
	NumberOfNotifiedThings *float64 `field:"optional" json:"numberOfNotifiedThings" yaml:"numberOfNotifiedThings"`
	// When this number of things have succeeded in their job execution, it will initiate an increase in the rollout rate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#number_of_succeeded_things IotOtaUpdate#number_of_succeeded_things}
	NumberOfSucceededThings *float64 `field:"optional" json:"numberOfSucceededThings" yaml:"numberOfSucceededThings"`
}

