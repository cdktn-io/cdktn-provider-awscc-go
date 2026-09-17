// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendraindex


type KendraIndexCapacityUnits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_index#query_capacity_units KendraIndex#query_capacity_units}.
	QueryCapacityUnits *float64 `field:"optional" json:"queryCapacityUnits" yaml:"queryCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_index#storage_capacity_units KendraIndex#storage_capacity_units}.
	StorageCapacityUnits *float64 `field:"optional" json:"storageCapacityUnits" yaml:"storageCapacityUnits"`
}

