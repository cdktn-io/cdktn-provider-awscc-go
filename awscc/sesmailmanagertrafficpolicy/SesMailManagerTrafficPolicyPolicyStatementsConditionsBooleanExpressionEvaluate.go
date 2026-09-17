// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagertrafficpolicy


type SesMailManagerTrafficPolicyPolicyStatementsConditionsBooleanExpressionEvaluate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_traffic_policy#analysis SesMailManagerTrafficPolicy#analysis}.
	Analysis *SesMailManagerTrafficPolicyPolicyStatementsConditionsBooleanExpressionEvaluateAnalysis `field:"optional" json:"analysis" yaml:"analysis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_traffic_policy#is_in_address_list SesMailManagerTrafficPolicy#is_in_address_list}.
	IsInAddressList *SesMailManagerTrafficPolicyPolicyStatementsConditionsBooleanExpressionEvaluateIsInAddressListStruct `field:"optional" json:"isInAddressList" yaml:"isInAddressList"`
}

