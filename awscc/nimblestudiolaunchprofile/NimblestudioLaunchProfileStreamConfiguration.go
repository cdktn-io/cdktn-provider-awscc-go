// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nimblestudiolaunchprofile


type NimblestudioLaunchProfileStreamConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_launch_profile#clipboard_mode NimblestudioLaunchProfile#clipboard_mode}.
	ClipboardMode *string `field:"required" json:"clipboardMode" yaml:"clipboardMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_launch_profile#ec_2_instance_types NimblestudioLaunchProfile#ec_2_instance_types}.
	Ec2InstanceTypes *[]*string `field:"required" json:"ec2InstanceTypes" yaml:"ec2InstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_launch_profile#streaming_image_ids NimblestudioLaunchProfile#streaming_image_ids}.
	StreamingImageIds *[]*string `field:"required" json:"streamingImageIds" yaml:"streamingImageIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_launch_profile#automatic_termination_mode NimblestudioLaunchProfile#automatic_termination_mode}.
	AutomaticTerminationMode *string `field:"optional" json:"automaticTerminationMode" yaml:"automaticTerminationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_launch_profile#max_session_length_in_minutes NimblestudioLaunchProfile#max_session_length_in_minutes}.
	MaxSessionLengthInMinutes *float64 `field:"optional" json:"maxSessionLengthInMinutes" yaml:"maxSessionLengthInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_launch_profile#max_stopped_session_length_in_minutes NimblestudioLaunchProfile#max_stopped_session_length_in_minutes}.
	MaxStoppedSessionLengthInMinutes *float64 `field:"optional" json:"maxStoppedSessionLengthInMinutes" yaml:"maxStoppedSessionLengthInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_launch_profile#session_backup NimblestudioLaunchProfile#session_backup}.
	SessionBackup *NimblestudioLaunchProfileStreamConfigurationSessionBackup `field:"optional" json:"sessionBackup" yaml:"sessionBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_launch_profile#session_persistence_mode NimblestudioLaunchProfile#session_persistence_mode}.
	SessionPersistenceMode *string `field:"optional" json:"sessionPersistenceMode" yaml:"sessionPersistenceMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_launch_profile#session_storage NimblestudioLaunchProfile#session_storage}.
	SessionStorage *NimblestudioLaunchProfileStreamConfigurationSessionStorage `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_launch_profile#volume_configuration NimblestudioLaunchProfile#volume_configuration}.
	VolumeConfiguration *NimblestudioLaunchProfileStreamConfigurationVolumeConfiguration `field:"optional" json:"volumeConfiguration" yaml:"volumeConfiguration"`
}

