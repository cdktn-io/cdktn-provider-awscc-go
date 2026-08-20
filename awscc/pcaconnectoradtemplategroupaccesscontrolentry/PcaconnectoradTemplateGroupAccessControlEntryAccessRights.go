// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplategroupaccesscontrolentry


type PcaconnectoradTemplateGroupAccessControlEntryAccessRights struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcaconnectorad_template_group_access_control_entry#auto_enroll PcaconnectoradTemplateGroupAccessControlEntry#auto_enroll}.
	AutoEnroll *string `field:"optional" json:"autoEnroll" yaml:"autoEnroll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcaconnectorad_template_group_access_control_entry#enroll PcaconnectoradTemplateGroupAccessControlEntry#enroll}.
	Enroll *string `field:"optional" json:"enroll" yaml:"enroll"`
}

