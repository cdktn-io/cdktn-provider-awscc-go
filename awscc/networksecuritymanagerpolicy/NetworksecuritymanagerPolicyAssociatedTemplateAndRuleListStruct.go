// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecuritymanagerpolicy


type NetworksecuritymanagerPolicyAssociatedTemplateAndRuleListStruct struct {
	// ARN of the associated rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/networksecuritymanager_policy#rule_arn NetworksecuritymanagerPolicy#rule_arn}
	RuleArn *string `field:"optional" json:"ruleArn" yaml:"ruleArn"`
	// ARN of the associated template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/networksecuritymanager_policy#template_arn NetworksecuritymanagerPolicy#template_arn}
	TemplateArn *string `field:"optional" json:"templateArn" yaml:"templateArn"`
}

