// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdashboard


type QuicksightDashboardParametersIntegerParameters struct {
	// <p>The name of the integer parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_dashboard#name QuicksightDashboard#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// <p>The values for the integer parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_dashboard#values QuicksightDashboard#values}
	Values *[]*float64 `field:"optional" json:"values" yaml:"values"`
}

