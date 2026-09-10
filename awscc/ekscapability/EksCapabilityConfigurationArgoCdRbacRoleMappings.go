// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscapability


type EksCapabilityConfigurationArgoCdRbacRoleMappings struct {
	// A list of IAM Identity Center identities (users or groups) that should be assigned this Argo CD role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_capability#identities EksCapability#identities}
	Identities interface{} `field:"optional" json:"identities" yaml:"identities"`
	// The Argo CD role to assign.
	//
	// Valid values are: ADMIN (full administrative access to Argo CD), EDITOR (edit access to Argo CD resources), or VIEWER (read-only access to Argo CD resources).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_capability#role EksCapability#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

