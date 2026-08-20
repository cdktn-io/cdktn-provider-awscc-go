// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivsplaybackkeypair

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IvsPlaybackKeyPairConfig struct {
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
	// An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource.
	//
	// The value does not need to be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_playback_key_pair#name IvsPlaybackKeyPair#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The public portion of a customer-generated key pair. This field is required to create the AWS::IVS::PlaybackKeyPair resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_playback_key_pair#public_key_material IvsPlaybackKeyPair#public_key_material}
	PublicKeyMaterial *string `field:"optional" json:"publicKeyMaterial" yaml:"publicKeyMaterial"`
	// A list of key-value pairs that contain metadata for the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_playback_key_pair#tags IvsPlaybackKeyPair#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

