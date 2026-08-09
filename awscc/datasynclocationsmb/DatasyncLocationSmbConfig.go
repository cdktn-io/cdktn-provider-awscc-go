// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationsmb

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatasyncLocationSmbConfig struct {
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
	// The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#agent_arns DatasyncLocationSmb#agent_arns}
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// The authentication mode used to determine identity of user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#authentication_type DatasyncLocationSmb#authentication_type}
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Specifies configuration information for a DataSync-managed secret, such as a password or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed AWS KMS key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#cmk_secret_config DatasyncLocationSmb#cmk_secret_config}
	CmkSecretConfig *DatasyncLocationSmbCmkSecretConfig `field:"optional" json:"cmkSecretConfig" yaml:"cmkSecretConfig"`
	// Specifies configuration information for a customer-managed secret, such as a password or set of credentials that DataSync uses to access a specific transfer location, and an IAM role that DataSync can assume and access the customer-managed secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#custom_secret_config DatasyncLocationSmb#custom_secret_config}
	CustomSecretConfig *DatasyncLocationSmbCustomSecretConfig `field:"optional" json:"customSecretConfig" yaml:"customSecretConfig"`
	// Specifies the IPv4 addresses for the DNS servers that your SMB file server belongs to.
	//
	// This parameter applies only if AuthenticationType is set to KERBEROS. If you have multiple domains in your environment, configuring this parameter makes sure that DataSync connects to the right SMB file server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#dns_ip_addresses DatasyncLocationSmb#dns_ip_addresses}
	DnsIpAddresses *[]*string `field:"optional" json:"dnsIpAddresses" yaml:"dnsIpAddresses"`
	// The name of the Windows domain that the SMB server belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#domain DatasyncLocationSmb#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The Base64 string representation of the Keytab file.
	//
	// Specifies your Kerberos key table (keytab) file, which includes mappings between your service principal name (SPN) and encryption keys. To avoid task execution errors, make sure that the SPN in the keytab file matches exactly what you specify for KerberosPrincipal and in your krb5.conf file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#kerberos_keytab DatasyncLocationSmb#kerberos_keytab}
	KerberosKeytab *string `field:"optional" json:"kerberosKeytab" yaml:"kerberosKeytab"`
	// The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket. Specifies a Kerberos configuration file (krb5.conf) that defines your Kerberos realm configuration. To avoid task execution errors, make sure that the service principal name (SPN) in the krb5.conf file matches exactly what you specify for KerberosPrincipal and in your keytab file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#kerberos_krb_5_conf DatasyncLocationSmb#kerberos_krb_5_conf}
	KerberosKrb5Conf *string `field:"optional" json:"kerberosKrb5Conf" yaml:"kerberosKrb5Conf"`
	// Specifies a service principal name (SPN), which is an identity in your Kerberos realm that has permission to access the files, folders, and file metadata in your SMB file server.
	//
	// SPNs are case sensitive and must include a prepended cifs/. For example, an SPN might look like cifs/kerberosuser@EXAMPLE.COM. Your task execution will fail if the SPN that you provide for this parameter doesn't match exactly what's in your keytab or krb5.conf files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#kerberos_principal DatasyncLocationSmb#kerberos_principal}
	KerberosPrincipal *string `field:"optional" json:"kerberosPrincipal" yaml:"kerberosPrincipal"`
	// The mount options used by DataSync to access the SMB server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#mount_options DatasyncLocationSmb#mount_options}
	MountOptions *DatasyncLocationSmbMountOptions `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#password DatasyncLocationSmb#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The name of the SMB server.
	//
	// This value is the IP address or Domain Name Service (DNS) name of the SMB server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#server_hostname DatasyncLocationSmb#server_hostname}
	ServerHostname *string `field:"optional" json:"serverHostname" yaml:"serverHostname"`
	// The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#subdirectory DatasyncLocationSmb#subdirectory}
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#tags DatasyncLocationSmb#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The user who can mount the share, has the permissions to access files and folders in the SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_smb#user DatasyncLocationSmb#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

