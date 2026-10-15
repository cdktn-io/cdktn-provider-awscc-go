// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagegatewaynfsfileshare


type StoragegatewayNfsFileShareNfsFileShareDefaults struct {
	// The Unix directory mode in the form nnnn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/storagegateway_nfs_file_share#directory_mode StoragegatewayNfsFileShare#directory_mode}
	DirectoryMode *string `field:"optional" json:"directoryMode" yaml:"directoryMode"`
	// The Unix file mode in the form nnnn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/storagegateway_nfs_file_share#file_mode StoragegatewayNfsFileShare#file_mode}
	FileMode *string `field:"optional" json:"fileMode" yaml:"fileMode"`
	// The default group ID for files in the file share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/storagegateway_nfs_file_share#group_id StoragegatewayNfsFileShare#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// The default owner ID for files in the file share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/storagegateway_nfs_file_share#owner_id StoragegatewayNfsFileShare#owner_id}
	OwnerId *float64 `field:"optional" json:"ownerId" yaml:"ownerId"`
}

