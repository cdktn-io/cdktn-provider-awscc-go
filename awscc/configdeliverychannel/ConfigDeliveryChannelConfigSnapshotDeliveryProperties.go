// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configdeliverychannel


type ConfigDeliveryChannelConfigSnapshotDeliveryProperties struct {
	// The frequency with which AWS Config delivers configuration snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/config_delivery_channel#delivery_frequency ConfigDeliveryChannel#delivery_frequency}
	DeliveryFrequency *string `field:"optional" json:"deliveryFrequency" yaml:"deliveryFrequency"`
}

