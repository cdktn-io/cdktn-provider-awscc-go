// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClustersEncryptionInTransit struct {
	// The type of encryption in transit to the Apache Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_replicator#encryption_type MskReplicator#encryption_type}
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The root CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_replicator#root_ca_certificate MskReplicator#root_ca_certificate}
	RootCaCertificate *string `field:"optional" json:"rootCaCertificate" yaml:"rootCaCertificate"`
}

