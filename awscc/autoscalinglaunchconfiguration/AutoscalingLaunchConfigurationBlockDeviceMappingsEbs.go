// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalinglaunchconfiguration


type AutoscalingLaunchConfigurationBlockDeviceMappingsEbs struct {
	// Indicates whether the volume is deleted on instance termination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/autoscaling_launch_configuration#delete_on_termination AutoscalingLaunchConfiguration#delete_on_termination}
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Specifies whether the volume should be encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/autoscaling_launch_configuration#encrypted AutoscalingLaunchConfiguration#encrypted}
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of input/output (I/O) operations per second (IOPS) to provision for the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/autoscaling_launch_configuration#iops AutoscalingLaunchConfiguration#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The snapshot ID of the volume to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/autoscaling_launch_configuration#snapshot_id AutoscalingLaunchConfiguration#snapshot_id}
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The throughput (MiBps) to provision for a gp3 volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/autoscaling_launch_configuration#throughput AutoscalingLaunchConfiguration#throughput}
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The volume size, in GiBs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/autoscaling_launch_configuration#volume_size AutoscalingLaunchConfiguration#volume_size}
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// The volume type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/autoscaling_launch_configuration#volume_type AutoscalingLaunchConfiguration#volume_type}
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

