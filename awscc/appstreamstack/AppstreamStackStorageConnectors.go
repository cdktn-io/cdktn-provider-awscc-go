// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamstack


type AppstreamStackStorageConnectors struct {
	// The type of storage connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appstream_stack#connector_type AppstreamStack#connector_type}
	ConnectorType *string `field:"optional" json:"connectorType" yaml:"connectorType"`
	// The names of the domains for the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appstream_stack#domains AppstreamStack#domains}
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// The ARN of the storage connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appstream_stack#resource_identifier AppstreamStack#resource_identifier}
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

