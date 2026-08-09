// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightanalysis


type QuicksightAnalysisSourceEntitySourceTemplate struct {
	// <p>The Amazon Resource Name (ARN) of the source template of an analysis.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_analysis#arn QuicksightAnalysis#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// <p>The dataset references of the source template of an analysis.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_analysis#data_set_references QuicksightAnalysis#data_set_references}
	DataSetReferences interface{} `field:"optional" json:"dataSetReferences" yaml:"dataSetReferences"`
}

