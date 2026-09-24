// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionVolumesFsxWindowsFileServerVolumeConfigurationAuthorizationConfig struct {
	// The authorization credential option to use.
	//
	// The authorization credential options can be provided using either the Amazon Resource Name (ARN) of an ASMlong secret or SSM Parameter Store parameter. The ARN refers to the stored credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_task_definition#credentials_parameter EcsTaskDefinition#credentials_parameter}
	CredentialsParameter *string `field:"optional" json:"credentialsParameter" yaml:"credentialsParameter"`
	// A fully qualified domain name hosted by an [](https://docs.aws.amazon.com/directoryservice/latest/admin-guide/directory_microsoft_ad.html) Managed Microsoft AD (Active Directory) or self-hosted AD on Amazon EC2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_task_definition#domain EcsTaskDefinition#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

