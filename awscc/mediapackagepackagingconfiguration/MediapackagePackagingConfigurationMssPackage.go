// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagepackagingconfiguration


type MediapackagePackagingConfigurationMssPackage struct {
	// A CMAF encryption configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackage_packaging_configuration#encryption MediapackagePackagingConfiguration#encryption}
	Encryption *MediapackagePackagingConfigurationMssPackageEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// A list of MSS manifest configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackage_packaging_configuration#mss_manifests MediapackagePackagingConfiguration#mss_manifests}
	MssManifests interface{} `field:"optional" json:"mssManifests" yaml:"mssManifests"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments will be rounded to the nearest multiple of the source fragment duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackage_packaging_configuration#segment_duration_seconds MediapackagePackagingConfiguration#segment_duration_seconds}
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
}

