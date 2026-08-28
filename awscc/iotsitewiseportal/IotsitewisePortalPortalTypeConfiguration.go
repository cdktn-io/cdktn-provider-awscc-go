// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseportal


type IotsitewisePortalPortalTypeConfiguration struct {
	// List of enabled Tools for a certain portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotsitewise_portal#portal_tools IotsitewisePortal#portal_tools}
	PortalTools *[]*string `field:"optional" json:"portalTools" yaml:"portalTools"`
}

