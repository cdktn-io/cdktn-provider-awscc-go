// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2trafficmirrorfilterrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2TrafficMirrorFilterRuleConfig struct {
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
	// The destination CIDR block to assign to the Traffic Mirror rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_traffic_mirror_filter_rule#destination_cidr_block Ec2TrafficMirrorFilterRule#destination_cidr_block}
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The action to take on the filtered traffic (accept/reject).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_traffic_mirror_filter_rule#rule_action Ec2TrafficMirrorFilterRule#rule_action}
	RuleAction *string `field:"required" json:"ruleAction" yaml:"ruleAction"`
	// The number of the Traffic Mirror rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_traffic_mirror_filter_rule#rule_number Ec2TrafficMirrorFilterRule#rule_number}
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// The source CIDR block to assign to the Traffic Mirror Filter rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_traffic_mirror_filter_rule#source_cidr_block Ec2TrafficMirrorFilterRule#source_cidr_block}
	SourceCidrBlock *string `field:"required" json:"sourceCidrBlock" yaml:"sourceCidrBlock"`
	// The direction of traffic (ingress/egress).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_traffic_mirror_filter_rule#traffic_direction Ec2TrafficMirrorFilterRule#traffic_direction}
	TrafficDirection *string `field:"required" json:"trafficDirection" yaml:"trafficDirection"`
	// The ID of the filter that this rule is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_traffic_mirror_filter_rule#traffic_mirror_filter_id Ec2TrafficMirrorFilterRule#traffic_mirror_filter_id}
	TrafficMirrorFilterId *string `field:"required" json:"trafficMirrorFilterId" yaml:"trafficMirrorFilterId"`
	// The description of the Traffic Mirror Filter rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_traffic_mirror_filter_rule#description Ec2TrafficMirrorFilterRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The destination port range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_traffic_mirror_filter_rule#destination_port_range Ec2TrafficMirrorFilterRule#destination_port_range}
	DestinationPortRange *Ec2TrafficMirrorFilterRuleDestinationPortRange `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// The number of protocol, for example 17 (UDP), to assign to the Traffic Mirror rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_traffic_mirror_filter_rule#protocol Ec2TrafficMirrorFilterRule#protocol}
	Protocol *float64 `field:"optional" json:"protocol" yaml:"protocol"`
	// The source port range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_traffic_mirror_filter_rule#source_port_range Ec2TrafficMirrorFilterRule#source_port_range}
	SourcePortRange *Ec2TrafficMirrorFilterRuleSourcePortRange `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
	// Any tags assigned to the Traffic Mirror Filter rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_traffic_mirror_filter_rule#tags Ec2TrafficMirrorFilterRule#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

