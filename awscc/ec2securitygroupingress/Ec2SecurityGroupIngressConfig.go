// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2securitygroupingress

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2SecurityGroupIngressConfig struct {
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
	// The IP protocol name (tcp, udp, icmp, icmpv6) or number (see Protocol Numbers).
	//
	// [VPC only] Use -1 to specify all protocols. When authorizing security group rules, specifying -1 or a protocol number other than tcp, udp, icmp, or icmpv6 allows traffic on all ports, regardless of any port range you specify. For tcp, udp, and icmp, you must specify a port range. For icmpv6, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_security_group_ingress#ip_protocol Ec2SecurityGroupIngress#ip_protocol}
	IpProtocol *string `field:"required" json:"ipProtocol" yaml:"ipProtocol"`
	// The IPv4 ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_security_group_ingress#cidr_ip Ec2SecurityGroupIngress#cidr_ip}
	CidrIp *string `field:"optional" json:"cidrIp" yaml:"cidrIp"`
	// [VPC only] The IPv6 ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_security_group_ingress#cidr_ipv_6 Ec2SecurityGroupIngress#cidr_ipv_6}
	CidrIpv6 *string `field:"optional" json:"cidrIpv6" yaml:"cidrIpv6"`
	// Updates the description of an ingress (inbound) security group rule.
	//
	// You can replace an existing description, or add a description to a rule that did not have one previously
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_security_group_ingress#description Ec2SecurityGroupIngress#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number.
	//
	// A value of -1 indicates all ICMP/ICMPv6 types. If you specify all ICMP/ICMPv6 types, you must specify all codes.
	//
	// Use this for ICMP and any protocol that uses ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_security_group_ingress#from_port Ec2SecurityGroupIngress#from_port}
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// The ID of the security group.
	//
	// You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.
	//
	// You must specify the GroupName property or the GroupId property. For security groups that are in a VPC, you must use the GroupId property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_security_group_ingress#group_id Ec2SecurityGroupIngress#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The name of the security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_security_group_ingress#group_name Ec2SecurityGroupIngress#group_name}
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// [EC2-VPC only] The ID of a prefix list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_security_group_ingress#source_prefix_list_id Ec2SecurityGroupIngress#source_prefix_list_id}
	SourcePrefixListId *string `field:"optional" json:"sourcePrefixListId" yaml:"sourcePrefixListId"`
	// The ID of the security group.
	//
	// You must specify either the security group ID or the security group name. For security groups in a nondefault VPC, you must specify the security group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_security_group_ingress#source_security_group_id Ec2SecurityGroupIngress#source_security_group_id}
	SourceSecurityGroupId *string `field:"optional" json:"sourceSecurityGroupId" yaml:"sourceSecurityGroupId"`
	// [EC2-Classic, default VPC] The name of the source security group.
	//
	// You must specify the GroupName property or the GroupId property. For security groups that are in a VPC, you must use the GroupId property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_security_group_ingress#source_security_group_name Ec2SecurityGroupIngress#source_security_group_name}
	SourceSecurityGroupName *string `field:"optional" json:"sourceSecurityGroupName" yaml:"sourceSecurityGroupName"`
	// [nondefault VPC] The AWS account ID that owns the source security group.
	//
	// You can't specify this property with an IP address range.
	//
	// If you specify SourceSecurityGroupName or SourceSecurityGroupId and that security group is owned by a different account than the account creating the stack, you must specify the SourceSecurityGroupOwnerId; otherwise, this property is optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_security_group_ingress#source_security_group_owner_id Ec2SecurityGroupIngress#source_security_group_owner_id}
	SourceSecurityGroupOwnerId *string `field:"optional" json:"sourceSecurityGroupOwnerId" yaml:"sourceSecurityGroupOwnerId"`
	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	//
	// A value of -1 indicates all ICMP/ICMPv6 codes for the specified ICMP type. If you specify all ICMP/ICMPv6 types, you must specify all codes.
	//
	// Use this for ICMP and any protocol that uses ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_security_group_ingress#to_port Ec2SecurityGroupIngress#to_port}
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

