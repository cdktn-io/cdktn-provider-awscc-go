// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetRowLevelPermissionDataSet struct {
	// <p>The Amazon Resource Name (ARN) of the permission dataset.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_set#arn QuicksightDataSet#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_set#format_version QuicksightDataSet#format_version}.
	FormatVersion *string `field:"optional" json:"formatVersion" yaml:"formatVersion"`
	// <p>The namespace associated with the row-level permissions dataset.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_set#namespace QuicksightDataSet#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_set#permission_policy QuicksightDataSet#permission_policy}.
	PermissionPolicy *string `field:"optional" json:"permissionPolicy" yaml:"permissionPolicy"`
}

