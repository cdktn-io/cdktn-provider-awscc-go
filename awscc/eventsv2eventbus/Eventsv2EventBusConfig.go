// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2eventbus

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Eventsv2EventBusConfig struct {
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
	// The name of the event bus.
	//
	// The first character must be alphanumeric; the remaining characters may also include '.', '-', and '_'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_bus#name Eventsv2EventBus#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the event bus. Control characters and Unicode line separators are not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_bus#description Eventsv2EventBus#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Encryption configuration for the event bus.
	//
	// The service stores and returns the customer managed key as its key ARN. The key ARN is recommended so that drift detection stays accurate. A key ID is also accepted and is matched to the returned key ARN when CloudFormation checks for drift; a key alias is accepted but can be reported as a value difference, because an alias cannot be matched to the key ARN it points to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_bus#encryption_configuration Eventsv2EventBus#encryption_configuration}
	EncryptionConfiguration *Eventsv2EventBusEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The event storage configuration for the event bus, which controls the number of days events are retained on the bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_bus#storage_configuration Eventsv2EventBus#storage_configuration}
	StorageConfiguration *Eventsv2EventBusStorageConfiguration `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
	// The tags assigned to the event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_bus#tags Eventsv2EventBus#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

