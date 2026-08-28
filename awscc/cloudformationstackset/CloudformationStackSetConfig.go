// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationstackset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudformationStackSetConfig struct {
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
	// Describes how the IAM roles required for stack set operations are created. By default, SELF-MANAGED is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#permission_model CloudformationStackSet#permission_model}
	PermissionModel *string `field:"required" json:"permissionModel" yaml:"permissionModel"`
	// The name to associate with the stack set.
	//
	// The name must be unique in the Region where you create your stack set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#stack_set_name CloudformationStackSet#stack_set_name}
	StackSetName *string `field:"required" json:"stackSetName" yaml:"stackSetName"`
	// The Amazon Resource Number (ARN) of the IAM role to use to create this stack set.
	//
	// Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#administration_role_arn CloudformationStackSet#administration_role_arn}
	AdministrationRoleArn *string `field:"optional" json:"administrationRoleArn" yaml:"administrationRoleArn"`
	// Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to the target organization or organizational unit (OU).
	//
	// Specify only if PermissionModel is SERVICE_MANAGED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#auto_deployment CloudformationStackSet#auto_deployment}
	AutoDeployment *CloudformationStackSetAutoDeployment `field:"optional" json:"autoDeployment" yaml:"autoDeployment"`
	// Specifies the AWS account that you are acting from.
	//
	// By default, SELF is specified. For self-managed permissions, specify SELF; for service-managed permissions, if you are signed in to the organization's management account, specify SELF. If you are signed in to a delegated administrator account, specify DELEGATED_ADMIN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#call_as CloudformationStackSet#call_as}
	CallAs *string `field:"optional" json:"callAs" yaml:"callAs"`
	// In some cases, you must explicitly acknowledge that your stack set template contains certain capabilities in order for AWS CloudFormation to create the stack set and related stack instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#capabilities CloudformationStackSet#capabilities}
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// A description of the stack set.
	//
	// You can use the description to identify the stack set's purpose or other important information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#description CloudformationStackSet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the IAM execution role to use to create the stack set.
	//
	// If you do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole role for the stack set operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#execution_role_name CloudformationStackSet#execution_role_name}
	ExecutionRoleName *string `field:"optional" json:"executionRoleName" yaml:"executionRoleName"`
	// Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#managed_execution CloudformationStackSet#managed_execution}
	ManagedExecution *CloudformationStackSetManagedExecution `field:"optional" json:"managedExecution" yaml:"managedExecution"`
	// The user-specified preferences for how AWS CloudFormation performs a stack set operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#operation_preferences CloudformationStackSet#operation_preferences}
	OperationPreferences *CloudformationStackSetOperationPreferences `field:"optional" json:"operationPreferences" yaml:"operationPreferences"`
	// The input parameters for the stack set template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#parameters CloudformationStackSet#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A group of stack instances with parameters in some specific accounts and regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#stack_instances_group CloudformationStackSet#stack_instances_group}
	StackInstancesGroup interface{} `field:"optional" json:"stackInstancesGroup" yaml:"stackInstancesGroup"`
	// The key-value pairs to associate with this stack set and the stacks created from it.
	//
	// AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#tags CloudformationStackSet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#template_body CloudformationStackSet#template_body}
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Location of file containing the template body.
	//
	// The URL must point to a template (max size: 460,800 bytes) that is located in an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#template_url CloudformationStackSet#template_url}
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
}

