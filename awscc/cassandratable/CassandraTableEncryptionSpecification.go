// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cassandratable


type CassandraTableEncryptionSpecification struct {
	// Server-side encryption type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cassandra_table#encryption_type CassandraTable#encryption_type}
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The AWS KMS customer master key (CMK) that should be used for the AWS KMS encryption.
	//
	// To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cassandra_table#kms_key_identifier CassandraTable#kms_key_identifier}
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
}

