// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package macieallowlist


type MacieAllowListCriteriaS3WordsListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/macie_allow_list#bucket_name MacieAllowList#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/macie_allow_list#object_key MacieAllowList#object_key}.
	ObjectKey *string `field:"optional" json:"objectKey" yaml:"objectKey"`
}

