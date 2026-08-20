// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment


type BatchComputeEnvironmentComputeResourcesLaunchTemplateOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#launch_template_id BatchComputeEnvironment#launch_template_id}.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#launch_template_name BatchComputeEnvironment#launch_template_name}.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#target_instance_types BatchComputeEnvironment#target_instance_types}.
	TargetInstanceTypes *[]*string `field:"optional" json:"targetInstanceTypes" yaml:"targetInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#userdata_type BatchComputeEnvironment#userdata_type}.
	UserdataType *string `field:"optional" json:"userdataType" yaml:"userdataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#version BatchComputeEnvironment#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

