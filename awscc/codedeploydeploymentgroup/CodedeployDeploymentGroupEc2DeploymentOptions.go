// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupEc2DeploymentOptions struct {
	// The Amazon EC2 Auto Scaling processes to suspend during an in-place deployment.
	//
	// CodeDeploy suspends these processes before it deploys to instances in an Auto Scaling group, and resumes them when the deployment succeeds, fails, or is stopped. Valid values are Terminate, AZRebalance, InstanceRefresh, and ScheduledActions. Process names are case sensitive. To stop suspending processes, remove the property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/codedeploy_deployment_group#auto_scaling_group_processes_to_suspend CodedeployDeploymentGroup#auto_scaling_group_processes_to_suspend}
	AutoScalingGroupProcessesToSuspend *[]*string `field:"optional" json:"autoScalingGroupProcessesToSuspend" yaml:"autoScalingGroupProcessesToSuspend"`
}

