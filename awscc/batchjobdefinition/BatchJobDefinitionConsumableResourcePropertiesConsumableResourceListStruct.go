// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionConsumableResourcePropertiesConsumableResourceListStruct struct {
	// The ARN of the consumable resource the job definition should consume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#consumable_resource BatchJobDefinition#consumable_resource}
	ConsumableResource *string `field:"optional" json:"consumableResource" yaml:"consumableResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#quantity BatchJobDefinition#quantity}.
	Quantity *float64 `field:"optional" json:"quantity" yaml:"quantity"`
}

