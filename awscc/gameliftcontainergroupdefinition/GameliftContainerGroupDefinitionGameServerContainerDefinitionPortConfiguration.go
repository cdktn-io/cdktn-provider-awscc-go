// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainergroupdefinition


type GameliftContainerGroupDefinitionGameServerContainerDefinitionPortConfiguration struct {
	// Specifies one or more ranges of ports on a container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#container_port_ranges GameliftContainerGroupDefinition#container_port_ranges}
	ContainerPortRanges interface{} `field:"optional" json:"containerPortRanges" yaml:"containerPortRanges"`
}

