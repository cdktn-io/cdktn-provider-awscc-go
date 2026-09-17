// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2trafficmirrortarget

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2TrafficMirrorTargetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The description of the Traffic Mirror target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_traffic_mirror_target#description Ec2TrafficMirrorTarget#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the Gateway Load Balancer endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_traffic_mirror_target#gateway_load_balancer_endpoint_id Ec2TrafficMirrorTarget#gateway_load_balancer_endpoint_id}
	GatewayLoadBalancerEndpointId *string `field:"optional" json:"gatewayLoadBalancerEndpointId" yaml:"gatewayLoadBalancerEndpointId"`
	// The network interface ID that is associated with the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_traffic_mirror_target#network_interface_id Ec2TrafficMirrorTarget#network_interface_id}
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_traffic_mirror_target#network_load_balancer_arn Ec2TrafficMirrorTarget#network_load_balancer_arn}
	NetworkLoadBalancerArn *string `field:"optional" json:"networkLoadBalancerArn" yaml:"networkLoadBalancerArn"`
	// The tags to assign to the Traffic Mirror target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_traffic_mirror_target#tags Ec2TrafficMirrorTarget#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

