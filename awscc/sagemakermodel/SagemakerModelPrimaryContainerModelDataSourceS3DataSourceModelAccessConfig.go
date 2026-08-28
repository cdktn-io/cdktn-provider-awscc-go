// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelPrimaryContainerModelDataSourceS3DataSourceModelAccessConfig struct {
	// Specifies agreement to the model end-user license agreement (EULA).
	//
	// The `AcceptEula` value must be explicitly defined as `True` in order to accept the EULA that this model requires. You are responsible for reviewing and complying with any applicable license terms and making sure they are acceptable for your use case before downloading or using a model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#accept_eula SagemakerModel#accept_eula}
	AcceptEula interface{} `field:"optional" json:"acceptEula" yaml:"acceptEula"`
}

