// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontfieldlevelencryptionconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudfrontFieldLevelEncryptionConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A unique value that identifies the create request.
	//
	// Any string unique within the account works; CloudFront refuses to change it once set, so supplying a different value replaces the configuration. It is also the only value CloudFront returns that identifies a configuration as this stack's, so cleaning up after a failed create depends on the template keeping it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudfront_field_level_encryption_config#caller_reference CloudfrontFieldLevelEncryptionConfig#caller_reference}
	CallerReference *string `field:"required" json:"callerReference" yaml:"callerReference"`
	// How CloudFront handles content types, including the profile paired with each one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudfront_field_level_encryption_config#content_type_profile_config CloudfrontFieldLevelEncryptionConfig#content_type_profile_config}
	ContentTypeProfileConfig *CloudfrontFieldLevelEncryptionConfigContentTypeProfileConfig `field:"required" json:"contentTypeProfileConfig" yaml:"contentTypeProfileConfig"`
	// How CloudFront handles query arguments, including the profile paired with each one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudfront_field_level_encryption_config#query_arg_profile_config CloudfrontFieldLevelEncryptionConfig#query_arg_profile_config}
	QueryArgProfileConfig *CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfig `field:"required" json:"queryArgProfileConfig" yaml:"queryArgProfileConfig"`
	// An optional comment describing the field-level encryption configuration.
	//
	// CloudFront stores an omitted comment and an explicitly empty one identically as an empty string, and Read normalizes that empty string back to absent -- so omitting Comment round-trips, while asking for an empty string does not and would report perpetual drift. `minLength: 1` refuses the empty string at template validation; omission is the way to express no comment. The service documents a maximum of 128 characters, and `maxLength` moves that refusal from the API to template validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudfront_field_level_encryption_config#comment CloudfrontFieldLevelEncryptionConfig#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

