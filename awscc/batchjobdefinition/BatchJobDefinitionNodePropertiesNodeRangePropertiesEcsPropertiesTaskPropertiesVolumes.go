// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesVolumes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#efs_volume_configuration BatchJobDefinition#efs_volume_configuration}.
	EfsVolumeConfiguration *BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesVolumesEfsVolumeConfiguration `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#host BatchJobDefinition#host}.
	Host *BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesVolumesHost `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#name BatchJobDefinition#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#s3_files_volume_configuration BatchJobDefinition#s3_files_volume_configuration}.
	S3FilesVolumeConfiguration *BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesVolumesS3FilesVolumeConfiguration `field:"optional" json:"s3FilesVolumeConfiguration" yaml:"s3FilesVolumeConfiguration"`
}

