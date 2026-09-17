// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type AwsccProviderEndpoints struct {
	// Use this to override the default Cloud Control API service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs#cloudcontrolapi AwsccProvider#cloudcontrolapi}
	Cloudcontrolapi *string `field:"optional" json:"cloudcontrolapi" yaml:"cloudcontrolapi"`
	// Use this to override the default IAM service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs#iam AwsccProvider#iam}
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
	// Use this to override the default SSO service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs#sso AwsccProvider#sso}
	Sso *string `field:"optional" json:"sso" yaml:"sso"`
	// Use this to override the default STS service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs#sts AwsccProvider#sts}
	Sts *string `field:"optional" json:"sts" yaml:"sts"`
}

