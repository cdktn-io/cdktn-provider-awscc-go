// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupBlueGreenDeploymentConfigurationDeploymentReadyOption struct {
	// Information about when to reroute traffic from an original environment to a replacement environment in a blue/green deployment.
	//
	// CONTINUE_DEPLOYMENT: Register new instances with the load balancer immediately after the new application revision is installed on the instances in the replacement environment. STOP_DEPLOYMENT: Do not register new instances with a load balancer unless traffic rerouting is started using ContinueDeployment . If traffic rerouting is not started before the end of the specified wait period, the deployment status is changed to Stopped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_group#action_on_timeout CodedeployDeploymentGroup#action_on_timeout}
	ActionOnTimeout *string `field:"optional" json:"actionOnTimeout" yaml:"actionOnTimeout"`
	// The number of minutes to wait before the status of a blue/green deployment is changed to Stopped if rerouting is not started manually.
	//
	// Applies only to the STOP_DEPLOYMENT option for actionOnTimeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_group#wait_time_in_minutes CodedeployDeploymentGroup#wait_time_in_minutes}
	WaitTimeInMinutes *float64 `field:"optional" json:"waitTimeInMinutes" yaml:"waitTimeInMinutes"`
}

