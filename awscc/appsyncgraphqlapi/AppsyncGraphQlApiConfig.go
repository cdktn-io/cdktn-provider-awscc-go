// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncgraphqlapi

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppsyncGraphQlApiConfig struct {
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
	// Security configuration for your GraphQL API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#authentication_type AppsyncGraphQlApi#authentication_type}
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The API name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#name AppsyncGraphQlApi#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of additional authentication providers for the GraphqlApi API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#additional_authentication_providers AppsyncGraphQlApi#additional_authentication_providers}
	AdditionalAuthenticationProviders interface{} `field:"optional" json:"additionalAuthenticationProviders" yaml:"additionalAuthenticationProviders"`
	// The value that indicates whether the GraphQL API is a standard API (GRAPHQL) or merged API (MERGED).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#api_type AppsyncGraphQlApi#api_type}
	ApiType *string `field:"optional" json:"apiType" yaml:"apiType"`
	// Enables and controls the enhanced metrics feature.
	//
	// Enhanced metrics emit granular data on API usage and performance such as AppSync request and error counts, latency, and cache hits/misses. All enhanced metric data is sent to your CloudWatch account, and you can configure the types of data that will be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#enhanced_metrics_config AppsyncGraphQlApi#enhanced_metrics_config}
	EnhancedMetricsConfig *AppsyncGraphQlApiEnhancedMetricsConfig `field:"optional" json:"enhancedMetricsConfig" yaml:"enhancedMetricsConfig"`
	// A map containing the list of resources with their properties and environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#environment_variables AppsyncGraphQlApi#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Sets the value of the GraphQL API to enable (ENABLED) or disable (DISABLED) introspection.
	//
	// If no value is provided, the introspection configuration will be set to ENABLED by default. This field will produce an error if the operation attempts to use the introspection feature while this field is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#introspection_config AppsyncGraphQlApi#introspection_config}
	IntrospectionConfig *string `field:"optional" json:"introspectionConfig" yaml:"introspectionConfig"`
	// A LambdaAuthorizerConfig holds configuration on how to authorize AWS AppSync API access when using the AWS_LAMBDA authorizer mode.
	//
	// Be aware that an AWS AppSync API may have only one Lambda authorizer configured at a time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#lambda_authorizer_config AppsyncGraphQlApi#lambda_authorizer_config}
	LambdaAuthorizerConfig *AppsyncGraphQlApiLambdaAuthorizerConfig `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// The Amazon CloudWatch Logs configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#log_config AppsyncGraphQlApi#log_config}
	LogConfig *AppsyncGraphQlApiLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// The AWS Identity and Access Management service role ARN for a merged API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#merged_api_execution_role_arn AppsyncGraphQlApi#merged_api_execution_role_arn}
	MergedApiExecutionRoleArn *string `field:"optional" json:"mergedApiExecutionRoleArn" yaml:"mergedApiExecutionRoleArn"`
	// The OpenID Connect configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#open_id_connect_config AppsyncGraphQlApi#open_id_connect_config}
	OpenIdConnectConfig *AppsyncGraphQlApiOpenIdConnectConfig `field:"optional" json:"openIdConnectConfig" yaml:"openIdConnectConfig"`
	// The owner contact information for an API resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#owner_contact AppsyncGraphQlApi#owner_contact}
	OwnerContact *string `field:"optional" json:"ownerContact" yaml:"ownerContact"`
	// The maximum depth a query can have in a single request.
	//
	// Depth refers to the amount of nested levels allowed in the body of query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#query_depth_limit AppsyncGraphQlApi#query_depth_limit}
	QueryDepthLimit *float64 `field:"optional" json:"queryDepthLimit" yaml:"queryDepthLimit"`
	// The maximum number of resolvers that can be invoked in a single request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#resolver_count_limit AppsyncGraphQlApi#resolver_count_limit}
	ResolverCountLimit *float64 `field:"optional" json:"resolverCountLimit" yaml:"resolverCountLimit"`
	// An arbitrary set of tags (key-value pairs) for this GraphQL API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#tags AppsyncGraphQlApi#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Optional authorization configuration for using Amazon Cognito user pools with your GraphQL endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#user_pool_config AppsyncGraphQlApi#user_pool_config}
	UserPoolConfig *AppsyncGraphQlApiUserPoolConfig `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
	// Sets the scope of the GraphQL API to public (GLOBAL) or private (PRIVATE).
	//
	// By default, the scope is set to Global if no value is provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#visibility AppsyncGraphQlApi#visibility}
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// A flag indicating whether to use AWS X-Ray tracing for this GraphqlApi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_graph_ql_api#xray_enabled AppsyncGraphQlApi#xray_enabled}
	XrayEnabled interface{} `field:"optional" json:"xrayEnabled" yaml:"xrayEnabled"`
}

