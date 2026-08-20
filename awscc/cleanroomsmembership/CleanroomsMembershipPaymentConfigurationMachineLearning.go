// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmembership


type CleanroomsMembershipPaymentConfigurationMachineLearning struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_membership#model_inference CleanroomsMembership#model_inference}.
	ModelInference *CleanroomsMembershipPaymentConfigurationMachineLearningModelInference `field:"optional" json:"modelInference" yaml:"modelInference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_membership#model_training CleanroomsMembership#model_training}.
	ModelTraining *CleanroomsMembershipPaymentConfigurationMachineLearningModelTraining `field:"optional" json:"modelTraining" yaml:"modelTraining"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_membership#synthetic_data_generation CleanroomsMembership#synthetic_data_generation}.
	SyntheticDataGeneration *CleanroomsMembershipPaymentConfigurationMachineLearningSyntheticDataGeneration `field:"optional" json:"syntheticDataGeneration" yaml:"syntheticDataGeneration"`
}

