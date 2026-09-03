// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferagreement

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TransferAgreementConfig struct {
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
	// Specifies the access role for the agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_agreement#access_role TransferAgreement#access_role}
	AccessRole *string `field:"required" json:"accessRole" yaml:"accessRole"`
	// A unique identifier for the local profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_agreement#local_profile_id TransferAgreement#local_profile_id}
	LocalProfileId *string `field:"required" json:"localProfileId" yaml:"localProfileId"`
	// A unique identifier for the partner profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_agreement#partner_profile_id TransferAgreement#partner_profile_id}
	PartnerProfileId *string `field:"required" json:"partnerProfileId" yaml:"partnerProfileId"`
	// A unique identifier for the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_agreement#server_id TransferAgreement#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Specifies the base directory for the agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_agreement#base_directory TransferAgreement#base_directory}
	BaseDirectory *string `field:"optional" json:"baseDirectory" yaml:"baseDirectory"`
	// Specifies a separate directory for each type of file to store for an AS2 message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_agreement#custom_directories TransferAgreement#custom_directories}
	CustomDirectories *TransferAgreementCustomDirectories `field:"optional" json:"customDirectories" yaml:"customDirectories"`
	// A textual description for the agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_agreement#description TransferAgreement#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether to enforce an AS2 message is signed for this agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_agreement#enforce_message_signing TransferAgreement#enforce_message_signing}
	EnforceMessageSigning *string `field:"optional" json:"enforceMessageSigning" yaml:"enforceMessageSigning"`
	// Specifies whether to preserve the filename received for this agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_agreement#preserve_filename TransferAgreement#preserve_filename}
	PreserveFilename *string `field:"optional" json:"preserveFilename" yaml:"preserveFilename"`
	// Specifies the status of the agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_agreement#status TransferAgreement#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Key-value pairs that can be used to group and search for agreements.
	//
	// Tags are metadata attached to agreements for any purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_agreement#tags TransferAgreement#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

