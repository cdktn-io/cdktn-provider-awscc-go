// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationMavenReference struct {
	// The artifact ID of the Maven reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kinesisanalyticsv2_application#artifact_id Kinesisanalyticsv2Application#artifact_id}
	ArtifactId *string `field:"optional" json:"artifactId" yaml:"artifactId"`
	// The group ID of the Maven reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kinesisanalyticsv2_application#group_id Kinesisanalyticsv2Application#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The version of the Maven reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kinesisanalyticsv2_application#version Kinesisanalyticsv2Application#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

