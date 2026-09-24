// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationobjectstorage

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatasyncLocationObjectStorageConfig struct {
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
	// Optional. The access key is used if credentials are required to access the self-managed object storage server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#access_key DatasyncLocationObjectStorage#access_key}
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Specifies the Amazon Resource Names (ARNs) of the DataSync agents that can connect with your object storage system.
	//
	// If you are setting up an agentless cross-cloud transfer, you do not need to specify a value for this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#agent_arns DatasyncLocationObjectStorage#agent_arns}
	AgentArns *[]*string `field:"optional" json:"agentArns" yaml:"agentArns"`
	// The name of the bucket on the self-managed object storage server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#bucket_name DatasyncLocationObjectStorage#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed AWS KMS key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#cmk_secret_config DatasyncLocationObjectStorage#cmk_secret_config}
	CmkSecretConfig *DatasyncLocationObjectStorageCmkSecretConfig `field:"optional" json:"cmkSecretConfig" yaml:"cmkSecretConfig"`
	// Specifies configuration information for a customer-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and an IAM role that DataSync can assume and access the customer-managed secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#custom_secret_config DatasyncLocationObjectStorage#custom_secret_config}
	CustomSecretConfig *DatasyncLocationObjectStorageCustomSecretConfig `field:"optional" json:"customSecretConfig" yaml:"customSecretConfig"`
	// Specifies the identity federation configuration that DataSync uses to access your object storage location using an OpenID Connect (OIDC) token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#federated_identity DatasyncLocationObjectStorage#federated_identity}
	FederatedIdentity *DatasyncLocationObjectStorageFederatedIdentity `field:"optional" json:"federatedIdentity" yaml:"federatedIdentity"`
	// Optional. The secret key is used if credentials are required to access the self-managed object storage server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#secret_key DatasyncLocationObjectStorage#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// X.509 PEM content containing a certificate authority or chain to trust.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#server_certificate DatasyncLocationObjectStorage#server_certificate}
	ServerCertificate *string `field:"optional" json:"serverCertificate" yaml:"serverCertificate"`
	// The name of the self-managed object storage server.
	//
	// This value is the IP address or Domain Name Service (DNS) name of the object storage server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#server_hostname DatasyncLocationObjectStorage#server_hostname}
	ServerHostname *string `field:"optional" json:"serverHostname" yaml:"serverHostname"`
	// The port that your self-managed server accepts inbound network traffic on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#server_port DatasyncLocationObjectStorage#server_port}
	ServerPort *float64 `field:"optional" json:"serverPort" yaml:"serverPort"`
	// The protocol that the object storage server uses to communicate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#server_protocol DatasyncLocationObjectStorage#server_protocol}
	ServerProtocol *string `field:"optional" json:"serverProtocol" yaml:"serverProtocol"`
	// The subdirectory in the self-managed object storage server that is used to read data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#subdirectory DatasyncLocationObjectStorage#subdirectory}
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#tags DatasyncLocationObjectStorage#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

