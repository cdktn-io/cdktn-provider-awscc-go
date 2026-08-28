// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotprovisioningtemplate


type IotProvisioningTemplatePreProvisioningHook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_provisioning_template#payload_version IotProvisioningTemplate#payload_version}.
	PayloadVersion *string `field:"optional" json:"payloadVersion" yaml:"payloadVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_provisioning_template#target_arn IotProvisioningTemplate#target_arn}.
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}

