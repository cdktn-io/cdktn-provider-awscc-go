// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventinvokeconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaEventInvokeConfigConfig struct {
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
	// The name of the Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_event_invoke_config#function_name LambdaEventInvokeConfig#function_name}
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The identifier of a version or alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_event_invoke_config#qualifier LambdaEventInvokeConfig#qualifier}
	Qualifier *string `field:"required" json:"qualifier" yaml:"qualifier"`
	// A destination for events after they have been sent to a function for processing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_event_invoke_config#destination_config LambdaEventInvokeConfig#destination_config}
	DestinationConfig *LambdaEventInvokeConfigDestinationConfig `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_event_invoke_config#maximum_event_age_in_seconds LambdaEventInvokeConfig#maximum_event_age_in_seconds}
	MaximumEventAgeInSeconds *float64 `field:"optional" json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_event_invoke_config#maximum_retry_attempts LambdaEventInvokeConfig#maximum_retry_attempts}
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

