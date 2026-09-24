// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package efsaccesspoint


type EfsAccessPointAccessPointTags struct {
	// The tag key (String). The key can't start with ``aws:``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/efs_access_point#key EfsAccessPoint#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/efs_access_point#value EfsAccessPoint#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

