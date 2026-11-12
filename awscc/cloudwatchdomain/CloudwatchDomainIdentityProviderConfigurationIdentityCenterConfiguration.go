// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchdomain


type CloudwatchDomainIdentityProviderConfigurationIdentityCenterConfiguration struct {
	// Identity Center instance ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain#identity_center_instance_arn CloudwatchDomain#identity_center_instance_arn}
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
}

