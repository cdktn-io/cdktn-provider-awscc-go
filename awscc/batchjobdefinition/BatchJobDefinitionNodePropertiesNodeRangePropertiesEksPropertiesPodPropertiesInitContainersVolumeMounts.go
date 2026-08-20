// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesInitContainersVolumeMounts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#mount_path BatchJobDefinition#mount_path}.
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#name BatchJobDefinition#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#read_only BatchJobDefinition#read_only}.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#sub_path BatchJobDefinition#sub_path}.
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
}

