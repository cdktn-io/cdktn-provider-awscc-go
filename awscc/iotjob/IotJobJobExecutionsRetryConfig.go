// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjob


type IotJobJobExecutionsRetryConfig struct {
	// The list of criteria that determines how many retries are allowed for each failure type for a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_job#criteria_list IotJob#criteria_list}
	CriteriaList interface{} `field:"optional" json:"criteriaList" yaml:"criteriaList"`
}

