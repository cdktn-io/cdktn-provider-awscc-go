// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagegatewaysmbfileshare


type StoragegatewaySmbFileShareCacheAttributes struct {
	// How long the directory contents cached by the gateway are considered valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/storagegateway_smb_file_share#cache_stale_timeout_in_seconds StoragegatewaySmbFileShare#cache_stale_timeout_in_seconds}
	CacheStaleTimeoutInSeconds *float64 `field:"optional" json:"cacheStaleTimeoutInSeconds" yaml:"cacheStaleTimeoutInSeconds"`
}

