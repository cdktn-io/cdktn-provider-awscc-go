// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule


type MediatailorPrefetchScheduleRetrieval struct {
	// The dynamic variables to use for substitution during prefetch requests to the ad decision server (ADS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#dynamic_variables MediatailorPrefetchSchedule#dynamic_variables}
	DynamicVariables *map[string]*string `field:"optional" json:"dynamicVariables" yaml:"dynamicVariables"`
	// The time when prefetch retrieval ends for the ad break, as an ISO 8601 date-time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#end_time MediatailorPrefetchSchedule#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The time when prefetch retrievals can start for this break, as an ISO 8601 date-time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#start_time MediatailorPrefetchSchedule#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#traffic_shaping_retrieval_window MediatailorPrefetchSchedule#traffic_shaping_retrieval_window}.
	TrafficShapingRetrievalWindow *MediatailorPrefetchScheduleRetrievalTrafficShapingRetrievalWindow `field:"optional" json:"trafficShapingRetrievalWindow" yaml:"trafficShapingRetrievalWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#traffic_shaping_tps_configuration MediatailorPrefetchSchedule#traffic_shaping_tps_configuration}.
	TrafficShapingTpsConfiguration *MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfiguration `field:"optional" json:"trafficShapingTpsConfiguration" yaml:"trafficShapingTpsConfiguration"`
	// Indicates the type of traffic shaping used to limit the number of requests to the ADS at one time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#traffic_shaping_type MediatailorPrefetchSchedule#traffic_shaping_type}
	TrafficShapingType *string `field:"optional" json:"trafficShapingType" yaml:"trafficShapingType"`
}

