// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottwinmakerscene

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IottwinmakerSceneConfig struct {
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
	// The relative path that specifies the location of the content definition file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iottwinmaker_scene#content_location IottwinmakerScene#content_location}
	ContentLocation *string `field:"required" json:"contentLocation" yaml:"contentLocation"`
	// The ID of the scene.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iottwinmaker_scene#scene_id IottwinmakerScene#scene_id}
	SceneId *string `field:"required" json:"sceneId" yaml:"sceneId"`
	// The ID of the scene.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iottwinmaker_scene#workspace_id IottwinmakerScene#workspace_id}
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// A list of capabilities that the scene uses to render.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iottwinmaker_scene#capabilities IottwinmakerScene#capabilities}
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// The description of the scene.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iottwinmaker_scene#description IottwinmakerScene#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A key-value pair of scene metadata for the scene.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iottwinmaker_scene#scene_metadata IottwinmakerScene#scene_metadata}
	SceneMetadata *map[string]*string `field:"optional" json:"sceneMetadata" yaml:"sceneMetadata"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iottwinmaker_scene#tags IottwinmakerScene#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

