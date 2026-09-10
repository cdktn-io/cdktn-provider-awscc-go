// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2volume

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2VolumeConfig struct {
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
	// Indicates whether the volume is auto-enabled for I/O operations.
	//
	// By default, EBS disables I/O to the volume from attached EC2 instances when it determines that a volume's data is potentially inconsistent. If the consistency of the volume is not a concern, and you prefer that the volume be made available immediately if it's impaired, you can configure the volume to automatically enable I/O.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#auto_enable_io Ec2Volume#auto_enable_io}
	AutoEnableIo interface{} `field:"optional" json:"autoEnableIo" yaml:"autoEnableIo"`
	// The ID of the Availability Zone in which to create the volume.
	//
	// For example, ``us-east-1a``.
	//  Either ``AvailabilityZone`` or ``AvailabilityZoneId`` must be specified, but not both.
	//  If you are creating a volume copy, omit this parameter. The volume copy is created in the same Availability Zone as the source volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#availability_zone Ec2Volume#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The ID of the Availability Zone in which to create the volume.
	//
	// For example, ``use1-az1``.
	//  Either ``AvailabilityZone`` or ``AvailabilityZoneId`` must be specified, but not both.
	//  If you are creating a volume copy, omit this parameter. The volume copy is created in the same Availability Zone as the source volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#availability_zone_id Ec2Volume#availability_zone_id}
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// Indicates whether the volume should be encrypted.
	//
	// The effect of setting the encryption state to ``true`` depends on the volume origin (new, from a snapshot, or from an existing volume), starting encryption state, ownership, and whether encryption by default is enabled. For more information, see [Encryption by default](https://docs.aws.amazon.com/ebs/latest/userguide/work-with-ebs-encr.html#encryption-by-default) in the *Amazon EBS User Guide*.
	//  If you are creating a volume copy, omit this parameter. The volume is automatically encrypted with the same KMS key as the source volume. You can't copy unencrypted volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#encrypted Ec2Volume#encrypted}
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// Required for ``io1`` and ``io2`` volumes. Optional for ``gp3`` volumes. Omit for all other volume types.
	//  Valid ranges:
	//   +  gp3: ``3,000``(*default*)``- 80,000`` IOPS
	//   +  io1: ``100 - 64,000`` IOPS
	//   +  io2: ``100 - 256,000`` IOPS
	//
	//   [Instances built on the Nitro System](https://docs.aws.amazon.com/ec2/latest/instancetypes/ec2-nitro-instances.html) can support up to 256,000 IOPS. Other instances can support up to 32,000 IOPS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#iops Ec2Volume#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The identifier of the kms-key-long to use for Amazon EBS encryption.
	//
	// If ``KmsKeyId`` is specified, the encrypted state must be ``true``.
	//  If you omit this property and your account is enabled for encryption by default, or *Encrypted* is set to ``true``, then the volume is encrypted using the default key specified for your account. If your account does not have a default key, then the volume is encrypted using the aws-managed-key.
	//  Alternatively, if you want to specify a different key, you can specify one of the following:
	//   +  Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.
	//   +  Key alias. Specify the alias for the key, prefixed with ``alias/``. For example, for a key with the alias ``my_cmk``, use ``alias/my_cmk``. Or to specify the aws-managed-key, use ``alias/aws/ebs``.
	//   +  Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//   +  Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	//  If you are creating a volume copy, omit this parameter. The volume is automatically encrypted with the same KMS key as the source volume. You can't copy unencrypted volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#kms_key_id Ec2Volume#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Indicates whether Amazon EBS Multi-Attach is enabled.
	//
	// CFNlong does not currently support updating a single-attach volume to be multi-attach enabled, updating a multi-attach enabled volume to be single-attach, or updating the size or number of I/O operations per second (IOPS) of a multi-attach enabled volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#multi_attach_enabled Ec2Volume#multi_attach_enabled}
	MultiAttachEnabled interface{} `field:"optional" json:"multiAttachEnabled" yaml:"multiAttachEnabled"`
	// The Amazon Resource Name (ARN) of the Outpost on which to create the volume.
	//
	// If you intend to use a volume with an instance running on an outpost, then you must create the volume on the same outpost as the instance. You can't use a volume created in an AWS Region with an instance on an AWS outpost, or the other way around.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#outpost_arn Ec2Volume#outpost_arn}
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// The size of the volume, in GiBs.
	//
	// +  Required for new empty volumes.
	//   +  Optional for volumes created from snapshots and volume copies. In this case, the size defaults to the size of the snapshot or source volume. You can optionally specify a size that is equal to or larger than the size of the source snapshot or volume.
	//
	//  Supported volume sizes:
	//   +  gp2: ``1 - 16,384`` GiB
	//   +  gp3: ``1 - 65,536`` GiB
	//   +  io1: ``4 - 16,384`` GiB
	//   +  io2: ``4 - 65,536`` GiB
	//   +  st1 and sc1: ``125 - 16,384`` GiB
	//   +  standard: ``1 - 1024`` GiB
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#size Ec2Volume#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The snapshot from which to create the volume.
	//
	// Only specify to create a volume from a snapshot. To create a new empty volume, omit this parameter and specify a value for ``Size`` instead. To create a volume copy, omit this parameter and specify ``SourceVolumeId`` instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#snapshot_id Ec2Volume#snapshot_id}
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The ID of the source EBS volume to copy.
	//
	// When specified, the volume is created as an exact copy of the specified volume. Only specify to create a volume copy. To create a new empty volume or to create a volume from a snapshot, omit this parameter,
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#source_volume_id Ec2Volume#source_volume_id}
	SourceVolumeId *string `field:"optional" json:"sourceVolumeId" yaml:"sourceVolumeId"`
	// The tags to apply to the volume during creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#tags Ec2Volume#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The throughput to provision for a volume, with a maximum of 2,000 MiB/s.
	//
	// This parameter is valid only for ``gp3`` volumes. The default value is 125.
	//  Valid Range: Minimum value of 125. Maximum value of 2000.
	//  The maximum ratio of throughput to IOPS is 0.25 MiB/s per IOPS. For example, a volume with 3,000 IOPS can have a maximum throughput of 750 MiB/s (3,000 x 0.25).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#throughput Ec2Volume#throughput}
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Specifies the Amazon EBS Provisioned Rate for Volume Initialization (volume initialization rate), in MiB/s, at which to download the snapshot blocks from Amazon S3 to the volume.
	//
	// This is also known as *volume initialization*. Specifying a volume initialization rate ensures that the volume is initialized at a predictable and consistent rate after creation.
	//  This parameter is supported only for volumes created from snapshots. Omit this parameter if:
	//   +  You want to create the volume using fast snapshot restore. You must specify a snapshot that is enabled for fast snapshot restore. In this case, the volume is fully initialized at creation.
	//   If you specify a snapshot that is enabled for fast snapshot restore and a volume initialization rate, the volume will be initialized at the specified rate instead of fast snapshot restore.
	//    +  You want to create a volume that is initialized at the default rate.
	//
	//  For more information, see [Initialize Amazon EBS volumes](https://docs.aws.amazon.com/ebs/latest/userguide/initalize-volume.html) in the *Amazon EC2 User Guide*.
	//  Valid range: 100 - 300 MiB/s
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#volume_initialization_rate Ec2Volume#volume_initialization_rate}
	VolumeInitializationRate *float64 `field:"optional" json:"volumeInitializationRate" yaml:"volumeInitializationRate"`
	// The volume type.
	//
	// This parameter can be one of the following values:
	//   +  General Purpose SSD: ``gp2`` | ``gp3``
	//   +  Provisioned IOPS SSD: ``io1`` | ``io2``
	//   +  Throughput Optimized HDD: ``st1``
	//   +  Cold HDD: ``sc1``
	//   +  Magnetic: ``standard``
	//
	//   Throughput Optimized HDD (``st1``) and Cold HDD (``sc1``) volumes can't be used as boot volumes.
	//   For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-volume-types.html) in the *Amazon EBS User Guide*.
	//  Default: ``gp2``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_volume#volume_type Ec2Volume#volume_type}
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

