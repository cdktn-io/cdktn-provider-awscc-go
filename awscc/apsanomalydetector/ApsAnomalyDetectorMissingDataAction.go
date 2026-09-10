// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsanomalydetector


type ApsAnomalyDetectorMissingDataAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/aps_anomaly_detector#mark_as_anomaly ApsAnomalyDetector#mark_as_anomaly}.
	MarkAsAnomaly interface{} `field:"optional" json:"markAsAnomaly" yaml:"markAsAnomaly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/aps_anomaly_detector#skip ApsAnomalyDetector#skip}.
	Skip interface{} `field:"optional" json:"skip" yaml:"skip"`
}

