// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.provider.AwsccProvider",
		reflect.TypeOf((*AwsccProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessKey", GoGetter: "AccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "accessKeyInput", GoGetter: "AccessKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRole", GoGetter: "AssumeRole"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleInput", GoGetter: "AssumeRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleWithWebIdentity", GoGetter: "AssumeRoleWithWebIdentity"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleWithWebIdentityInput", GoGetter: "AssumeRoleWithWebIdentityInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberProperty{JsiiProperty: "endpointsInput", GoGetter: "EndpointsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "httpProxy", GoGetter: "HttpProxy"},
			_jsii_.MemberProperty{JsiiProperty: "httpProxyInput", GoGetter: "HttpProxyInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpsProxy", GoGetter: "HttpsProxy"},
			_jsii_.MemberProperty{JsiiProperty: "httpsProxyInput", GoGetter: "HttpsProxyInput"},
			_jsii_.MemberProperty{JsiiProperty: "insecure", GoGetter: "Insecure"},
			_jsii_.MemberProperty{JsiiProperty: "insecureInput", GoGetter: "InsecureInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetries", GoGetter: "MaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesInput", GoGetter: "MaxRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "noProxy", GoGetter: "NoProxy"},
			_jsii_.MemberProperty{JsiiProperty: "noProxyInput", GoGetter: "NoProxyInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "profile", GoGetter: "Profile"},
			_jsii_.MemberProperty{JsiiProperty: "profileInput", GoGetter: "ProfileInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessKey", GoMethod: "ResetAccessKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssumeRole", GoMethod: "ResetAssumeRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssumeRoleWithWebIdentity", GoMethod: "ResetAssumeRoleWithWebIdentity"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndpoints", GoMethod: "ResetEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpProxy", GoMethod: "ResetHttpProxy"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpsProxy", GoMethod: "ResetHttpsProxy"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecure", GoMethod: "ResetInsecure"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRetries", GoMethod: "ResetMaxRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoProxy", GoMethod: "ResetNoProxy"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProfile", GoMethod: "ResetProfile"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleArn", GoMethod: "ResetRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretKey", GoMethod: "ResetSecretKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharedConfigFiles", GoMethod: "ResetSharedConfigFiles"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharedCredentialsFiles", GoMethod: "ResetSharedCredentialsFiles"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipMedatadataApiCheck", GoMethod: "ResetSkipMedatadataApiCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipMetadataApiCheck", GoMethod: "ResetSkipMetadataApiCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserAgent", GoMethod: "ResetUserAgent"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "secretKey", GoGetter: "SecretKey"},
			_jsii_.MemberProperty{JsiiProperty: "secretKeyInput", GoGetter: "SecretKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "sharedConfigFiles", GoGetter: "SharedConfigFiles"},
			_jsii_.MemberProperty{JsiiProperty: "sharedConfigFilesInput", GoGetter: "SharedConfigFilesInput"},
			_jsii_.MemberProperty{JsiiProperty: "sharedCredentialsFiles", GoGetter: "SharedCredentialsFiles"},
			_jsii_.MemberProperty{JsiiProperty: "sharedCredentialsFilesInput", GoGetter: "SharedCredentialsFilesInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipMedatadataApiCheck", GoGetter: "SkipMedatadataApiCheck"},
			_jsii_.MemberProperty{JsiiProperty: "skipMedatadataApiCheckInput", GoGetter: "SkipMedatadataApiCheckInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipMetadataApiCheck", GoGetter: "SkipMetadataApiCheck"},
			_jsii_.MemberProperty{JsiiProperty: "skipMetadataApiCheckInput", GoGetter: "SkipMetadataApiCheckInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userAgent", GoGetter: "UserAgent"},
			_jsii_.MemberProperty{JsiiProperty: "userAgentInput", GoGetter: "UserAgentInput"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_AwsccProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.provider.AwsccProviderAssumeRole",
		reflect.TypeOf((*AwsccProviderAssumeRole)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.provider.AwsccProviderAssumeRoleWithWebIdentity",
		reflect.TypeOf((*AwsccProviderAssumeRoleWithWebIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.provider.AwsccProviderConfig",
		reflect.TypeOf((*AwsccProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.provider.AwsccProviderEndpoints",
		reflect.TypeOf((*AwsccProviderEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.provider.AwsccProviderUserAgent",
		reflect.TypeOf((*AwsccProviderUserAgent)(nil)).Elem(),
	)
}
