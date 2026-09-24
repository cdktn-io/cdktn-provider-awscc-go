// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtraildashboard


type CloudtrailDashboardRefreshScheduleFrequency struct {
	// The frequency unit. Supported values are HOURS and DAYS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudtrail_dashboard#unit CloudtrailDashboard#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// The frequency value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudtrail_dashboard#value CloudtrailDashboard#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

