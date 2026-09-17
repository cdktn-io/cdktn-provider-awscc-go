// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightanalysis


type QuicksightAnalysisErrors struct {
	// <p>The message associated with the analysis error.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_analysis#message QuicksightAnalysis#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_analysis#type QuicksightAnalysis#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

