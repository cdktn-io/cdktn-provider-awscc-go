// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagertrafficpolicy


type SesMailManagerTrafficPolicyPolicyStatementsConditionsBooleanExpressionEvaluateIsInAddressListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_traffic_policy#address_lists SesMailManagerTrafficPolicy#address_lists}.
	AddressLists *[]*string `field:"optional" json:"addressLists" yaml:"addressLists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_traffic_policy#attribute SesMailManagerTrafficPolicy#attribute}.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

