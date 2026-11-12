// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateAwsJobAbortConfigAbortCriteriaListStruct struct {
	// The type of job action to take to initiate the job abort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#action IotOtaUpdate#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The type of job execution failures that can initiate a job abort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#failure_type IotOtaUpdate#failure_type}
	FailureType *string `field:"optional" json:"failureType" yaml:"failureType"`
	// The minimum number of things which must receive job execution notifications before the job can be aborted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#min_number_of_executed_things IotOtaUpdate#min_number_of_executed_things}
	MinNumberOfExecutedThings *float64 `field:"optional" json:"minNumberOfExecutedThings" yaml:"minNumberOfExecutedThings"`
	// The minimum percentage of job execution failures that must occur to initiate the job abort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#threshold_percentage IotOtaUpdate#threshold_percentage}
	ThresholdPercentage *float64 `field:"optional" json:"thresholdPercentage" yaml:"thresholdPercentage"`
}

