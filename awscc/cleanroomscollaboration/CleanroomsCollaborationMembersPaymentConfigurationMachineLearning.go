// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomscollaboration


type CleanroomsCollaborationMembersPaymentConfigurationMachineLearning struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanrooms_collaboration#model_inference CleanroomsCollaboration#model_inference}.
	ModelInference *CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelInference `field:"optional" json:"modelInference" yaml:"modelInference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanrooms_collaboration#model_training CleanroomsCollaboration#model_training}.
	ModelTraining *CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelTraining `field:"optional" json:"modelTraining" yaml:"modelTraining"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanrooms_collaboration#synthetic_data_generation CleanroomsCollaboration#synthetic_data_generation}.
	SyntheticDataGeneration *CleanroomsCollaborationMembersPaymentConfigurationMachineLearningSyntheticDataGeneration `field:"optional" json:"syntheticDataGeneration" yaml:"syntheticDataGeneration"`
}

