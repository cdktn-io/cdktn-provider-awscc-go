// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpuReferences struct {
	// The instance family to refer.
	//
	// Ensure that you specify the correct family name. For example, C6i and C6g are valid values, but C6 is not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#instance_family Ec2LaunchTemplate#instance_family}
	InstanceFamily *string `field:"optional" json:"instanceFamily" yaml:"instanceFamily"`
}

