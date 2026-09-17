// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionEcsPropertiesTaskPropertiesRuntimePlatform struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/batch_job_definition#cpu_architecture BatchJobDefinition#cpu_architecture}.
	CpuArchitecture *string `field:"optional" json:"cpuArchitecture" yaml:"cpuArchitecture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/batch_job_definition#operating_system_family BatchJobDefinition#operating_system_family}.
	OperatingSystemFamily *string `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
}

