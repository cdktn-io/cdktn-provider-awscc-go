// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmopsitem

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsmOpsItemConfig struct {
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
	// The description of the OpsItem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_ops_item#description SsmOpsItem#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The origin of the OpsItem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_ops_item#source SsmOpsItem#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// The title of the OpsItem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_ops_item#title SsmOpsItem#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// The category of the OpsItem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_ops_item#category SsmOpsItem#category}
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The priority of the OpsItem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_ops_item#priority SsmOpsItem#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The severity of the OpsItem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_ops_item#severity SsmOpsItem#severity}
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
	// Tags for the OpsItem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_ops_item#tags SsmOpsItem#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

