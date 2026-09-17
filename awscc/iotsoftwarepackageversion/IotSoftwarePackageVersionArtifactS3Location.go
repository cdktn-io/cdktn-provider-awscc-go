// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsoftwarepackageversion


type IotSoftwarePackageVersionArtifactS3Location struct {
	// The S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_software_package_version#bucket IotSoftwarePackageVersion#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_software_package_version#key IotSoftwarePackageVersion#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The S3 version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_software_package_version#version IotSoftwarePackageVersion#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

