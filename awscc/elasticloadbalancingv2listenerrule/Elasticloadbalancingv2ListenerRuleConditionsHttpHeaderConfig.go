// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listenerrule


type Elasticloadbalancingv2ListenerRuleConditionsHttpHeaderConfig struct {
	// The name of the HTTP header field.
	//
	// The maximum size is 40 characters. The header name is case insensitive. The allowed characters are specified by RFC 7230. Wildcards are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener_rule#http_header_name Elasticloadbalancingv2ListenerRule#http_header_name}
	HttpHeaderName *string `field:"optional" json:"httpHeaderName" yaml:"httpHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener_rule#regex_values Elasticloadbalancingv2ListenerRule#regex_values}.
	RegexValues *[]*string `field:"optional" json:"regexValues" yaml:"regexValues"`
	// The strings to compare against the value of the HTTP header.
	//
	// The maximum length of each string is 128 characters. The comparison strings are case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).
	//  If the same header appears multiple times in the request, we search them in order until a match is found.
	//  If you specify multiple strings, the condition is satisfied if one of the strings matches the value of the HTTP header. To require that all of the strings are a match, create one condition per string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener_rule#values Elasticloadbalancingv2ListenerRule#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

