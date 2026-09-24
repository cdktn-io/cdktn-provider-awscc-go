// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationchangeset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudformationChangeSetConfig struct {
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
	// The name of the change set. Must be unique among all change sets associated with the specified stack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#change_set_name CloudformationChangeSet#change_set_name}
	ChangeSetName *string `field:"required" json:"changeSetName" yaml:"changeSetName"`
	// The name or unique ID of the stack for which you are creating a change set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#stack_name CloudformationChangeSet#stack_name}
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// The capabilities that are allowed in the stack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#capabilities CloudformationChangeSet#capabilities}
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// The type of change set operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#change_set_type CloudformationChangeSet#change_set_type}
	ChangeSetType *string `field:"optional" json:"changeSetType" yaml:"changeSetType"`
	// Determines how CloudFormation handles configuration drift during deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#deployment_mode CloudformationChangeSet#deployment_mode}
	DeploymentMode *string `field:"optional" json:"deploymentMode" yaml:"deploymentMode"`
	// A description to help you identify this change set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#description CloudformationChangeSet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates if the change set imports resources that already exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#import_existing_resources CloudformationChangeSet#import_existing_resources}
	ImportExistingResources interface{} `field:"optional" json:"importExistingResources" yaml:"importExistingResources"`
	// Creates a change set for all nested stacks specified in the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#include_nested_stacks CloudformationChangeSet#include_nested_stacks}
	IncludeNestedStacks interface{} `field:"optional" json:"includeNestedStacks" yaml:"includeNestedStacks"`
	// The ARNs of Amazon SNS topics that CloudFormation associates with the stack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#notification_ar_ns CloudformationChangeSet#notification_ar_ns}
	NotificationArNs *[]*string `field:"optional" json:"notificationArNs" yaml:"notificationArNs"`
	// Determines what action will be taken if stack creation fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#on_stack_failure CloudformationChangeSet#on_stack_failure}
	OnStackFailure *string `field:"optional" json:"onStackFailure" yaml:"onStackFailure"`
	// The ARN of an IAM role that CloudFormation assumes when executing the change set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#role_arn CloudformationChangeSet#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Key-value pairs to associate with the change set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#tags CloudformationChangeSet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A structure that contains the body of the revised template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#template_body CloudformationChangeSet#template_body}
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// The URL of the file that contains the revised template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#template_url CloudformationChangeSet#template_url}
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
	// Whether to reuse the template associated with the stack to create the change set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_change_set#use_previous_template CloudformationChangeSet#use_previous_template}
	UsePreviousTemplate interface{} `field:"optional" json:"usePreviousTemplate" yaml:"usePreviousTemplate"`
}

