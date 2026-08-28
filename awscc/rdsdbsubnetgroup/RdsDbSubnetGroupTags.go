// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbsubnetgroup


type RdsDbSubnetGroupTags struct {
	// A key is the required name of the tag.
	//
	// The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_subnet_group#key RdsDbSubnetGroup#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A value is the optional value of the tag.
	//
	// The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_subnet_group#value RdsDbSubnetGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

