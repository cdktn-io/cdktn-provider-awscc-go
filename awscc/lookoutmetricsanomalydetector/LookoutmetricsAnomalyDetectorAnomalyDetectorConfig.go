// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lookoutmetricsanomalydetector


type LookoutmetricsAnomalyDetectorAnomalyDetectorConfig struct {
	// Frequency of anomaly detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lookoutmetrics_anomaly_detector#anomaly_detector_frequency LookoutmetricsAnomalyDetector#anomaly_detector_frequency}
	AnomalyDetectorFrequency *string `field:"required" json:"anomalyDetectorFrequency" yaml:"anomalyDetectorFrequency"`
}

