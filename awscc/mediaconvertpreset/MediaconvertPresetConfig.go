// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconvertpreset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconvertPresetConfig struct {
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
	// Specify, in JSON format, the transcoding job settings for this output preset.
	//
	// This specification must conform to the AWS Elemental MediaConvert job validation. For information about forming this specification, see the Remarks section later in this topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconvert_preset#settings_json MediaconvertPreset#settings_json}
	SettingsJson *string `field:"required" json:"settingsJson" yaml:"settingsJson"`
	// The new category for the preset, if you are changing it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconvert_preset#category MediaconvertPreset#category}
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The new description for the preset, if you are changing it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconvert_preset#description MediaconvertPreset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the preset that you are modifying.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconvert_preset#name MediaconvertPreset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconvert_preset#tags MediaconvertPreset#tags}.
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
}

