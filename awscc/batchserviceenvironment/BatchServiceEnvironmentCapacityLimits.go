// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchserviceenvironment


type BatchServiceEnvironmentCapacityLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/batch_service_environment#capacity_unit BatchServiceEnvironment#capacity_unit}.
	CapacityUnit *string `field:"optional" json:"capacityUnit" yaml:"capacityUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/batch_service_environment#max_capacity BatchServiceEnvironment#max_capacity}.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
}

