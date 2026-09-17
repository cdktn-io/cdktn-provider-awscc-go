// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowTriggerConfig struct {
	// Trigger type of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#trigger_type AppflowFlow#trigger_type}
	TriggerType *string `field:"required" json:"triggerType" yaml:"triggerType"`
	// Details required based on the type of trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#trigger_properties AppflowFlow#trigger_properties}
	TriggerProperties *AppflowFlowTriggerConfigTriggerProperties `field:"optional" json:"triggerProperties" yaml:"triggerProperties"`
}

