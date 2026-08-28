// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerartifact


type SagemakerArtifactSource struct {
	// The URI of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_artifact#source_uri SagemakerArtifact#source_uri}
	SourceUri *string `field:"required" json:"sourceUri" yaml:"sourceUri"`
	// A list of source types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_artifact#source_types SagemakerArtifact#source_types}
	SourceTypes interface{} `field:"optional" json:"sourceTypes" yaml:"sourceTypes"`
}

