// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabricrespondergateway


type RtbfabricResponderGatewayManagedEndpointConfigurationAutoScalingGroupsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_responder_gateway#auto_scaling_group_name_list RtbfabricResponderGateway#auto_scaling_group_name_list}.
	AutoScalingGroupNameList *[]*string `field:"optional" json:"autoScalingGroupNameList" yaml:"autoScalingGroupNameList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_responder_gateway#health_check_config RtbfabricResponderGateway#health_check_config}.
	HealthCheckConfig *RtbfabricResponderGatewayManagedEndpointConfigurationAutoScalingGroupsConfigurationHealthCheckConfig `field:"optional" json:"healthCheckConfig" yaml:"healthCheckConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_responder_gateway#role_arn RtbfabricResponderGateway#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

