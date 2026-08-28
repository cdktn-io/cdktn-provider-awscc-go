// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayauthorizer

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApigatewayAuthorizerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_authorizer#name ApigatewayAuthorizer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_authorizer#rest_api_id ApigatewayAuthorizer#rest_api_id}.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_authorizer#type ApigatewayAuthorizer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_authorizer#authorizer_credentials ApigatewayAuthorizer#authorizer_credentials}.
	AuthorizerCredentials *string `field:"optional" json:"authorizerCredentials" yaml:"authorizerCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_authorizer#authorizer_result_ttl_in_seconds ApigatewayAuthorizer#authorizer_result_ttl_in_seconds}.
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_authorizer#authorizer_uri ApigatewayAuthorizer#authorizer_uri}.
	AuthorizerUri *string `field:"optional" json:"authorizerUri" yaml:"authorizerUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_authorizer#auth_type ApigatewayAuthorizer#auth_type}.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_authorizer#identity_source ApigatewayAuthorizer#identity_source}.
	IdentitySource *string `field:"optional" json:"identitySource" yaml:"identitySource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_authorizer#identity_validation_expression ApigatewayAuthorizer#identity_validation_expression}.
	IdentityValidationExpression *string `field:"optional" json:"identityValidationExpression" yaml:"identityValidationExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_authorizer#provider_ar_ns ApigatewayAuthorizer#provider_ar_ns}.
	ProviderArNs *[]*string `field:"optional" json:"providerArNs" yaml:"providerArNs"`
}

