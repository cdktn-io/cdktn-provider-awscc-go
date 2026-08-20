// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listenerrule


type Elasticloadbalancingv2ListenerRuleActionsJwtValidationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticloadbalancingv2_listener_rule#additional_claims Elasticloadbalancingv2ListenerRule#additional_claims}.
	AdditionalClaims interface{} `field:"optional" json:"additionalClaims" yaml:"additionalClaims"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticloadbalancingv2_listener_rule#issuer Elasticloadbalancingv2ListenerRule#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticloadbalancingv2_listener_rule#jwks_endpoint Elasticloadbalancingv2ListenerRule#jwks_endpoint}.
	JwksEndpoint *string `field:"optional" json:"jwksEndpoint" yaml:"jwksEndpoint"`
}

