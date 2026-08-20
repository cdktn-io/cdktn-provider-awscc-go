// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord


type AgentregistryRegistryRecordDescriptorsMcpServerSourceFromUrlCredentialProviderConfigurationsCredentialProviderIamCredentialProvider struct {
	// The SigV4 signing region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/agentregistry_registry_record#region AgentregistryRegistryRecord#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The ARN of the IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/agentregistry_registry_record#role_arn AgentregistryRegistryRecord#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The SigV4 signing service name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/agentregistry_registry_record#service AgentregistryRegistryRecord#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

