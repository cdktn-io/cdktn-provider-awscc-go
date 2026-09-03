// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesspartneraccount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotwirelessPartnerAccountConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Whether the partner account is linked to the AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_partner_account#account_linked IotwirelessPartnerAccount#account_linked}
	AccountLinked interface{} `field:"optional" json:"accountLinked" yaml:"accountLinked"`
	// The partner account ID to disassociate from the AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_partner_account#partner_account_id IotwirelessPartnerAccount#partner_account_id}
	PartnerAccountId *string `field:"optional" json:"partnerAccountId" yaml:"partnerAccountId"`
	// The partner type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_partner_account#partner_type IotwirelessPartnerAccount#partner_type}
	PartnerType *string `field:"optional" json:"partnerType" yaml:"partnerType"`
	// The Sidewalk account credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_partner_account#sidewalk IotwirelessPartnerAccount#sidewalk}
	Sidewalk *IotwirelessPartnerAccountSidewalk `field:"optional" json:"sidewalk" yaml:"sidewalk"`
	// The Sidewalk account credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_partner_account#sidewalk_response IotwirelessPartnerAccount#sidewalk_response}
	SidewalkResponse *IotwirelessPartnerAccountSidewalkResponse `field:"optional" json:"sidewalkResponse" yaml:"sidewalkResponse"`
	// The Sidewalk account credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_partner_account#sidewalk_update IotwirelessPartnerAccount#sidewalk_update}
	SidewalkUpdate *IotwirelessPartnerAccountSidewalkUpdate `field:"optional" json:"sidewalkUpdate" yaml:"sidewalkUpdate"`
	// A list of key-value pairs that contain metadata for the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_partner_account#tags IotwirelessPartnerAccount#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

