// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewjob


type DatabrewJobRecipe struct {
	// Recipe name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/databrew_job#name DatabrewJob#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Recipe version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/databrew_job#version DatabrewJob#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

