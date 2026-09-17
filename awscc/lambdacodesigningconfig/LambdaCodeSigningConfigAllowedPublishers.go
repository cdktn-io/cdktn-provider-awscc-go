// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdacodesigningconfig


type LambdaCodeSigningConfigAllowedPublishers struct {
	// List of Signing profile version Arns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_code_signing_config#signing_profile_version_arns LambdaCodeSigningConfig#signing_profile_version_arns}
	SigningProfileVersionArns *[]*string `field:"required" json:"signingProfileVersionArns" yaml:"signingProfileVersionArns"`
}

