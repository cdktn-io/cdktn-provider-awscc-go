// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mpaidentitysource


type MpaIdentitySourceIdentitySourceParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mpa_identity_source#iam_identity_center MpaIdentitySource#iam_identity_center}.
	IamIdentityCenter *MpaIdentitySourceIdentitySourceParametersIamIdentityCenter `field:"required" json:"iamIdentityCenter" yaml:"iamIdentityCenter"`
}

