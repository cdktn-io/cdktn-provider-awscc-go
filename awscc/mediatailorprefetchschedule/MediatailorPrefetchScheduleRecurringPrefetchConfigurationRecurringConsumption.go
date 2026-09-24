// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule


type MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumption struct {
	// The configuration for the dynamic variables that determine which ad breaks that MediaTailor inserts prefetched ads in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediatailor_prefetch_schedule#avail_matching_criteria MediatailorPrefetchSchedule#avail_matching_criteria}
	AvailMatchingCriteria interface{} `field:"optional" json:"availMatchingCriteria" yaml:"availMatchingCriteria"`
	// The number of seconds that an ad is available for insertion after it was prefetched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediatailor_prefetch_schedule#retrieved_ad_expiration_seconds MediatailorPrefetchSchedule#retrieved_ad_expiration_seconds}
	RetrievedAdExpirationSeconds *float64 `field:"optional" json:"retrievedAdExpirationSeconds" yaml:"retrievedAdExpirationSeconds"`
}

