// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisedataset


type IotsitewiseDatasetDatasetSourceSourceDetail struct {
	// Contains details about the Kendra dataset source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_dataset#kendra IotsitewiseDataset#kendra}
	Kendra *IotsitewiseDatasetDatasetSourceSourceDetailKendra `field:"optional" json:"kendra" yaml:"kendra"`
}

