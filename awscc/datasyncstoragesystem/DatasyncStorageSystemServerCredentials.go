// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasyncstoragesystem


type DatasyncStorageSystemServerCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_storage_system#password DatasyncStorageSystem#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_storage_system#username DatasyncStorageSystem#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

