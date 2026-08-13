// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDomainSettings struct {
	// A collection of settings that are required to start docker-proxy server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#docker_settings SagemakerDomain#docker_settings}
	DockerSettings *SagemakerDomainDomainSettingsDockerSettings `field:"optional" json:"dockerSettings" yaml:"dockerSettings"`
	// The configuration for attaching a SageMaker user profile name to the execution role as a sts:SourceIdentity key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#execution_role_identity_config SagemakerDomain#execution_role_identity_config}
	ExecutionRoleIdentityConfig *string `field:"optional" json:"executionRoleIdentityConfig" yaml:"executionRoleIdentityConfig"`
	// A setting to indicate if IPv6 routing should be enabled along with IPv4 or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#ip_address_type SagemakerDomain#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// A collection of settings that update the current configuration for the RStudioServerPro Domain-level app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#r_studio_server_pro_domain_settings SagemakerDomain#r_studio_server_pro_domain_settings}
	RStudioServerProDomainSettings *SagemakerDomainDomainSettingsRStudioServerProDomainSettings `field:"optional" json:"rStudioServerProDomainSettings" yaml:"rStudioServerProDomainSettings"`
	// The security groups for the Amazon Virtual Private Cloud that the Domain uses for communication between Domain-level apps and user apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#security_group_ids SagemakerDomain#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A collection of settings that apply to an Amazon SageMaker AI domain when you use it in Amazon SageMaker Unified Studio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#unified_studio_settings SagemakerDomain#unified_studio_settings}
	UnifiedStudioSettings *SagemakerDomainDomainSettingsUnifiedStudioSettings `field:"optional" json:"unifiedStudioSettings" yaml:"unifiedStudioSettings"`
}

