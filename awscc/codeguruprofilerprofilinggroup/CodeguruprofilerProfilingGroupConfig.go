// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codeguruprofilerprofilinggroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CodeguruprofilerProfilingGroupConfig struct {
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
	// The name of the profiling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codeguruprofiler_profiling_group#profiling_group_name CodeguruprofilerProfilingGroup#profiling_group_name}
	ProfilingGroupName *string `field:"required" json:"profilingGroupName" yaml:"profilingGroupName"`
	// The agent permissions attached to this profiling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codeguruprofiler_profiling_group#agent_permissions CodeguruprofilerProfilingGroup#agent_permissions}
	AgentPermissions *CodeguruprofilerProfilingGroupAgentPermissions `field:"optional" json:"agentPermissions" yaml:"agentPermissions"`
	// Configuration for Notification Channels for Anomaly Detection feature in CodeGuru Profiler which enables customers to detect anomalies in the application profile for those methods that represent the highest proportion of CPU time or latency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codeguruprofiler_profiling_group#anomaly_detection_notification_configuration CodeguruprofilerProfilingGroup#anomaly_detection_notification_configuration}
	AnomalyDetectionNotificationConfiguration interface{} `field:"optional" json:"anomalyDetectionNotificationConfiguration" yaml:"anomalyDetectionNotificationConfiguration"`
	// The compute platform of the profiling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codeguruprofiler_profiling_group#compute_platform CodeguruprofilerProfilingGroup#compute_platform}
	ComputePlatform *string `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// The tags associated with a profiling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codeguruprofiler_profiling_group#tags CodeguruprofilerProfilingGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

