// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountaccessapplication


type AccountaccessApplicationIdentitySource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/accountaccess_application#identity_center AccountaccessApplication#identity_center}.
	IdentityCenter *AccountaccessApplicationIdentitySourceIdentityCenter `field:"required" json:"identityCenter" yaml:"identityCenter"`
}

