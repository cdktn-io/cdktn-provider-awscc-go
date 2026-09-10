// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule


type MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindow struct {
	// The amount of time, in seconds, that MediaTailor spreads prefetch requests to the ADS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_prefetch_schedule#retrieval_window_duration_seconds MediatailorPrefetchSchedule#retrieval_window_duration_seconds}
	RetrievalWindowDurationSeconds *float64 `field:"optional" json:"retrievalWindowDurationSeconds" yaml:"retrievalWindowDurationSeconds"`
}

