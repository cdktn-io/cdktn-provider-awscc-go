// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorReplicationInfoListConsumerGroupReplication struct {
	// List of regular expression patterns indicating the consumer groups to copy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_replicator#consumer_groups_to_replicate MskReplicator#consumer_groups_to_replicate}
	ConsumerGroupsToReplicate *[]*string `field:"required" json:"consumerGroupsToReplicate" yaml:"consumerGroupsToReplicate"`
	// The consumer group offset synchronization mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_replicator#consumer_group_offset_sync_mode MskReplicator#consumer_group_offset_sync_mode}
	ConsumerGroupOffsetSyncMode *string `field:"optional" json:"consumerGroupOffsetSyncMode" yaml:"consumerGroupOffsetSyncMode"`
	// List of regular expression patterns indicating the consumer groups that should not be replicated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_replicator#consumer_groups_to_exclude MskReplicator#consumer_groups_to_exclude}
	ConsumerGroupsToExclude *[]*string `field:"optional" json:"consumerGroupsToExclude" yaml:"consumerGroupsToExclude"`
	// Whether to periodically check for new consumer groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_replicator#detect_and_copy_new_consumer_groups MskReplicator#detect_and_copy_new_consumer_groups}
	DetectAndCopyNewConsumerGroups interface{} `field:"optional" json:"detectAndCopyNewConsumerGroups" yaml:"detectAndCopyNewConsumerGroups"`
	// Whether to periodically write the translated offsets to __consumer_offsets topic in target cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_replicator#synchronise_consumer_group_offsets MskReplicator#synchronise_consumer_group_offsets}
	SynchroniseConsumerGroupOffsets interface{} `field:"optional" json:"synchroniseConsumerGroupOffsets" yaml:"synchroniseConsumerGroupOffsets"`
}

