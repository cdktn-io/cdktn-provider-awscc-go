// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package internetmonitormonitor


type InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/internetmonitor_monitor#health_score_threshold InternetmonitorMonitor#health_score_threshold}.
	HealthScoreThreshold *float64 `field:"optional" json:"healthScoreThreshold" yaml:"healthScoreThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/internetmonitor_monitor#min_traffic_impact InternetmonitorMonitor#min_traffic_impact}.
	MinTrafficImpact *float64 `field:"optional" json:"minTrafficImpact" yaml:"minTrafficImpact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/internetmonitor_monitor#status InternetmonitorMonitor#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

