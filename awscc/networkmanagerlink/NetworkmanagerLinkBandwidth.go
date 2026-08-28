// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerlink


type NetworkmanagerLinkBandwidth struct {
	// Download speed in Mbps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_link#download_speed NetworkmanagerLink#download_speed}
	DownloadSpeed *float64 `field:"optional" json:"downloadSpeed" yaml:"downloadSpeed"`
	// Upload speed in Mbps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_link#upload_speed NetworkmanagerLink#upload_speed}
	UploadSpeed *float64 `field:"optional" json:"uploadSpeed" yaml:"uploadSpeed"`
}

