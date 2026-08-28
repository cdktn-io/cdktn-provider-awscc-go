// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listener


type Elasticloadbalancingv2ListenerDefaultActionsJwtValidationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener#additional_claims Elasticloadbalancingv2Listener#additional_claims}.
	AdditionalClaims interface{} `field:"optional" json:"additionalClaims" yaml:"additionalClaims"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener#issuer Elasticloadbalancingv2Listener#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener#jwks_endpoint Elasticloadbalancingv2Listener#jwks_endpoint}.
	JwksEndpoint *string `field:"optional" json:"jwksEndpoint" yaml:"jwksEndpoint"`
}

