// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2loggingconfiguration


type Wafv2LoggingConfigurationRedactedFields struct {
	// Inspect the HTTP method.
	//
	// The method indicates the type of operation that the request is asking the origin to perform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wafv2_logging_configuration#method Wafv2LoggingConfiguration#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Inspect the query string.
	//
	// This is the part of a URL that appears after a ? character, if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wafv2_logging_configuration#query_string Wafv2LoggingConfiguration#query_string}
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// Inspect a single header.
	//
	// Provide the name of the header to inspect, for example, User-Agent or Referer. This setting isn't case sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wafv2_logging_configuration#single_header Wafv2LoggingConfiguration#single_header}
	SingleHeader *Wafv2LoggingConfigurationRedactedFieldsSingleHeader `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// Inspect the request URI path.
	//
	// This is the part of a web request that identifies a resource, for example, /images/daily-ad.jpg.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wafv2_logging_configuration#uri_path Wafv2LoggingConfiguration#uri_path}
	UriPath *string `field:"optional" json:"uriPath" yaml:"uriPath"`
}

