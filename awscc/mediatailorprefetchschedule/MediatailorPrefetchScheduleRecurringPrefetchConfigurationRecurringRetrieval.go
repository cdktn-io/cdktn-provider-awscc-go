// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule


type MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrieval struct {
	// The number of seconds that MediaTailor waits after an ad avail before prefetching ads for the next avail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediatailor_prefetch_schedule#delay_after_avail_end_seconds MediatailorPrefetchSchedule#delay_after_avail_end_seconds}
	DelayAfterAvailEndSeconds *float64 `field:"optional" json:"delayAfterAvailEndSeconds" yaml:"delayAfterAvailEndSeconds"`
	// The dynamic variables to use for substitution during prefetch requests to the ADS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediatailor_prefetch_schedule#dynamic_variables MediatailorPrefetchSchedule#dynamic_variables}
	DynamicVariables *map[string]*string `field:"optional" json:"dynamicVariables" yaml:"dynamicVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediatailor_prefetch_schedule#traffic_shaping_retrieval_window MediatailorPrefetchSchedule#traffic_shaping_retrieval_window}.
	TrafficShapingRetrievalWindow *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindow `field:"optional" json:"trafficShapingRetrievalWindow" yaml:"trafficShapingRetrievalWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediatailor_prefetch_schedule#traffic_shaping_tps_configuration MediatailorPrefetchSchedule#traffic_shaping_tps_configuration}.
	TrafficShapingTpsConfiguration *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingTpsConfiguration `field:"optional" json:"trafficShapingTpsConfiguration" yaml:"trafficShapingTpsConfiguration"`
	// Indicates the type of traffic shaping used to limit the number of requests to the ADS at one time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediatailor_prefetch_schedule#traffic_shaping_type MediatailorPrefetchSchedule#traffic_shaping_type}
	TrafficShapingType *string `field:"optional" json:"trafficShapingType" yaml:"trafficShapingType"`
}

