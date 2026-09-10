// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package interconnectconnection


type InterconnectConnectionRemoteAccount struct {
	// The identifier of the remote account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/interconnect_connection#identifier InterconnectConnection#identifier}
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

