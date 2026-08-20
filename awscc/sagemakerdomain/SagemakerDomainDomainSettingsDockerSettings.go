// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDomainSettingsDockerSettings struct {
	// The flag to enable/disable docker-proxy server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#enable_docker_access SagemakerDomain#enable_docker_access}
	EnableDockerAccess *string `field:"optional" json:"enableDockerAccess" yaml:"enableDockerAccess"`
	// A list of account id's that would be used to pull images from in VpcOnly mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#vpc_only_trusted_accounts SagemakerDomain#vpc_only_trusted_accounts}
	VpcOnlyTrustedAccounts *[]*string `field:"optional" json:"vpcOnlyTrustedAccounts" yaml:"vpcOnlyTrustedAccounts"`
}

