// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncgraphqlapi


type AppsyncGraphQlApiEnhancedMetricsConfig struct {
	// Controls how data source metrics will be emitted to CloudWatch. Data source metrics include:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appsync_graph_ql_api#data_source_level_metrics_behavior AppsyncGraphQlApi#data_source_level_metrics_behavior}
	DataSourceLevelMetricsBehavior *string `field:"optional" json:"dataSourceLevelMetricsBehavior" yaml:"dataSourceLevelMetricsBehavior"`
	// Controls how operation metrics will be emitted to CloudWatch. Operation metrics include:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appsync_graph_ql_api#operation_level_metrics_config AppsyncGraphQlApi#operation_level_metrics_config}
	OperationLevelMetricsConfig *string `field:"optional" json:"operationLevelMetricsConfig" yaml:"operationLevelMetricsConfig"`
	// Controls how resolver metrics will be emitted to CloudWatch. Resolver metrics include:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appsync_graph_ql_api#resolver_level_metrics_behavior AppsyncGraphQlApi#resolver_level_metrics_behavior}
	ResolverLevelMetricsBehavior *string `field:"optional" json:"resolverLevelMetricsBehavior" yaml:"resolverLevelMetricsBehavior"`
}

