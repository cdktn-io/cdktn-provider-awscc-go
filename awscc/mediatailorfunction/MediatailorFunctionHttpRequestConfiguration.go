// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorfunction


type MediatailorFunctionHttpRequestConfiguration struct {
	// The body of the HTTP request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediatailor_function#body MediatailorFunction#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// A map of HTTP headers to include in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediatailor_function#headers MediatailorFunction#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// The HTTP method type for the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediatailor_function#method_type MediatailorFunction#method_type}
	MethodType *string `field:"optional" json:"methodType" yaml:"methodType"`
	// A map of output key-value pairs.
	//
	// Keys must start with session., temp., avail., scte., or be a valid adsRequest directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediatailor_function#output MediatailorFunction#output}
	Output *map[string]*string `field:"optional" json:"output" yaml:"output"`
	// The timeout in milliseconds for the HTTP request. Maximum value is 2000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediatailor_function#request_timeout_milliseconds MediatailorFunction#request_timeout_milliseconds}
	RequestTimeoutMilliseconds *float64 `field:"optional" json:"requestTimeoutMilliseconds" yaml:"requestTimeoutMilliseconds"`
	// The runtime environment for the function expression language.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediatailor_function#runtime MediatailorFunction#runtime}
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// The URL endpoint for the HTTP request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediatailor_function#url MediatailorFunction#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

