// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputDestinations struct {
	// The system-generated static IP address of endpoint. It remains fixed for the lifetime of the input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#ip MedialiveInput#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// The ID of the attached network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#network MedialiveInput#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// If the push input has an input location of ON-PREM it's a requirement to specify what the route of the input is going to be on the customer local network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#network_routes MedialiveInput#network_routes}
	NetworkRoutes interface{} `field:"optional" json:"networkRoutes" yaml:"networkRoutes"`
	// The port number for the input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#port MedialiveInput#port}
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The static IP address for the input destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#static_ip_address MedialiveInput#static_ip_address}
	StaticIpAddress *string `field:"optional" json:"staticIpAddress" yaml:"staticIpAddress"`
	// The stream name for the input destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#stream_name MedialiveInput#stream_name}
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
	// This represents the endpoint that the customer stream will be pushed to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#url MedialiveInput#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The properties for a VPC type input destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#vpc MedialiveInput#vpc}
	Vpc *MedialiveInputDestinationsVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

