// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule


type MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfiguration struct {
	// The expected peak number of concurrent viewers for your content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediatailor_prefetch_schedule#peak_concurrent_users MediatailorPrefetchSchedule#peak_concurrent_users}
	PeakConcurrentUsers *float64 `field:"optional" json:"peakConcurrentUsers" yaml:"peakConcurrentUsers"`
	// The maximum number of transactions per second (TPS) that your ad decision server (ADS) can handle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediatailor_prefetch_schedule#peak_tps MediatailorPrefetchSchedule#peak_tps}
	PeakTps *float64 `field:"optional" json:"peakTps" yaml:"peakTps"`
}

