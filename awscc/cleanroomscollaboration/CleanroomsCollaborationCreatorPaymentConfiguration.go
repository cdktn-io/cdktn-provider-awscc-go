// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomscollaboration


type CleanroomsCollaborationCreatorPaymentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_collaboration#job_compute CleanroomsCollaboration#job_compute}.
	JobCompute *CleanroomsCollaborationCreatorPaymentConfigurationJobCompute `field:"optional" json:"jobCompute" yaml:"jobCompute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_collaboration#machine_learning CleanroomsCollaboration#machine_learning}.
	MachineLearning *CleanroomsCollaborationCreatorPaymentConfigurationMachineLearning `field:"optional" json:"machineLearning" yaml:"machineLearning"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_collaboration#query_compute CleanroomsCollaboration#query_compute}.
	QueryCompute *CleanroomsCollaborationCreatorPaymentConfigurationQueryCompute `field:"optional" json:"queryCompute" yaml:"queryCompute"`
}

