// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemon


type EcsDaemonDeploymentConfigurationAlarms struct {
	// The CloudWatch alarm names to monitor during a daemon deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_daemon#alarm_names EcsDaemon#alarm_names}
	AlarmNames *[]*string `field:"optional" json:"alarmNames" yaml:"alarmNames"`
	// Determines whether to use the CloudWatch alarm option in the daemon deployment process. The default value is ``false``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_daemon#enable EcsDaemon#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

