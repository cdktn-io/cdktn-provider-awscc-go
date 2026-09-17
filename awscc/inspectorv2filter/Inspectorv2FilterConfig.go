// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2filter

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Inspectorv2FilterConfig struct {
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
	// Findings filter action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_filter#filter_action Inspectorv2Filter#filter_action}
	FilterAction *string `field:"required" json:"filterAction" yaml:"filterAction"`
	// Findings filter criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_filter#filter_criteria Inspectorv2Filter#filter_criteria}
	FilterCriteria *Inspectorv2FilterFilterCriteria `field:"required" json:"filterCriteria" yaml:"filterCriteria"`
	// Findings filter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_filter#name Inspectorv2Filter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Findings filter description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_filter#description Inspectorv2Filter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_filter#tags Inspectorv2Filter#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

