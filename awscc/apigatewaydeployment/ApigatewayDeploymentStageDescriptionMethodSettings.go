// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewaydeployment


type ApigatewayDeploymentStageDescriptionMethodSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#cache_data_encrypted ApigatewayDeployment#cache_data_encrypted}.
	CacheDataEncrypted interface{} `field:"optional" json:"cacheDataEncrypted" yaml:"cacheDataEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#cache_ttl_in_seconds ApigatewayDeployment#cache_ttl_in_seconds}.
	CacheTtlInSeconds *float64 `field:"optional" json:"cacheTtlInSeconds" yaml:"cacheTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#caching_enabled ApigatewayDeployment#caching_enabled}.
	CachingEnabled interface{} `field:"optional" json:"cachingEnabled" yaml:"cachingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#data_trace_enabled ApigatewayDeployment#data_trace_enabled}.
	DataTraceEnabled interface{} `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// The HTTP method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#http_method ApigatewayDeployment#http_method}
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#logging_level ApigatewayDeployment#logging_level}.
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#metrics_enabled ApigatewayDeployment#metrics_enabled}.
	MetricsEnabled interface{} `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// The resource path for this method.
	//
	// Forward slashes (``/``) are encoded as ``~1`` and the initial slash must include a forward slash. For example, the path value ``/resource/subresource`` must be encoded as ``/~1resource~1subresource``. To specify the root path, use only a slash (``/``).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#resource_path ApigatewayDeployment#resource_path}
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#throttling_burst_limit ApigatewayDeployment#throttling_burst_limit}.
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#throttling_rate_limit ApigatewayDeployment#throttling_rate_limit}.
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

