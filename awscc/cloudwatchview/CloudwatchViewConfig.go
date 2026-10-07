// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchview

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchViewConfig struct {
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
	// The SQL query that defines the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloudwatch_view#definition CloudwatchView#definition}
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// A description of the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloudwatch_view#description CloudwatchView#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the view.
	//
	// Must begin with the "view." prefix followed by lowercase alphanumeric characters, hyphens, and underscores. View names must be unique within the account and region. If not specified, a name is generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloudwatch_view#name CloudwatchView#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloudwatch_view#tags CloudwatchView#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

