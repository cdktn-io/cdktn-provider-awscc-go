// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionspolicy


type VerifiedpermissionsPolicyDefinitionTemplateLinked struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/verifiedpermissions_policy#policy_template_id VerifiedpermissionsPolicy#policy_template_id}.
	PolicyTemplateId *string `field:"optional" json:"policyTemplateId" yaml:"policyTemplateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/verifiedpermissions_policy#principal VerifiedpermissionsPolicy#principal}.
	Principal *VerifiedpermissionsPolicyDefinitionTemplateLinkedPrincipal `field:"optional" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/verifiedpermissions_policy#resource VerifiedpermissionsPolicy#resource}.
	Resource *VerifiedpermissionsPolicyDefinitionTemplateLinkedResource `field:"optional" json:"resource" yaml:"resource"`
}

