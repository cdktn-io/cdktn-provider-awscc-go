// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointSegmentEncryption struct {
	// <p>Excludes SEIG and SGPD boxes from segment metadata in CMAF containers.</p> <p>When set to <code>true</code>, MediaPackage omits these DRM metadata boxes from CMAF segments, which can improve compatibility with certain devices and players that don't support these boxes.</p> <p>Important considerations:</p> <ul> <li> <p>This setting only affects CMAF container formats</p> </li> <li> <p>Key rotation can still be handled through media playlist signaling</p> </li> <li> <p>PSSH and TENC boxes remain unaffected</p> </li> <li> <p>Default behavior is preserved when this setting is disabled</p> </li> </ul> <p>Valid values: <code>true</code> | <code>false</code> </p> <p>Default: <code>false</code> </p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#cmaf_exclude_segment_drm_metadata Mediapackagev2OriginEndpoint#cmaf_exclude_segment_drm_metadata}
	CmafExcludeSegmentDrmMetadata interface{} `field:"optional" json:"cmafExcludeSegmentDrmMetadata" yaml:"cmafExcludeSegmentDrmMetadata"`
	// <p>A 128-bit, 16-byte hex value represented by a 32-character string, used in conjunction with the key for encrypting content.
	//
	// If you don't specify a value, then MediaPackage creates the constant initialization vector (IV).</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#constant_initialization_vector Mediapackagev2OriginEndpoint#constant_initialization_vector}
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// <p>The encryption type.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#encryption_method Mediapackagev2OriginEndpoint#encryption_method}
	EncryptionMethod *Mediapackagev2OriginEndpointSegmentEncryptionEncryptionMethod `field:"optional" json:"encryptionMethod" yaml:"encryptionMethod"`
	// <p>The frequency (in seconds) of key changes for live workflows, in which content is streamed real time.
	//
	// The service retrieves content keys before the live content begins streaming, and then retrieves them as needed over the lifetime of the workflow. By default, key rotation is set to 300 seconds (5 minutes), the minimum rotation interval, which is equivalent to setting it to 300. If you don't enter an interval, content keys aren't rotated.</p> <p>The following example setting causes the service to rotate keys every thirty minutes: <code>1800</code> </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#key_rotation_interval_seconds Mediapackagev2OriginEndpoint#key_rotation_interval_seconds}
	KeyRotationIntervalSeconds *float64 `field:"optional" json:"keyRotationIntervalSeconds" yaml:"keyRotationIntervalSeconds"`
	// <p>The parameters for the SPEKE key provider.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#speke_key_provider Mediapackagev2OriginEndpoint#speke_key_provider}
	SpekeKeyProvider *Mediapackagev2OriginEndpointSegmentEncryptionSpekeKeyProvider `field:"optional" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
}

