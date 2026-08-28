// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmlconfiguredmodelalgorithm

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsmlConfiguredModelAlgorithmConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm#name CleanroomsmlConfiguredModelAlgorithm#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm#role_arn CleanroomsmlConfiguredModelAlgorithm#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm#description CleanroomsmlConfiguredModelAlgorithm#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm#inference_container_config CleanroomsmlConfiguredModelAlgorithm#inference_container_config}.
	InferenceContainerConfig *CleanroomsmlConfiguredModelAlgorithmInferenceContainerConfig `field:"optional" json:"inferenceContainerConfig" yaml:"inferenceContainerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm#kms_key_arn CleanroomsmlConfiguredModelAlgorithm#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms-ml configured model algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm#tags CleanroomsmlConfiguredModelAlgorithm#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm#training_container_config CleanroomsmlConfiguredModelAlgorithm#training_container_config}.
	TrainingContainerConfig *CleanroomsmlConfiguredModelAlgorithmTrainingContainerConfig `field:"optional" json:"trainingContainerConfig" yaml:"trainingContainerConfig"`
}

