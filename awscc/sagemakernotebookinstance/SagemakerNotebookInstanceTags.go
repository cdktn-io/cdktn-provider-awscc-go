// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakernotebookinstance


type SagemakerNotebookInstanceTags struct {
	// The tag key. Tag keys must be unique per resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_notebook_instance#key SagemakerNotebookInstance#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_notebook_instance#value SagemakerNotebookInstance#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

