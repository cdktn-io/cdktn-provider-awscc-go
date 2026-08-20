// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdevice


type SagemakerDeviceDevice struct {
	// Description of the device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_device#description SagemakerDevice#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_device#device_name SagemakerDevice#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// AWS Internet of Things (IoT) object name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_device#iot_thing_name SagemakerDevice#iot_thing_name}
	IotThingName *string `field:"optional" json:"iotThingName" yaml:"iotThingName"`
}

