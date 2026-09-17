// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instanceconnectendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2InstanceConnectEndpointConfig struct {
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
	// The ID of the subnet in which the EC2 Instance Connect Endpoint was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_instance_connect_endpoint#subnet_id Ec2InstanceConnectEndpoint#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The client token of the instance connect endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_instance_connect_endpoint#client_token Ec2InstanceConnectEndpoint#client_token}
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// Indicates whether your client's IP address is preserved as the source when you connect to a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_instance_connect_endpoint#preserve_client_ip Ec2InstanceConnectEndpoint#preserve_client_ip}
	PreserveClientIp interface{} `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// The security groups associated with the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_instance_connect_endpoint#security_group_ids Ec2InstanceConnectEndpoint#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The tags assigned to the EC2 Instance Connect Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_instance_connect_endpoint#tags Ec2InstanceConnectEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

