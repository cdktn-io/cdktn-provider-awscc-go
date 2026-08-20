// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontpublickey


type CloudfrontPublicKeyPublicKeyConfig struct {
	// A string included in the request to help make sure that the request can't be replayed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_public_key#caller_reference CloudfrontPublicKey#caller_reference}
	CallerReference *string `field:"required" json:"callerReference" yaml:"callerReference"`
	// The public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html), or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_public_key#encoded_key CloudfrontPublicKey#encoded_key}
	EncodedKey *string `field:"required" json:"encodedKey" yaml:"encodedKey"`
	// A name to help identify the public key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_public_key#name CloudfrontPublicKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A comment to describe the public key. The comment cannot be longer than 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_public_key#comment CloudfrontPublicKey#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

