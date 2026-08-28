// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesswirelessdevice


type IotwirelessWirelessDeviceLoRaWanAbpV10X struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#dev_addr IotwirelessWirelessDevice#dev_addr}.
	DevAddr *string `field:"optional" json:"devAddr" yaml:"devAddr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#session_keys IotwirelessWirelessDevice#session_keys}.
	SessionKeys *IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeys `field:"optional" json:"sessionKeys" yaml:"sessionKeys"`
}

