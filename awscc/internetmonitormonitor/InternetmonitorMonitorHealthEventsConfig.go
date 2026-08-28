// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package internetmonitormonitor


type InternetmonitorMonitorHealthEventsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/internetmonitor_monitor#availability_local_health_events_config InternetmonitorMonitor#availability_local_health_events_config}.
	AvailabilityLocalHealthEventsConfig *InternetmonitorMonitorHealthEventsConfigAvailabilityLocalHealthEventsConfig `field:"optional" json:"availabilityLocalHealthEventsConfig" yaml:"availabilityLocalHealthEventsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/internetmonitor_monitor#availability_score_threshold InternetmonitorMonitor#availability_score_threshold}.
	AvailabilityScoreThreshold *float64 `field:"optional" json:"availabilityScoreThreshold" yaml:"availabilityScoreThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/internetmonitor_monitor#performance_local_health_events_config InternetmonitorMonitor#performance_local_health_events_config}.
	PerformanceLocalHealthEventsConfig *InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfig `field:"optional" json:"performanceLocalHealthEventsConfig" yaml:"performanceLocalHealthEventsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/internetmonitor_monitor#performance_score_threshold InternetmonitorMonitor#performance_score_threshold}.
	PerformanceScoreThreshold *float64 `field:"optional" json:"performanceScoreThreshold" yaml:"performanceScoreThreshold"`
}

