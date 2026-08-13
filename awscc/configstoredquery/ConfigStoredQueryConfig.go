// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configstoredquery

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigStoredQueryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/config_stored_query#query_expression ConfigStoredQuery#query_expression}.
	QueryExpression *string `field:"required" json:"queryExpression" yaml:"queryExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/config_stored_query#query_name ConfigStoredQuery#query_name}.
	QueryName *string `field:"required" json:"queryName" yaml:"queryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/config_stored_query#query_description ConfigStoredQuery#query_description}.
	QueryDescription *string `field:"optional" json:"queryDescription" yaml:"queryDescription"`
	// The tags for the stored query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/config_stored_query#tags ConfigStoredQuery#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

