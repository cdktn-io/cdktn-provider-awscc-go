// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfigurationRulesSourceSelectionCriteria struct {
	// A filter that you can specify for selection for modifications on replicas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3_bucket#replica_modifications S3Bucket#replica_modifications}
	ReplicaModifications *S3BucketReplicationConfigurationRulesSourceSelectionCriteriaReplicaModifications `field:"optional" json:"replicaModifications" yaml:"replicaModifications"`
	// A container for filter information for the selection of Amazon S3 objects encrypted with AWS KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3_bucket#sse_kms_encrypted_objects S3Bucket#sse_kms_encrypted_objects}
	SseKmsEncryptedObjects *S3BucketReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects `field:"optional" json:"sseKmsEncryptedObjects" yaml:"sseKmsEncryptedObjects"`
}

