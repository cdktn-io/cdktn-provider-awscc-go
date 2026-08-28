// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdashboard


type QuicksightDashboardParametersDecimalParameters struct {
	// <p>A display name for the decimal parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_dashboard#name QuicksightDashboard#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// <p>The values for the decimal parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_dashboard#values QuicksightDashboard#values}
	Values *[]*float64 `field:"optional" json:"values" yaml:"values"`
}

