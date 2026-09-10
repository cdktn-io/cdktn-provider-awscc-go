// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsanomalydetector


type ApsAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAbove struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/aps_anomaly_detector#amount ApsAnomalyDetector#amount}.
	Amount *float64 `field:"optional" json:"amount" yaml:"amount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/aps_anomaly_detector#ratio ApsAnomalyDetector#ratio}.
	Ratio *float64 `field:"optional" json:"ratio" yaml:"ratio"`
}

