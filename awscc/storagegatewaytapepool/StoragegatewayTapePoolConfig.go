// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagegatewaytapepool

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StoragegatewayTapePoolConfig struct {
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
	// The name of the custom tape pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/storagegateway_tape_pool#pool_name StoragegatewayTapePool#pool_name}
	PoolName *string `field:"required" json:"poolName" yaml:"poolName"`
	// The storage class associated with the custom pool (S3 Glacier or S3 Glacier Deep Archive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/storagegateway_tape_pool#storage_class StoragegatewayTapePool#storage_class}
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Tape retention lock time in days (up to 36,500 days / 100 years).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/storagegateway_tape_pool#retention_lock_time_in_days StoragegatewayTapePool#retention_lock_time_in_days}
	RetentionLockTimeInDays *float64 `field:"optional" json:"retentionLockTimeInDays" yaml:"retentionLockTimeInDays"`
	// Tape retention lock type. Governance mode allows authorized removal; compliance mode prevents all removal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/storagegateway_tape_pool#retention_lock_type StoragegatewayTapePool#retention_lock_type}
	RetentionLockType *string `field:"optional" json:"retentionLockType" yaml:"retentionLockType"`
	// A list of up to 50 tags for the tape pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/storagegateway_tape_pool#tags StoragegatewayTapePool#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

