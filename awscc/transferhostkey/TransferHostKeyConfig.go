// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferhostkey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TransferHostKeyConfig struct {
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
	// The identifier of the server that contains the host key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/transfer_host_key#server_id TransferHostKey#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// The text description for this host key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/transfer_host_key#description TransferHostKey#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The private key portion of an SSH key pair. Transfer Family accepts RSA, ECDSA, and ED25519 keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/transfer_host_key#host_key_body TransferHostKey#host_key_body}
	HostKeyBody *string `field:"optional" json:"hostKeyBody" yaml:"hostKeyBody"`
	// Key-value pairs that can be used to group and search for host keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/transfer_host_key#tags TransferHostKey#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

