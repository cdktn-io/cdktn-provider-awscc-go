// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncfunctionconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppsyncFunctionConfigurationConfig struct {
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
	// The AWS AppSync GraphQL API that you want to attach using this function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#api_id AppsyncFunctionConfiguration#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The name of data source this function will attach.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#data_source_name AppsyncFunctionConfiguration#data_source_name}
	DataSourceName *string `field:"required" json:"dataSourceName" yaml:"dataSourceName"`
	// The name of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#name AppsyncFunctionConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The resolver code that contains the request and response functions.
	//
	// When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#code AppsyncFunctionConfiguration#code}
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The Amazon S3 endpoint (where the code is located??).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#code_s3_location AppsyncFunctionConfiguration#code_s3_location}
	CodeS3Location *string `field:"optional" json:"codeS3Location" yaml:"codeS3Location"`
	// The function description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#description AppsyncFunctionConfiguration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version of the request mapping template. Currently, only the 2018-05-29 version of the template is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#function_version AppsyncFunctionConfiguration#function_version}
	FunctionVersion *string `field:"optional" json:"functionVersion" yaml:"functionVersion"`
	// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a BatchInvoke operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#max_batch_size AppsyncFunctionConfiguration#max_batch_size}
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#request_mapping_template AppsyncFunctionConfiguration#request_mapping_template}
	RequestMappingTemplate *string `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// Describes a Sync configuration for a resolver.
	//
	// Contains information on which Conflict Detection, as well as Resolution strategy, should be performed when the resolver is invoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#request_mapping_template_s3_location AppsyncFunctionConfiguration#request_mapping_template_s3_location}
	RequestMappingTemplateS3Location *string `field:"optional" json:"requestMappingTemplateS3Location" yaml:"requestMappingTemplateS3Location"`
	// The Function response mapping template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#response_mapping_template AppsyncFunctionConfiguration#response_mapping_template}
	ResponseMappingTemplate *string `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The location of a response mapping template in an Amazon S3 bucket.
	//
	// Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#response_mapping_template_s3_location AppsyncFunctionConfiguration#response_mapping_template_s3_location}
	ResponseMappingTemplateS3Location *string `field:"optional" json:"responseMappingTemplateS3Location" yaml:"responseMappingTemplateS3Location"`
	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function.
	//
	// Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#runtime AppsyncFunctionConfiguration#runtime}
	Runtime *AppsyncFunctionConfigurationRuntime `field:"optional" json:"runtime" yaml:"runtime"`
	// Describes a Sync configuration for a resolver.
	//
	// Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration#sync_config AppsyncFunctionConfiguration#sync_config}
	SyncConfig *AppsyncFunctionConfigurationSyncConfig `field:"optional" json:"syncConfig" yaml:"syncConfig"`
}

