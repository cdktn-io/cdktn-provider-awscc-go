// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenapreparedstatement

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AthenaPreparedStatementConfig struct {
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
	// The query string for the prepared statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_prepared_statement#query_statement AthenaPreparedStatement#query_statement}
	QueryStatement *string `field:"required" json:"queryStatement" yaml:"queryStatement"`
	// The name of the prepared statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_prepared_statement#statement_name AthenaPreparedStatement#statement_name}
	StatementName *string `field:"required" json:"statementName" yaml:"statementName"`
	// The name of the workgroup to which the prepared statement belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_prepared_statement#work_group AthenaPreparedStatement#work_group}
	WorkGroup *string `field:"required" json:"workGroup" yaml:"workGroup"`
	// The description of the prepared statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_prepared_statement#description AthenaPreparedStatement#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

