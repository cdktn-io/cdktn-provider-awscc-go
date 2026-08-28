// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationazureblob


type DatasyncLocationAzureBlobTags struct {
	// The key for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datasync_location_azure_blob#key DatasyncLocationAzureBlob#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datasync_location_azure_blob#value DatasyncLocationAzureBlob#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

