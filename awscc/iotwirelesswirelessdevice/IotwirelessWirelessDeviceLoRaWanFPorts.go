// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesswirelessdevice


type IotwirelessWirelessDeviceLoRaWanFPorts struct {
	// A list of optional LoRaWAN application information, which can be used for geolocation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_wireless_device#applications IotwirelessWirelessDevice#applications}
	Applications interface{} `field:"optional" json:"applications" yaml:"applications"`
}

