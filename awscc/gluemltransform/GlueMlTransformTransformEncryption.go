// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluemltransform


type GlueMlTransformTransformEncryption struct {
	// The encryption-at-rest settings of the transform that apply to accessing user data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_ml_transform#ml_user_data_encryption GlueMlTransform#ml_user_data_encryption}
	MlUserDataEncryption *GlueMlTransformTransformEncryptionMlUserDataEncryption `field:"optional" json:"mlUserDataEncryption" yaml:"mlUserDataEncryption"`
	// The name of the security configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_ml_transform#task_run_security_configuration_name GlueMlTransform#task_run_security_configuration_name}
	TaskRunSecurityConfigurationName *string `field:"optional" json:"taskRunSecurityConfigurationName" yaml:"taskRunSecurityConfigurationName"`
}

