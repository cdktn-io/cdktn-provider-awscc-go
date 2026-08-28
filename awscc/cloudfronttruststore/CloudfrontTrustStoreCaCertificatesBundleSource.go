// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfronttruststore


type CloudfrontTrustStoreCaCertificatesBundleSource struct {
	// The CA certificates bundle location in Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_trust_store#ca_certificates_bundle_s3_location CloudfrontTrustStore#ca_certificates_bundle_s3_location}
	CaCertificatesBundleS3Location *CloudfrontTrustStoreCaCertificatesBundleSourceCaCertificatesBundleS3Location `field:"optional" json:"caCertificatesBundleS3Location" yaml:"caCertificatesBundleS3Location"`
}

