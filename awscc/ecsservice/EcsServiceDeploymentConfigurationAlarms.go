// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfigurationAlarms struct {
	// One or more CloudWatch alarm names. Use a "," to separate the alarms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#alarm_names EcsService#alarm_names}
	AlarmNames *[]*string `field:"optional" json:"alarmNames" yaml:"alarmNames"`
	// Determines whether to use the CloudWatch alarm option in the service deployment process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#enable EcsService#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Determines whether to configure Amazon ECS to roll back the service if a service deployment fails.
	//
	// If rollback is used, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#rollback EcsService#rollback}
	Rollback interface{} `field:"optional" json:"rollback" yaml:"rollback"`
}

