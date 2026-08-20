// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceMonitoringMetricConfigurations struct {
	// The list of metric names to configure. The supported metric names are ``CPUUtilization`` and ``MemoryUtilization``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_service#metric_names EcsService#metric_names}
	MetricNames *[]*string `field:"optional" json:"metricNames" yaml:"metricNames"`
	// The resolution, in seconds, at which to collect the metrics. The valid values are ``20`` and ``60``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_service#resolution_seconds EcsService#resolution_seconds}
	ResolutionSeconds *float64 `field:"optional" json:"resolutionSeconds" yaml:"resolutionSeconds"`
}

