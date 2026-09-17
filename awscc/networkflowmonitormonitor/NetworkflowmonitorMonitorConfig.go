// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkflowmonitormonitor

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkflowmonitorMonitorConfig struct {
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
	// The local resources to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkflowmonitor_monitor#local_resources NetworkflowmonitorMonitor#local_resources}
	LocalResources interface{} `field:"required" json:"localResources" yaml:"localResources"`
	// The name of the monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkflowmonitor_monitor#monitor_name NetworkflowmonitorMonitor#monitor_name}
	MonitorName *string `field:"required" json:"monitorName" yaml:"monitorName"`
	// The remote resources to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkflowmonitor_monitor#remote_resources NetworkflowmonitorMonitor#remote_resources}
	RemoteResources interface{} `field:"optional" json:"remoteResources" yaml:"remoteResources"`
	// The Amazon Resource Name (ARN) of the scope for the monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkflowmonitor_monitor#scope_arn NetworkflowmonitorMonitor#scope_arn}
	ScopeArn *string `field:"optional" json:"scopeArn" yaml:"scopeArn"`
	// The tags for the monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkflowmonitor_monitor#tags NetworkflowmonitorMonitor#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

