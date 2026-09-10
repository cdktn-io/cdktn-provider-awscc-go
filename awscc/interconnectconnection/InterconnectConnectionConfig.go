// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package interconnectconnection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type InterconnectConnectionConfig struct {
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
	// The logical attachment point in your AWS network where the managed connection will be connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/interconnect_connection#attach_point InterconnectConnection#attach_point}
	AttachPoint *InterconnectConnectionAttachPoint `field:"required" json:"attachPoint" yaml:"attachPoint"`
	// The activation key for accepting a connection proposal from a partner CSP. Mutually exclusive with EnvironmentId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/interconnect_connection#activation_key InterconnectConnection#activation_key}
	ActivationKey *string `field:"optional" json:"activationKey" yaml:"activationKey"`
	// The bandwidth of the connection (e.g., 50Mbps, 1Gbps). Required when creating a connection through AWS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/interconnect_connection#bandwidth InterconnectConnection#bandwidth}
	Bandwidth *string `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// A description of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/interconnect_connection#description InterconnectConnection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the environment for the connection. Required when creating a connection through AWS. Mutually exclusive with ActivationKey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/interconnect_connection#environment_id InterconnectConnection#environment_id}
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
	// The remote account identifier for the connection. Required when creating a connection through AWS. Replaces RemoteOwnerAccount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/interconnect_connection#remote_account InterconnectConnection#remote_account}
	RemoteAccount *InterconnectConnectionRemoteAccount `field:"optional" json:"remoteAccount" yaml:"remoteAccount"`
	// Deprecated. Use RemoteAccount instead. The account ID of the remote owner. Required when creating a connection through AWS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/interconnect_connection#remote_owner_account InterconnectConnection#remote_owner_account}
	RemoteOwnerAccount *string `field:"optional" json:"remoteOwnerAccount" yaml:"remoteOwnerAccount"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/interconnect_connection#tags InterconnectConnection#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

