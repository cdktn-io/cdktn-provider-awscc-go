// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberRetryPolicy struct {
	// The maximum age of an event in seconds, 60-86400 (24 hours).
	//
	// When an event reaches this age, retries stop; if OnFailureConfiguration is set, the event is delivered to that destination, otherwise it is dropped. The default is 300.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#max_event_age_in_seconds Eventsv2Subscriber#max_event_age_in_seconds}
	MaxEventAgeInSeconds *float64 `field:"optional" json:"maxEventAgeInSeconds" yaml:"maxEventAgeInSeconds"`
	// The maximum number of retry attempts, 0-185.
	//
	// When the attempts are exhausted, retries stop; if OnFailureConfiguration is set, the event is delivered to that destination, otherwise it is dropped. The default is 5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#max_retry_attempts Eventsv2Subscriber#max_retry_attempts}
	MaxRetryAttempts *float64 `field:"optional" json:"maxRetryAttempts" yaml:"maxRetryAttempts"`
	// Which errors are retried. ALL retries all errors. The default is ALL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#retry_strategy Eventsv2Subscriber#retry_strategy}
	RetryStrategy *string `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
}

