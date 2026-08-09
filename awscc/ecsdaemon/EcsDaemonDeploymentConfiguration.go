// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemon


type EcsDaemonDeploymentConfiguration struct {
	// The CloudWatch alarm configuration for the daemon deployment.
	//
	// When alarms are triggered during a deployment, the deployment can be automatically rolled back.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_daemon#alarms EcsDaemon#alarms}
	Alarms *EcsDaemonDeploymentConfigurationAlarms `field:"optional" json:"alarms" yaml:"alarms"`
	// The amount of time (in minutes) to wait after a successful deployment step before proceeding.
	//
	// This allows time to monitor for issues before continuing. The default value is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_daemon#bake_time_in_minutes EcsDaemon#bake_time_in_minutes}
	BakeTimeInMinutes *float64 `field:"optional" json:"bakeTimeInMinutes" yaml:"bakeTimeInMinutes"`
	// The percentage of container instances to drain simultaneously during a daemon deployment. Valid values are between 0.0 and 100.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_daemon#drain_percent EcsDaemon#drain_percent}
	DrainPercent *float64 `field:"optional" json:"drainPercent" yaml:"drainPercent"`
}

