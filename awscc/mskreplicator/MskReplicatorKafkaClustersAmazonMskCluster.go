// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClustersAmazonMskCluster struct {
	// The ARN of an Amazon MSK cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_replicator#msk_cluster_arn MskReplicator#msk_cluster_arn}
	MskClusterArn *string `field:"optional" json:"mskClusterArn" yaml:"mskClusterArn"`
}

