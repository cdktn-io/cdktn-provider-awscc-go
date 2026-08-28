// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance


type Ec2InstanceLicenseSpecifications struct {
	// The Amazon Resource Name (ARN) of the license configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance#license_configuration_arn Ec2Instance#license_configuration_arn}
	LicenseConfigurationArn *string `field:"optional" json:"licenseConfigurationArn" yaml:"licenseConfigurationArn"`
}

