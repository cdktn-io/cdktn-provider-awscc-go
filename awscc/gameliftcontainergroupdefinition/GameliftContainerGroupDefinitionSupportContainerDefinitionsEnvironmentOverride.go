// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainergroupdefinition


type GameliftContainerGroupDefinitionSupportContainerDefinitionsEnvironmentOverride struct {
	// The environment variable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/gamelift_container_group_definition#name GameliftContainerGroupDefinition#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The environment variable value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/gamelift_container_group_definition#value GameliftContainerGroupDefinition#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

