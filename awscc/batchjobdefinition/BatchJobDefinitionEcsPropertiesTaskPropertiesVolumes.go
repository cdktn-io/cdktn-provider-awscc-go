// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionEcsPropertiesTaskPropertiesVolumes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#efs_volume_configuration BatchJobDefinition#efs_volume_configuration}.
	EfsVolumeConfiguration *BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfiguration `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#host BatchJobDefinition#host}.
	Host *BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesHost `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#name BatchJobDefinition#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#s3_files_volume_configuration BatchJobDefinition#s3_files_volume_configuration}.
	S3FilesVolumeConfiguration *BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesS3FilesVolumeConfiguration `field:"optional" json:"s3FilesVolumeConfiguration" yaml:"s3FilesVolumeConfiguration"`
}

