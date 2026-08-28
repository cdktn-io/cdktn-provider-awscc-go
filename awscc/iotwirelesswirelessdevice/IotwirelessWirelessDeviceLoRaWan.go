// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesswirelessdevice


type IotwirelessWirelessDeviceLoRaWan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#abp_v10_x IotwirelessWirelessDevice#abp_v10_x}.
	AbpV10X *IotwirelessWirelessDeviceLoRaWanAbpV10X `field:"optional" json:"abpV10X" yaml:"abpV10X"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#abp_v11 IotwirelessWirelessDevice#abp_v11}.
	AbpV11 *IotwirelessWirelessDeviceLoRaWanAbpV11 `field:"optional" json:"abpV11" yaml:"abpV11"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#dev_eui IotwirelessWirelessDevice#dev_eui}.
	DevEui *string `field:"optional" json:"devEui" yaml:"devEui"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#device_profile_id IotwirelessWirelessDevice#device_profile_id}.
	DeviceProfileId *string `field:"optional" json:"deviceProfileId" yaml:"deviceProfileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#f_ports IotwirelessWirelessDevice#f_ports}.
	FPorts *IotwirelessWirelessDeviceLoRaWanFPorts `field:"optional" json:"fPorts" yaml:"fPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#otaa_v10_x IotwirelessWirelessDevice#otaa_v10_x}.
	OtaaV10X *IotwirelessWirelessDeviceLoRaWanOtaaV10X `field:"optional" json:"otaaV10X" yaml:"otaaV10X"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#otaa_v11 IotwirelessWirelessDevice#otaa_v11}.
	OtaaV11 *IotwirelessWirelessDeviceLoRaWanOtaaV11 `field:"optional" json:"otaaV11" yaml:"otaaV11"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device#service_profile_id IotwirelessWirelessDevice#service_profile_id}.
	ServiceProfileId *string `field:"optional" json:"serviceProfileId" yaml:"serviceProfileId"`
}

