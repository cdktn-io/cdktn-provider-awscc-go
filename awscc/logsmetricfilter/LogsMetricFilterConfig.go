// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsmetricfilter

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogsMetricFilterConfig struct {
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
	// A filter pattern for extracting metric data out of ingested log events.
	//
	// For more information, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_metric_filter#filter_pattern LogsMetricFilter#filter_pattern}
	FilterPattern *string `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The name of an existing log group that you want to associate with this metric filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_metric_filter#log_group_name LogsMetricFilter#log_group_name}
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The metric transformations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_metric_filter#metric_transformations LogsMetricFilter#metric_transformations}
	MetricTransformations interface{} `field:"required" json:"metricTransformations" yaml:"metricTransformations"`
	// This parameter is valid only for log groups that have an active log transformer.
	//
	// For more information about log transformers, see [PutTransformer](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html).
	//  If this value is ``true``, the metric filter is applied on the transformed version of the log events instead of the original ingested log events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_metric_filter#apply_on_transformed_logs LogsMetricFilter#apply_on_transformed_logs}
	ApplyOnTransformedLogs interface{} `field:"optional" json:"applyOnTransformedLogs" yaml:"applyOnTransformedLogs"`
	// The list of system fields that are emitted as additional dimensions in the generated metrics.
	//
	// Returns the ``emitSystemFieldDimensions`` value if it was specified when the metric filter was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_metric_filter#emit_system_field_dimensions LogsMetricFilter#emit_system_field_dimensions}
	EmitSystemFieldDimensions *[]*string `field:"optional" json:"emitSystemFieldDimensions" yaml:"emitSystemFieldDimensions"`
	// The filter expression that specifies which log events are processed by this metric filter based on system fields.
	//
	// Returns the ``fieldSelectionCriteria`` value if it was specified when the metric filter was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_metric_filter#field_selection_criteria LogsMetricFilter#field_selection_criteria}
	FieldSelectionCriteria *string `field:"optional" json:"fieldSelectionCriteria" yaml:"fieldSelectionCriteria"`
	// The name of the metric filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_metric_filter#filter_name LogsMetricFilter#filter_name}
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
}

