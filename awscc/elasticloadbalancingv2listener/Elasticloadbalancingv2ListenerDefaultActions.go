// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listener


type Elasticloadbalancingv2ListenerDefaultActions struct {
	// The type of action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticloadbalancingv2_listener#type Elasticloadbalancingv2Listener#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// [HTTPS listeners] Information for using Amazon Cognito to authenticate users. Specify only when ``Type`` is ``authenticate-cognito``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticloadbalancingv2_listener#authenticate_cognito_config Elasticloadbalancingv2Listener#authenticate_cognito_config}
	AuthenticateCognitoConfig *Elasticloadbalancingv2ListenerDefaultActionsAuthenticateCognitoConfig `field:"optional" json:"authenticateCognitoConfig" yaml:"authenticateCognitoConfig"`
	// [HTTPS listeners] Information about an identity provider that is compliant with OpenID Connect (OIDC).
	//
	// Specify only when ``Type`` is ``authenticate-oidc``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticloadbalancingv2_listener#authenticate_oidc_config Elasticloadbalancingv2Listener#authenticate_oidc_config}
	AuthenticateOidcConfig *Elasticloadbalancingv2ListenerDefaultActionsAuthenticateOidcConfig `field:"optional" json:"authenticateOidcConfig" yaml:"authenticateOidcConfig"`
	// [Application Load Balancer] Information for creating an action that returns a custom HTTP response.
	//
	// Specify only when ``Type`` is ``fixed-response``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticloadbalancingv2_listener#fixed_response_config Elasticloadbalancingv2Listener#fixed_response_config}
	FixedResponseConfig *Elasticloadbalancingv2ListenerDefaultActionsFixedResponseConfig `field:"optional" json:"fixedResponseConfig" yaml:"fixedResponseConfig"`
	// Information for creating an action that distributes requests among multiple target groups.
	//
	// Specify only when ``Type`` is ``forward``.
	//  If you specify both ``ForwardConfig`` and ``TargetGroupArn``, you can specify only one target group using ``ForwardConfig`` and it must be the same target group specified in ``TargetGroupArn``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticloadbalancingv2_listener#forward_config Elasticloadbalancingv2Listener#forward_config}
	ForwardConfig *Elasticloadbalancingv2ListenerDefaultActionsForwardConfig `field:"optional" json:"forwardConfig" yaml:"forwardConfig"`
	// [HTTPS listeners] Information for validating JWT access tokens in client requests. Specify only when ``Type`` is ``jwt-validation``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticloadbalancingv2_listener#jwt_validation_config Elasticloadbalancingv2Listener#jwt_validation_config}
	JwtValidationConfig *Elasticloadbalancingv2ListenerDefaultActionsJwtValidationConfig `field:"optional" json:"jwtValidationConfig" yaml:"jwtValidationConfig"`
	// The order for the action.
	//
	// This value is required for rules with multiple actions. The action with the lowest value for order is performed first.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticloadbalancingv2_listener#order Elasticloadbalancingv2Listener#order}
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// [Application Load Balancer] Information for creating a redirect action. Specify only when ``Type`` is ``redirect``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticloadbalancingv2_listener#redirect_config Elasticloadbalancingv2Listener#redirect_config}
	RedirectConfig *Elasticloadbalancingv2ListenerDefaultActionsRedirectConfig `field:"optional" json:"redirectConfig" yaml:"redirectConfig"`
	// The Amazon Resource Name (ARN) of the target group.
	//
	// Specify only when ``Type`` is ``forward`` and you want to route to a single target group. To route to multiple target groups, you must use ``ForwardConfig`` instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticloadbalancingv2_listener#target_group_arn Elasticloadbalancingv2Listener#target_group_arn}
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

