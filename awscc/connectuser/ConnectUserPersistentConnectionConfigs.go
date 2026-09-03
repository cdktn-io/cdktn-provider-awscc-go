// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuser


type ConnectUserPersistentConnectionConfigs struct {
	// The channels that agents can handle in the Contact Control Panel (CCP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user#channel ConnectUser#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// The Persistent Connection setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user#persistent_connection ConnectUser#persistent_connection}
	PersistentConnection interface{} `field:"optional" json:"persistentConnection" yaml:"persistentConnection"`
}

