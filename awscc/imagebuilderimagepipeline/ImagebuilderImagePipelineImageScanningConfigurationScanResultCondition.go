// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagepipeline


type ImagebuilderImagePipelineImageScanningConfigurationScanResultCondition struct {
	// Specifies whether only findings that have a fix available count toward MaximumSeverityLabel.
	//
	// When false, Image Builder counts every finding. The default is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/imagebuilder_image_pipeline#fix_available ImagebuilderImagePipeline#fix_available}
	FixAvailable interface{} `field:"optional" json:"fixAvailable" yaml:"fixAvailable"`
	// The maximum severity allowed for findings on the output AMI.
	//
	// A finding above this severity fails the image, unless FixAvailable excludes it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/imagebuilder_image_pipeline#maximum_severity_label ImagebuilderImagePipeline#maximum_severity_label}
	MaximumSeverityLabel *string `field:"optional" json:"maximumSeverityLabel" yaml:"maximumSeverityLabel"`
}

