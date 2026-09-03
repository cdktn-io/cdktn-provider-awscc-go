// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerproject


type SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParameters struct {
	// The parameter key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_project#key SagemakerProject#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_project#value SagemakerProject#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

