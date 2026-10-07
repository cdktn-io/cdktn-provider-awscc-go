// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontfieldlevelencryptionprofile


type CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfig struct {
	// A unique value that identifies the creation request.
	//
	// Caller references are unique within an AWS account and cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/cloudfront_field_level_encryption_profile#caller_reference CloudfrontFieldLevelEncryptionProfile#caller_reference}
	CallerReference *string `field:"required" json:"callerReference" yaml:"callerReference"`
	// The encryption entities of the field-level encryption profile. At least one entity is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/cloudfront_field_level_encryption_profile#encryption_entities CloudfrontFieldLevelEncryptionProfile#encryption_entities}
	EncryptionEntities interface{} `field:"required" json:"encryptionEntities" yaml:"encryptionEntities"`
	// The name of the field-level encryption profile. Names are unique within an AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/cloudfront_field_level_encryption_profile#name CloudfrontFieldLevelEncryptionProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional comment describing the field-level encryption profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/cloudfront_field_level_encryption_profile#comment CloudfrontFieldLevelEncryptionProfile#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

