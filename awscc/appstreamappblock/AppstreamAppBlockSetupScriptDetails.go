// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamappblock


type AppstreamAppBlockSetupScriptDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appstream_app_block#executable_parameters AppstreamAppBlock#executable_parameters}.
	ExecutableParameters *string `field:"optional" json:"executableParameters" yaml:"executableParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appstream_app_block#executable_path AppstreamAppBlock#executable_path}.
	ExecutablePath *string `field:"optional" json:"executablePath" yaml:"executablePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appstream_app_block#script_s3_location AppstreamAppBlock#script_s3_location}.
	ScriptS3Location *AppstreamAppBlockSetupScriptDetailsScriptS3Location `field:"optional" json:"scriptS3Location" yaml:"scriptS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appstream_app_block#timeout_in_seconds AppstreamAppBlock#timeout_in_seconds}.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

