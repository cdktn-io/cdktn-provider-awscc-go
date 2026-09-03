// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package robomakerrobotapplication


type RobomakerRobotApplicationRobotSoftwareSuite struct {
	// The name of robot software suite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/robomaker_robot_application#name RobomakerRobotApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version of robot software suite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/robomaker_robot_application#version RobomakerRobotApplication#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

