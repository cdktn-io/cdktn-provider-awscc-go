// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationDestinationEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_input#endpoint_lifecycle_management MediaconnectRouterInput#endpoint_lifecycle_management}.
	EndpointLifecycleManagement *string `field:"optional" json:"endpointLifecycleManagement" yaml:"endpointLifecycleManagement"`
	// The encryption settings for an endpoint connection.
	//
	// This is a union type, so exactly one of its members is set: secretsManager when you store the encryption key or passphrase in Secrets Manager, or automatic when you let MediaConnect determine the encryption key when the endpoint connects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_input#source_transit_decryption MediaconnectRouterInput#source_transit_decryption}
	SourceTransitDecryption *MediaconnectRouterInputConfigurationDestinationEndpointSourceTransitDecryption `field:"optional" json:"sourceTransitDecryption" yaml:"sourceTransitDecryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_input#source_transit_decryption_type MediaconnectRouterInput#source_transit_decryption_type}.
	SourceTransitDecryptionType *string `field:"optional" json:"sourceTransitDecryptionType" yaml:"sourceTransitDecryptionType"`
}

