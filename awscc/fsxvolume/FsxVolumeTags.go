// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeTags struct {
	// A value that specifies the TagKey, the name of the tag.
	//
	// Tag keys must be unique for the resource to which they are attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#key FsxVolume#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A value that specifies the TagValue, the value assigned to the corresponding tag key.
	//
	// Tag values can be null and don't have to be unique in a tag set. For example, you can have a key-value pair in a tag set of finances : April and also of payroll : April.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#value FsxVolume#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

