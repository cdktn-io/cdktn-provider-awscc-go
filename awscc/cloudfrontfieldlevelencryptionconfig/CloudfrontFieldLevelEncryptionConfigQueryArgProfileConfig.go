// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontfieldlevelencryptionconfig


type CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfig struct {
	// Whether CloudFront forwards a request when the query argument names a profile that does not exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudfront_field_level_encryption_config#forward_when_query_arg_profile_is_unknown CloudfrontFieldLevelEncryptionConfig#forward_when_query_arg_profile_is_unknown}
	ForwardWhenQueryArgProfileIsUnknown interface{} `field:"required" json:"forwardWhenQueryArgProfileIsUnknown" yaml:"forwardWhenQueryArgProfileIsUnknown"`
	// The query arguments paired with a profile.
	//
	// Required because the service returns an empty list rather than preserving absence, so an omitted value cannot round-trip. An empty list is accepted and is preserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudfront_field_level_encryption_config#query_arg_profiles CloudfrontFieldLevelEncryptionConfig#query_arg_profiles}
	QueryArgProfiles interface{} `field:"required" json:"queryArgProfiles" yaml:"queryArgProfiles"`
}

