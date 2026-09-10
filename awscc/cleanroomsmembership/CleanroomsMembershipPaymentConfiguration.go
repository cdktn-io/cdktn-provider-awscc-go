// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmembership


type CleanroomsMembershipPaymentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cleanrooms_membership#job_compute CleanroomsMembership#job_compute}.
	JobCompute *CleanroomsMembershipPaymentConfigurationJobCompute `field:"optional" json:"jobCompute" yaml:"jobCompute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cleanrooms_membership#machine_learning CleanroomsMembership#machine_learning}.
	MachineLearning *CleanroomsMembershipPaymentConfigurationMachineLearning `field:"optional" json:"machineLearning" yaml:"machineLearning"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cleanrooms_membership#query_compute CleanroomsMembership#query_compute}.
	QueryCompute *CleanroomsMembershipPaymentConfigurationQueryCompute `field:"optional" json:"queryCompute" yaml:"queryCompute"`
}

