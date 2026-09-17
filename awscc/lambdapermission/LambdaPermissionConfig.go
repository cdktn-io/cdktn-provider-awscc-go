// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdapermission

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaPermissionConfig struct {
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
	// The action that the principal can use on the function. For example, ``lambda:InvokeFunction`` or ``lambda:GetFunction``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_permission#action LambdaPermission#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The name or ARN of the Lambda function, version, or alias.
	//
	// **Name formats**
	//  +  *Function name* ? ``my-function`` (name-only), ``my-function:v1`` (with alias).
	//   +  *Function ARN* ? ``arn:aws:lambda:us-west-2:123456789012:function:my-function``.
	//   +  *Partial ARN* ? ``123456789012:function:my-function``.
	//
	//  You can append a version number or alias to any of the formats. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_permission#function_name LambdaPermission#function_name}
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The AWS-service, AWS-account, IAM user, or IAM role that invokes the function.
	//
	// If you specify a service, use ``SourceArn`` or ``SourceAccount`` to limit who can invoke the function through that service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_permission#principal LambdaPermission#principal}
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// For Alexa Smart Home functions, a token that the invoker must supply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_permission#event_source_token LambdaPermission#event_source_token}
	EventSourceToken *string `field:"optional" json:"eventSourceToken" yaml:"eventSourceToken"`
	// The type of authentication that your function URL uses.
	//
	// Set to ``AWS_IAM`` if you want to restrict access to authenticated users only. Set to ``NONE`` if you want to bypass IAM authentication to create a public endpoint. For more information, see [Control access to Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_permission#function_url_auth_type LambdaPermission#function_url_auth_type}
	FunctionUrlAuthType *string `field:"optional" json:"functionUrlAuthType" yaml:"functionUrlAuthType"`
	// Indicates whether the permission applies when the function is invoked through a function URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_permission#invoked_via_function_url LambdaPermission#invoked_via_function_url}
	InvokedViaFunctionUrl interface{} `field:"optional" json:"invokedViaFunctionUrl" yaml:"invokedViaFunctionUrl"`
	// The identifier for your organization in AOlong. Use this to grant permissions to all the AWS-accounts under this organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_permission#principal_org_id LambdaPermission#principal_org_id}
	PrincipalOrgId *string `field:"optional" json:"principalOrgId" yaml:"principalOrgId"`
	// For AWS-service, the ID of the AWS-account that owns the resource.
	//
	// Use this together with ``SourceArn`` to ensure that the specified account owns the resource. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_permission#source_account LambdaPermission#source_account}
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
	// For AWS-services, the ARN of the AWS resource that invokes the function.
	//
	// For example, an Amazon S3 bucket or Amazon SNS topic.
	//  Note that Lambda configures the comparison using the ``StringLike`` operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_permission#source_arn LambdaPermission#source_arn}
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

