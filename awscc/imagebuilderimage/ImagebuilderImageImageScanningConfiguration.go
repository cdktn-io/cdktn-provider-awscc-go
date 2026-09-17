// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimage


type ImagebuilderImageImageScanningConfiguration struct {
	// Contains ECR settings for vulnerability scans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_image#ecr_configuration ImagebuilderImage#ecr_configuration}
	EcrConfiguration *ImagebuilderImageImageScanningConfigurationEcrConfiguration `field:"optional" json:"ecrConfiguration" yaml:"ecrConfiguration"`
	// This sets whether Image Builder keeps a snapshot of the vulnerability scans that Amazon Inspector runs against the test instance when you create a new image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_image#image_scanning_enabled ImagebuilderImage#image_scanning_enabled}
	ImageScanningEnabled interface{} `field:"optional" json:"imageScanningEnabled" yaml:"imageScanningEnabled"`
}

