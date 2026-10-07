// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberInvokeConfigurationKinesisParameters struct {
	// An explicit hash key that overrides the partition key's shard assignment. Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#explicit_hash_key Eventsv2Subscriber#explicit_hash_key}
	ExplicitHashKey *string `field:"optional" json:"explicitHashKey" yaml:"explicitHashKey"`
	// The partition key that determines which shard each record is written to.
	//
	// Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#partition_key Eventsv2Subscriber#partition_key}
	PartitionKey *string `field:"optional" json:"partitionKey" yaml:"partitionKey"`
}

