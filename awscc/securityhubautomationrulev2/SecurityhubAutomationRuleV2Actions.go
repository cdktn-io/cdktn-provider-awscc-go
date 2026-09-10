// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2


type SecurityhubAutomationRuleV2Actions struct {
	// The category of action to be executed by the automation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_automation_rule_v2#type SecurityhubAutomationRuleV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The settings for integrating automation rule actions with external systems or service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_automation_rule_v2#external_integration_configuration SecurityhubAutomationRuleV2#external_integration_configuration}
	ExternalIntegrationConfiguration *SecurityhubAutomationRuleV2ActionsExternalIntegrationConfiguration `field:"optional" json:"externalIntegrationConfiguration" yaml:"externalIntegrationConfiguration"`
	// The changes to be applied to fields in a security finding when an automation rule is triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_automation_rule_v2#finding_fields_update SecurityhubAutomationRuleV2#finding_fields_update}
	FindingFieldsUpdate *SecurityhubAutomationRuleV2ActionsFindingFieldsUpdate `field:"optional" json:"findingFieldsUpdate" yaml:"findingFieldsUpdate"`
}

