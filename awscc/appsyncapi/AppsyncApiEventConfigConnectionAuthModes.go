// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncapi


type AppsyncApiEventConfigConnectionAuthModes struct {
	// Security configuration for your AppSync API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appsync_api#auth_type AppsyncApi#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
}

