package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointDashManifestsProgramInformation struct {
	// <p>A copyright statement about the content.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackagev2_origin_endpoint#copyright Mediapackagev2OriginEndpoint#copyright}
	Copyright *string `field:"optional" json:"copyright" yaml:"copyright"`
	// <p>The language code for this manifest.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackagev2_origin_endpoint#language_code Mediapackagev2OriginEndpoint#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// <p>An absolute URL that contains more information about this content.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackagev2_origin_endpoint#more_information_url Mediapackagev2OriginEndpoint#more_information_url}
	MoreInformationUrl *string `field:"optional" json:"moreInformationUrl" yaml:"moreInformationUrl"`
	// <p>Information about the content provider.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackagev2_origin_endpoint#source Mediapackagev2OriginEndpoint#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// <p>The title for the manifest.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackagev2_origin_endpoint#title Mediapackagev2OriginEndpoint#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

