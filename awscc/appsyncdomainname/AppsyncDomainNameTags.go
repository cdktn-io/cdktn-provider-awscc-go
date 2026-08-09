// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncdomainname


type AppsyncDomainNameTags struct {
	// A string used to identify this tag.
	//
	// You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_domain_name#key AppsyncDomainName#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A string containing the value for this tag.
	//
	// You can specify a maximum of 256 characters for a tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_domain_name#value AppsyncDomainName#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

