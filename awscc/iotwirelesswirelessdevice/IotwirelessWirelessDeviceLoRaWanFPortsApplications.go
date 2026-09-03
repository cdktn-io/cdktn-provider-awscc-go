// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesswirelessdevice


type IotwirelessWirelessDeviceLoRaWanFPortsApplications struct {
	// The name of the position data destination that describes the AWS IoT rule that processes the device's position data for use by AWS IoT Core for LoRaWAN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_wireless_device#destination_name IotwirelessWirelessDevice#destination_name}
	DestinationName *string `field:"optional" json:"destinationName" yaml:"destinationName"`
	// The Fport value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_wireless_device#f_port IotwirelessWirelessDevice#f_port}
	FPort *float64 `field:"optional" json:"fPort" yaml:"fPort"`
	// Application type, which can be specified to obtain real-time position information of your LoRaWAN device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_wireless_device#type IotwirelessWirelessDevice#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

