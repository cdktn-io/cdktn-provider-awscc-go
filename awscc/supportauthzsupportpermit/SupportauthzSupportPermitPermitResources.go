// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package supportauthzsupportpermit


type SupportauthzSupportPermitPermitResources struct {
	// Applies to all resources in the region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/supportauthz_support_permit#all_resources_in_region SupportauthzSupportPermit#all_resources_in_region}
	AllResourcesInRegion *string `field:"optional" json:"allResourcesInRegion" yaml:"allResourcesInRegion"`
	// An explicit list of resource ARNs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/supportauthz_support_permit#resources SupportauthzSupportPermit#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

