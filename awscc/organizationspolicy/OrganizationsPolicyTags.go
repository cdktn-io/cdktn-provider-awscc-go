// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organizationspolicy


type OrganizationsPolicyTags struct {
	// The key identifier, or name, of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/organizations_policy#key OrganizationsPolicy#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The string value that's associated with the key of the tag.
	//
	// You can set the value of a tag to an empty string, but you can't set the value of a tag to null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/organizations_policy#value OrganizationsPolicy#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

