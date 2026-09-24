// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpu struct {
	// A list of references to be used as baseline for the CPU performance.
	//
	// Currently, you can only specify a single reference across different instance type variations such as CPU manufacturers, architectures etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_launch_template#references Ec2LaunchTemplate#references}
	References interface{} `field:"optional" json:"references" yaml:"references"`
}

