// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupTriggerConfigurations struct {
	// The event type or types that trigger notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codedeploy_deployment_group#trigger_events CodedeployDeploymentGroup#trigger_events}
	TriggerEvents *[]*string `field:"optional" json:"triggerEvents" yaml:"triggerEvents"`
	// The name of the notification trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codedeploy_deployment_group#trigger_name CodedeployDeploymentGroup#trigger_name}
	TriggerName *string `field:"optional" json:"triggerName" yaml:"triggerName"`
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic through which notifications about deployment or instance events are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codedeploy_deployment_group#trigger_target_arn CodedeployDeploymentGroup#trigger_target_arn}
	TriggerTargetArn *string `field:"optional" json:"triggerTargetArn" yaml:"triggerTargetArn"`
}

