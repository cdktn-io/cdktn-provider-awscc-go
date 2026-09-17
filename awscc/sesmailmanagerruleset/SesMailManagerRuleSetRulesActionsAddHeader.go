// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagerruleset


type SesMailManagerRuleSetRulesActionsAddHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_rule_set#header_name SesMailManagerRuleSet#header_name}.
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_rule_set#header_value SesMailManagerRuleSet#header_value}.
	HeaderValue *string `field:"optional" json:"headerValue" yaml:"headerValue"`
}

