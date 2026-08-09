// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncresolver

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppsyncResolverConfig struct {
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
	// The APSYlong GraphQL API to which you want to attach this resolver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#api_id AppsyncResolver#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The GraphQL field on a type that invokes the resolver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#field_name AppsyncResolver#field_name}
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// The GraphQL type that invokes this resolver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#type_name AppsyncResolver#type_name}
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// The caching configuration for the resolver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#caching_config AppsyncResolver#caching_config}
	CachingConfig *AppsyncResolverCachingConfig `field:"optional" json:"cachingConfig" yaml:"cachingConfig"`
	// The ``resolver`` code that contains the request and response functions.
	//
	// When code is used, the ``runtime`` is required. The runtime value must be ``APPSYNC_JS``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#code AppsyncResolver#code}
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The Amazon S3 endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#code_s3_location AppsyncResolver#code_s3_location}
	CodeS3Location *string `field:"optional" json:"codeS3Location" yaml:"codeS3Location"`
	// The resolver data source name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#data_source_name AppsyncResolver#data_source_name}
	DataSourceName *string `field:"optional" json:"dataSourceName" yaml:"dataSourceName"`
	// The resolver type.
	//
	// +  *UNIT*: A UNIT resolver type. A UNIT resolver is the default resolver type. You can use a UNIT resolver to run a GraphQL query against a single data source.
	//   +  *PIPELINE*: A PIPELINE resolver type. You can use a PIPELINE resolver to invoke a series of ``Function`` objects in a serial manner. You can use a pipeline resolver to run a GraphQL query against multiple data sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#kind AppsyncResolver#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// The maximum number of resolver request inputs that will be sent to a single LAMlong function in a ``BatchInvoke`` operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#max_batch_size AppsyncResolver#max_batch_size}
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// Enables or disables enhanced resolver metrics for specified resolvers.
	//
	// Note that ``MetricsConfig`` won't be used unless the ``resolverLevelMetricsBehavior`` value is set to ``PER_RESOLVER_METRICS``. If the ``resolverLevelMetricsBehavior`` is set to ``FULL_REQUEST_RESOLVER_METRICS`` instead, ``MetricsConfig`` will be ignored. However, you can still set its value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#metrics_config AppsyncResolver#metrics_config}
	MetricsConfig *string `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// Functions linked with the pipeline resolver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#pipeline_config AppsyncResolver#pipeline_config}
	PipelineConfig *AppsyncResolverPipelineConfig `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// The request mapping template.
	//
	// Request mapping templates are optional when using a Lambda data source. For all other data sources, a request mapping template is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#request_mapping_template AppsyncResolver#request_mapping_template}
	RequestMappingTemplate *string `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// The location of a request mapping template in an S3 bucket.
	//
	// Use this if you want to provision with a template file in S3 rather than embedding it in your CFNshort template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#request_mapping_template_s3_location AppsyncResolver#request_mapping_template_s3_location}
	RequestMappingTemplateS3Location *string `field:"optional" json:"requestMappingTemplateS3Location" yaml:"requestMappingTemplateS3Location"`
	// The response mapping template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#response_mapping_template AppsyncResolver#response_mapping_template}
	ResponseMappingTemplate *string `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The location of a response mapping template in an S3 bucket.
	//
	// Use this if you want to provision with a template file in S3 rather than embedding it in your CFNshort template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#response_mapping_template_s3_location AppsyncResolver#response_mapping_template_s3_location}
	ResponseMappingTemplateS3Location *string `field:"optional" json:"responseMappingTemplateS3Location" yaml:"responseMappingTemplateS3Location"`
	// Describes a runtime used by an APSYlong resolver or APSYlong function.
	//
	// Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#runtime AppsyncResolver#runtime}
	Runtime *AppsyncResolverRuntime `field:"optional" json:"runtime" yaml:"runtime"`
	// The ``SyncConfig`` for a resolver attached to a versioned data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_resolver#sync_config AppsyncResolver#sync_config}
	SyncConfig *AppsyncResolverSyncConfig `field:"optional" json:"syncConfig" yaml:"syncConfig"`
}

