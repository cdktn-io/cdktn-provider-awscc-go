// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewproject


type DatabrewProjectSample struct {
	// Sample size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/databrew_project#size DatabrewProject#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Sample type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/databrew_project#type DatabrewProject#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

