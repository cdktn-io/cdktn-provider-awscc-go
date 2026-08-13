// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainergroupdefinition


type GameliftContainerGroupDefinitionSupportContainerDefinitionsLinuxCapabilities struct {
	// The list of Linux capabilities to add to the container's default configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_container_group_definition#include GameliftContainerGroupDefinition#include}
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

