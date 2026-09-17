// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagepipeline


type ImagebuilderImagePipelineImageScanningConfiguration struct {
	// Contains ECR settings for vulnerability scans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_image_pipeline#ecr_configuration ImagebuilderImagePipeline#ecr_configuration}
	EcrConfiguration *ImagebuilderImagePipelineImageScanningConfigurationEcrConfiguration `field:"optional" json:"ecrConfiguration" yaml:"ecrConfiguration"`
	// This sets whether Image Builder keeps a snapshot of the vulnerability scans that Amazon Inspector runs against the test instance when you create a new image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_image_pipeline#image_scanning_enabled ImagebuilderImagePipeline#image_scanning_enabled}
	ImageScanningEnabled interface{} `field:"optional" json:"imageScanningEnabled" yaml:"imageScanningEnabled"`
}

