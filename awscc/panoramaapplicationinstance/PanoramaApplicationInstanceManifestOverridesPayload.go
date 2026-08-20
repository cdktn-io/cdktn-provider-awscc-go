// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package panoramaapplicationinstance


type PanoramaApplicationInstanceManifestOverridesPayload struct {
	// The overrides document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/panorama_application_instance#payload_data PanoramaApplicationInstance#payload_data}
	PayloadData *string `field:"optional" json:"payloadData" yaml:"payloadData"`
}

