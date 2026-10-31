// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchdomain


type CloudwatchDomainIdentityProviderConfiguration struct {
	// Identity Center configuration. Required when identityProviders includes IDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/cloudwatch_domain#identity_center_configuration CloudwatchDomain#identity_center_configuration}
	IdentityCenterConfiguration *CloudwatchDomainIdentityProviderConfigurationIdentityCenterConfiguration `field:"optional" json:"identityCenterConfiguration" yaml:"identityCenterConfiguration"`
}

