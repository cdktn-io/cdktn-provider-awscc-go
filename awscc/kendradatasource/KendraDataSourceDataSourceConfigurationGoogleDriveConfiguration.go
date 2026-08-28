// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceDataSourceConfigurationGoogleDriveConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_data_source#exclude_mime_types KendraDataSource#exclude_mime_types}.
	ExcludeMimeTypes *[]*string `field:"optional" json:"excludeMimeTypes" yaml:"excludeMimeTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_data_source#exclude_shared_drives KendraDataSource#exclude_shared_drives}.
	ExcludeSharedDrives *[]*string `field:"optional" json:"excludeSharedDrives" yaml:"excludeSharedDrives"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_data_source#exclude_user_accounts KendraDataSource#exclude_user_accounts}.
	ExcludeUserAccounts *[]*string `field:"optional" json:"excludeUserAccounts" yaml:"excludeUserAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_data_source#exclusion_patterns KendraDataSource#exclusion_patterns}.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_data_source#field_mappings KendraDataSource#field_mappings}.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_data_source#inclusion_patterns KendraDataSource#inclusion_patterns}.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_data_source#secret_arn KendraDataSource#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

