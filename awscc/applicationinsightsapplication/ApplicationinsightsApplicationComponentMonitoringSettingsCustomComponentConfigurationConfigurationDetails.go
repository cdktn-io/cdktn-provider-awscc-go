// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetails struct {
	// A list of metrics to monitor for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/applicationinsights_application#alarm_metrics ApplicationinsightsApplication#alarm_metrics}
	AlarmMetrics interface{} `field:"optional" json:"alarmMetrics" yaml:"alarmMetrics"`
	// A list of alarms to monitor for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/applicationinsights_application#alarms ApplicationinsightsApplication#alarms}
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// The HA cluster Prometheus Exporter settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/applicationinsights_application#ha_cluster_prometheus_exporter ApplicationinsightsApplication#ha_cluster_prometheus_exporter}
	HaClusterPrometheusExporter *ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsHaClusterPrometheusExporter `field:"optional" json:"haClusterPrometheusExporter" yaml:"haClusterPrometheusExporter"`
	// The HANA DB Prometheus Exporter settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/applicationinsights_application#hana_prometheus_exporter ApplicationinsightsApplication#hana_prometheus_exporter}
	HanaPrometheusExporter *ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsHanaPrometheusExporter `field:"optional" json:"hanaPrometheusExporter" yaml:"hanaPrometheusExporter"`
	// The JMX Prometheus Exporter settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/applicationinsights_application#jmx_prometheus_exporter ApplicationinsightsApplication#jmx_prometheus_exporter}
	JmxPrometheusExporter *ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporter `field:"optional" json:"jmxPrometheusExporter" yaml:"jmxPrometheusExporter"`
	// A list of logs to monitor for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/applicationinsights_application#logs ApplicationinsightsApplication#logs}
	Logs interface{} `field:"optional" json:"logs" yaml:"logs"`
	// The NetWeaver Prometheus Exporter settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/applicationinsights_application#net_weaver_prometheus_exporter ApplicationinsightsApplication#net_weaver_prometheus_exporter}
	NetWeaverPrometheusExporter *ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsNetWeaverPrometheusExporter `field:"optional" json:"netWeaverPrometheusExporter" yaml:"netWeaverPrometheusExporter"`
	// A list of processes to monitor for the component. Only Windows EC2 instances can have a processes section.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/applicationinsights_application#processes ApplicationinsightsApplication#processes}
	Processes interface{} `field:"optional" json:"processes" yaml:"processes"`
	// The SQL Prometheus Exporter settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/applicationinsights_application#sql_server_prometheus_exporter ApplicationinsightsApplication#sql_server_prometheus_exporter}
	SqlServerPrometheusExporter *ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsSqlServerPrometheusExporter `field:"optional" json:"sqlServerPrometheusExporter" yaml:"sqlServerPrometheusExporter"`
	// A list of Windows Events to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/applicationinsights_application#windows_events ApplicationinsightsApplication#windows_events}
	WindowsEvents interface{} `field:"optional" json:"windowsEvents" yaml:"windowsEvents"`
}

