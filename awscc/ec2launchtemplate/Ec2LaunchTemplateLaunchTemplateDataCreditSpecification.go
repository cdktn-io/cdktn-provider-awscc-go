// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataCreditSpecification struct {
	// The credit option for CPU usage of a T instance.  Valid values: ``standard`` | ``unlimited``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#cpu_credits Ec2LaunchTemplate#cpu_credits}
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
}

