// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainergroupdefinition


type GameliftContainerGroupDefinitionSupportContainerDefinitionsHealthCheck struct {
	// A string array representing the command that the container runs to determine if it is healthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#command GameliftContainerGroupDefinition#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// How often (in seconds) the health is checked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#interval GameliftContainerGroupDefinition#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// How many times the process manager will retry the command after a timeout.
	//
	// (The first run of the command does not count as a retry.)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#retries GameliftContainerGroupDefinition#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The optional grace period (in seconds) to give a container time to boostrap before teh health check is declared failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#start_period GameliftContainerGroupDefinition#start_period}
	StartPeriod *float64 `field:"optional" json:"startPeriod" yaml:"startPeriod"`
	// How many seconds the process manager allows the command to run before canceling it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#timeout GameliftContainerGroupDefinition#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

