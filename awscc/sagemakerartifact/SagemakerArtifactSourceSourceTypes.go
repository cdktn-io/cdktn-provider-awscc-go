// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerartifact


type SagemakerArtifactSourceSourceTypes struct {
	// The type of ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_artifact#source_id_type SagemakerArtifact#source_id_type}
	SourceIdType *string `field:"optional" json:"sourceIdType" yaml:"sourceIdType"`
	// The ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_artifact#value SagemakerArtifact#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

