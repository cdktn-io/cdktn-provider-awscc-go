// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagegatewaytape

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StoragegatewayTapeConfig struct {
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
	// The Amazon Resource Name (ARN) of the Tape Gateway that hosts the virtual tape.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/storagegateway_tape#gateway_arn StoragegatewayTape#gateway_arn}
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// The size, in bytes, of the virtual tape that you want to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/storagegateway_tape#tape_size_in_bytes StoragegatewayTape#tape_size_in_bytes}
	TapeSizeInBytes *float64 `field:"required" json:"tapeSizeInBytes" yaml:"tapeSizeInBytes"`
	// Set to true to use Amazon S3 server-side encryption with your own KMS key, or false to use a key managed by Amazon S3.
	//
	// Optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/storagegateway_tape#kms_encrypted StoragegatewayTape#kms_encrypted}
	KmsEncrypted interface{} `field:"optional" json:"kmsEncrypted" yaml:"kmsEncrypted"`
	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for Amazon S3 server-side encryption.
	//
	// This value must be set if KMSEncrypted is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/storagegateway_tape#kms_key StoragegatewayTape#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The ID of the pool that you want to add your tape to for archiving.
	//
	// Tapes in this pool are archived in the S3 storage class that is associated with the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/storagegateway_tape#pool_id StoragegatewayTape#pool_id}
	PoolId *string `field:"optional" json:"poolId" yaml:"poolId"`
	// A list of up to 50 tags to assign to the virtual tape.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/storagegateway_tape#tags StoragegatewayTape#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The barcode that you want to assign to the virtual tape.
	//
	// Barcodes cannot be reused, even after a tape is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/storagegateway_tape#tape_barcode StoragegatewayTape#tape_barcode}
	TapeBarcode *string `field:"optional" json:"tapeBarcode" yaml:"tapeBarcode"`
	// Set to true to create a write-once-read-many (WORM) virtual tape.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/storagegateway_tape#worm StoragegatewayTape#worm}
	Worm interface{} `field:"optional" json:"worm" yaml:"worm"`
}

