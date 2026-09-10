// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerartifact

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerArtifactConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The artifact type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_artifact#artifact_type SagemakerArtifact#artifact_type}
	ArtifactType *string `field:"required" json:"artifactType" yaml:"artifactType"`
	// The source of the artifact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_artifact#source SagemakerArtifact#source}
	Source *SagemakerArtifactSource `field:"required" json:"source" yaml:"source"`
	// The name of the artifact. Must be unique to your account in an AWS Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_artifact#artifact_name SagemakerArtifact#artifact_name}
	ArtifactName *string `field:"optional" json:"artifactName" yaml:"artifactName"`
	// Metadata properties of the tracking entity, trial, or trial component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_artifact#metadata_properties SagemakerArtifact#metadata_properties}
	MetadataProperties *SagemakerArtifactMetadataProperties `field:"optional" json:"metadataProperties" yaml:"metadataProperties"`
	// A list of properties to add to the artifact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_artifact#properties SagemakerArtifact#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// A list of tags to apply to the artifact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_artifact#tags SagemakerArtifact#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

