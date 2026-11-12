// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateAwsJobAbortConfig struct {
	// The list of criteria that determine when and how to abort the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#abort_criteria_list IotOtaUpdate#abort_criteria_list}
	AbortCriteriaList interface{} `field:"optional" json:"abortCriteriaList" yaml:"abortCriteriaList"`
}

