// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesVolumes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#empty_dir BatchJobDefinition#empty_dir}.
	EmptyDir *BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#host_path BatchJobDefinition#host_path}.
	HostPath *BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesVolumesHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#name BatchJobDefinition#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#persistent_volume_claim BatchJobDefinition#persistent_volume_claim}.
	PersistentVolumeClaim *BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesVolumesPersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#secret BatchJobDefinition#secret}.
	Secret *BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
}

