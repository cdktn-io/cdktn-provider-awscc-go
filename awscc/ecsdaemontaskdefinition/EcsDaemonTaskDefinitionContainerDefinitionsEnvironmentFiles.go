// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFiles struct {
	// The file type to use. Environment files are objects in Amazon S3. The only supported value is ``s3``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#type EcsDaemonTaskDefinition#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The Amazon Resource Name (ARN) of the Amazon S3 object containing the environment variable file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#value EcsDaemonTaskDefinition#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

