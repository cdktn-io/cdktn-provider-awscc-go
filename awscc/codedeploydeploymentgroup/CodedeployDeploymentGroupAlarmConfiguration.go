// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupAlarmConfiguration struct {
	// A list of alarms configured for the deployment or deployment group. A maximum of 10 alarms can be added.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codedeploy_deployment_group#alarms CodedeployDeploymentGroup#alarms}
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// Indicates whether the alarm configuration is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codedeploy_deployment_group#enabled CodedeployDeploymentGroup#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicates whether a deployment should continue if information about the current state of alarms cannot be retrieved from Amazon CloudWatch.
	//
	// The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codedeploy_deployment_group#ignore_poll_alarm_failure CodedeployDeploymentGroup#ignore_poll_alarm_failure}
	IgnorePollAlarmFailure interface{} `field:"optional" json:"ignorePollAlarmFailure" yaml:"ignorePollAlarmFailure"`
}

