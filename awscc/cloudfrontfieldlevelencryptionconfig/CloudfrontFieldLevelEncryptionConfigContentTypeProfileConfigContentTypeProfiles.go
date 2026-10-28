// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontfieldlevelencryptionconfig


type CloudfrontFieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfiles struct {
	// The content type CloudFront encrypts. The service accepts only application/x-www-form-urlencoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudfront_field_level_encryption_config#content_type CloudfrontFieldLevelEncryptionConfig#content_type}
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// The encryption format CloudFront uses for this content type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudfront_field_level_encryption_config#format CloudfrontFieldLevelEncryptionConfig#format}
	Format *string `field:"required" json:"format" yaml:"format"`
	// The identifier of the field-level encryption profile CloudFront uses for this content type.
	//
	// Required, so that it cannot be omitted: the service accepts a content-type profile with the member absent on create and refuses one on update, and an update sends the whole profile list, so a configuration created that way could never be updated again. Supply an empty string to configure a content type and forward it WITHOUT field-level encryption. That mode is available through this resource and round-trips: create, update and read all accept and return the empty string, and it is the service's own encoding for the state -- asked for a content-type list with no items, CloudFront returns one item whose ProfileId is empty. Note that this is not the same as omitting the content type from the configuration, because ForwardWhenContentTypeIsUnknown governs only content types absent from it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudfront_field_level_encryption_config#profile_id CloudfrontFieldLevelEncryptionConfig#profile_id}
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
}

