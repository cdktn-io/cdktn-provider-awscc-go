// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationstack

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudformationStackConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#stack_name CloudformationStack#stack_name}.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#capabilities CloudformationStack#capabilities}.
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#description CloudformationStack#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#disable_rollback CloudformationStack#disable_rollback}.
	DisableRollback interface{} `field:"optional" json:"disableRollback" yaml:"disableRollback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#enable_termination_protection CloudformationStack#enable_termination_protection}.
	EnableTerminationProtection interface{} `field:"optional" json:"enableTerminationProtection" yaml:"enableTerminationProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#notification_ar_ns CloudformationStack#notification_ar_ns}.
	NotificationArNs *[]*string `field:"optional" json:"notificationArNs" yaml:"notificationArNs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#parameters CloudformationStack#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#role_arn CloudformationStack#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#stack_policy_body CloudformationStack#stack_policy_body}.
	StackPolicyBody *string `field:"optional" json:"stackPolicyBody" yaml:"stackPolicyBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#stack_policy_url CloudformationStack#stack_policy_url}.
	StackPolicyUrl *string `field:"optional" json:"stackPolicyUrl" yaml:"stackPolicyUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#stack_status_reason CloudformationStack#stack_status_reason}.
	StackStatusReason *string `field:"optional" json:"stackStatusReason" yaml:"stackStatusReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#tags CloudformationStack#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#template_body CloudformationStack#template_body}.
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#template_url CloudformationStack#template_url}.
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack#timeout_in_minutes CloudformationStack#timeout_in_minutes}.
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

