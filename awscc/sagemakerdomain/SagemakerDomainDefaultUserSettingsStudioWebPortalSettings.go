// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsStudioWebPortalSettings struct {
	// Applications supported in Studio that are hidden from the Studio left navigation pane.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#hidden_app_types SagemakerDomain#hidden_app_types}
	HiddenAppTypes *[]*string `field:"optional" json:"hiddenAppTypes" yaml:"hiddenAppTypes"`
	// The instance types you are hiding from the Studio user interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#hidden_instance_types SagemakerDomain#hidden_instance_types}
	HiddenInstanceTypes *[]*string `field:"optional" json:"hiddenInstanceTypes" yaml:"hiddenInstanceTypes"`
	// The machine learning tools that are hidden from the Studio left navigation pane.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#hidden_ml_tools SagemakerDomain#hidden_ml_tools}
	HiddenMlTools *[]*string `field:"optional" json:"hiddenMlTools" yaml:"hiddenMlTools"`
	// The version aliases you are hiding from the Studio user interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#hidden_sage_maker_image_version_aliases SagemakerDomain#hidden_sage_maker_image_version_aliases}
	HiddenSageMakerImageVersionAliases interface{} `field:"optional" json:"hiddenSageMakerImageVersionAliases" yaml:"hiddenSageMakerImageVersionAliases"`
}

