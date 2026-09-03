// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncapi


type AppsyncApiTags struct {
	// A string used to identify this tag. You can specify a maximum of 128 characters for a tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appsync_api#key AppsyncApi#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A string containing the value for this tag.
	//
	// You can specify a maximum of 256 characters for a tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appsync_api#value AppsyncApi#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

