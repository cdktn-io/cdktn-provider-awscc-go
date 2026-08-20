// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package xraysamplingrule


type XraySamplingRuleSamplingRuleUpdate struct {
	// Matches attributes derived from the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/xray_sampling_rule#attributes XraySamplingRule#attributes}
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/xray_sampling_rule#fixed_rate XraySamplingRule#fixed_rate}
	FixedRate *float64 `field:"optional" json:"fixedRate" yaml:"fixedRate"`
	// Matches the hostname from a request URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/xray_sampling_rule#host XraySamplingRule#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Matches the HTTP method from a request URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/xray_sampling_rule#http_method XraySamplingRule#http_method}
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// The priority of the sampling rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/xray_sampling_rule#priority XraySamplingRule#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate.
	//
	// The reservoir is not used directly by services, but applies to all services using the rule collectively.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/xray_sampling_rule#reservoir_size XraySamplingRule#reservoir_size}
	ReservoirSize *float64 `field:"optional" json:"reservoirSize" yaml:"reservoirSize"`
	// Matches the ARN of the AWS resource on which the service runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/xray_sampling_rule#resource_arn XraySamplingRule#resource_arn}
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/xray_sampling_rule#rule_arn XraySamplingRule#rule_arn}
	RuleArn *string `field:"optional" json:"ruleArn" yaml:"ruleArn"`
	// The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/xray_sampling_rule#rule_name XraySamplingRule#rule_name}
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/xray_sampling_rule#sampling_rate_boost XraySamplingRule#sampling_rate_boost}.
	SamplingRateBoost *XraySamplingRuleSamplingRuleUpdateSamplingRateBoost `field:"optional" json:"samplingRateBoost" yaml:"samplingRateBoost"`
	// Matches the name that the service uses to identify itself in segments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/xray_sampling_rule#service_name XraySamplingRule#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Matches the origin that the service uses to identify its type in segments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/xray_sampling_rule#service_type XraySamplingRule#service_type}
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
	// Matches the path from a request URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/xray_sampling_rule#url_path XraySamplingRule#url_path}
	UrlPath *string `field:"optional" json:"urlPath" yaml:"urlPath"`
}

