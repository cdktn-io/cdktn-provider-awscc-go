// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfiguration struct {
	// Set this to either `UDF` or `DEPENDENCY_JAR`.
	//
	// `UDF` stands for user-defined functions. This type of artifact must be in an S3 bucket. A `DEPENDENCY_JAR` can be in either Maven or an S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisanalyticsv2_application#artifact_type Kinesisanalyticsv2Application#artifact_type}
	ArtifactType *string `field:"optional" json:"artifactType" yaml:"artifactType"`
	// The parameters required to fully specify a Maven reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisanalyticsv2_application#maven_reference Kinesisanalyticsv2Application#maven_reference}
	MavenReference *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationMavenReference `field:"optional" json:"mavenReference" yaml:"mavenReference"`
	// The location of the custom artifacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisanalyticsv2_application#s3_content_location Kinesisanalyticsv2Application#s3_content_location}
	S3ContentLocation *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCustomArtifactsConfigurationS3ContentLocation `field:"optional" json:"s3ContentLocation" yaml:"s3ContentLocation"`
}

