// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogcloudformationprovisionedproduct


type ServicecatalogCloudformationProvisionedProductProvisioningPreferences struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicecatalog_cloudformation_provisioned_product#stack_set_accounts ServicecatalogCloudformationProvisionedProduct#stack_set_accounts}.
	StackSetAccounts *[]*string `field:"optional" json:"stackSetAccounts" yaml:"stackSetAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicecatalog_cloudformation_provisioned_product#stack_set_failure_tolerance_count ServicecatalogCloudformationProvisionedProduct#stack_set_failure_tolerance_count}.
	StackSetFailureToleranceCount *float64 `field:"optional" json:"stackSetFailureToleranceCount" yaml:"stackSetFailureToleranceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicecatalog_cloudformation_provisioned_product#stack_set_failure_tolerance_percentage ServicecatalogCloudformationProvisionedProduct#stack_set_failure_tolerance_percentage}.
	StackSetFailureTolerancePercentage *float64 `field:"optional" json:"stackSetFailureTolerancePercentage" yaml:"stackSetFailureTolerancePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicecatalog_cloudformation_provisioned_product#stack_set_max_concurrency_count ServicecatalogCloudformationProvisionedProduct#stack_set_max_concurrency_count}.
	StackSetMaxConcurrencyCount *float64 `field:"optional" json:"stackSetMaxConcurrencyCount" yaml:"stackSetMaxConcurrencyCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicecatalog_cloudformation_provisioned_product#stack_set_max_concurrency_percentage ServicecatalogCloudformationProvisionedProduct#stack_set_max_concurrency_percentage}.
	StackSetMaxConcurrencyPercentage *float64 `field:"optional" json:"stackSetMaxConcurrencyPercentage" yaml:"stackSetMaxConcurrencyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicecatalog_cloudformation_provisioned_product#stack_set_operation_type ServicecatalogCloudformationProvisionedProduct#stack_set_operation_type}.
	StackSetOperationType *string `field:"optional" json:"stackSetOperationType" yaml:"stackSetOperationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicecatalog_cloudformation_provisioned_product#stack_set_regions ServicecatalogCloudformationProvisionedProduct#stack_set_regions}.
	StackSetRegions *[]*string `field:"optional" json:"stackSetRegions" yaml:"stackSetRegions"`
}

