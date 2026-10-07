// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontfieldlevelencryptionprofile


type CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigEncryptionEntities struct {
	// The request-body field names to encrypt.
	//
	// A pattern is either a full field name or leading characters followed by a wildcard (*). Patterns are case-sensitive and must not overlap.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/cloudfront_field_level_encryption_profile#field_patterns CloudfrontFieldLevelEncryptionProfile#field_patterns}
	FieldPatterns *[]*string `field:"required" json:"fieldPatterns" yaml:"fieldPatterns"`
	// The provider associated with the public key.
	//
	// The same value must be supplied with the private key for an application to decrypt the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/cloudfront_field_level_encryption_profile#provider_id CloudfrontFieldLevelEncryptionProfile#provider_id}
	ProviderId *string `field:"required" json:"providerId" yaml:"providerId"`
	// The identifier of the CloudFront public key used to encrypt the fields that match the patterns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/cloudfront_field_level_encryption_profile#public_key_id CloudfrontFieldLevelEncryptionProfile#public_key_id}
	PublicKeyId *string `field:"required" json:"publicKeyId" yaml:"publicKeyId"`
}

