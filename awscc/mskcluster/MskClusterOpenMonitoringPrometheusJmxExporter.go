// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterOpenMonitoringPrometheusJmxExporter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_cluster#enabled_in_broker MskCluster#enabled_in_broker}.
	EnabledInBroker interface{} `field:"optional" json:"enabledInBroker" yaml:"enabledInBroker"`
}

