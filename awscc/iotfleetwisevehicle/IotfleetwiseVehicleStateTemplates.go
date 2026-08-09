// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisevehicle


type IotfleetwiseVehicleStateTemplates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotfleetwise_vehicle#identifier IotfleetwiseVehicle#identifier}.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotfleetwise_vehicle#state_template_update_strategy IotfleetwiseVehicle#state_template_update_strategy}.
	StateTemplateUpdateStrategy *IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategy `field:"optional" json:"stateTemplateUpdateStrategy" yaml:"stateTemplateUpdateStrategy"`
}

