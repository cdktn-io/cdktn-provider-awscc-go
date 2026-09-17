// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule


type MediatailorPrefetchScheduleRecurringPrefetchConfiguration struct {
	// The end time for the window that MediaTailor prefetches and inserts ads in a live event, as an ISO 8601 date-time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#end_time MediatailorPrefetchSchedule#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#recurring_consumption MediatailorPrefetchSchedule#recurring_consumption}.
	RecurringConsumption *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumption `field:"optional" json:"recurringConsumption" yaml:"recurringConsumption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#recurring_retrieval MediatailorPrefetchSchedule#recurring_retrieval}.
	RecurringRetrieval *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrieval `field:"optional" json:"recurringRetrieval" yaml:"recurringRetrieval"`
	// The start time for the window that MediaTailor prefetches and inserts ads in a live event, as an ISO 8601 date-time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#start_time MediatailorPrefetchSchedule#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

