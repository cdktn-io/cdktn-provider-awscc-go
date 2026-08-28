// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewaypolicytableentry

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2TransitGatewayPolicyTableEntryConfig struct {
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
	// The policy rule associated with the entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_policy_table_entry#policy_rule Ec2TransitGatewayPolicyTableEntry#policy_rule}
	PolicyRule *Ec2TransitGatewayPolicyTableEntryPolicyRule `field:"required" json:"policyRule" yaml:"policyRule"`
	// The rule number for the policy table entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_policy_table_entry#policy_rule_number Ec2TransitGatewayPolicyTableEntry#policy_rule_number}
	PolicyRuleNumber *string `field:"required" json:"policyRuleNumber" yaml:"policyRuleNumber"`
	// The ID of the target route table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_policy_table_entry#target_route_table_id Ec2TransitGatewayPolicyTableEntry#target_route_table_id}
	TargetRouteTableId *string `field:"required" json:"targetRouteTableId" yaml:"targetRouteTableId"`
	// The ID of the transit gateway policy table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_policy_table_entry#transit_gateway_policy_table_id Ec2TransitGatewayPolicyTableEntry#transit_gateway_policy_table_id}
	TransitGatewayPolicyTableId *string `field:"required" json:"transitGatewayPolicyTableId" yaml:"transitGatewayPolicyTableId"`
}

