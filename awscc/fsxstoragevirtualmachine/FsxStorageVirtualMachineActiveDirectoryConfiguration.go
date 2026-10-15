// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxstoragevirtualmachine


type FsxStorageVirtualMachineActiveDirectoryConfiguration struct {
	// The NetBIOS name of the Active Directory computer object that will be created for your SVM.
	//
	// Minimum values: 1. Maximum values: 15. Pattern: `^[^\u0000\u0085\u2028\u2029\r\n]{1,255}$`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/fsx_storage_virtual_machine#net_bios_name FsxStorageVirtualMachine#net_bios_name}
	NetBiosName *string `field:"optional" json:"netBiosName" yaml:"netBiosName"`
	// The configuration that Amazon FSx uses to join the ONTAP storage virtual machine (SVM) to your self-managed (including on-premises) Microsoft Active Directory directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/fsx_storage_virtual_machine#self_managed_active_directory_configuration FsxStorageVirtualMachine#self_managed_active_directory_configuration}
	SelfManagedActiveDirectoryConfiguration *FsxStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration `field:"optional" json:"selfManagedActiveDirectoryConfiguration" yaml:"selfManagedActiveDirectoryConfiguration"`
}

