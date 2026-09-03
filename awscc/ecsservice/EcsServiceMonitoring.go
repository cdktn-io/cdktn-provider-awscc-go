// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceMonitoring struct {
	// The list of metric configurations for the service monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecs_service#metric_configurations EcsService#metric_configurations}
	MetricConfigurations interface{} `field:"optional" json:"metricConfigurations" yaml:"metricConfigurations"`
}

