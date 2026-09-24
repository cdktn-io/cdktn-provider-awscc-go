// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsanomalydetector


type ApsAnomalyDetectorConfigurationRandomCutForest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_anomaly_detector#query ApsAnomalyDetector#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_anomaly_detector#ignore_near_expected_from_above ApsAnomalyDetector#ignore_near_expected_from_above}.
	IgnoreNearExpectedFromAbove *ApsAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAbove `field:"optional" json:"ignoreNearExpectedFromAbove" yaml:"ignoreNearExpectedFromAbove"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_anomaly_detector#ignore_near_expected_from_below ApsAnomalyDetector#ignore_near_expected_from_below}.
	IgnoreNearExpectedFromBelow *ApsAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelow `field:"optional" json:"ignoreNearExpectedFromBelow" yaml:"ignoreNearExpectedFromBelow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_anomaly_detector#sample_size ApsAnomalyDetector#sample_size}.
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_anomaly_detector#shingle_size ApsAnomalyDetector#shingle_size}.
	ShingleSize *float64 `field:"optional" json:"shingleSize" yaml:"shingleSize"`
}

