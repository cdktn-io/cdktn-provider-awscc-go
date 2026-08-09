// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsdeliverydestination


type LogsDeliveryDestinationDeliveryDestinationPolicy struct {
	// The name of the delivery destination to assign this policy to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/logs_delivery_destination#delivery_destination_name LogsDeliveryDestination#delivery_destination_name}
	DeliveryDestinationName *string `field:"optional" json:"deliveryDestinationName" yaml:"deliveryDestinationName"`
	// The contents of the policy attached to the delivery destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/logs_delivery_destination#delivery_destination_policy LogsDeliveryDestination#delivery_destination_policy}
	DeliveryDestinationPolicy *string `field:"optional" json:"deliveryDestinationPolicy" yaml:"deliveryDestinationPolicy"`
}

