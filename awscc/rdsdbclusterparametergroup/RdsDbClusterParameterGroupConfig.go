// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbclusterparametergroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsDbClusterParameterGroupConfig struct {
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
	// The description for the DB cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_cluster_parameter_group#description RdsDbClusterParameterGroup#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The DB cluster parameter group family name.
	//
	// A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a database engine and engine version compatible with that DB cluster parameter group family.
	//   *Aurora MySQL*
	//  Example: ``aurora-mysql5.7``, ``aurora-mysql8.0``
	//   *Aurora PostgreSQL*
	//  Example: ``aurora-postgresql14``
	//   *RDS for MySQL*
	//  Example: ``mysql8.0``
	//   *RDS for PostgreSQL*
	//  Example: ``postgres13``
	//  To list all of the available parameter group families for a DB engine, use the following command:
	//   ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine <engine>``
	//  For example, to list all of the available parameter group families for the Aurora PostgreSQL DB engine, use the following command:
	//   ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine aurora-postgresql``
	//   The output contains duplicates.
	//   The following are the valid DB engine values:
	//   +   ``aurora-mysql``
	//   +   ``aurora-postgresql``
	//   +   ``mysql``
	//   +   ``postgres``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_cluster_parameter_group#family RdsDbClusterParameterGroup#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// Provides a list of parameters for the DB cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_cluster_parameter_group#parameters RdsDbClusterParameterGroup#parameters}
	Parameters *string `field:"required" json:"parameters" yaml:"parameters"`
	// The name of the DB cluster parameter group.
	//
	// Constraints:
	//   +  Must not match the name of an existing DB cluster parameter group.
	//
	//   This value is stored as a lowercase string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_cluster_parameter_group#db_cluster_parameter_group_name RdsDbClusterParameterGroup#db_cluster_parameter_group_name}
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// Tags to assign to the DB cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_cluster_parameter_group#tags RdsDbClusterParameterGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

