// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskExcludes struct {
	// The type of filter rule to apply. AWS DataSync only supports the SIMPLE_PATTERN rule type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datasync_task#filter_type DatasyncTask#filter_type}
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
	// A single filter string that consists of the patterns to include or exclude. The patterns are delimited by "|".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datasync_task#value DatasyncTask#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

