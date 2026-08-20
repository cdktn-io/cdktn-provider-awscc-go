// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2


type SecurityhubAutomationRuleV2Criteria struct {
	// The filtering conditions that align with OCSF standards.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityhub_automation_rule_v2#ocsf_finding_criteria SecurityhubAutomationRuleV2#ocsf_finding_criteria}
	OcsfFindingCriteria *SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteria `field:"optional" json:"ocsfFindingCriteria" yaml:"ocsfFindingCriteria"`
}

