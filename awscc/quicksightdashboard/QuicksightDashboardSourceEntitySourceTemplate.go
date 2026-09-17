// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdashboard


type QuicksightDashboardSourceEntitySourceTemplate struct {
	// <p>The Amazon Resource Name (ARN) of the resource.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_dashboard#arn QuicksightDashboard#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// <p>Dataset references.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_dashboard#data_set_references QuicksightDashboard#data_set_references}
	DataSetReferences interface{} `field:"optional" json:"dataSetReferences" yaml:"dataSetReferences"`
}

