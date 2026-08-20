// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjob


type IotJobAbortConfigCriteriaListStruct struct {
	// The type of job action to take to initiate the job abort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_job#action IotJob#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The type of job execution failures that can initiate a job abort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_job#failure_type IotJob#failure_type}
	FailureType *string `field:"optional" json:"failureType" yaml:"failureType"`
	// The minimum number of things which must receive job execution notifications before the job can be aborted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_job#min_number_of_executed_things IotJob#min_number_of_executed_things}
	MinNumberOfExecutedThings *float64 `field:"optional" json:"minNumberOfExecutedThings" yaml:"minNumberOfExecutedThings"`
	// The minimum percentage of job execution failures that must occur to initiate the job abort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_job#threshold_percentage IotJob#threshold_percentage}
	ThresholdPercentage *float64 `field:"optional" json:"thresholdPercentage" yaml:"thresholdPercentage"`
}

