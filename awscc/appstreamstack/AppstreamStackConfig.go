// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamstack

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppstreamStackConfig struct {
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
	// The list of virtual private cloud (VPC) interface endpoint objects.
	//
	// Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#access_endpoints AppstreamStack#access_endpoints}
	AccessEndpoints interface{} `field:"optional" json:"accessEndpoints" yaml:"accessEndpoints"`
	// The configuration for agent access on the stack. If specified, agent access is enabled for the stack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#agent_access_config AppstreamStack#agent_access_config}
	AgentAccessConfig *AppstreamStackAgentAccessConfig `field:"optional" json:"agentAccessConfig" yaml:"agentAccessConfig"`
	// The persistent application settings for users of the stack.
	//
	// When these settings are enabled, changes that users make to applications and Windows settings are automatically saved after each session and applied to the next session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#application_settings AppstreamStack#application_settings}
	ApplicationSettings *AppstreamStackApplicationSettings `field:"optional" json:"applicationSettings" yaml:"applicationSettings"`
	// The stack attributes to delete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#attributes_to_delete AppstreamStack#attributes_to_delete}
	AttributesToDelete *[]*string `field:"optional" json:"attributesToDelete" yaml:"attributesToDelete"`
	// The content redirection settings for the stack.
	//
	// These settings control URL redirection between the streaming session and the local device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#content_redirection AppstreamStack#content_redirection}
	ContentRedirection *AppstreamStackContentRedirection `field:"optional" json:"contentRedirection" yaml:"contentRedirection"`
	// This parameter has been deprecated. Deletes the storage connectors currently enabled for the stack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#delete_storage_connectors AppstreamStack#delete_storage_connectors}
	DeleteStorageConnectors interface{} `field:"optional" json:"deleteStorageConnectors" yaml:"deleteStorageConnectors"`
	// The description to display.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#description AppstreamStack#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The stack name to display.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#display_name AppstreamStack#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#embed_host_domains AppstreamStack#embed_host_domains}
	EmbedHostDomains *[]*string `field:"optional" json:"embedHostDomains" yaml:"embedHostDomains"`
	// The URL that users are redirected to after they click the Send Feedback link.
	//
	// If no URL is specified, no Send Feedback link is displayed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#feedback_url AppstreamStack#feedback_url}
	FeedbackUrl *string `field:"optional" json:"feedbackUrl" yaml:"feedbackUrl"`
	// The name of the stack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#name AppstreamStack#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The URL that users are redirected to after their streaming session ends.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#redirect_url AppstreamStack#redirect_url}
	RedirectUrl *string `field:"optional" json:"redirectUrl" yaml:"redirectUrl"`
	// The storage connectors to enable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#storage_connectors AppstreamStack#storage_connectors}
	StorageConnectors interface{} `field:"optional" json:"storageConnectors" yaml:"storageConnectors"`
	// The streaming protocol that you want your stack to prefer.
	//
	// This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#streaming_experience_settings AppstreamStack#streaming_experience_settings}
	StreamingExperienceSettings *AppstreamStackStreamingExperienceSettings `field:"optional" json:"streamingExperienceSettings" yaml:"streamingExperienceSettings"`
	// An array of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#tags AppstreamStack#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appstream_stack#user_settings AppstreamStack#user_settings}
	UserSettings interface{} `field:"optional" json:"userSettings" yaml:"userSettings"`
}

