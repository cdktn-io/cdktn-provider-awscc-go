// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimage


type ImagebuilderImageImagePipelineExecutionSettings struct {
	// The deployment ID of the pipeline, used to trigger new image pipeline executions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/imagebuilder_image#deployment_id ImagebuilderImage#deployment_id}
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// Whether to trigger the image pipeline when the pipeline is updated. False by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/imagebuilder_image#on_update ImagebuilderImage#on_update}
	OnUpdate interface{} `field:"optional" json:"onUpdate" yaml:"onUpdate"`
}

