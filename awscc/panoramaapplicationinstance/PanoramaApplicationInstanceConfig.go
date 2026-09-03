// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package panoramaapplicationinstance

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PanoramaApplicationInstanceConfig struct {
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
	// The device's ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/panorama_application_instance#default_runtime_context_device PanoramaApplicationInstance#default_runtime_context_device}
	DefaultRuntimeContextDevice *string `field:"required" json:"defaultRuntimeContextDevice" yaml:"defaultRuntimeContextDevice"`
	// The application's manifest document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/panorama_application_instance#manifest_payload PanoramaApplicationInstance#manifest_payload}
	ManifestPayload *PanoramaApplicationInstanceManifestPayload `field:"required" json:"manifestPayload" yaml:"manifestPayload"`
	// The ID of an application instance to replace with the new instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/panorama_application_instance#application_instance_id_to_replace PanoramaApplicationInstance#application_instance_id_to_replace}
	ApplicationInstanceIdToReplace *string `field:"optional" json:"applicationInstanceIdToReplace" yaml:"applicationInstanceIdToReplace"`
	// A description for the application instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/panorama_application_instance#description PanoramaApplicationInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Setting overrides for the application manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/panorama_application_instance#manifest_overrides_payload PanoramaApplicationInstance#manifest_overrides_payload}
	ManifestOverridesPayload *PanoramaApplicationInstanceManifestOverridesPayload `field:"optional" json:"manifestOverridesPayload" yaml:"manifestOverridesPayload"`
	// A name for the application instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/panorama_application_instance#name PanoramaApplicationInstance#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of a runtime role for the application instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/panorama_application_instance#runtime_role_arn PanoramaApplicationInstance#runtime_role_arn}
	RuntimeRoleArn *string `field:"optional" json:"runtimeRoleArn" yaml:"runtimeRoleArn"`
	// Tags for the application instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/panorama_application_instance#tags PanoramaApplicationInstance#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

