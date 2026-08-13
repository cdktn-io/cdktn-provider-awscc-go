// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionEcsProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/batch_job_definition#task_properties BatchJobDefinition#task_properties}.
	TaskProperties interface{} `field:"optional" json:"taskProperties" yaml:"taskProperties"`
}

