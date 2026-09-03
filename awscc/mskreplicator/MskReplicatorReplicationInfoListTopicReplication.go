// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorReplicationInfoListTopicReplication struct {
	// List of regular expression patterns indicating the topics to copy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#topics_to_replicate MskReplicator#topics_to_replicate}
	TopicsToReplicate *[]*string `field:"required" json:"topicsToReplicate" yaml:"topicsToReplicate"`
	// Whether to periodically configure remote topic ACLs to match their corresponding upstream topics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#copy_access_control_lists_for_topics MskReplicator#copy_access_control_lists_for_topics}
	CopyAccessControlListsForTopics interface{} `field:"optional" json:"copyAccessControlListsForTopics" yaml:"copyAccessControlListsForTopics"`
	// Whether to periodically configure remote topics to match their corresponding upstream topics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#copy_topic_configurations MskReplicator#copy_topic_configurations}
	CopyTopicConfigurations interface{} `field:"optional" json:"copyTopicConfigurations" yaml:"copyTopicConfigurations"`
	// Whether to periodically check for new topics and partitions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#detect_and_copy_new_topics MskReplicator#detect_and_copy_new_topics}
	DetectAndCopyNewTopics interface{} `field:"optional" json:"detectAndCopyNewTopics" yaml:"detectAndCopyNewTopics"`
	// Configuration for specifying the position in the topics to start replicating from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#starting_position MskReplicator#starting_position}
	StartingPosition *MskReplicatorReplicationInfoListTopicReplicationStartingPosition `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// Configuration for specifying replicated topic names should be the same as their corresponding upstream topics or prefixed with source cluster alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#topic_name_configuration MskReplicator#topic_name_configuration}
	TopicNameConfiguration *MskReplicatorReplicationInfoListTopicReplicationTopicNameConfiguration `field:"optional" json:"topicNameConfiguration" yaml:"topicNameConfiguration"`
	// List of regular expression patterns indicating the topics that should not be replicated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#topics_to_exclude MskReplicator#topics_to_exclude}
	TopicsToExclude *[]*string `field:"optional" json:"topicsToExclude" yaml:"topicsToExclude"`
}

