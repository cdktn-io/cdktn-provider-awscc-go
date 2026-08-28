// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionEksProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#pod_properties BatchJobDefinition#pod_properties}.
	PodProperties *BatchJobDefinitionEksPropertiesPodProperties `field:"optional" json:"podProperties" yaml:"podProperties"`
}

