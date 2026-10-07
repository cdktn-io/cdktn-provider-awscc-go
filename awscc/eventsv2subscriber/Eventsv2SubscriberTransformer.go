// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberTransformer struct {
	// The JSONata expression configuration. Required when Type is JSONATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#jsonata_configuration Eventsv2Subscriber#jsonata_configuration}
	JsonataConfiguration *Eventsv2SubscriberTransformerJsonataConfiguration `field:"optional" json:"jsonataConfiguration" yaml:"jsonataConfiguration"`
	// The transform type: RAW delivers the event payload only;
	//
	// WITH_METADATA delivers the event with its metadata envelope; JSONATA delivers the output of the JSONata expression in JsonataConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#type Eventsv2Subscriber#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

