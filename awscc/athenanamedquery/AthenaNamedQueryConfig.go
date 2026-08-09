// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenanamedquery

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AthenaNamedQueryConfig struct {
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
	// The database to which the query belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/athena_named_query#database AthenaNamedQuery#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The contents of the query with all query statements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/athena_named_query#query_string AthenaNamedQuery#query_string}
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// The query description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/athena_named_query#description AthenaNamedQuery#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The query name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/athena_named_query#name AthenaNamedQuery#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the workgroup that contains the named query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/athena_named_query#work_group AthenaNamedQuery#work_group}
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

