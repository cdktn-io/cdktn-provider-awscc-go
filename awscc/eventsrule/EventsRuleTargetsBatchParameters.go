// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsrule


type EventsRuleTargetsBatchParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/events_rule#array_properties EventsRule#array_properties}.
	ArrayProperties *EventsRuleTargetsBatchParametersArrayProperties `field:"optional" json:"arrayProperties" yaml:"arrayProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/events_rule#job_definition EventsRule#job_definition}.
	JobDefinition *string `field:"optional" json:"jobDefinition" yaml:"jobDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/events_rule#job_name EventsRule#job_name}.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/events_rule#retry_strategy EventsRule#retry_strategy}.
	RetryStrategy *EventsRuleTargetsBatchParametersRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
}

