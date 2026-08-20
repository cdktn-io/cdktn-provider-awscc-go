// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjob


type IotJobAbortConfig struct {
	// The list of criteria that determine when and how to abort the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_job#criteria_list IotJob#criteria_list}
	CriteriaList interface{} `field:"optional" json:"criteriaList" yaml:"criteriaList"`
}

