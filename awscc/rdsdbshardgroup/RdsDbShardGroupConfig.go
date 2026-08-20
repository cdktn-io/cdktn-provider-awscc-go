// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbshardgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsDbShardGroupConfig struct {
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
	// The name of the primary DB cluster for the DB shard group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_shard_group#db_cluster_identifier RdsDbShardGroup#db_cluster_identifier}
	DbClusterIdentifier *string `field:"required" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The maximum capacity of the DB shard group in Aurora capacity units (ACUs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_shard_group#max_acu RdsDbShardGroup#max_acu}
	MaxAcu *float64 `field:"required" json:"maxAcu" yaml:"maxAcu"`
	// Specifies whether to create standby standby DB data access shard for the DB shard group.
	//
	// Valid values are the following:
	//   +  0 - Creates a DB shard group without a standby DB data access shard. This is the default value.
	//   +  1 - Creates a DB shard group with a standby DB data access shard in a different Availability Zone (AZ).
	//   +  2 - Creates a DB shard group with two standby DB data access shard in two different AZs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_shard_group#compute_redundancy RdsDbShardGroup#compute_redundancy}
	ComputeRedundancy *float64 `field:"optional" json:"computeRedundancy" yaml:"computeRedundancy"`
	// The name of the DB shard group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_shard_group#db_shard_group_identifier RdsDbShardGroup#db_shard_group_identifier}
	DbShardGroupIdentifier *string `field:"optional" json:"dbShardGroupIdentifier" yaml:"dbShardGroupIdentifier"`
	// The minimum capacity of the DB shard group in Aurora capacity units (ACUs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_shard_group#min_acu RdsDbShardGroup#min_acu}
	MinAcu *float64 `field:"optional" json:"minAcu" yaml:"minAcu"`
	// Specifies whether the DB shard group is publicly accessible.
	//
	// When the DB shard group is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB shard group's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB shard group's VPC. Access to the DB shard group is ultimately controlled by the security group it uses. That public access is not permitted if the security group assigned to the DB shard group doesn't permit it.
	//  When the DB shard group isn't publicly accessible, it is an internal DB shard group with a DNS name that resolves to a private IP address.
	//  Default: The default behavior varies depending on whether ``DBSubnetGroupName`` is specified.
	//  If ``DBSubnetGroupName`` isn't specified, and ``PubliclyAccessible`` isn't specified, the following applies:
	//   +  If the default VPC in the target Region doesn?t have an internet gateway attached to it, the DB shard group is private.
	//   +  If the default VPC in the target Region has an internet gateway attached to it, the DB shard group is public.
	//
	//  If ``DBSubnetGroupName`` is specified, and ``PubliclyAccessible`` isn't specified, the following applies:
	//   +  If the subnets are part of a VPC that doesn?t have an internet gateway attached to it, the DB shard group is private.
	//   +  If the subnets are part of a VPC that has an internet gateway attached to it, the DB shard group is public.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_shard_group#publicly_accessible RdsDbShardGroup#publicly_accessible}
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the DB shard group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_shard_group#tags RdsDbShardGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

