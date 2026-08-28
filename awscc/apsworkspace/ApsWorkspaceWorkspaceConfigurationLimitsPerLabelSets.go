// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsworkspace


type ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSets struct {
	// An array of series labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aps_workspace#label_set ApsWorkspace#label_set}
	LabelSet interface{} `field:"optional" json:"labelSet" yaml:"labelSet"`
	// Limits that can be applied to a label set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aps_workspace#limits ApsWorkspace#limits}
	Limits *ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsLimits `field:"optional" json:"limits" yaml:"limits"`
}

