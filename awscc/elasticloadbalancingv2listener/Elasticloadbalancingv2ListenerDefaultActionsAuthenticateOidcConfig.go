// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listener


type Elasticloadbalancingv2ListenerDefaultActionsAuthenticateOidcConfig struct {
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_listener#authentication_request_extra_params Elasticloadbalancingv2Listener#authentication_request_extra_params}
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The authorization endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_listener#authorization_endpoint Elasticloadbalancingv2Listener#authorization_endpoint}
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The OAuth 2.0 client identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_listener#client_id Elasticloadbalancingv2Listener#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The OAuth 2.0 client secret. This parameter is required if you are creating a rule. If you are modifying a rule, you can omit this parameter if you set ``UseExistingClientSecret`` to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_listener#client_secret Elasticloadbalancingv2Listener#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The OIDC issuer identifier of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_listener#issuer Elasticloadbalancingv2Listener#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// The behavior if the user is not authenticated.
	//
	// The following are possible values:
	//   +  deny```` - Return an HTTP 401 Unauthorized error.
	//   +  allow```` - Allow the request to be forwarded to the target.
	//   +  authenticate```` - Redirect the request to the IdP authorization endpoint. This is the default value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_listener#on_unauthenticated_request Elasticloadbalancingv2Listener#on_unauthenticated_request}
	OnUnauthenticatedRequest *string `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP.
	//
	// The default is ``openid``.
	//  To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_listener#scope Elasticloadbalancingv2Listener#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information. The default is AWSELBAuthSessionCookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_listener#session_cookie_name Elasticloadbalancingv2Listener#session_cookie_name}
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session, in seconds. The default is 604800 seconds (7 days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_listener#session_timeout Elasticloadbalancingv2Listener#session_timeout}
	SessionTimeout *string `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// The token endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_listener#token_endpoint Elasticloadbalancingv2Listener#token_endpoint}
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// Indicates whether to use the existing client secret when modifying a rule.
	//
	// If you are creating a rule, you can omit this parameter or set it to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_listener#use_existing_client_secret Elasticloadbalancingv2Listener#use_existing_client_secret}
	UseExistingClientSecret interface{} `field:"optional" json:"useExistingClientSecret" yaml:"useExistingClientSecret"`
	// The user info endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_listener#user_info_endpoint Elasticloadbalancingv2Listener#user_info_endpoint}
	UserInfoEndpoint *string `field:"optional" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
}

