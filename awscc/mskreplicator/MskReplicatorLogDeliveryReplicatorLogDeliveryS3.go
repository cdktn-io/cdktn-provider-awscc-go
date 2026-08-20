// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorLogDeliveryReplicatorLogDeliveryS3 struct {
	// The S3 bucket that is the destination for log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_replicator#bucket MskReplicator#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Whether log delivery to S3 is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_replicator#enabled MskReplicator#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The S3 prefix that is the destination for log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_replicator#prefix MskReplicator#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

