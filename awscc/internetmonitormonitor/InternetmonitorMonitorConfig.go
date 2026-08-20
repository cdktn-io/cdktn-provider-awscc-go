// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package internetmonitormonitor

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type InternetmonitorMonitorConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/internetmonitor_monitor#monitor_name InternetmonitorMonitor#monitor_name}.
	MonitorName *string `field:"required" json:"monitorName" yaml:"monitorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/internetmonitor_monitor#health_events_config InternetmonitorMonitor#health_events_config}.
	HealthEventsConfig *InternetmonitorMonitorHealthEventsConfig `field:"optional" json:"healthEventsConfig" yaml:"healthEventsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/internetmonitor_monitor#include_linked_accounts InternetmonitorMonitor#include_linked_accounts}.
	IncludeLinkedAccounts interface{} `field:"optional" json:"includeLinkedAccounts" yaml:"includeLinkedAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/internetmonitor_monitor#internet_measurements_log_delivery InternetmonitorMonitor#internet_measurements_log_delivery}.
	InternetMeasurementsLogDelivery *InternetmonitorMonitorInternetMeasurementsLogDelivery `field:"optional" json:"internetMeasurementsLogDelivery" yaml:"internetMeasurementsLogDelivery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/internetmonitor_monitor#linked_account_id InternetmonitorMonitor#linked_account_id}.
	LinkedAccountId *string `field:"optional" json:"linkedAccountId" yaml:"linkedAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/internetmonitor_monitor#max_city_networks_to_monitor InternetmonitorMonitor#max_city_networks_to_monitor}.
	MaxCityNetworksToMonitor *float64 `field:"optional" json:"maxCityNetworksToMonitor" yaml:"maxCityNetworksToMonitor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/internetmonitor_monitor#resources InternetmonitorMonitor#resources}.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/internetmonitor_monitor#resources_to_add InternetmonitorMonitor#resources_to_add}.
	ResourcesToAdd *[]*string `field:"optional" json:"resourcesToAdd" yaml:"resourcesToAdd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/internetmonitor_monitor#resources_to_remove InternetmonitorMonitor#resources_to_remove}.
	ResourcesToRemove *[]*string `field:"optional" json:"resourcesToRemove" yaml:"resourcesToRemove"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/internetmonitor_monitor#status InternetmonitorMonitor#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/internetmonitor_monitor#tags InternetmonitorMonitor#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/internetmonitor_monitor#traffic_percentage_to_monitor InternetmonitorMonitor#traffic_percentage_to_monitor}.
	TrafficPercentageToMonitor *float64 `field:"optional" json:"trafficPercentageToMonitor" yaml:"trafficPercentageToMonitor"`
}

