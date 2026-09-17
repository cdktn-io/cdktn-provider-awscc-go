// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorPrefetchScheduleConfig struct {
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
	// The name to assign to the prefetch schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#name MediatailorPrefetchSchedule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the playback configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#playback_configuration_name MediatailorPrefetchSchedule#playback_configuration_name}
	PlaybackConfigurationName *string `field:"required" json:"playbackConfigurationName" yaml:"playbackConfigurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#consumption MediatailorPrefetchSchedule#consumption}.
	Consumption *MediatailorPrefetchScheduleConsumption `field:"optional" json:"consumption" yaml:"consumption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#recurring_prefetch_configuration MediatailorPrefetchSchedule#recurring_prefetch_configuration}.
	RecurringPrefetchConfiguration *MediatailorPrefetchScheduleRecurringPrefetchConfiguration `field:"optional" json:"recurringPrefetchConfiguration" yaml:"recurringPrefetchConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#retrieval MediatailorPrefetchSchedule#retrieval}.
	Retrieval *MediatailorPrefetchScheduleRetrieval `field:"optional" json:"retrieval" yaml:"retrieval"`
	// The frequency that MediaTailor creates prefetch schedules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#schedule_type MediatailorPrefetchSchedule#schedule_type}
	ScheduleType *string `field:"optional" json:"scheduleType" yaml:"scheduleType"`
	// An optional stream identifier that MediaTailor uses to prefetch ads for multiple streams that use the same playback configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#stream_id MediatailorPrefetchSchedule#stream_id}
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// The tags assigned to the prefetch schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#tags MediatailorPrefetchSchedule#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

