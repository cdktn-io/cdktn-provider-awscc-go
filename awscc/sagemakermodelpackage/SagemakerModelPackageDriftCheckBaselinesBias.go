// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageDriftCheckBaselinesBias struct {
	// Represents a File Source Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#config_file SagemakerModelPackage#config_file}
	ConfigFile *SagemakerModelPackageDriftCheckBaselinesBiasConfigFile `field:"optional" json:"configFile" yaml:"configFile"`
	// Represents a Metric Source Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#post_training_constraints SagemakerModelPackage#post_training_constraints}
	PostTrainingConstraints *SagemakerModelPackageDriftCheckBaselinesBiasPostTrainingConstraints `field:"optional" json:"postTrainingConstraints" yaml:"postTrainingConstraints"`
	// Represents a Metric Source Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#pre_training_constraints SagemakerModelPackage#pre_training_constraints}
	PreTrainingConstraints *SagemakerModelPackageDriftCheckBaselinesBiasPreTrainingConstraints `field:"optional" json:"preTrainingConstraints" yaml:"preTrainingConstraints"`
}

