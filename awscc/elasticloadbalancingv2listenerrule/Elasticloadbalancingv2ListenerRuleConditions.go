// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listenerrule


type Elasticloadbalancingv2ListenerRuleConditions struct {
	// The name of the field.
	//
	// The possible values are:
	//   +  ``http-header`` ? [ALB] Matches on an HTTP header field.
	//   +  ``http-request-method`` ? [ALB] Matches on the HTTP request method.
	//   +  ``host-header`` ? [ALB] Matches on the host header.
	//   +  ``path-pattern`` ? [ALB] Matches on the URL path of the request.
	//   +  ``query-string`` ? [ALB] Matches on a query string parameter.
	//   +  ``source-ip`` ? [ALB, NLB] Matches on the source IP address. For ALB, use ``SourceIpConfig`` with ``Values`` to specify CIDR ranges. For NLB, use ``SourceIpConfig`` with ``IpAddressType`` to match the IP address type (``ipv4`` or ``ipv6``).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elasticloadbalancingv2_listener_rule#field Elasticloadbalancingv2ListenerRule#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Information for a host header condition. Specify only when ``Field`` is ``host-header``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elasticloadbalancingv2_listener_rule#host_header_config Elasticloadbalancingv2ListenerRule#host_header_config}
	HostHeaderConfig *Elasticloadbalancingv2ListenerRuleConditionsHostHeaderConfig `field:"optional" json:"hostHeaderConfig" yaml:"hostHeaderConfig"`
	// Information for an HTTP header condition. Specify only when ``Field`` is ``http-header``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elasticloadbalancingv2_listener_rule#http_header_config Elasticloadbalancingv2ListenerRule#http_header_config}
	HttpHeaderConfig *Elasticloadbalancingv2ListenerRuleConditionsHttpHeaderConfig `field:"optional" json:"httpHeaderConfig" yaml:"httpHeaderConfig"`
	// Information for an HTTP method condition. Specify only when ``Field`` is ``http-request-method``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elasticloadbalancingv2_listener_rule#http_request_method_config Elasticloadbalancingv2ListenerRule#http_request_method_config}
	HttpRequestMethodConfig *Elasticloadbalancingv2ListenerRuleConditionsHttpRequestMethodConfig `field:"optional" json:"httpRequestMethodConfig" yaml:"httpRequestMethodConfig"`
	// Information for a path pattern condition. Specify only when ``Field`` is ``path-pattern``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elasticloadbalancingv2_listener_rule#path_pattern_config Elasticloadbalancingv2ListenerRule#path_pattern_config}
	PathPatternConfig *Elasticloadbalancingv2ListenerRuleConditionsPathPatternConfig `field:"optional" json:"pathPatternConfig" yaml:"pathPatternConfig"`
	// Information for a query string condition. Specify only when ``Field`` is ``query-string``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elasticloadbalancingv2_listener_rule#query_string_config Elasticloadbalancingv2ListenerRule#query_string_config}
	QueryStringConfig *Elasticloadbalancingv2ListenerRuleConditionsQueryStringConfig `field:"optional" json:"queryStringConfig" yaml:"queryStringConfig"`
	// The regular expressions to match against the condition field.
	//
	// The maximum length of each string is 128 characters. Specify only when ``Field`` is ``http-header``, ``host-header``, or ``path-pattern``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elasticloadbalancingv2_listener_rule#regex_values Elasticloadbalancingv2ListenerRule#regex_values}
	RegexValues *[]*string `field:"optional" json:"regexValues" yaml:"regexValues"`
	// Information for a source IP condition. Specify only when ``Field`` is ``source-ip``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elasticloadbalancingv2_listener_rule#source_ip_config Elasticloadbalancingv2ListenerRule#source_ip_config}
	SourceIpConfig *Elasticloadbalancingv2ListenerRuleConditionsSourceIpConfig `field:"optional" json:"sourceIpConfig" yaml:"sourceIpConfig"`
	// The condition value.
	//
	// Specify only when ``Field`` is ``host-header`` or ``path-pattern``. Alternatively, to specify multiple host names or multiple path patterns, use ``HostHeaderConfig`` or ``PathPatternConfig``.
	//  If ``Field`` is ``host-header`` and you're not using ``HostHeaderConfig``, you can specify a single host name (for example, my.example.com). A host name is case insensitive, can be up to 128 characters in length, and can contain any of the following characters.
	//   +  A-Z, a-z, 0-9
	//   +  - .
	//   +  * (matches 0 or more characters)
	//   +  ? (matches exactly 1 character)
	//
	//  If ``Field`` is ``path-pattern`` and you're not using ``PathPatternConfig``, you can specify a single path pattern (for example, /img/*). A path pattern is case-sensitive, can be up to 128 characters in length, and can contain any of the following characters.
	//   +  A-Z, a-z, 0-9
	//   +  _ - . $ / ~ " ' @ : +
	//   +  & (using &amp;)
	//   +  * (matches 0 or more characters)
	//   +  ? (matches exactly 1 character)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elasticloadbalancingv2_listener_rule#values Elasticloadbalancingv2ListenerRule#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

