// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule


type MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionAvailMatchingCriteria struct {
	// The dynamic variable(s) that MediaTailor should use as avail matching criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#dynamic_variable MediatailorPrefetchSchedule#dynamic_variable}
	DynamicVariable *string `field:"optional" json:"dynamicVariable" yaml:"dynamicVariable"`
	// For the DynamicVariable specified in AvailMatchingCriteria, the Operator that is used for the comparison.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_prefetch_schedule#operator MediatailorPrefetchSchedule#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

