// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppsyncDataSourceConfig struct {
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
	// Unique AWS AppSync GraphQL API identifier where this data source will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#api_id AppsyncDataSource#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Friendly name for you to identify your AppSync data source after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#name AppsyncDataSource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#type AppsyncDataSource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#description AppsyncDataSource#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// AWS Region and TableName for an Amazon DynamoDB table in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#dynamo_db_config AppsyncDataSource#dynamo_db_config}
	DynamoDbConfig *AppsyncDataSourceDynamoDbConfig `field:"optional" json:"dynamoDbConfig" yaml:"dynamoDbConfig"`
	// AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
	//
	// As of September 2021, Amazon Elasticsearch Service is Amazon OpenSearch Service. This property is deprecated. For new data sources, use OpenSearchServiceConfig to specify an OpenSearch Service data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#elasticsearch_config AppsyncDataSource#elasticsearch_config}
	ElasticsearchConfig *AppsyncDataSourceElasticsearchConfig `field:"optional" json:"elasticsearchConfig" yaml:"elasticsearchConfig"`
	// ARN for the EventBridge bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#event_bridge_config AppsyncDataSource#event_bridge_config}
	EventBridgeConfig *AppsyncDataSourceEventBridgeConfig `field:"optional" json:"eventBridgeConfig" yaml:"eventBridgeConfig"`
	// Endpoints for an HTTP data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#http_config AppsyncDataSource#http_config}
	HttpConfig *AppsyncDataSourceHttpConfig `field:"optional" json:"httpConfig" yaml:"httpConfig"`
	// An ARN of a Lambda function in valid ARN format.
	//
	// This can be the ARN of a Lambda function that exists in the current account or in another account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#lambda_config AppsyncDataSource#lambda_config}
	LambdaConfig *AppsyncDataSourceLambdaConfig `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#metrics_config AppsyncDataSource#metrics_config}.
	MetricsConfig *string `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#open_search_service_config AppsyncDataSource#open_search_service_config}
	OpenSearchServiceConfig *AppsyncDataSourceOpenSearchServiceConfig `field:"optional" json:"openSearchServiceConfig" yaml:"openSearchServiceConfig"`
	// Relational Database configuration of the relational database data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#relational_database_config AppsyncDataSource#relational_database_config}
	RelationalDatabaseConfig *AppsyncDataSourceRelationalDatabaseConfig `field:"optional" json:"relationalDatabaseConfig" yaml:"relationalDatabaseConfig"`
	// The AWS Identity and Access Management service role ARN for the data source.
	//
	// The system assumes this role when accessing the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_data_source#service_role_arn AppsyncDataSource#service_role_arn}
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
}

