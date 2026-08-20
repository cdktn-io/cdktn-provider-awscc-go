// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesswirelessdevice


type IotwirelessWirelessDeviceLoRaWanOtaaV11 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotwireless_wireless_device#app_key IotwirelessWirelessDevice#app_key}.
	AppKey *string `field:"optional" json:"appKey" yaml:"appKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotwireless_wireless_device#join_eui IotwirelessWirelessDevice#join_eui}.
	JoinEui *string `field:"optional" json:"joinEui" yaml:"joinEui"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotwireless_wireless_device#nwk_key IotwirelessWirelessDevice#nwk_key}.
	NwkKey *string `field:"optional" json:"nwkKey" yaml:"nwkKey"`
}

