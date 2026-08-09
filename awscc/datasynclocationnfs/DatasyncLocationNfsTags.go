// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationnfs


type DatasyncLocationNfsTags struct {
	// The key for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_nfs#key DatasyncLocationNfs#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_nfs#value DatasyncLocationNfs#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

