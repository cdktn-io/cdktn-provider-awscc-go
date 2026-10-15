// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxstoragevirtualmachine

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FsxStorageVirtualMachineConfig struct {
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
	// Specifies the FSx for ONTAP file system on which to create the SVM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/fsx_storage_virtual_machine#file_system_id FsxStorageVirtualMachine#file_system_id}
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// The name of the SVM. Pattern: `^[^\u0000\u0085\u2028\u2029\r\n]{1,47}$`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/fsx_storage_virtual_machine#name FsxStorageVirtualMachine#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Describes the Microsoft Active Directory configuration to which the SVM is joined, if applicable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/fsx_storage_virtual_machine#active_directory_configuration FsxStorageVirtualMachine#active_directory_configuration}
	ActiveDirectoryConfiguration *FsxStorageVirtualMachineActiveDirectoryConfiguration `field:"optional" json:"activeDirectoryConfiguration" yaml:"activeDirectoryConfiguration"`
	// The security style of the root volume of the SVM. Allowed values: UNIX | NTFS | MIXED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/fsx_storage_virtual_machine#root_volume_security_style FsxStorageVirtualMachine#root_volume_security_style}
	RootVolumeSecurityStyle *string `field:"optional" json:"rootVolumeSecurityStyle" yaml:"rootVolumeSecurityStyle"`
	// Specifies the password to use when logging on to the SVM using a secure shell (SSH) connection to the SVM's management endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/fsx_storage_virtual_machine#svm_admin_password FsxStorageVirtualMachine#svm_admin_password}
	SvmAdminPassword *string `field:"optional" json:"svmAdminPassword" yaml:"svmAdminPassword"`
	// A list of Tag values, with a maximum of 50 elements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/fsx_storage_virtual_machine#tags FsxStorageVirtualMachine#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

