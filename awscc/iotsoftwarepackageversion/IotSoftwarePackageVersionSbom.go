// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsoftwarepackageversion


type IotSoftwarePackageVersionSbom struct {
	// The Amazon S3 location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_software_package_version#s3_location IotSoftwarePackageVersion#s3_location}
	S3Location *IotSoftwarePackageVersionSbomS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

