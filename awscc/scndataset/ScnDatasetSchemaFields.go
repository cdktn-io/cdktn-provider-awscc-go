// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataset


type ScnDatasetSchemaFields struct {
	// Indicate if the field is required or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scn_dataset#is_required ScnDataset#is_required}
	IsRequired interface{} `field:"optional" json:"isRequired" yaml:"isRequired"`
	// The dataset field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scn_dataset#name ScnDataset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The dataset field type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scn_dataset#type ScnDataset#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

