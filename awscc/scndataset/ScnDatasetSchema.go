// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataset


type ScnDatasetSchema struct {
	// The list of field details of the dataset schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/scn_dataset#fields ScnDataset#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The name of the dataset schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/scn_dataset#name ScnDataset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The list of primary key fields for the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/scn_dataset#primary_keys ScnDataset#primary_keys}
	PrimaryKeys interface{} `field:"optional" json:"primaryKeys" yaml:"primaryKeys"`
}

