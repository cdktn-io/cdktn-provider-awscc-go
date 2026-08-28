// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectmetric

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectMetricConfig struct {
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
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_metric#instance_arn ConnectMetric#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The description of the custom metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_metric#description ConnectMetric#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The calculation configuration for the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_metric#metric_calculation ConnectMetric#metric_calculation}
	MetricCalculation *ConnectMetricMetricCalculation `field:"optional" json:"metricCalculation" yaml:"metricCalculation"`
	// The name of the custom metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_metric#name ConnectMetric#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates how to classify a positive trend in metric data on the UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_metric#positive_trend_indicator ConnectMetric#positive_trend_indicator}
	PositiveTrendIndicator *string `field:"optional" json:"positiveTrendIndicator" yaml:"positiveTrendIndicator"`
	// The status of the custom metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_metric#status ConnectMetric#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_metric#tags ConnectMetric#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Display unit for the metric data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_metric#unit ConnectMetric#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

