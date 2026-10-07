// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmonitormonitor

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmonitorMonitorConfig struct {
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
	// The name identifying the monitor.
	//
	// It can contain only letters, underscores (_), or dashes (-), and can be up to 200 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/networkmonitor_monitor#monitor_name NetworkmonitorMonitor#monitor_name}
	MonitorName *string `field:"required" json:"monitorName" yaml:"monitorName"`
	// The time, in seconds, that metrics are aggregated and sent to Amazon CloudWatch.
	//
	// Valid values are either 30 or 60.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/networkmonitor_monitor#aggregation_period NetworkmonitorMonitor#aggregation_period}
	AggregationPeriod *float64 `field:"optional" json:"aggregationPeriod" yaml:"aggregationPeriod"`
	// The tags for the monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/networkmonitor_monitor#tags NetworkmonitorMonitor#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

