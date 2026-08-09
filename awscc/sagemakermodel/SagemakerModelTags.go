// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelTags struct {
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model#key SagemakerModel#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag key. Tag keys must be unique per resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model#value SagemakerModel#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

