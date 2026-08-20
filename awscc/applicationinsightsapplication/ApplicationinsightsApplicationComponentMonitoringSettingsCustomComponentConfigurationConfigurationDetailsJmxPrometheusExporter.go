// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporter struct {
	// Java agent host port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationinsights_application#host_port ApplicationinsightsApplication#host_port}
	HostPort *string `field:"optional" json:"hostPort" yaml:"hostPort"`
	// JMX service URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationinsights_application#jmxurl ApplicationinsightsApplication#jmxurl}
	Jmxurl *string `field:"optional" json:"jmxurl" yaml:"jmxurl"`
	// Prometheus exporter port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationinsights_application#prometheus_port ApplicationinsightsApplication#prometheus_port}
	PrometheusPort *string `field:"optional" json:"prometheusPort" yaml:"prometheusPort"`
}

