// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsreplicationsubnetgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsReplicationSubnetGroupConfig struct {
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
	// The description for the subnet group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_subnet_group#replication_subnet_group_description DmsReplicationSubnetGroup#replication_subnet_group_description}
	ReplicationSubnetGroupDescription *string `field:"required" json:"replicationSubnetGroupDescription" yaml:"replicationSubnetGroupDescription"`
	// One or more subnet IDs to be assigned to the replication subnet group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_subnet_group#subnet_ids DmsReplicationSubnetGroup#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The name for the replication subnet group. This value is stored as a lowercase string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_subnet_group#replication_subnet_group_identifier DmsReplicationSubnetGroup#replication_subnet_group_identifier}
	ReplicationSubnetGroupIdentifier *string `field:"optional" json:"replicationSubnetGroupIdentifier" yaml:"replicationSubnetGroupIdentifier"`
	// One or more tags to be assigned to the replication subnet group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_subnet_group#tags DmsReplicationSubnetGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

