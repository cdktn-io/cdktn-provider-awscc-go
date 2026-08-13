// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagertrafficpolicy


type SesMailManagerTrafficPolicyPolicyStatementsConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ses_mail_manager_traffic_policy#boolean_expression SesMailManagerTrafficPolicy#boolean_expression}.
	BooleanExpression *SesMailManagerTrafficPolicyPolicyStatementsConditionsBooleanExpression `field:"optional" json:"booleanExpression" yaml:"booleanExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ses_mail_manager_traffic_policy#ip_expression SesMailManagerTrafficPolicy#ip_expression}.
	IpExpression *SesMailManagerTrafficPolicyPolicyStatementsConditionsIpExpression `field:"optional" json:"ipExpression" yaml:"ipExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ses_mail_manager_traffic_policy#ipv_6_expression SesMailManagerTrafficPolicy#ipv_6_expression}.
	Ipv6Expression *SesMailManagerTrafficPolicyPolicyStatementsConditionsIpv6Expression `field:"optional" json:"ipv6Expression" yaml:"ipv6Expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ses_mail_manager_traffic_policy#string_expression SesMailManagerTrafficPolicy#string_expression}.
	StringExpression *SesMailManagerTrafficPolicyPolicyStatementsConditionsStringExpression `field:"optional" json:"stringExpression" yaml:"stringExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ses_mail_manager_traffic_policy#tls_expression SesMailManagerTrafficPolicy#tls_expression}.
	TlsExpression *SesMailManagerTrafficPolicyPolicyStatementsConditionsTlsExpression `field:"optional" json:"tlsExpression" yaml:"tlsExpression"`
}

