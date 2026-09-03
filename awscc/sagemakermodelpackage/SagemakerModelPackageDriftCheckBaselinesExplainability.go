// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageDriftCheckBaselinesExplainability struct {
	// Represents a File Source Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_model_package#config_file SagemakerModelPackage#config_file}
	ConfigFile *SagemakerModelPackageDriftCheckBaselinesExplainabilityConfigFile `field:"optional" json:"configFile" yaml:"configFile"`
	// Represents a Metric Source Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_model_package#constraints SagemakerModelPackage#constraints}
	Constraints *SagemakerModelPackageDriftCheckBaselinesExplainabilityConstraints `field:"optional" json:"constraints" yaml:"constraints"`
}

