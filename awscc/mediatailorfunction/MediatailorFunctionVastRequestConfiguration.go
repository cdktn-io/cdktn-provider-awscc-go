// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorfunction


type MediatailorFunctionVastRequestConfiguration struct {
	// An expression that evaluates to the request body, for example to send an OpenRTB bid request.
	//
	// The expression can be up to 100,000 characters, and the body after evaluation can be up to 64 KB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#body MediatailorFunction#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// A map of HTTP header names to expression values.
	//
	// MediaTailor evaluates each header value expression at runtime and includes the result in the outbound request. Headers beginning with X-Amz- are reserved by the service, and method override headers are not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#headers MediatailorFunction#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#method_type MediatailorFunction#method_type}.
	MethodType *string `field:"optional" json:"methodType" yaml:"methodType"`
	// A map of output bindings.
	//
	// Each key is a namespaced output path (such as temp.wrappedAds), and each value is an expression that MediaTailor evaluates at runtime. Output expressions in a VAST_REQUEST function can reference the response object, which exposes response.parsedAds, the ads parsed from the VAST response after schema validation and wrapper resolution, and response.statusCode. For more information about expression syntax, see JSONata expression reference (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-jsonata.html) in the MediaTailor User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#output MediatailorFunction#output}
	Output *map[string]*string `field:"optional" json:"output" yaml:"output"`
	// The maximum time, in milliseconds, that MediaTailor waits for a response from the VAST endpoint.
	//
	// The timeout covers the entire response, including any wrapper redirects that MediaTailor follows. If the call exceeds this timeout, MediaTailor proceeds with an empty ad list and continues output expression evaluation. Valid values are 100 to 2000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#request_timeout_milliseconds MediatailorFunction#request_timeout_milliseconds}
	RequestTimeoutMilliseconds *float64 `field:"optional" json:"requestTimeoutMilliseconds" yaml:"requestTimeoutMilliseconds"`
	// The expression language used to evaluate expressions in the function configuration. Set this to JSONATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#runtime MediatailorFunction#runtime}
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// An expression that evaluates to the VAST endpoint URL.
	//
	// Use {%...%} delimiters for dynamic expressions. A literal value must be an https:// URL. The expression can be up to 25,000 characters, and the URL after evaluation can be up to 2,048 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#url MediatailorFunction#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

