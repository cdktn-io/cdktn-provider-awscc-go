// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchdomainfororganization


type CloudwatchDomainForOrganizationIdentityProviderConfiguration struct {
	// Identity Center configuration. Required when identityProviders includes IDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain_for_organization#identity_center_configuration CloudwatchDomainForOrganization#identity_center_configuration}
	IdentityCenterConfiguration *CloudwatchDomainForOrganizationIdentityProviderConfigurationIdentityCenterConfiguration `field:"optional" json:"identityCenterConfiguration" yaml:"identityCenterConfiguration"`
}

