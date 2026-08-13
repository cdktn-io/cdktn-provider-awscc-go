// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoListStruct struct {
	// The path used by a load balancer to route production traffic when an Amazon ECS deployment is complete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codedeploy_deployment_group#prod_traffic_route CodedeployDeploymentGroup#prod_traffic_route}
	ProdTrafficRoute *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoListProdTrafficRoute `field:"optional" json:"prodTrafficRoute" yaml:"prodTrafficRoute"`
	// One pair of target groups.
	//
	// One is associated with the original task set. The second is associated with the task set that serves traffic after the deployment is complete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codedeploy_deployment_group#target_groups CodedeployDeploymentGroup#target_groups}
	TargetGroups interface{} `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// An optional path used by a load balancer to route test traffic after an Amazon ECS deployment.
	//
	// Validation can occur while test traffic is served during a deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codedeploy_deployment_group#test_traffic_route CodedeployDeploymentGroup#test_traffic_route}
	TestTrafficRoute *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoListTestTrafficRoute `field:"optional" json:"testTrafficRoute" yaml:"testTrafficRoute"`
}

