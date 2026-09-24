// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamstack


type AppstreamStackContentRedirection struct {
	// The URL redirection configuration from the streaming session host to the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appstream_stack#host_to_client AppstreamStack#host_to_client}
	HostToClient *AppstreamStackContentRedirectionHostToClient `field:"optional" json:"hostToClient" yaml:"hostToClient"`
}

