// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftlocation


type GameliftLocationTags struct {
	// The key name of the tag.
	//
	// You can specify a value that is 1 to 128 Unicode characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_location#key GameliftLocation#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_location#value GameliftLocation#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

