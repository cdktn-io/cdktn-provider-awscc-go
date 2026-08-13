// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2securitygroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2SecurityGroupConfig struct {
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
	// A description for the security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_security_group#group_description Ec2SecurityGroup#group_description}
	GroupDescription *string `field:"required" json:"groupDescription" yaml:"groupDescription"`
	// The name of the security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_security_group#group_name Ec2SecurityGroup#group_name}
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// [VPC only] The outbound rules associated with the security group.
	//
	// There is a short interruption during which you cannot connect to the security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_security_group#security_group_egress Ec2SecurityGroup#security_group_egress}
	SecurityGroupEgress interface{} `field:"optional" json:"securityGroupEgress" yaml:"securityGroupEgress"`
	// The inbound rules associated with the security group.
	//
	// There is a short interruption during which you cannot connect to the security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_security_group#security_group_ingress Ec2SecurityGroup#security_group_ingress}
	SecurityGroupIngress interface{} `field:"optional" json:"securityGroupIngress" yaml:"securityGroupIngress"`
	// Any tags assigned to the security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_security_group#tags Ec2SecurityGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC for the security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_security_group#vpc_id Ec2SecurityGroup#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

