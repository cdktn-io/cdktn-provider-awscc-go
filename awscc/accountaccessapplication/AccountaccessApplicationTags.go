// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountaccessapplication


type AccountaccessApplicationTags struct {
	// The key name of the tag.
	//
	// You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/accountaccess_application#key AccountaccessApplication#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/accountaccess_application#value AccountaccessApplication#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

