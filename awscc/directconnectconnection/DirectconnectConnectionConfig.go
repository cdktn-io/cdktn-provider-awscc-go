// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directconnectconnection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DirectconnectConnectionConfig struct {
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
	// The bandwidth of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_connection#bandwidth DirectconnectConnection#bandwidth}
	Bandwidth *string `field:"required" json:"bandwidth" yaml:"bandwidth"`
	// The name of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_connection#connection_name DirectconnectConnection#connection_name}
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// The location of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_connection#location DirectconnectConnection#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID or ARN of the LAG to associate the connection with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_connection#lag_id DirectconnectConnection#lag_id}
	LagId *string `field:"optional" json:"lagId" yaml:"lagId"`
	// The name of the service provider associated with the requested connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_connection#provider_name DirectconnectConnection#provider_name}
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// Indicates whether you want the connection to support MAC Security (MACsec).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_connection#request_mac_sec DirectconnectConnection#request_mac_sec}
	RequestMacSec interface{} `field:"optional" json:"requestMacSec" yaml:"requestMacSec"`
	// The tags associated with the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_connection#tags DirectconnectConnection#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

