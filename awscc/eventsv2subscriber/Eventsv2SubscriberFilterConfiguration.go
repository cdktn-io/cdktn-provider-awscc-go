// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberFilterConfiguration struct {
	// The list of filters, 1-50 entries. An event must match every filter to be delivered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#filters Eventsv2Subscriber#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// The filter language. The default is EVENT_BRIDGE_PATTERN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#language Eventsv2Subscriber#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
}

