// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonepolicygrant


type DatazonePolicyGrantDetail struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_policy_grant#add_to_project_member_pool DatazonePolicyGrant#add_to_project_member_pool}.
	AddToProjectMemberPool *DatazonePolicyGrantDetailAddToProjectMemberPool `field:"optional" json:"addToProjectMemberPool" yaml:"addToProjectMemberPool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_policy_grant#create_asset_type DatazonePolicyGrant#create_asset_type}.
	CreateAssetType *DatazonePolicyGrantDetailCreateAssetType `field:"optional" json:"createAssetType" yaml:"createAssetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_policy_grant#create_domain_unit DatazonePolicyGrant#create_domain_unit}.
	CreateDomainUnit *DatazonePolicyGrantDetailCreateDomainUnit `field:"optional" json:"createDomainUnit" yaml:"createDomainUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_policy_grant#create_environment DatazonePolicyGrant#create_environment}.
	CreateEnvironment *string `field:"optional" json:"createEnvironment" yaml:"createEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_policy_grant#create_environment_from_blueprint DatazonePolicyGrant#create_environment_from_blueprint}.
	CreateEnvironmentFromBlueprint *string `field:"optional" json:"createEnvironmentFromBlueprint" yaml:"createEnvironmentFromBlueprint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_policy_grant#create_environment_profile DatazonePolicyGrant#create_environment_profile}.
	CreateEnvironmentProfile *DatazonePolicyGrantDetailCreateEnvironmentProfile `field:"optional" json:"createEnvironmentProfile" yaml:"createEnvironmentProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_policy_grant#create_form_type DatazonePolicyGrant#create_form_type}.
	CreateFormType *DatazonePolicyGrantDetailCreateFormType `field:"optional" json:"createFormType" yaml:"createFormType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_policy_grant#create_glossary DatazonePolicyGrant#create_glossary}.
	CreateGlossary *DatazonePolicyGrantDetailCreateGlossary `field:"optional" json:"createGlossary" yaml:"createGlossary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_policy_grant#create_project DatazonePolicyGrant#create_project}.
	CreateProject *DatazonePolicyGrantDetailCreateProject `field:"optional" json:"createProject" yaml:"createProject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_policy_grant#create_project_from_project_profile DatazonePolicyGrant#create_project_from_project_profile}.
	CreateProjectFromProjectProfile *DatazonePolicyGrantDetailCreateProjectFromProjectProfile `field:"optional" json:"createProjectFromProjectProfile" yaml:"createProjectFromProjectProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_policy_grant#delegate_create_environment_profile DatazonePolicyGrant#delegate_create_environment_profile}.
	DelegateCreateEnvironmentProfile *string `field:"optional" json:"delegateCreateEnvironmentProfile" yaml:"delegateCreateEnvironmentProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_policy_grant#override_domain_unit_owners DatazonePolicyGrant#override_domain_unit_owners}.
	OverrideDomainUnitOwners *DatazonePolicyGrantDetailOverrideDomainUnitOwners `field:"optional" json:"overrideDomainUnitOwners" yaml:"overrideDomainUnitOwners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_policy_grant#override_project_owners DatazonePolicyGrant#override_project_owners}.
	OverrideProjectOwners *DatazonePolicyGrantDetailOverrideProjectOwners `field:"optional" json:"overrideProjectOwners" yaml:"overrideProjectOwners"`
}

