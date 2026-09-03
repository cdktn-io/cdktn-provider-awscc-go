// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferconnector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TransferConnectorConfig struct {
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
	// Specifies the access role for the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_connector#access_role TransferConnector#access_role}
	AccessRole *string `field:"required" json:"accessRole" yaml:"accessRole"`
	// Configuration for an AS2 connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_connector#as_2_config TransferConnector#as_2_config}
	As2Config *TransferConnectorAs2Config `field:"optional" json:"as2Config" yaml:"as2Config"`
	// Egress configuration for the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_connector#egress_config TransferConnector#egress_config}
	EgressConfig *TransferConnectorEgressConfig `field:"optional" json:"egressConfig" yaml:"egressConfig"`
	// Specifies the egress type for the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_connector#egress_type TransferConnector#egress_type}
	EgressType *string `field:"optional" json:"egressType" yaml:"egressType"`
	// IP address type for Connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_connector#ip_address_type TransferConnector#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Specifies the logging role for the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_connector#logging_role TransferConnector#logging_role}
	LoggingRole *string `field:"optional" json:"loggingRole" yaml:"loggingRole"`
	// Security policy for SFTP Connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_connector#security_policy_name TransferConnector#security_policy_name}
	SecurityPolicyName *string `field:"optional" json:"securityPolicyName" yaml:"securityPolicyName"`
	// Configuration for an SFTP connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_connector#sftp_config TransferConnector#sftp_config}
	SftpConfig *TransferConnectorSftpConfig `field:"optional" json:"sftpConfig" yaml:"sftpConfig"`
	// Key-value pairs that can be used to group and search for connectors.
	//
	// Tags are metadata attached to connectors for any purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_connector#tags TransferConnector#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// URL for Connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/transfer_connector#url TransferConnector#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

