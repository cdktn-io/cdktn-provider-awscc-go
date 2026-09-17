// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2egressonlyinternetgateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2EgressOnlyInternetGatewayConfig struct {
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
	// The ID of the VPC for which to create the egress-only internet gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_egress_only_internet_gateway#vpc_id Ec2EgressOnlyInternetGateway#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Any tags assigned to the egress only internet gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_egress_only_internet_gateway#tags Ec2EgressOnlyInternetGateway#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

