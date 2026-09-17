// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsNetWeaverPrometheusExporter struct {
	// SAP instance numbers for ASCS, ERS, and App Servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationinsights_application#instance_numbers ApplicationinsightsApplication#instance_numbers}
	InstanceNumbers *[]*string `field:"optional" json:"instanceNumbers" yaml:"instanceNumbers"`
	// Prometheus exporter port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationinsights_application#prometheus_port ApplicationinsightsApplication#prometheus_port}
	PrometheusPort *string `field:"optional" json:"prometheusPort" yaml:"prometheusPort"`
	// SAP NetWeaver SID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationinsights_application#sapsid ApplicationinsightsApplication#sapsid}
	Sapsid *string `field:"optional" json:"sapsid" yaml:"sapsid"`
}

