// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset


type IotanalyticsDatasetLateDataRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_dataset#rule_configuration IotanalyticsDataset#rule_configuration}.
	RuleConfiguration *IotanalyticsDatasetLateDataRulesRuleConfiguration `field:"optional" json:"ruleConfiguration" yaml:"ruleConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_dataset#rule_name IotanalyticsDataset#rule_name}.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
}

