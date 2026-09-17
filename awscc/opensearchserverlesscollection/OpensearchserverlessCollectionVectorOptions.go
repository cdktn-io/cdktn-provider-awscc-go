// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserverlesscollection


type OpensearchserverlessCollectionVectorOptions struct {
	// Indicates whether GPU acceleration is enabled for vector indexing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/opensearchserverless_collection#serverless_vector_acceleration OpensearchserverlessCollection#serverless_vector_acceleration}
	ServerlessVectorAcceleration *string `field:"optional" json:"serverlessVectorAcceleration" yaml:"serverlessVectorAcceleration"`
}

