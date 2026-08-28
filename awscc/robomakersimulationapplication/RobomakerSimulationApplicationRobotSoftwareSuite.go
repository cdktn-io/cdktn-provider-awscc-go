// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package robomakersimulationapplication


type RobomakerSimulationApplicationRobotSoftwareSuite struct {
	// The name of the robot software suite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/robomaker_simulation_application#name RobomakerSimulationApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version of the robot software suite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/robomaker_simulation_application#version RobomakerSimulationApplication#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

