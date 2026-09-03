// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailcontainer


type LightsailContainerPrivateRegistryAccessEcrImagePullerRole struct {
	// A Boolean value that indicates whether to activate the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_container#is_active LightsailContainer#is_active}
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
}

