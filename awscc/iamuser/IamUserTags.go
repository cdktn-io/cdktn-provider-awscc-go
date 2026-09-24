// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamuser


type IamUserTags struct {
	// The key name that can be used to look up or retrieve the associated value.
	//
	// For example, ``Department`` or ``Cost Center`` are common choices.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iam_user#key IamUser#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value associated with this tag.
	//
	// For example, tags with a key name of ``Department`` could have values such as ``Human Resources``, ``Accounting``, and ``Support``. Tags with a key name of ``Cost Center`` might have values that consist of the number associated with the different cost centers in your company. Typically, many resources have tags with the same key name but with different values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iam_user#value IamUser#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

