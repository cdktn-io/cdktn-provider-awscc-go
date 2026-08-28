// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelensgroup


type S3StorageLensGroupFilterOrMatchObjectSize struct {
	// Minimum object size to which the rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens_group#bytes_greater_than S3StorageLensGroup#bytes_greater_than}
	BytesGreaterThan *float64 `field:"optional" json:"bytesGreaterThan" yaml:"bytesGreaterThan"`
	// Maximum object size to which the rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens_group#bytes_less_than S3StorageLensGroup#bytes_less_than}
	BytesLessThan *float64 `field:"optional" json:"bytesLessThan" yaml:"bytesLessThan"`
}

