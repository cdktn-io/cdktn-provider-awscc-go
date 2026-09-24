// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package snstopic


type SnsTopicSubscription struct {
	// The endpoint that receives notifications from the SNS topic.
	//
	// The endpoint value depends on the protocol that you specify. For more information, see the ``Endpoint`` parameter of the ``Subscribe`` action in the *API Reference*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sns_topic#endpoint SnsTopic#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The subscription's protocol. For more information, see the ``Protocol`` parameter of the ``Subscribe`` action in the *API Reference*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sns_topic#protocol SnsTopic#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

