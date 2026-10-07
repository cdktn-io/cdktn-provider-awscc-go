// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2eventbus


type Eventsv2EventBusStorageConfiguration struct {
	// The number of days events are retained on the event bus for replay, 1-365.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_bus#retention_period_in_days Eventsv2EventBus#retention_period_in_days}
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
}

