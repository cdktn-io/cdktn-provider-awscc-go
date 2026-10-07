// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorfunction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorFunctionConfig struct {
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
	// The unique identifier for the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_function#function_id MediatailorFunction#function_id}
	FunctionId *string `field:"required" json:"functionId" yaml:"functionId"`
	// The type of the function. Determines which configuration object is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_function#function_type MediatailorFunction#function_type}
	FunctionType *string `field:"required" json:"functionType" yaml:"functionType"`
	// The configuration for an AWS_SERVICE_REQUEST function.
	//
	// Contains the target service, target Region, and request parameters that the function uses to call an AWS service API. For more information, see AWS_SERVICE_REQUEST (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-types-aws-service-request.html) in the MediaTailor User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_function#aws_service_request_configuration MediatailorFunction#aws_service_request_configuration}
	AwsServiceRequestConfiguration *MediatailorFunctionAwsServiceRequestConfiguration `field:"optional" json:"awsServiceRequestConfiguration" yaml:"awsServiceRequestConfiguration"`
	// The configuration for a CONCURRENT_EXECUTOR function. Required when FunctionType is CONCURRENT_EXECUTOR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_function#concurrent_executor_configuration MediatailorFunction#concurrent_executor_configuration}
	ConcurrentExecutorConfiguration *MediatailorFunctionConcurrentExecutorConfiguration `field:"optional" json:"concurrentExecutorConfiguration" yaml:"concurrentExecutorConfiguration"`
	// Configuration for custom output functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_function#custom_output_configuration MediatailorFunction#custom_output_configuration}
	CustomOutputConfiguration *MediatailorFunctionCustomOutputConfiguration `field:"optional" json:"customOutputConfiguration" yaml:"customOutputConfiguration"`
	// A description of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_function#description MediatailorFunction#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration for HTTP request functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_function#http_request_configuration MediatailorFunction#http_request_configuration}
	HttpRequestConfiguration *MediatailorFunctionHttpRequestConfiguration `field:"optional" json:"httpRequestConfiguration" yaml:"httpRequestConfiguration"`
	// The configuration for a SEQUENTIAL_EXECUTOR function.
	//
	// A SEQUENTIAL_EXECUTOR runs an ordered list of child functions one at a time, passing data between them. For more information about functions, see Working with functions (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions.html) in the MediaTailor User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_function#sequential_executor_configuration MediatailorFunction#sequential_executor_configuration}
	SequentialExecutorConfiguration *MediatailorFunctionSequentialExecutorConfiguration `field:"optional" json:"sequentialExecutorConfiguration" yaml:"sequentialExecutorConfiguration"`
	// The tags to assign to the function resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_function#tags MediatailorFunction#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The configuration for a VAST_REQUEST function.
	//
	// Specifies the HTTP method, URL, headers, body, timeout, and output expressions for a request to a VAST endpoint. MediaTailor parses the response as VAST and resolves wrapper redirects, then makes the parsed ads available to the function's output expressions. For more information, see Function types and composition (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-types.html) in the MediaTailor User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_function#vast_request_configuration MediatailorFunction#vast_request_configuration}
	VastRequestConfiguration *MediatailorFunctionVastRequestConfiguration `field:"optional" json:"vastRequestConfiguration" yaml:"vastRequestConfiguration"`
}

