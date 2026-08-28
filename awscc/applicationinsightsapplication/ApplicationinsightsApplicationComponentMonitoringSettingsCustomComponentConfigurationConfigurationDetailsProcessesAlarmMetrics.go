// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsProcessesAlarmMetrics struct {
	// The name of the metric to be monitored for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/applicationinsights_application#alarm_metric_name ApplicationinsightsApplication#alarm_metric_name}
	AlarmMetricName *string `field:"optional" json:"alarmMetricName" yaml:"alarmMetricName"`
}

