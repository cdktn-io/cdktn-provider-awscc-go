// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightanalysis


type QuicksightAnalysisSourceEntitySourceTemplateDataSetReferences struct {
	// <p>Dataset Amazon Resource Name (ARN).</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_analysis#data_set_arn QuicksightAnalysis#data_set_arn}
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
	// <p>Dataset placeholder.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_analysis#data_set_placeholder QuicksightAnalysis#data_set_placeholder}
	DataSetPlaceholder *string `field:"optional" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

