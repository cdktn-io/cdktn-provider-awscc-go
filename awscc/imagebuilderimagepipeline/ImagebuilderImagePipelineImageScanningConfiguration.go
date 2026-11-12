// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagepipeline


type ImagebuilderImagePipelineImageScanningConfiguration struct {
	// Contains ECR settings for vulnerability scans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/imagebuilder_image_pipeline#ecr_configuration ImagebuilderImagePipeline#ecr_configuration}
	EcrConfiguration *ImagebuilderImagePipelineImageScanningConfigurationEcrConfiguration `field:"optional" json:"ecrConfiguration" yaml:"ecrConfiguration"`
	// This sets whether Image Builder keeps a snapshot of the vulnerability scans that Amazon Inspector runs against the test instance when you create a new image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/imagebuilder_image_pipeline#image_scanning_enabled ImagebuilderImagePipeline#image_scanning_enabled}
	ImageScanningEnabled interface{} `field:"optional" json:"imageScanningEnabled" yaml:"imageScanningEnabled"`
	// The type of vulnerability scan that Image Builder runs for AMI output when ImageScanningEnabled is true.
	//
	// Specify a single value. If you omit ImageScanningTypes, Image Builder uses TEST_INSTANCE_CVE. With TEST_INSTANCE_CVE, Amazon Inspector scans the test instance during image tests, and you must activate Amazon Inspector for Amazon EC2 scanning in your account. With ON_DEMAND_AMI_CVE, Amazon Inspector starts an on-demand scan of the output AMI during image tests. You must build from an image recipe, but you do not need to activate Amazon Inspector in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/imagebuilder_image_pipeline#image_scanning_types ImagebuilderImagePipeline#image_scanning_types}
	ImageScanningTypes *[]*string `field:"optional" json:"imageScanningTypes" yaml:"imageScanningTypes"`
	// The condition that the output AMI's scan results must meet for the image to pass.
	//
	// You can set ScanResultCondition only when ImageScanningTypes is ON_DEMAND_AMI_CVE. If you omit it, Image Builder starts the scan but does not wait for it or use its results to pass or fail the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/imagebuilder_image_pipeline#scan_result_condition ImagebuilderImagePipeline#scan_result_condition}
	ScanResultCondition *ImagebuilderImagePipelineImageScanningConfigurationScanResultCondition `field:"optional" json:"scanResultCondition" yaml:"scanResultCondition"`
}

