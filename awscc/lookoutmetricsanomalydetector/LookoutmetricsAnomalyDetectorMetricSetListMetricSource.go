// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lookoutmetricsanomalydetector


type LookoutmetricsAnomalyDetectorMetricSetListMetricSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lookoutmetrics_anomaly_detector#app_flow_config LookoutmetricsAnomalyDetector#app_flow_config}.
	AppFlowConfig *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceAppFlowConfig `field:"optional" json:"appFlowConfig" yaml:"appFlowConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lookoutmetrics_anomaly_detector#cloudwatch_config LookoutmetricsAnomalyDetector#cloudwatch_config}.
	CloudwatchConfig *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceCloudwatchConfig `field:"optional" json:"cloudwatchConfig" yaml:"cloudwatchConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lookoutmetrics_anomaly_detector#rds_source_config LookoutmetricsAnomalyDetector#rds_source_config}.
	RdsSourceConfig *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfig `field:"optional" json:"rdsSourceConfig" yaml:"rdsSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lookoutmetrics_anomaly_detector#redshift_source_config LookoutmetricsAnomalyDetector#redshift_source_config}.
	RedshiftSourceConfig *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRedshiftSourceConfig `field:"optional" json:"redshiftSourceConfig" yaml:"redshiftSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lookoutmetrics_anomaly_detector#s3_source_config LookoutmetricsAnomalyDetector#s3_source_config}.
	S3SourceConfig *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfig `field:"optional" json:"s3SourceConfig" yaml:"s3SourceConfig"`
}

