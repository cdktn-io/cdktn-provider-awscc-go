// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#container_path BatchJobDefinition#container_path}.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#host_path BatchJobDefinition#host_path}.
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#permissions BatchJobDefinition#permissions}.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
}

