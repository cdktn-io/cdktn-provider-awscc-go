// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsworkspace


type ApsWorkspaceQueryLoggingConfigurationDestinationsFilters struct {
	// Query logs with QSP above this limit are vended.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_workspace#qsp_threshold ApsWorkspace#qsp_threshold}
	QspThreshold *float64 `field:"optional" json:"qspThreshold" yaml:"qspThreshold"`
}

