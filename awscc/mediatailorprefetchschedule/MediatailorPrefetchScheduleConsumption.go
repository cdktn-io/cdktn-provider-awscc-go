// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule


type MediatailorPrefetchScheduleConsumption struct {
	// If you only want MediaTailor to insert prefetched ads into avails that match specific dynamic variables, set the avail matching criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediatailor_prefetch_schedule#avail_matching_criteria MediatailorPrefetchSchedule#avail_matching_criteria}
	AvailMatchingCriteria interface{} `field:"optional" json:"availMatchingCriteria" yaml:"availMatchingCriteria"`
	// The time when MediaTailor no longer considers the prefetched ads for use in an ad break, as an ISO 8601 date-time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediatailor_prefetch_schedule#end_time MediatailorPrefetchSchedule#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The time when prefetched ads are considered for use in an ad break, as an ISO 8601 date-time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediatailor_prefetch_schedule#start_time MediatailorPrefetchSchedule#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

