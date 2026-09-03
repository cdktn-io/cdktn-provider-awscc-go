// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightanalysis


type QuicksightAnalysisParametersDecimalParameters struct {
	// <p>A display name for the decimal parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_analysis#name QuicksightAnalysis#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// <p>The values for the decimal parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_analysis#values QuicksightAnalysis#values}
	Values *[]*float64 `field:"optional" json:"values" yaml:"values"`
}

