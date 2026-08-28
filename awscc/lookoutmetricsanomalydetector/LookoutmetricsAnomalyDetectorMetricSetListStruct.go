// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lookoutmetricsanomalydetector


type LookoutmetricsAnomalyDetectorMetricSetListStruct struct {
	// Metrics captured by this MetricSet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lookoutmetrics_anomaly_detector#metric_list LookoutmetricsAnomalyDetector#metric_list}
	MetricList interface{} `field:"required" json:"metricList" yaml:"metricList"`
	// The name of the MetricSet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lookoutmetrics_anomaly_detector#metric_set_name LookoutmetricsAnomalyDetector#metric_set_name}
	MetricSetName *string `field:"required" json:"metricSetName" yaml:"metricSetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lookoutmetrics_anomaly_detector#metric_source LookoutmetricsAnomalyDetector#metric_source}.
	MetricSource *LookoutmetricsAnomalyDetectorMetricSetListMetricSource `field:"required" json:"metricSource" yaml:"metricSource"`
	// Dimensions for this MetricSet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lookoutmetrics_anomaly_detector#dimension_list LookoutmetricsAnomalyDetector#dimension_list}
	DimensionList *[]*string `field:"optional" json:"dimensionList" yaml:"dimensionList"`
	// A description for the MetricSet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lookoutmetrics_anomaly_detector#metric_set_description LookoutmetricsAnomalyDetector#metric_set_description}
	MetricSetDescription *string `field:"optional" json:"metricSetDescription" yaml:"metricSetDescription"`
	// A frequency period to aggregate the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lookoutmetrics_anomaly_detector#metric_set_frequency LookoutmetricsAnomalyDetector#metric_set_frequency}
	MetricSetFrequency *string `field:"optional" json:"metricSetFrequency" yaml:"metricSetFrequency"`
	// Offset, in seconds, between the frequency interval and the time at which the metrics are available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lookoutmetrics_anomaly_detector#offset LookoutmetricsAnomalyDetector#offset}
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lookoutmetrics_anomaly_detector#timestamp_column LookoutmetricsAnomalyDetector#timestamp_column}.
	TimestampColumn *LookoutmetricsAnomalyDetectorMetricSetListTimestampColumn `field:"optional" json:"timestampColumn" yaml:"timestampColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lookoutmetrics_anomaly_detector#timezone LookoutmetricsAnomalyDetector#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

