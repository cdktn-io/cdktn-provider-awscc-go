// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetmetric

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotFleetMetricConfig struct {
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
	// The name of the fleet metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_fleet_metric#metric_name IotFleetMetric#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The aggregation field to perform aggregation and metric emission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_fleet_metric#aggregation_field IotFleetMetric#aggregation_field}
	AggregationField *string `field:"optional" json:"aggregationField" yaml:"aggregationField"`
	// Aggregation types supported by Fleet Indexing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_fleet_metric#aggregation_type IotFleetMetric#aggregation_type}
	AggregationType *IotFleetMetricAggregationType `field:"optional" json:"aggregationType" yaml:"aggregationType"`
	// The description of a fleet metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_fleet_metric#description IotFleetMetric#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The index name of a fleet metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_fleet_metric#index_name IotFleetMetric#index_name}
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// The period of metric emission in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_fleet_metric#period IotFleetMetric#period}
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The Fleet Indexing query used by a fleet metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_fleet_metric#query_string IotFleetMetric#query_string}
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// The version of a Fleet Indexing query used by a fleet metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_fleet_metric#query_version IotFleetMetric#query_version}
	QueryVersion *string `field:"optional" json:"queryVersion" yaml:"queryVersion"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_fleet_metric#tags IotFleetMetric#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The unit of data points emitted by a fleet metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_fleet_metric#unit IotFleetMetric#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

