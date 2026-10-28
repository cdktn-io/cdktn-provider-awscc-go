// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagegatewaynfsfileshare

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StoragegatewayNfsFileShareConfig struct {
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
	// The ARN of the S3 File Gateway that hosts the share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#gateway_arn StoragegatewayNfsFileShare#gateway_arn}
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// The ARN of the S3 bucket, or bucket prefix, the share exposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#location_arn StoragegatewayNfsFileShare#location_arn}
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
	// The ARN of the IAM role the gateway assumes when it accesses the S3 location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#role StoragegatewayNfsFileShare#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The ARN of the CloudWatch log group that receives the share's audit logs.
	//
	// Removing this property from a template detaches audit logging: the update sends the empty string, which the service treats as the reset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#audit_destination_arn StoragegatewayNfsFileShare#audit_destination_arn}
	AuditDestinationArn *string `field:"optional" json:"auditDestinationArn" yaml:"auditDestinationArn"`
	// The Region of the S3 bucket, required when the bucket is reached through a VPC endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#bucket_region StoragegatewayNfsFileShare#bucket_region}
	BucketRegion *string `field:"optional" json:"bucketRegion" yaml:"bucketRegion"`
	// Refresh cache information for the file share.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#cache_attributes StoragegatewayNfsFileShare#cache_attributes}
	CacheAttributes *StoragegatewayNfsFileShareCacheAttributes `field:"optional" json:"cacheAttributes" yaml:"cacheAttributes"`
	// The IPv4 or IPv6 addresses and CIDR blocks permitted to mount the NFS share.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#client_list StoragegatewayNfsFileShare#client_list}
	ClientList *[]*string `field:"optional" json:"clientList" yaml:"clientList"`
	// The S3 storage class for new objects.
	//
	// The service accepts S3_STANDARD, S3_INTELLIGENT_TIERING, S3_STANDARD_IA and S3_ONEZONE_IA; it is a length-bounded string in the API model rather than an enumeration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#default_storage_class StoragegatewayNfsFileShare#default_storage_class}
	DefaultStorageClass *string `field:"optional" json:"defaultStorageClass" yaml:"defaultStorageClass"`
	// The server-side encryption the gateway applies to objects it writes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#encryption_type StoragegatewayNfsFileShare#encryption_type}
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The name of the file share, which defaults to the bucket name when omitted.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#file_share_name StoragegatewayNfsFileShare#file_share_name}
	FileShareName *string `field:"optional" json:"fileShareName" yaml:"fileShareName"`
	// Whether the gateway infers the MIME type of uploaded objects from their extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#guess_mime_type_enabled StoragegatewayNfsFileShare#guess_mime_type_enabled}
	GuessMimeTypeEnabled interface{} `field:"optional" json:"guessMimeTypeEnabled" yaml:"guessMimeTypeEnabled"`
	// Deprecated by the service in favour of EncryptionType; true selects KMS encryption. Retained because it remains readable and settable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#kms_encrypted StoragegatewayNfsFileShare#kms_encrypted}
	KmsEncrypted interface{} `field:"optional" json:"kmsEncrypted" yaml:"kmsEncrypted"`
	// The KMS key used for encryption, when EncryptionType selects KMS.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#kms_key StoragegatewayNfsFileShare#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Default values for the file and directory permissions of an NFS share.
	//
	// Removing this property from a template leaves the live values unchanged rather than clearing them, because the service has no reset for it: a create that omits the object leaves it absent, but the update path merges member by member keyed on the same internal sentinel that reports absence, so no request can restore that state. Set the members you want rather than relying on removal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#nfs_file_share_defaults StoragegatewayNfsFileShare#nfs_file_share_defaults}
	NfsFileShareDefaults *StoragegatewayNfsFileShareNfsFileShareDefaults `field:"optional" json:"nfsFileShareDefaults" yaml:"nfsFileShareDefaults"`
	// The upload-notification settling period, as a JSON document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#notification_policy StoragegatewayNfsFileShare#notification_policy}
	NotificationPolicy *string `field:"optional" json:"notificationPolicy" yaml:"notificationPolicy"`
	// The canned S3 ACL applied to objects the gateway writes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#object_acl StoragegatewayNfsFileShare#object_acl}
	ObjectAcl *string `field:"optional" json:"objectAcl" yaml:"objectAcl"`
	// Whether clients may write to the share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#read_only StoragegatewayNfsFileShare#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Whether the bucket is a Requester Pays bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#requester_pays StoragegatewayNfsFileShare#requester_pays}
	RequesterPays interface{} `field:"optional" json:"requesterPays" yaml:"requesterPays"`
	// How the NFS share maps client user IDs.
	//
	// The service accepts RootSquash, NoSquash and AllSquash; it is a length-bounded string in the API model rather than an enumeration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#squash StoragegatewayNfsFileShare#squash}
	Squash *string `field:"optional" json:"squash" yaml:"squash"`
	// Tags assigned to the file share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#tags StoragegatewayNfsFileShare#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The DNS name of the S3 VPC endpoint the gateway uses to reach the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_nfs_file_share#vpc_endpoint_dns_name StoragegatewayNfsFileShare#vpc_endpoint_dns_name}
	VpcEndpointDnsName *string `field:"optional" json:"vpcEndpointDnsName" yaml:"vpcEndpointDnsName"`
}

