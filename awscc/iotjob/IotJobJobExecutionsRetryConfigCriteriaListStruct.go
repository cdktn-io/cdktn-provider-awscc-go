// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjob


type IotJobJobExecutionsRetryConfigCriteriaListStruct struct {
	// The type of job execution failures that can initiate a job retry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#failure_type IotJob#failure_type}
	FailureType *string `field:"optional" json:"failureType" yaml:"failureType"`
	// The number of retries allowed for a failure type for the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#number_of_retries IotJob#number_of_retries}
	NumberOfRetries *float64 `field:"optional" json:"numberOfRetries" yaml:"numberOfRetries"`
}

