// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtrailtrail


type CloudtrailTrailEventSelectors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudtrail_trail#data_resources CloudtrailTrail#data_resources}.
	DataResources interface{} `field:"optional" json:"dataResources" yaml:"dataResources"`
	// An optional list of service event sources from which you do not want management events to be logged on your trail.
	//
	// In this release, the list can be empty (disables the filter), or it can filter out AWS Key Management Service events by containing "kms.amazonaws.com". By default, ExcludeManagementEventSources is empty, and AWS KMS events are included in events that are logged to your trail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudtrail_trail#exclude_management_event_sources CloudtrailTrail#exclude_management_event_sources}
	ExcludeManagementEventSources *[]*string `field:"optional" json:"excludeManagementEventSources" yaml:"excludeManagementEventSources"`
	// Specify if you want your event selector to include management events for your trail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudtrail_trail#include_management_events CloudtrailTrail#include_management_events}
	IncludeManagementEvents interface{} `field:"optional" json:"includeManagementEvents" yaml:"includeManagementEvents"`
	// Specify if you want your trail to log read-only events, write-only events, or all.
	//
	// For example, the EC2 GetConsoleOutput is a read-only API operation and RunInstances is a write-only API operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudtrail_trail#read_write_type CloudtrailTrail#read_write_type}
	ReadWriteType *string `field:"optional" json:"readWriteType" yaml:"readWriteType"`
}

