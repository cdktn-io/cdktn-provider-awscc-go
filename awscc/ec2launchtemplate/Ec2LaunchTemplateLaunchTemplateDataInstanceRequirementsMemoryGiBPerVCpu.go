// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpu struct {
	// The maximum amount of memory per vCPU, in GiB. To specify no maximum limit, omit this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_launch_template#max Ec2LaunchTemplate#max}
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of memory per vCPU, in GiB. To specify no minimum limit, omit this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_launch_template#min Ec2LaunchTemplate#min}
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

