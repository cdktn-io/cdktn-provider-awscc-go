// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamuser


type IamUserLoginProfile struct {
	// The user's password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iam_user#password IamUser#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Specifies whether the user is required to set a new password on next sign-in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iam_user#password_reset_required IamUser#password_reset_required}
	PasswordResetRequired interface{} `field:"optional" json:"passwordResetRequired" yaml:"passwordResetRequired"`
}

