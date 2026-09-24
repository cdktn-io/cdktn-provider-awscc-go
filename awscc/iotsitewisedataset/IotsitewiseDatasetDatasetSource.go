// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisedataset


type IotsitewiseDatasetDatasetSource struct {
	// The details of the dataset source associated with the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_dataset#source_detail IotsitewiseDataset#source_detail}
	SourceDetail *IotsitewiseDatasetDatasetSourceSourceDetail `field:"optional" json:"sourceDetail" yaml:"sourceDetail"`
	// The format of the dataset source associated with the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_dataset#source_format IotsitewiseDataset#source_format}
	SourceFormat *string `field:"optional" json:"sourceFormat" yaml:"sourceFormat"`
	// The type of data source for the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_dataset#source_type IotsitewiseDataset#source_type}
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

