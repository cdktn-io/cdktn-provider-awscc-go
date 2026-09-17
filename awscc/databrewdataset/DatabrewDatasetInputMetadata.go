// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewdataset


type DatabrewDatasetInputMetadata struct {
	// Arn of the source of the dataset. For e.g.: AppFlow Flow ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/databrew_dataset#source_arn DatabrewDataset#source_arn}
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

