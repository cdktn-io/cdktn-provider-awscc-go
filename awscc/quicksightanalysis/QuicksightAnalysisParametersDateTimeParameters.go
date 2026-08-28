// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightanalysis


type QuicksightAnalysisParametersDateTimeParameters struct {
	// <p>A display name for the date-time parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_analysis#name QuicksightAnalysis#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// <p>The values for the date-time parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_analysis#values QuicksightAnalysis#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

