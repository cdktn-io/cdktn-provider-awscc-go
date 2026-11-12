// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfigurationSourceEndpoint struct {
	// The encryption settings for an endpoint connection.
	//
	// This is a union type, so exactly one of its members is set: secretsManager when you store the encryption key or passphrase in Secrets Manager, or automatic when you let MediaConnect determine the encryption key when the endpoint connects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_output#destination_transit_encryption MediaconnectRouterOutput#destination_transit_encryption}
	DestinationTransitEncryption *MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryption `field:"optional" json:"destinationTransitEncryption" yaml:"destinationTransitEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_output#destination_transit_encryption_type MediaconnectRouterOutput#destination_transit_encryption_type}.
	DestinationTransitEncryptionType *string `field:"optional" json:"destinationTransitEncryptionType" yaml:"destinationTransitEncryptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_output#endpoint_lifecycle_management MediaconnectRouterOutput#endpoint_lifecycle_management}.
	EndpointLifecycleManagement *string `field:"optional" json:"endpointLifecycleManagement" yaml:"endpointLifecycleManagement"`
}

