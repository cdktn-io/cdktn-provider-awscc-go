// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsanomalydetector


type ApsAnomalyDetectorConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_anomaly_detector#random_cut_forest ApsAnomalyDetector#random_cut_forest}.
	RandomCutForest *ApsAnomalyDetectorConfigurationRandomCutForest `field:"required" json:"randomCutForest" yaml:"randomCutForest"`
}

