// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionspolicy


type VerifiedpermissionsPolicyDefinitionStatic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/verifiedpermissions_policy#description VerifiedpermissionsPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/verifiedpermissions_policy#statement VerifiedpermissionsPolicy#statement}.
	Statement *string `field:"optional" json:"statement" yaml:"statement"`
}

