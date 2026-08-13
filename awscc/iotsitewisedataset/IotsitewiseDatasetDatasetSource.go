// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisedataset


type IotsitewiseDatasetDatasetSource struct {
	// The format of the dataset source associated with the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_dataset#source_format IotsitewiseDataset#source_format}
	SourceFormat *string `field:"required" json:"sourceFormat" yaml:"sourceFormat"`
	// The type of data source for the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_dataset#source_type IotsitewiseDataset#source_type}
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// The details of the dataset source associated with the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_dataset#source_detail IotsitewiseDataset#source_detail}
	SourceDetail *IotsitewiseDatasetDatasetSourceSourceDetail `field:"optional" json:"sourceDetail" yaml:"sourceDetail"`
}

