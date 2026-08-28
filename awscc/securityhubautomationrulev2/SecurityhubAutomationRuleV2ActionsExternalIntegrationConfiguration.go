// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2


type SecurityhubAutomationRuleV2ActionsExternalIntegrationConfiguration struct {
	// The ARN of the connector that establishes the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityhub_automation_rule_v2#connector_arn SecurityhubAutomationRuleV2#connector_arn}
	ConnectorArn *string `field:"optional" json:"connectorArn" yaml:"connectorArn"`
}

