// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesswirelessdeviceimporttask


type IotwirelessWirelessDeviceImportTaskSidewalk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_wireless_device_import_task#device_creation_file IotwirelessWirelessDeviceImportTask#device_creation_file}.
	DeviceCreationFile *string `field:"optional" json:"deviceCreationFile" yaml:"deviceCreationFile"`
	// sidewalk role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_wireless_device_import_task#role IotwirelessWirelessDeviceImportTask#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_wireless_device_import_task#sidewalk_manufacturing_sn IotwirelessWirelessDeviceImportTask#sidewalk_manufacturing_sn}.
	SidewalkManufacturingSn *string `field:"optional" json:"sidewalkManufacturingSn" yaml:"sidewalkManufacturingSn"`
}

