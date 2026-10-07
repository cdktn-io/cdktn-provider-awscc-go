// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorfunction


type MediatailorFunctionAwsServiceRequestConfiguration struct {
	// An expression that evaluates to the request body for the AWS service API call.
	//
	// The body must conform to the input format that the target service operation expects. Applies only when the target operation accepts a request body. The maximum size after evaluation is 64 KB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#body MediatailorFunction#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// A map of HTTP header names to expression values.
	//
	// MediaTailor evaluates each header value expression at runtime and includes the result in the outbound request to the AWS service. Use this to pass any headers required by the target service operation. You can include a maximum of 50 headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#headers MediatailorFunction#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#method_type MediatailorFunction#method_type}.
	MethodType *string `field:"optional" json:"methodType" yaml:"methodType"`
	// A map of output bindings.
	//
	// Each key is a namespaced output path, such as player_params.device_type. Each value is an expression that MediaTailor evaluates at runtime and can reference the response object from the target service. For more information, see JSONata expression reference (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-jsonata.html) in the MediaTailor User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#output MediatailorFunction#output}
	Output *map[string]*string `field:"optional" json:"output" yaml:"output"`
	// The maximum time, in milliseconds, that MediaTailor waits for a response from the AWS service.
	//
	// If the call exceeds this timeout, MediaTailor sets the response status code to null and proceeds with output expression evaluation. Valid values are 100 to 2000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#request_timeout_milliseconds MediatailorFunction#request_timeout_milliseconds}
	RequestTimeoutMilliseconds *float64 `field:"optional" json:"requestTimeoutMilliseconds" yaml:"requestTimeoutMilliseconds"`
	// The expression language used to evaluate expressions in the function configuration. Set this to JSONATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#runtime MediatailorFunction#runtime}
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// The AWS Region for the target service.
	//
	// Specify a static Region code (for example, us-east-1) or a JSONata expression that resolves to a Region code at runtime (for example, {%inference.region%}).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#target_region MediatailorFunction#target_region}
	TargetRegion *string `field:"optional" json:"targetRegion" yaml:"targetRegion"`
	// The AWS service to call. Valid value: elemental-inference (AWS Elemental Inference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#target_service MediatailorFunction#target_service}
	TargetService *string `field:"optional" json:"targetService" yaml:"targetService"`
	// An expression that evaluates to the endpoint URL for the target AWS service API operation.
	//
	// Use {%...%} delimiters for dynamic expressions. The URL must correspond to a valid endpoint for the service specified in TargetService. The maximum length after evaluation is 2,048 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#url MediatailorFunction#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

