// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontfieldlevelencryptionconfig


type CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfiles struct {
	// The identifier of the field-level encryption profile CloudFront uses for this query argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudfront_field_level_encryption_config#profile_id CloudfrontFieldLevelEncryptionConfig#profile_id}
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// The query argument that CloudFront encrypts with the paired profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudfront_field_level_encryption_config#query_arg CloudfrontFieldLevelEncryptionConfig#query_arg}
	QueryArg *string `field:"required" json:"queryArg" yaml:"queryArg"`
}

