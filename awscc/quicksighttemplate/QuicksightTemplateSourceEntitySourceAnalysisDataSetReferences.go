// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttemplate


type QuicksightTemplateSourceEntitySourceAnalysisDataSetReferences struct {
	// <p>Dataset Amazon Resource Name (ARN).</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_template#data_set_arn QuicksightTemplate#data_set_arn}
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
	// <p>Dataset placeholder.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_template#data_set_placeholder QuicksightTemplate#data_set_placeholder}
	DataSetPlaceholder *string `field:"optional" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

