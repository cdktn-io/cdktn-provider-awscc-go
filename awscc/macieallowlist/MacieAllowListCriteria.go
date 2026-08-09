// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package macieallowlist


type MacieAllowListCriteria struct {
	// The S3 object key for the AllowList.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/macie_allow_list#regex MacieAllowList#regex}
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// The S3 location for the AllowList.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/macie_allow_list#s3_words_list MacieAllowList#s3_words_list}
	S3WordsList *MacieAllowListCriteriaS3WordsListStruct `field:"optional" json:"s3WordsList" yaml:"s3WordsList"`
}

