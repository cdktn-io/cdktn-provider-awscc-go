// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesstaskdefinition


type IotwirelessTaskDefinitionUpdateLoRaWan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_task_definition#current_version IotwirelessTaskDefinition#current_version}.
	CurrentVersion *IotwirelessTaskDefinitionUpdateLoRaWanCurrentVersion `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_task_definition#sig_key_crc IotwirelessTaskDefinition#sig_key_crc}.
	SigKeyCrc *float64 `field:"optional" json:"sigKeyCrc" yaml:"sigKeyCrc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_task_definition#update_signature IotwirelessTaskDefinition#update_signature}.
	UpdateSignature *string `field:"optional" json:"updateSignature" yaml:"updateSignature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_task_definition#update_version IotwirelessTaskDefinition#update_version}.
	UpdateVersion *IotwirelessTaskDefinitionUpdateLoRaWanUpdateVersion `field:"optional" json:"updateVersion" yaml:"updateVersion"`
}

