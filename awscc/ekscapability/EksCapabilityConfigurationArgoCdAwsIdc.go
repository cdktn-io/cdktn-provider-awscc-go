// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscapability


type EksCapabilityConfigurationArgoCdAwsIdc struct {
	// The ARN of the IAM Identity Center instance to use for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_capability#idc_instance_arn EksCapability#idc_instance_arn}
	IdcInstanceArn *string `field:"optional" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// The Region where your IAM Identity Center instance is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_capability#idc_region EksCapability#idc_region}
	IdcRegion *string `field:"optional" json:"idcRegion" yaml:"idcRegion"`
}

