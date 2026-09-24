// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainergroupdefinition


type GameliftContainerGroupDefinitionSupportContainerDefinitionsPortConfigurationContainerPortRanges struct {
	// A starting value for the range of allowed port numbers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_container_group_definition#from_port GameliftContainerGroupDefinition#from_port}
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Defines the protocol of these ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_container_group_definition#protocol GameliftContainerGroupDefinition#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// An ending value for the range of allowed port numbers.
	//
	// Port numbers are end-inclusive. This value must be equal to or greater than FromPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_container_group_definition#to_port GameliftContainerGroupDefinition#to_port}
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

