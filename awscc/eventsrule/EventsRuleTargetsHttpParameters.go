// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsrule


type EventsRuleTargetsHttpParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/events_rule#header_parameters EventsRule#header_parameters}.
	HeaderParameters *map[string]*string `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/events_rule#path_parameter_values EventsRule#path_parameter_values}.
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/events_rule#query_string_parameters EventsRule#query_string_parameters}.
	QueryStringParameters *map[string]*string `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

