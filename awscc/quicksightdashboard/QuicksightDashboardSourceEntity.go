// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdashboard


type QuicksightDashboardSourceEntity struct {
	// <p>Dashboard source template.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_dashboard#source_template QuicksightDashboard#source_template}
	SourceTemplate *QuicksightDashboardSourceEntitySourceTemplate `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

