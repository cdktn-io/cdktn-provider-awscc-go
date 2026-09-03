// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionContainerDefinitionsFirelensConfiguration struct {
	// The options to use when configuring the log router.
	//
	// This field is optional and can be used to specify a custom configuration file or to add additional metadata, such as the task, task definition, cluster, and container instance details to the log event. If specified, the syntax to use is ``"options":{"enable-ecs-log-metadata":"true|false","config-file-type:"s3|file","config-file-value":"arn:aws:s3:::mybucket/fluent.conf|filepath"}``. For more information, see [Creating a task definition that uses a FireLens configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef) in the *Amazon Elastic Container Service Developer Guide*.
	//   Tasks hosted on FARGATElong only support the ``file`` configuration file type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecs_daemon_task_definition#options EcsDaemonTaskDefinition#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// The log router to use. The valid values are ``fluentd`` or ``fluentbit``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecs_daemon_task_definition#type EcsDaemonTaskDefinition#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

