// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberTransformerJsonataConfiguration struct {
	// The JSONata expression that transforms the event, enclosed in {% %} delimiters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#expression Eventsv2Subscriber#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

