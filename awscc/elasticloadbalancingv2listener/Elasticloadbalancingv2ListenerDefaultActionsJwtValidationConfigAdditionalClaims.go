// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listener


type Elasticloadbalancingv2ListenerDefaultActionsJwtValidationConfigAdditionalClaims struct {
	// The format of the claim value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticloadbalancingv2_listener#format Elasticloadbalancingv2Listener#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The name of the claim. You can't specify ``exp``, ``iss``, ``nbf``, or ``iat`` because we validate them by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticloadbalancingv2_listener#name Elasticloadbalancingv2Listener#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The claim value.
	//
	// The maximum size of the list is 10. Each value can be up to 256 characters in length. If the format is ``space-separated-values``, the values can't include spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticloadbalancingv2_listener#values Elasticloadbalancingv2Listener#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

