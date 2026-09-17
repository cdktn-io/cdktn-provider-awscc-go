// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationfsxontap


type DatasyncLocationFsxOntapProtocolSmb struct {
	// Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed AWS KMS key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_fsx_ontap#cmk_secret_config DatasyncLocationFsxOntap#cmk_secret_config}
	CmkSecretConfig *DatasyncLocationFsxOntapProtocolSmbCmkSecretConfig `field:"optional" json:"cmkSecretConfig" yaml:"cmkSecretConfig"`
	// Specifies configuration information for a customer-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and an IAM role that DataSync can assume and access the customer-managed secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_fsx_ontap#custom_secret_config DatasyncLocationFsxOntap#custom_secret_config}
	CustomSecretConfig *DatasyncLocationFsxOntapProtocolSmbCustomSecretConfig `field:"optional" json:"customSecretConfig" yaml:"customSecretConfig"`
	// The name of the Windows domain that the SMB server belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_fsx_ontap#domain DatasyncLocationFsxOntap#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The mount options used by DataSync to access the SMB server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_fsx_ontap#mount_options DatasyncLocationFsxOntap#mount_options}
	MountOptions *DatasyncLocationFsxOntapProtocolSmbMountOptions `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_fsx_ontap#password DatasyncLocationFsxOntap#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The user who can mount the share, has the permissions to access files and folders in the SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_fsx_ontap#user DatasyncLocationFsxOntap#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

