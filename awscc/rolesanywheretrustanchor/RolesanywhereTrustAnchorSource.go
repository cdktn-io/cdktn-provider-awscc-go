// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rolesanywheretrustanchor


type RolesanywhereTrustAnchorSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rolesanywhere_trust_anchor#source_data RolesanywhereTrustAnchor#source_data}.
	SourceData *RolesanywhereTrustAnchorSourceSourceData `field:"required" json:"sourceData" yaml:"sourceData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rolesanywhere_trust_anchor#source_type RolesanywhereTrustAnchor#source_type}.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
}

