// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelessfuotatask


type IotwirelessFuotaTaskLoRaWan struct {
	// FUOTA task LoRaWAN RF region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_fuota_task#rf_region IotwirelessFuotaTask#rf_region}
	RfRegion *string `field:"required" json:"rfRegion" yaml:"rfRegion"`
}

