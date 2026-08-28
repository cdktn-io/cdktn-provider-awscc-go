// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagepipeline


type ImagebuilderImagePipelineWorkflowsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_pipeline#name ImagebuilderImagePipeline#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_pipeline#value ImagebuilderImagePipeline#value}.
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

