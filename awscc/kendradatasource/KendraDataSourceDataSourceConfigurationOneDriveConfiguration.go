// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceDataSourceConfigurationOneDriveConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#disable_local_groups KendraDataSource#disable_local_groups}.
	DisableLocalGroups interface{} `field:"optional" json:"disableLocalGroups" yaml:"disableLocalGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#exclusion_patterns KendraDataSource#exclusion_patterns}.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#field_mappings KendraDataSource#field_mappings}.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#inclusion_patterns KendraDataSource#inclusion_patterns}.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#one_drive_users KendraDataSource#one_drive_users}.
	OneDriveUsers *KendraDataSourceDataSourceConfigurationOneDriveConfigurationOneDriveUsers `field:"optional" json:"oneDriveUsers" yaml:"oneDriveUsers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#secret_arn KendraDataSource#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#tenant_domain KendraDataSource#tenant_domain}.
	TenantDomain *string `field:"optional" json:"tenantDomain" yaml:"tenantDomain"`
}

