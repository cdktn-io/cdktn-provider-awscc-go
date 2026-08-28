// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instanceconnectendpoint


type Ec2InstanceConnectEndpointTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance_connect_endpoint#key Ec2InstanceConnectEndpoint#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance_connect_endpoint#value Ec2InstanceConnectEndpoint#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

