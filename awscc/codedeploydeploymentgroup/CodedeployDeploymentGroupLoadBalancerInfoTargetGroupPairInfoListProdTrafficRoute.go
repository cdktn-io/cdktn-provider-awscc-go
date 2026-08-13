// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoListProdTrafficRoute struct {
	// The Amazon Resource Name (ARN) of one listener.
	//
	// The listener identifies the route between a target group and a load balancer. This is an array of strings with a maximum size of one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codedeploy_deployment_group#listener_arns CodedeployDeploymentGroup#listener_arns}
	ListenerArns *[]*string `field:"optional" json:"listenerArns" yaml:"listenerArns"`
}

