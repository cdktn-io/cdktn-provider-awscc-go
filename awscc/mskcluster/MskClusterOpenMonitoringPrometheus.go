// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterOpenMonitoringPrometheus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_cluster#jmx_exporter MskCluster#jmx_exporter}.
	JmxExporter *MskClusterOpenMonitoringPrometheusJmxExporter `field:"optional" json:"jmxExporter" yaml:"jmxExporter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_cluster#node_exporter MskCluster#node_exporter}.
	NodeExporter *MskClusterOpenMonitoringPrometheusNodeExporter `field:"optional" json:"nodeExporter" yaml:"nodeExporter"`
}

