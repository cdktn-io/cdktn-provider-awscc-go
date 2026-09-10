// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelessmulticastgroup


type IotwirelessMulticastGroupLoRaWan struct {
	// Multicast group LoRaWAN DL Class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotwireless_multicast_group#dl_class IotwirelessMulticastGroup#dl_class}
	DlClass *string `field:"required" json:"dlClass" yaml:"dlClass"`
	// Multicast group LoRaWAN RF region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotwireless_multicast_group#rf_region IotwirelessMulticastGroup#rf_region}
	RfRegion *string `field:"required" json:"rfRegion" yaml:"rfRegion"`
}

