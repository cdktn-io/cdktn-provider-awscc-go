// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsworkspace


type ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsLimits struct {
	// The maximum number of active series that can be ingested for this label set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aps_workspace#max_series ApsWorkspace#max_series}
	MaxSeries *float64 `field:"optional" json:"maxSeries" yaml:"maxSeries"`
}

