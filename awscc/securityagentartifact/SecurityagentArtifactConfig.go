// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentartifact

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityagentArtifactConfig struct {
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
	// The unique identifier of the agent space to add the artifact to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityagent_artifact#agent_space_id SecurityagentArtifact#agent_space_id}
	AgentSpaceId *string `field:"required" json:"agentSpaceId" yaml:"agentSpaceId"`
	// The file type of the artifact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityagent_artifact#artifact_type SecurityagentArtifact#artifact_type}
	ArtifactType *string `field:"required" json:"artifactType" yaml:"artifactType"`
	// The file name of the artifact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityagent_artifact#file_name SecurityagentArtifact#file_name}
	FileName *string `field:"required" json:"fileName" yaml:"fileName"`
	// The binary content of the artifact to upload, encoded as a Base64 string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityagent_artifact#artifact_content SecurityagentArtifact#artifact_content}
	ArtifactContent *string `field:"optional" json:"artifactContent" yaml:"artifactContent"`
}

