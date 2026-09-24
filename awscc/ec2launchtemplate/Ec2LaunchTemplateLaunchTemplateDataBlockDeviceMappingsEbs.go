// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsEbs struct {
	// Indicates whether the EBS volume is deleted on instance termination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_launch_template#delete_on_termination Ec2LaunchTemplate#delete_on_termination}
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Indicates whether the EBS volume is encrypted.
	//
	// Encrypted volumes can only be attached to instances that support Amazon EBS encryption. If you are creating a volume from a snapshot, you can't specify an encryption value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_launch_template#encrypted Ec2LaunchTemplate#encrypted}
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of I/O operations per second (IOPS).
	//
	// For ``gp3``, ``io1``, and ``io2`` volumes, this represents the number of IOPS that are provisioned for the volume. For ``gp2`` volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.
	//  The following are the supported values for each volume type:
	//   +   ``gp3``: 3,000 - 16,000 IOPS
	//   +   ``io1``: 100 - 64,000 IOPS
	//   +   ``io2``: 100 - 256,000 IOPS
	//
	//  For ``io2`` volumes, you can achieve up to 256,000 IOPS on [instances built on the Nitro System](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances). On other instances, you can achieve performance up to 32,000 IOPS.
	//  This parameter is supported for ``io1``, ``io2``, and ``gp3`` volumes only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_launch_template#iops Ec2LaunchTemplate#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The ARN of the symmetric KMSlong (KMS) CMK used for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_launch_template#kms_key_id Ec2LaunchTemplate#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The ID of the snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_launch_template#snapshot_id Ec2LaunchTemplate#snapshot_id}
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The throughput to provision for a ``gp3`` volume, with a maximum of 1,000 MiB/s.
	//
	// Valid Range: Minimum value of 125. Maximum value of 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_launch_template#throughput Ec2LaunchTemplate#throughput}
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The size of the volume, in GiBs.
	//
	// You must specify either a snapshot ID or a volume size. The following are the supported volumes sizes for each volume type:
	//   +   ``gp2`` and ``gp3``: 1 - 16,384 GiB
	//   +   ``io1``: 4 - 16,384 GiB
	//   +   ``io2``: 4 - 65,536 GiB
	//   +   ``st1`` and ``sc1``: 125 - 16,384 GiB
	//   +   ``standard``: 1 - 1024 GiB
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_launch_template#volume_size Ec2LaunchTemplate#volume_size}
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// The volume type. For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-volume-types.html) in the *Amazon EBS User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_launch_template#volume_type Ec2LaunchTemplate#volume_type}
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

