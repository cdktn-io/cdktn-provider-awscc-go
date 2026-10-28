// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontfieldlevelencryptionconfig


type CloudfrontFieldLevelEncryptionConfigContentTypeProfileConfig struct {
	// The content types paired with a profile.
	//
	// Required, and at least one item, because the service synthesizes one default item when the list is absent or empty, so neither state round-trips.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudfront_field_level_encryption_config#content_type_profiles CloudfrontFieldLevelEncryptionConfig#content_type_profiles}
	ContentTypeProfiles interface{} `field:"required" json:"contentTypeProfiles" yaml:"contentTypeProfiles"`
	// Whether CloudFront forwards a request when the content type names a profile that does not exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudfront_field_level_encryption_config#forward_when_content_type_is_unknown CloudfrontFieldLevelEncryptionConfig#forward_when_content_type_is_unknown}
	ForwardWhenContentTypeIsUnknown interface{} `field:"required" json:"forwardWhenContentTypeIsUnknown" yaml:"forwardWhenContentTypeIsUnknown"`
}

