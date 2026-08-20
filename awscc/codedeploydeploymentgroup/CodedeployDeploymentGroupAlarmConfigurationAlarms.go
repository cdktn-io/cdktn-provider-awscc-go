// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupAlarmConfigurationAlarms struct {
	// The name of the alarm.
	//
	// Maximum length is 255 characters. Each alarm name can be used only once in a list of alarms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_group#name CodedeployDeploymentGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

