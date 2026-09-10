// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataHibernationOptions struct {
	// If you set this parameter to ``true``, the instance is enabled for hibernation.  Default: ``false``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#configured Ec2LaunchTemplate#configured}
	Configured interface{} `field:"optional" json:"configured" yaml:"configured"`
}

