// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionConfigurations struct {
	// The classification of the connection configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_connection#classification DatazoneConnection#classification}
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Property Map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_connection#properties DatazoneConnection#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

