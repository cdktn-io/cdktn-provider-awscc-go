// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataMonitoring struct {
	// Specify ``true`` to enable detailed monitoring. Otherwise, basic monitoring is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_launch_template#enabled Ec2LaunchTemplate#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

