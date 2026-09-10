// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasyncagent

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatasyncAgentConfig struct {
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
	// Activation key of the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datasync_agent#activation_key DatasyncAgent#activation_key}
	ActivationKey *string `field:"optional" json:"activationKey" yaml:"activationKey"`
	// The name configured for the agent. Text reference used to identify the agent in the console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datasync_agent#agent_name DatasyncAgent#agent_name}
	AgentName *string `field:"optional" json:"agentName" yaml:"agentName"`
	// The ARNs of the security group used to protect your data transfer task subnets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datasync_agent#security_group_arns DatasyncAgent#security_group_arns}
	SecurityGroupArns *[]*string `field:"optional" json:"securityGroupArns" yaml:"securityGroupArns"`
	// The ARNs of the subnets in which DataSync will create elastic network interfaces for each data transfer task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datasync_agent#subnet_arns DatasyncAgent#subnet_arns}
	SubnetArns *[]*string `field:"optional" json:"subnetArns" yaml:"subnetArns"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datasync_agent#tags DatasyncAgent#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC endpoint that the agent has access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datasync_agent#vpc_endpoint_id DatasyncAgent#vpc_endpoint_id}
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

