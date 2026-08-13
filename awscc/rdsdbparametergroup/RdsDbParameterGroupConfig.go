// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbparametergroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsDbParameterGroupConfig struct {
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
	// Provides the customer-specified description for this DB parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rds_db_parameter_group#description RdsDbParameterGroup#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The DB parameter group family name.
	//
	// A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a database engine and engine version compatible with that DB parameter group family.
	//  To list all of the available parameter group families for a DB engine, use the following command:
	//   ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine <engine>``
	//  For example, to list all of the available parameter group families for the MySQL DB engine, use the following command:
	//   ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine mysql``
	//   The output contains duplicates.
	//   The following are the valid DB engine values:
	//   +   ``aurora-mysql``
	//   +   ``aurora-postgresql``
	//   +   ``db2-ae``
	//   +   ``db2-ce``
	//   +   ``db2-se``
	//   +   ``mysql``
	//   +   ``oracle-ee``
	//   +   ``oracle-ee-cdb``
	//   +   ``oracle-se2``
	//   +   ``oracle-se2-cdb``
	//   +   ``postgres``
	//   +   ``sqlserver-ee``
	//   +   ``sqlserver-se``
	//   +   ``sqlserver-ex``
	//   +   ``sqlserver-web``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rds_db_parameter_group#family RdsDbParameterGroup#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// The name of the DB parameter group.
	//
	// Constraints:
	//   +  Must be 1 to 255 letters, numbers, or hyphens.
	//   +  First character must be a letter
	//   +  Can't end with a hyphen or contain two consecutive hyphens
	//
	//  If you don't specify a value for ``DBParameterGroupName`` property, a name is automatically created for the DB parameter group.
	//   This value is stored as a lowercase string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rds_db_parameter_group#db_parameter_group_name RdsDbParameterGroup#db_parameter_group_name}
	DbParameterGroupName *string `field:"optional" json:"dbParameterGroupName" yaml:"dbParameterGroupName"`
	// A mapping of parameter names and values for the parameter update.
	//
	// You must specify at least one parameter name and value.
	//  For more information about parameter groups, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide*, or [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.
	//   AWS CloudFormation doesn't support specifying an apply method for each individual parameter. The default apply method for each parameter is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rds_db_parameter_group#parameters RdsDbParameterGroup#parameters}
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// Tags to assign to the DB parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rds_db_parameter_group#tags RdsDbParameterGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

