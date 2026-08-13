// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationguardhook

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudformationGuardHookConfig struct {
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
	// The typename alias for the hook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_guard_hook#alias CloudformationGuardHook#alias}
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The execution role ARN assumed by hooks to read Guard rules from S3 and write Guard outputs to S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_guard_hook#execution_role CloudformationGuardHook#execution_role}
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// S3 Source Location for the Guard files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_guard_hook#rule_location CloudformationGuardHook#rule_location}
	RuleLocation *CloudformationGuardHookRuleLocation `field:"required" json:"ruleLocation" yaml:"ruleLocation"`
	// Which operations should this Hook run against? Resource changes, stacks or change sets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_guard_hook#target_operations CloudformationGuardHook#target_operations}
	TargetOperations *[]*string `field:"required" json:"targetOperations" yaml:"targetOperations"`
	// Attribute to specify CloudFormation behavior on hook failure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_guard_hook#failure_mode CloudformationGuardHook#failure_mode}
	FailureMode *string `field:"optional" json:"failureMode" yaml:"failureMode"`
	// Attribute to specify which stacks this hook applies to or should get invoked for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_guard_hook#hook_status CloudformationGuardHook#hook_status}
	HookStatus *string `field:"optional" json:"hookStatus" yaml:"hookStatus"`
	// S3 Bucket where the guard validate report will be uploaded to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_guard_hook#log_bucket CloudformationGuardHook#log_bucket}
	LogBucket *string `field:"optional" json:"logBucket" yaml:"logBucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_guard_hook#options CloudformationGuardHook#options}.
	Options *CloudformationGuardHookOptions `field:"optional" json:"options" yaml:"options"`
	// Filters to allow hooks to target specific stack attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_guard_hook#stack_filters CloudformationGuardHook#stack_filters}
	StackFilters *CloudformationGuardHookStackFilters `field:"optional" json:"stackFilters" yaml:"stackFilters"`
	// Attribute to specify which targets should invoke the hook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_guard_hook#target_filters CloudformationGuardHook#target_filters}
	TargetFilters *CloudformationGuardHookTargetFilters `field:"optional" json:"targetFilters" yaml:"targetFilters"`
}

