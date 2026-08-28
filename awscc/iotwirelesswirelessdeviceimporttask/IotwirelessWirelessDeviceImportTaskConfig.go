// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesswirelessdeviceimporttask

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotwirelessWirelessDeviceImportTaskConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Destination Name for import task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device_import_task#destination_name IotwirelessWirelessDeviceImportTask#destination_name}
	DestinationName *string `field:"required" json:"destinationName" yaml:"destinationName"`
	// sidewalk contain file for created device and role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device_import_task#sidewalk IotwirelessWirelessDeviceImportTask#sidewalk}
	Sidewalk *IotwirelessWirelessDeviceImportTaskSidewalk `field:"required" json:"sidewalk" yaml:"sidewalk"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_wireless_device_import_task#tags IotwirelessWirelessDeviceImportTask#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

