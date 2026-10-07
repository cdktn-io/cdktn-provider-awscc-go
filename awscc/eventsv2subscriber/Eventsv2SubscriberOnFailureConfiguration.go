// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberOnFailureConfiguration struct {
	// The ARN of the destination that receives events that could not be delivered.
	//
	// An Amazon SQS queue is the supported destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#arn Eventsv2Subscriber#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

