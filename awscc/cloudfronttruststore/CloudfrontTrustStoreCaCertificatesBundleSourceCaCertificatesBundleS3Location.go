// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfronttruststore


type CloudfrontTrustStoreCaCertificatesBundleSourceCaCertificatesBundleS3Location struct {
	// The S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_trust_store#bucket CloudfrontTrustStore#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The location's key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_trust_store#key CloudfrontTrustStore#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The location's Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_trust_store#region CloudfrontTrustStore#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The location's version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_trust_store#version CloudfrontTrustStore#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

