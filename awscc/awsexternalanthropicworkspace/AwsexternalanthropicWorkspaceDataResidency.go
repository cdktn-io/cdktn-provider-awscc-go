// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsexternalanthropicworkspace


type AwsexternalanthropicWorkspaceDataResidency struct {
	// Permitted inference geo values. Omit to allow all geos (the service default of 'unrestricted'); otherwise list specific geos.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/awsexternalanthropic_workspace#allowed_inference_geos AwsexternalanthropicWorkspace#allowed_inference_geos}
	AllowedInferenceGeos *[]*string `field:"optional" json:"allowedInferenceGeos" yaml:"allowedInferenceGeos"`
	// Default inference geo applied when requests omit the parameter.
	//
	// Defaults to 'global' if omitted. Must be a member of AllowedInferenceGeos unless AllowedInferenceGeos is omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/awsexternalanthropic_workspace#default_inference_geo AwsexternalanthropicWorkspace#default_inference_geo}
	DefaultInferenceGeo *string `field:"optional" json:"defaultInferenceGeo" yaml:"defaultInferenceGeo"`
	// Geographic region for workspace data storage. Immutable after creation. Defaults to 'us' if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/awsexternalanthropic_workspace#workspace_geo AwsexternalanthropicWorkspace#workspace_geo}
	WorkspaceGeo *string `field:"optional" json:"workspaceGeo" yaml:"workspaceGeo"`
}

