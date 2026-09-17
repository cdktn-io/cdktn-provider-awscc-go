// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClustersVpcConfig struct {
	// The AWS security groups to associate with the elastic network interfaces in order to specify what the replicator has access to.
	//
	// If a security group is not specified, the default security group associated with the VPC is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_replicator#security_group_ids MskReplicator#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The list of subnets to connect to in the virtual private cloud (VPC).
	//
	// AWS creates elastic network interfaces inside these subnets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_replicator#subnet_ids MskReplicator#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

