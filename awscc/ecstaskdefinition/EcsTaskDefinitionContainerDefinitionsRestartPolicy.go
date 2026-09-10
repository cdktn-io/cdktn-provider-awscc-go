// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsRestartPolicy struct {
	// Specifies whether a restart policy is enabled for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_definition#enabled EcsTaskDefinition#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of exit codes that Amazon ECS will ignore and not attempt a restart on.
	//
	// You can specify a maximum of 50 container exit codes. By default, Amazon ECS does not ignore any exit codes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_definition#ignored_exit_codes EcsTaskDefinition#ignored_exit_codes}
	IgnoredExitCodes *[]*float64 `field:"optional" json:"ignoredExitCodes" yaml:"ignoredExitCodes"`
	// A period of time (in seconds) that the container must run for before a restart can be attempted.
	//
	// A container can be restarted only once every ``restartAttemptPeriod`` seconds. If a container isn't able to run for this time period and exits early, it will not be restarted. You can set a minimum ``restartAttemptPeriod`` of 60 seconds and a maximum ``restartAttemptPeriod`` of 1800 seconds. By default, a container must run for 300 seconds before it can be restarted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_definition#restart_attempt_period EcsTaskDefinition#restart_attempt_period}
	RestartAttemptPeriod *float64 `field:"optional" json:"restartAttemptPeriod" yaml:"restartAttemptPeriod"`
}

