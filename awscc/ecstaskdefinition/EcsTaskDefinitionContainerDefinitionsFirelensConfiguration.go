// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsFirelensConfiguration struct {
	// The options to use when configuring the log router.
	//
	// This field is optional and can be used to add additional metadata, such as the task, task definition, cluster, and container instance details to the log event.
	//   If specified, valid option keys are:
	//   +  ``enable-ecs-log-metadata``, which can be ``true`` or ``false``
	//   +  ``config-file-type``, which can be ``s3`` or ``file``
	//   +  ``config-file-value``, which is either an S3 ARN or a file path
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_task_definition#options EcsTaskDefinition#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// The log router to use. The valid values are ``fluentd`` or ``fluentbit``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_task_definition#type EcsTaskDefinition#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

