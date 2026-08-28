// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewjob


type DatabrewJobJobSample struct {
	// Sample configuration mode for profile jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/databrew_job#mode DatabrewJob#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Sample configuration size for profile jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/databrew_job#size DatabrewJob#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

