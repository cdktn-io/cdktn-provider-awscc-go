// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationEngineVersion struct {
	// The engine version requested by the user.
	//
	// Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/athena_work_group#selected_engine_version AthenaWorkGroup#selected_engine_version}
	SelectedEngineVersion *string `field:"optional" json:"selectedEngineVersion" yaml:"selectedEngineVersion"`
}

