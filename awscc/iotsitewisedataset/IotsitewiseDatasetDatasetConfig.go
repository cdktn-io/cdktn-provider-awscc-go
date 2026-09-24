// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisedataset


type IotsitewiseDatasetDatasetConfig struct {
	// The session configuration for a SESSION dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_dataset#session IotsitewiseDataset#session}
	Session *IotsitewiseDatasetDatasetConfigSession `field:"optional" json:"session" yaml:"session"`
}

