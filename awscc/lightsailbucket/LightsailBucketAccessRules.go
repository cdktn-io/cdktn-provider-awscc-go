// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailbucket


type LightsailBucketAccessRules struct {
	// A Boolean value that indicates whether the access control list (ACL) permissions that are applied to individual objects override the getObject option that is currently specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_bucket#allow_public_overrides LightsailBucket#allow_public_overrides}
	AllowPublicOverrides interface{} `field:"optional" json:"allowPublicOverrides" yaml:"allowPublicOverrides"`
	// Specifies the anonymous access to all objects in a bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_bucket#get_object LightsailBucket#get_object}
	FetchObject *string `field:"optional" json:"fetchObject" yaml:"fetchObject"`
}

