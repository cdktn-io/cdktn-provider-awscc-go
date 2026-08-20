// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferserver


type TransferServerS3StorageOptions struct {
	// Indicates whether optimization to directory listing on S3 servers is used. Disabled by default for compatibility.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/transfer_server#directory_listing_optimization TransferServer#directory_listing_optimization}
	DirectoryListingOptimization *string `field:"optional" json:"directoryListingOptimization" yaml:"directoryListingOptimization"`
}

