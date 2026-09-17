// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package drslaunchconfigurationtemplate


type DrsLaunchConfigurationTemplateLicensing struct {
	// Whether to enable Bring your own license or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/drs_launch_configuration_template#os_byol DrsLaunchConfigurationTemplate#os_byol}
	OsByol interface{} `field:"optional" json:"osByol" yaml:"osByol"`
}

