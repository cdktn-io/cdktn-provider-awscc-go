// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplate


type WisdomMessageTemplateDefaultAttributes struct {
	// The agent attributes that are used with the message template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_message_template#agent_attributes WisdomMessageTemplate#agent_attributes}
	AgentAttributes *WisdomMessageTemplateDefaultAttributesAgentAttributes `field:"optional" json:"agentAttributes" yaml:"agentAttributes"`
	// The custom attributes that are used with the message template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_message_template#custom_attributes WisdomMessageTemplate#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The customer profile attributes that are used with the message template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_message_template#customer_profile_attributes WisdomMessageTemplate#customer_profile_attributes}
	CustomerProfileAttributes *WisdomMessageTemplateDefaultAttributesCustomerProfileAttributes `field:"optional" json:"customerProfileAttributes" yaml:"customerProfileAttributes"`
	// The system attributes that are used with the message template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_message_template#system_attributes WisdomMessageTemplate#system_attributes}
	SystemAttributes *WisdomMessageTemplateDefaultAttributesSystemAttributes `field:"optional" json:"systemAttributes" yaml:"systemAttributes"`
}

