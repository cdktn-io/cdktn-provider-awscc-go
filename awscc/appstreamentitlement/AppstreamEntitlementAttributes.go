// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamentitlement


type AppstreamEntitlementAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appstream_entitlement#name AppstreamEntitlement#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appstream_entitlement#value AppstreamEntitlement#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

