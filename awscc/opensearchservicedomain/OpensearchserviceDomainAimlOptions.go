// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchservicedomain


type OpensearchserviceDomainAimlOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#s3_vectors_engine OpensearchserviceDomain#s3_vectors_engine}.
	S3VectorsEngine *OpensearchserviceDomainAimlOptionsS3VectorsEngine `field:"optional" json:"s3VectorsEngine" yaml:"s3VectorsEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#serverless_vector_acceleration OpensearchserviceDomain#serverless_vector_acceleration}.
	ServerlessVectorAcceleration *OpensearchserviceDomainAimlOptionsServerlessVectorAcceleration `field:"optional" json:"serverlessVectorAcceleration" yaml:"serverlessVectorAcceleration"`
}

