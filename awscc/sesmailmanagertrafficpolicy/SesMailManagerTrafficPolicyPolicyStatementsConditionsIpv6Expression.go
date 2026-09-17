// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagertrafficpolicy


type SesMailManagerTrafficPolicyPolicyStatementsConditionsIpv6Expression struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_traffic_policy#evaluate SesMailManagerTrafficPolicy#evaluate}.
	Evaluate *SesMailManagerTrafficPolicyPolicyStatementsConditionsIpv6ExpressionEvaluate `field:"optional" json:"evaluate" yaml:"evaluate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_traffic_policy#operator SesMailManagerTrafficPolicy#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_traffic_policy#values SesMailManagerTrafficPolicy#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

