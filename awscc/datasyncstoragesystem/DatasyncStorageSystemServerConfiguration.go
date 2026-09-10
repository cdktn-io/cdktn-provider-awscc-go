// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasyncstoragesystem


type DatasyncStorageSystemServerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datasync_storage_system#server_hostname DatasyncStorageSystem#server_hostname}.
	ServerHostname *string `field:"required" json:"serverHostname" yaml:"serverHostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datasync_storage_system#server_port DatasyncStorageSystem#server_port}.
	ServerPort *float64 `field:"optional" json:"serverPort" yaml:"serverPort"`
}

