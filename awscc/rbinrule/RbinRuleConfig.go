// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rbinrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RbinRuleConfig struct {
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
	// The resource type retained by the retention rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rbin_rule#resource_type RbinRule#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Information about the retention period for which the retention rule is to retain resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rbin_rule#retention_period RbinRule#retention_period}
	RetentionPeriod *RbinRuleRetentionPeriod `field:"required" json:"retentionPeriod" yaml:"retentionPeriod"`
	// The description of the retention rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rbin_rule#description RbinRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Information about the exclude resource tags used to identify resources that are excluded by the retention rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rbin_rule#exclude_resource_tags RbinRule#exclude_resource_tags}
	ExcludeResourceTags interface{} `field:"optional" json:"excludeResourceTags" yaml:"excludeResourceTags"`
	// Information about the retention rule lock configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rbin_rule#lock_configuration RbinRule#lock_configuration}
	LockConfiguration *RbinRuleLockConfiguration `field:"optional" json:"lockConfiguration" yaml:"lockConfiguration"`
	// Information about the resource tags used to identify resources that are retained by the retention rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rbin_rule#resource_tags RbinRule#resource_tags}
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// The state of the retention rule. Only retention rules that are in the available state retain resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rbin_rule#status RbinRule#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Information about the tags assigned to the retention rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rbin_rule#tags RbinRule#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

