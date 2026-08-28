// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionEcsPropertiesTaskPropertiesContainersLinuxParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#devices BatchJobDefinition#devices}.
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#init_process_enabled BatchJobDefinition#init_process_enabled}.
	InitProcessEnabled interface{} `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#max_swap BatchJobDefinition#max_swap}.
	MaxSwap *float64 `field:"optional" json:"maxSwap" yaml:"maxSwap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#shared_memory_size BatchJobDefinition#shared_memory_size}.
	SharedMemorySize *float64 `field:"optional" json:"sharedMemorySize" yaml:"sharedMemorySize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#swappiness BatchJobDefinition#swappiness}.
	Swappiness *float64 `field:"optional" json:"swappiness" yaml:"swappiness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#tmpfs BatchJobDefinition#tmpfs}.
	Tmpfs interface{} `field:"optional" json:"tmpfs" yaml:"tmpfs"`
}

