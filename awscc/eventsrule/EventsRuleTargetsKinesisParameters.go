// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsrule


type EventsRuleTargetsKinesisParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_rule#partition_key_path EventsRule#partition_key_path}.
	PartitionKeyPath *string `field:"optional" json:"partitionKeyPath" yaml:"partitionKeyPath"`
}

