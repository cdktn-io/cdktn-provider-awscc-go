// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MedialiveInputConfig struct {
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
	// A list of the destinations of the input (PUSH-type).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#destinations MedialiveInput#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Settings for the input devices.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#input_devices MedialiveInput#input_devices}
	InputDevices interface{} `field:"optional" json:"inputDevices" yaml:"inputDevices"`
	// The location of this input.
	//
	// AWS, for an input existing in the AWS Cloud, On-Prem for
	// an input in a customer network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#input_network_location MedialiveInput#input_network_location}
	InputNetworkLocation *string `field:"optional" json:"inputNetworkLocation" yaml:"inputNetworkLocation"`
	// A list of security groups referenced by IDs to attach to the input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#input_security_groups MedialiveInput#input_security_groups}
	InputSecurityGroups *[]*string `field:"optional" json:"inputSecurityGroups" yaml:"inputSecurityGroups"`
	// A list of MediaConnect Flows for this input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#media_connect_flows MedialiveInput#media_connect_flows}
	MediaConnectFlows interface{} `field:"optional" json:"mediaConnectFlows" yaml:"mediaConnectFlows"`
	// Multicast Input settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#multicast_settings MedialiveInput#multicast_settings}
	MulticastSettings *MedialiveInputMulticastSettings `field:"optional" json:"multicastSettings" yaml:"multicastSettings"`
	// The user-assigned name (This is a mutable value).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#name MedialiveInput#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the role this input assumes during and after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#role_arn MedialiveInput#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Information about the MediaConnect Router association for this input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#router_settings MedialiveInput#router_settings}
	RouterSettings *MedialiveInputRouterSettings `field:"optional" json:"routerSettings" yaml:"routerSettings"`
	// SDI Sources for this Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#sdi_sources MedialiveInput#sdi_sources}
	SdiSources *[]*string `field:"optional" json:"sdiSources" yaml:"sdiSources"`
	// Include this parameter if the input is a SMPTE 2110 input, to identify the stream sources for this input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#smpte_2110_receiver_group_settings MedialiveInput#smpte_2110_receiver_group_settings}
	Smpte2110ReceiverGroupSettings *MedialiveInputSmpte2110ReceiverGroupSettings `field:"optional" json:"smpte2110ReceiverGroupSettings" yaml:"smpte2110ReceiverGroupSettings"`
	// A list of the sources of the input (PULL-type).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#sources MedialiveInput#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// The settings associated with an SRT input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#srt_settings MedialiveInput#srt_settings}
	SrtSettings *MedialiveInputSrtSettings `field:"optional" json:"srtSettings" yaml:"srtSettings"`
	// A collection of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#tags MedialiveInput#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The different types of inputs that AWS Elemental MediaLive supports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#type MedialiveInput#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Settings for a private VPC Input.
	//
	// When this property is specified, the input destination addresses will be created in a VPC rather than with public Internet addresses.
	// This property requires setting the roleArn property on Input creation.
	// Not compatible with the inputSecurityGroups property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#vpc MedialiveInput#vpc}
	Vpc *MedialiveInputVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

