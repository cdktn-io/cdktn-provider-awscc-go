// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagegatewaysmbfileshare

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StoragegatewaySmbFileShareConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#gateway_arn StoragegatewaySmbFileShare#gateway_arn}
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// The ARN of the S3 bucket, or bucket prefix, the share exposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#location_arn StoragegatewaySmbFileShare#location_arn}
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
	// The ARN of the IAM role the gateway assumes when it accesses the S3 location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#role StoragegatewaySmbFileShare#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Whether clients see only the files and folders they may access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#access_based_enumeration StoragegatewaySmbFileShare#access_based_enumeration}
	AccessBasedEnumeration interface{} `field:"optional" json:"accessBasedEnumeration" yaml:"accessBasedEnumeration"`
	// Users and groups granted administrative privileges on the SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#admin_user_list StoragegatewaySmbFileShare#admin_user_list}
	AdminUserList *[]*string `field:"optional" json:"adminUserList" yaml:"adminUserList"`
	// The ARN of the CloudWatch log group that receives the share's audit logs.
	//
	// Removing this property from a template detaches audit logging: the update sends the empty string, which the service treats as the reset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#audit_destination_arn StoragegatewaySmbFileShare#audit_destination_arn}
	AuditDestinationArn *string `field:"optional" json:"auditDestinationArn" yaml:"auditDestinationArn"`
	// The authentication the SMB share uses.
	//
	// The service accepts ActiveDirectory and GuestAccess; it is a length-bounded string in the API model rather than an enumeration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#authentication StoragegatewaySmbFileShare#authentication}
	Authentication *string `field:"optional" json:"authentication" yaml:"authentication"`
	// The Region of the S3 bucket, required when the bucket is reached through a VPC endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#bucket_region StoragegatewaySmbFileShare#bucket_region}
	BucketRegion *string `field:"optional" json:"bucketRegion" yaml:"bucketRegion"`
	// Refresh cache information for the file share.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#cache_attributes StoragegatewaySmbFileShare#cache_attributes}
	CacheAttributes *StoragegatewaySmbFileShareCacheAttributes `field:"optional" json:"cacheAttributes" yaml:"cacheAttributes"`
	// Whether the SMB share is case sensitive or defers to the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#case_sensitivity StoragegatewaySmbFileShare#case_sensitivity}
	CaseSensitivity *string `field:"optional" json:"caseSensitivity" yaml:"caseSensitivity"`
	// The S3 storage class for new objects.
	//
	// The service accepts S3_STANDARD, S3_INTELLIGENT_TIERING, S3_STANDARD_IA and S3_ONEZONE_IA; it is a length-bounded string in the API model rather than an enumeration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#default_storage_class StoragegatewaySmbFileShare#default_storage_class}
	DefaultStorageClass *string `field:"optional" json:"defaultStorageClass" yaml:"defaultStorageClass"`
	// The server-side encryption the gateway applies to objects it writes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#encryption_type StoragegatewaySmbFileShare#encryption_type}
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The name of the file share, which defaults to the bucket name when omitted.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#file_share_name StoragegatewaySmbFileShare#file_share_name}
	FileShareName *string `field:"optional" json:"fileShareName" yaml:"fileShareName"`
	// Whether the gateway infers the MIME type of uploaded objects from their extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#guess_mime_type_enabled StoragegatewaySmbFileShare#guess_mime_type_enabled}
	GuessMimeTypeEnabled interface{} `field:"optional" json:"guessMimeTypeEnabled" yaml:"guessMimeTypeEnabled"`
	// Users and groups denied access to the SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#invalid_user_list StoragegatewaySmbFileShare#invalid_user_list}
	InvalidUserList *[]*string `field:"optional" json:"invalidUserList" yaml:"invalidUserList"`
	// Deprecated by the service in favour of EncryptionType; true selects KMS encryption. Retained because it remains readable and settable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#kms_encrypted StoragegatewaySmbFileShare#kms_encrypted}
	KmsEncrypted interface{} `field:"optional" json:"kmsEncrypted" yaml:"kmsEncrypted"`
	// The KMS key used for encryption, when EncryptionType selects KMS.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#kms_key StoragegatewaySmbFileShare#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The upload-notification settling period, as a JSON document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#notification_policy StoragegatewaySmbFileShare#notification_policy}
	NotificationPolicy *string `field:"optional" json:"notificationPolicy" yaml:"notificationPolicy"`
	// The canned S3 ACL applied to objects the gateway writes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#object_acl StoragegatewaySmbFileShare#object_acl}
	ObjectAcl *string `field:"optional" json:"objectAcl" yaml:"objectAcl"`
	// Whether the SMB share grants opportunistic locks.
	//
	// Removing this property from a template leaves the live value unchanged rather than restoring the created default, because that default depends on the gateway's software version -- opportunistic locking is enabled by default only on appliances at 1.19.0.0 or above -- and the version is not part of this resource. Set the value explicitly rather than relying on removal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#oplocks_enabled StoragegatewaySmbFileShare#oplocks_enabled}
	OplocksEnabled interface{} `field:"optional" json:"oplocksEnabled" yaml:"oplocksEnabled"`
	// Whether clients may write to the share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#read_only StoragegatewaySmbFileShare#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Whether the bucket is a Requester Pays bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#requester_pays StoragegatewaySmbFileShare#requester_pays}
	RequesterPays interface{} `field:"optional" json:"requesterPays" yaml:"requesterPays"`
	// Whether the SMB share uses Windows ACLs rather than POSIX permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#smbacl_enabled StoragegatewaySmbFileShare#smbacl_enabled}
	SmbaclEnabled interface{} `field:"optional" json:"smbaclEnabled" yaml:"smbaclEnabled"`
	// Tags assigned to the file share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#tags StoragegatewaySmbFileShare#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Users and groups permitted to access the SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#valid_user_list StoragegatewaySmbFileShare#valid_user_list}
	ValidUserList *[]*string `field:"optional" json:"validUserList" yaml:"validUserList"`
	// The DNS name of the S3 VPC endpoint the gateway uses to reach the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/storagegateway_smb_file_share#vpc_endpoint_dns_name StoragegatewaySmbFileShare#vpc_endpoint_dns_name}
	VpcEndpointDnsName *string `field:"optional" json:"vpcEndpointDnsName" yaml:"vpcEndpointDnsName"`
}

