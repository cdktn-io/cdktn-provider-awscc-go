// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxstoragevirtualmachine


type FsxStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration struct {
	// A list of IP addresses of DNS servers or domain controllers in the self-managed AD directory.
	//
	// Minimum: 1. Maximum: 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/fsx_storage_virtual_machine#dns_ips FsxStorageVirtualMachine#dns_ips}
	DnsIps *[]*string `field:"optional" json:"dnsIps" yaml:"dnsIps"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret containing the self-managed Active Directory domain join service account credentials.
	//
	// Minimum: 64. Maximum: 1024.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/fsx_storage_virtual_machine#domain_join_service_account_secret FsxStorageVirtualMachine#domain_join_service_account_secret}
	DomainJoinServiceAccountSecret *string `field:"optional" json:"domainJoinServiceAccountSecret" yaml:"domainJoinServiceAccountSecret"`
	// The fully qualified domain name of the self-managed AD directory, such as corp.example.com. Minimum: 1. Maximum: 255. Pattern: `^[^\u0000\u0085\u2028\u2029\r\n]{1,255}$`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/fsx_storage_virtual_machine#domain_name FsxStorageVirtualMachine#domain_name}
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The name of the domain group whose members are granted administrative privileges for the file system.
	//
	// Minimum: 1. Maximum: 256. Pattern: `^[^\u0000\u0085\u2028\u2029\r\n]{1,256}$`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/fsx_storage_virtual_machine#file_system_administrators_group FsxStorageVirtualMachine#file_system_administrators_group}
	FileSystemAdministratorsGroup *string `field:"optional" json:"fileSystemAdministratorsGroup" yaml:"fileSystemAdministratorsGroup"`
	// The fully qualified distinguished name of the organizational unit within your self-managed AD directory.
	//
	// Minimum: 1. Maximum: 2000. Pattern: `^[^\u0000\u0085\u2028\u2029\r\n]{1,2000}$`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/fsx_storage_virtual_machine#organizational_unit_distinguished_name FsxStorageVirtualMachine#organizational_unit_distinguished_name}
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
	// The password for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	//
	// Minimum: 1. Maximum: 256. Pattern: `^.{1,256}$`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/fsx_storage_virtual_machine#password FsxStorageVirtualMachine#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The user name for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	//
	// Minimum: 1. Maximum: 256. Pattern: `^[^\u0000\u0085\u2028\u2029\r\n]{1,256}$`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/fsx_storage_virtual_machine#user_name FsxStorageVirtualMachine#user_name}
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

