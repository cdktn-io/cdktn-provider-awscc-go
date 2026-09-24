// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointSegmentEncryptionSpekeKeyProvider struct {
	// <p>The ARN for the certificate that you imported to Amazon Web Services Certificate Manager to add content key encryption to this endpoint.
	//
	// For this feature to work, your DRM key provider must support content key encryption.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediapackagev2_origin_endpoint#certificate_arn Mediapackagev2OriginEndpoint#certificate_arn}
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// <p>The DRM solution provider you're using to protect your content during distribution.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediapackagev2_origin_endpoint#drm_systems Mediapackagev2OriginEndpoint#drm_systems}
	DrmSystems *[]*string `field:"optional" json:"drmSystems" yaml:"drmSystems"`
	// <p>Configure one or more content encryption keys for your endpoints that use SPEKE Version 2.0. The encryption contract defines which content keys are used to encrypt the audio and video tracks in your stream. To configure the encryption contract, specify which audio and video encryption presets to use.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediapackagev2_origin_endpoint#encryption_contract_configuration Mediapackagev2OriginEndpoint#encryption_contract_configuration}
	EncryptionContractConfiguration *Mediapackagev2OriginEndpointSegmentEncryptionSpekeKeyProviderEncryptionContractConfiguration `field:"optional" json:"encryptionContractConfiguration" yaml:"encryptionContractConfiguration"`
	// <p>The unique identifier for the content.
	//
	// The service sends this to the key server to identify the current endpoint. How unique you make this depends on how fine-grained you want access controls to be. The service does not permit you to use the same ID for two simultaneous encryption processes. The resource ID is also known as the content ID.</p> <p>The following example shows a resource ID: <code>MovieNight20171126093045</code> </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediapackagev2_origin_endpoint#resource_id Mediapackagev2OriginEndpoint#resource_id}
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// <p>The ARN for the IAM role granted by the key provider that provides access to the key provider API.
	//
	// This role must have a trust policy that allows MediaPackage to assume the role, and it must have a sufficient permissions policy to allow access to the specific key retrieval URL. Get this from your DRM solution provider.</p> <p>Valid format: <code>arn:aws:iam::{accountID}:role/{name}</code>. The following example shows a role ARN: <code>arn:aws:iam::444455556666:role/SpekeAccess</code> </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediapackagev2_origin_endpoint#role_arn Mediapackagev2OriginEndpoint#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// <p>The URL of the API Gateway proxy that you set up to talk to your key server.
	//
	// The API Gateway proxy must reside in the same AWS Region as MediaPackage and must start with https://.</p> <p>The following example shows a URL: <code>https://1wm2dx1f33.execute-api.us-west-2.amazonaws.com/SpekeSample/copyProtection</code> </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediapackagev2_origin_endpoint#url Mediapackagev2OriginEndpoint#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

