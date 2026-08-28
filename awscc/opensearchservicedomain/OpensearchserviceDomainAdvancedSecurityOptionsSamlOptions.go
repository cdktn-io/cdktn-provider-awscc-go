// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchservicedomain


type OpensearchserviceDomainAdvancedSecurityOptionsSamlOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#enabled OpensearchserviceDomain#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#idp OpensearchserviceDomain#idp}.
	Idp *OpensearchserviceDomainAdvancedSecurityOptionsSamlOptionsIdp `field:"optional" json:"idp" yaml:"idp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#master_backend_role OpensearchserviceDomain#master_backend_role}.
	MasterBackendRole *string `field:"optional" json:"masterBackendRole" yaml:"masterBackendRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#master_user_name OpensearchserviceDomain#master_user_name}.
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#roles_key OpensearchserviceDomain#roles_key}.
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#session_timeout_minutes OpensearchserviceDomain#session_timeout_minutes}.
	SessionTimeoutMinutes *float64 `field:"optional" json:"sessionTimeoutMinutes" yaml:"sessionTimeoutMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#subject_key OpensearchserviceDomain#subject_key}.
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

