// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightspace


type QuicksightSpacePermissions struct {
	// The list of actions granted to the principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_space#actions QuicksightSpace#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// The ARN of the principal (user or group) receiving the permission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_space#principal QuicksightSpace#principal}
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
}

