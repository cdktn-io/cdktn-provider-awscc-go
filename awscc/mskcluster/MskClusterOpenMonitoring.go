// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterOpenMonitoring struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/msk_cluster#prometheus MskCluster#prometheus}.
	Prometheus *MskClusterOpenMonitoringPrometheus `field:"optional" json:"prometheus" yaml:"prometheus"`
}

