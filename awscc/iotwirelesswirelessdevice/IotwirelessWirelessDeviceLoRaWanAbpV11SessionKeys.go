// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesswirelessdevice


type IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeys struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#app_s_key IotwirelessWirelessDevice#app_s_key}.
	AppSKey *string `field:"optional" json:"appSKey" yaml:"appSKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#f_nwk_s_int_key IotwirelessWirelessDevice#f_nwk_s_int_key}.
	FNwkSIntKey *string `field:"optional" json:"fNwkSIntKey" yaml:"fNwkSIntKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#nwk_s_enc_key IotwirelessWirelessDevice#nwk_s_enc_key}.
	NwkSEncKey *string `field:"optional" json:"nwkSEncKey" yaml:"nwkSEncKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#s_nwk_s_int_key IotwirelessWirelessDevice#s_nwk_s_int_key}.
	SNwkSIntKey *string `field:"optional" json:"sNwkSIntKey" yaml:"sNwkSIntKey"`
}

