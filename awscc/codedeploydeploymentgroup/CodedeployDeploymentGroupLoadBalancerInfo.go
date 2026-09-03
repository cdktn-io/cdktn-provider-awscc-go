// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupLoadBalancerInfo struct {
	// An array that contains information about the load balancers to use for load balancing in a deployment.
	//
	// If you're using Classic Load Balancers, specify those load balancers in this array.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codedeploy_deployment_group#elb_info_list CodedeployDeploymentGroup#elb_info_list}
	ElbInfoList interface{} `field:"optional" json:"elbInfoList" yaml:"elbInfoList"`
	// An array that contains information about the target groups to use for load balancing in a deployment.
	//
	// If you're using Application Load Balancers and Network Load Balancers, specify their associated target groups in this array.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codedeploy_deployment_group#target_group_info_list CodedeployDeploymentGroup#target_group_info_list}
	TargetGroupInfoList interface{} `field:"optional" json:"targetGroupInfoList" yaml:"targetGroupInfoList"`
	// The target group pair information. This is an array of TargeGroupPairInfo objects with a maximum size of one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codedeploy_deployment_group#target_group_pair_info_list CodedeployDeploymentGroup#target_group_pair_info_list}
	TargetGroupPairInfoList interface{} `field:"optional" json:"targetGroupPairInfoList" yaml:"targetGroupPairInfoList"`
}

