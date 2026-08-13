// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcloudfrontoriginaccessidentity


type CloudfrontCloudfrontOriginAccessIdentityCloudfrontOriginAccessIdentityConfig struct {
	// A comment to describe the origin access identity. The comment cannot be longer than 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudfront_cloudfront_origin_access_identity#comment CloudfrontCloudfrontOriginAccessIdentity#comment}
	Comment *string `field:"required" json:"comment" yaml:"comment"`
}

