// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchdomainfororganization


type CloudwatchDomainForOrganizationIdentityProviderConfigurationIdentityCenterConfiguration struct {
	// Identity Center instance ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain_for_organization#identity_center_instance_arn CloudwatchDomainForOrganization#identity_center_instance_arn}
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
}

