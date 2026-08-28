// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentials struct {
	// The Amazon Resource Name (ARN) of the secret containing the private repository credentials.
	//
	// When you use the Amazon ECS API, CLI, or AWS SDK, if the secret exists in the same Region as the task that you're launching then you can use either the full ARN or the name of the secret. When you use the AWS Management Console, you must specify the full ARN of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_daemon_task_definition#credentials_parameter EcsDaemonTaskDefinition#credentials_parameter}
	CredentialsParameter *string `field:"optional" json:"credentialsParameter" yaml:"credentialsParameter"`
}

