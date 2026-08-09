// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomain",
		reflect.TypeOf((*SagemakerDomain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "appNetworkAccessType", GoGetter: "AppNetworkAccessType"},
			_jsii_.MemberProperty{JsiiProperty: "appNetworkAccessTypeInput", GoGetter: "AppNetworkAccessTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "appSecurityGroupManagement", GoGetter: "AppSecurityGroupManagement"},
			_jsii_.MemberProperty{JsiiProperty: "appSecurityGroupManagementInput", GoGetter: "AppSecurityGroupManagementInput"},
			_jsii_.MemberProperty{JsiiProperty: "authMode", GoGetter: "AuthMode"},
			_jsii_.MemberProperty{JsiiProperty: "authModeInput", GoGetter: "AuthModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultSpaceSettings", GoGetter: "DefaultSpaceSettings"},
			_jsii_.MemberProperty{JsiiProperty: "defaultSpaceSettingsInput", GoGetter: "DefaultSpaceSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUserSettings", GoGetter: "DefaultUserSettings"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUserSettingsInput", GoGetter: "DefaultUserSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "domainArn", GoGetter: "DomainArn"},
			_jsii_.MemberProperty{JsiiProperty: "domainId", GoGetter: "DomainId"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "domainNameInput", GoGetter: "DomainNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "domainSettings", GoGetter: "DomainSettings"},
			_jsii_.MemberProperty{JsiiProperty: "domainSettingsInput", GoGetter: "DomainSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "homeEfsFileSystemCreation", GoGetter: "HomeEfsFileSystemCreation"},
			_jsii_.MemberProperty{JsiiProperty: "homeEfsFileSystemCreationInput", GoGetter: "HomeEfsFileSystemCreationInput"},
			_jsii_.MemberProperty{JsiiProperty: "homeEfsFileSystemId", GoGetter: "HomeEfsFileSystemId"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultSpaceSettings", GoMethod: "PutDefaultSpaceSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultUserSettings", GoMethod: "PutDefaultUserSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putDomainSettings", GoMethod: "PutDomainSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putTags", GoMethod: "PutTags"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppNetworkAccessType", GoMethod: "ResetAppNetworkAccessType"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppSecurityGroupManagement", GoMethod: "ResetAppSecurityGroupManagement"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultSpaceSettings", GoMethod: "ResetDefaultSpaceSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomainSettings", GoMethod: "ResetDomainSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetHomeEfsFileSystemCreation", GoMethod: "ResetHomeEfsFileSystemCreation"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnetIds", GoMethod: "ResetSubnetIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagPropagation", GoMethod: "ResetTagPropagation"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcId", GoMethod: "ResetVpcId"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdForDomainBoundary", GoGetter: "SecurityGroupIdForDomainBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "singleSignOnApplicationArn", GoGetter: "SingleSignOnApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "singleSignOnManagedApplicationInstanceId", GoGetter: "SingleSignOnManagedApplicationInstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdsInput", GoGetter: "SubnetIdsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tagPropagation", GoGetter: "TagPropagation"},
			_jsii_.MemberProperty{JsiiProperty: "tagPropagationInput", GoGetter: "TagPropagationInput"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcIdInput", GoGetter: "VpcIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomain{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainConfig",
		reflect.TypeOf((*SagemakerDomainConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettings",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigs",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsEfsFileSystemConfig",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsEfsFileSystemConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsEfsFileSystemConfigOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsEfsFileSystemConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemId", GoGetter: "FileSystemId"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemIdInput", GoGetter: "FileSystemIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemPath", GoGetter: "FileSystemPath"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemPathInput", GoGetter: "FileSystemPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileSystemId", GoMethod: "ResetFileSystemId"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileSystemPath", GoMethod: "ResetFileSystemPath"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsEfsFileSystemConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsFsxLustreFileSystemConfig",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsFsxLustreFileSystemConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsFsxLustreFileSystemConfigOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsFsxLustreFileSystemConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemId", GoGetter: "FileSystemId"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemIdInput", GoGetter: "FileSystemIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemPath", GoGetter: "FileSystemPath"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemPathInput", GoGetter: "FileSystemPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileSystemId", GoMethod: "ResetFileSystemId"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileSystemPath", GoMethod: "ResetFileSystemPath"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsFsxLustreFileSystemConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsList",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "efsFileSystemConfig", GoGetter: "EfsFileSystemConfig"},
			_jsii_.MemberProperty{JsiiProperty: "efsFileSystemConfigInput", GoGetter: "EfsFileSystemConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "fsxLustreFileSystemConfig", GoGetter: "FsxLustreFileSystemConfig"},
			_jsii_.MemberProperty{JsiiProperty: "fsxLustreFileSystemConfigInput", GoGetter: "FsxLustreFileSystemConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putEfsFileSystemConfig", GoMethod: "PutEfsFileSystemConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putFsxLustreFileSystemConfig", GoMethod: "PutFsxLustreFileSystemConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putS3FileSystemConfig", GoMethod: "PutS3FileSystemConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetEfsFileSystemConfig", GoMethod: "ResetEfsFileSystemConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetFsxLustreFileSystemConfig", GoMethod: "ResetFsxLustreFileSystemConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3FileSystemConfig", GoMethod: "ResetS3FileSystemConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3FileSystemConfig", GoGetter: "S3FileSystemConfig"},
			_jsii_.MemberProperty{JsiiProperty: "s3FileSystemConfigInput", GoGetter: "S3FileSystemConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsS3FileSystemConfig",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsS3FileSystemConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsS3FileSystemConfigOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsS3FileSystemConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mountPath", GoGetter: "MountPath"},
			_jsii_.MemberProperty{JsiiProperty: "mountPathInput", GoGetter: "MountPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMountPath", GoMethod: "ResetMountPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3Uri", GoMethod: "ResetS3Uri"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3Uri", GoGetter: "S3Uri"},
			_jsii_.MemberProperty{JsiiProperty: "s3UriInput", GoGetter: "S3UriInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsS3FileSystemConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsCustomPosixUserConfig",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsCustomPosixUserConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsCustomPosixUserConfigOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsCustomPosixUserConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "gid", GoGetter: "Gid"},
			_jsii_.MemberProperty{JsiiProperty: "gidInput", GoGetter: "GidInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetGid", GoMethod: "ResetGid"},
			_jsii_.MemberMethod{JsiiMethod: "resetUid", GoMethod: "ResetUid"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "uid", GoGetter: "Uid"},
			_jsii_.MemberProperty{JsiiProperty: "uidInput", GoGetter: "UidInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomPosixUserConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettings",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagement",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagement)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettings",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idleTimeoutInMinutes", GoGetter: "IdleTimeoutInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "idleTimeoutInMinutesInput", GoGetter: "IdleTimeoutInMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleManagement", GoGetter: "LifecycleManagement"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleManagementInput", GoGetter: "LifecycleManagementInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxIdleTimeoutInMinutes", GoGetter: "MaxIdleTimeoutInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "maxIdleTimeoutInMinutesInput", GoGetter: "MaxIdleTimeoutInMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "minIdleTimeoutInMinutes", GoGetter: "MinIdleTimeoutInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "minIdleTimeoutInMinutesInput", GoGetter: "MinIdleTimeoutInMinutesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdleTimeoutInMinutes", GoMethod: "ResetIdleTimeoutInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleManagement", GoMethod: "ResetLifecycleManagement"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxIdleTimeoutInMinutes", GoMethod: "ResetMaxIdleTimeoutInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinIdleTimeoutInMinutes", GoMethod: "ResetMinIdleTimeoutInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idleSettings", GoGetter: "IdleSettings"},
			_jsii_.MemberProperty{JsiiProperty: "idleSettingsInput", GoGetter: "IdleSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdleSettings", GoMethod: "PutIdleSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdleSettings", GoMethod: "ResetIdleSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCodeRepositories",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCodeRepositories)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCodeRepositoriesList",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCodeRepositoriesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCodeRepositoriesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCodeRepositoriesOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCodeRepositoriesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryUrl", GoGetter: "RepositoryUrl"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryUrlInput", GoGetter: "RepositoryUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRepositoryUrl", GoMethod: "ResetRepositoryUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCodeRepositoriesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCustomImages",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCustomImages)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCustomImagesList",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCustomImagesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCustomImagesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCustomImagesOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCustomImagesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigName", GoGetter: "AppImageConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigNameInput", GoGetter: "AppImageConfigNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "imageName", GoGetter: "ImageName"},
			_jsii_.MemberProperty{JsiiProperty: "imageNameInput", GoGetter: "ImageNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionNumber", GoGetter: "ImageVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionNumberInput", GoGetter: "ImageVersionNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppImageConfigName", GoMethod: "ResetAppImageConfigName"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageName", GoMethod: "ResetImageName"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageVersionNumber", GoMethod: "ResetImageVersionNumber"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsCustomImagesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageArn", GoMethod: "ResetSageMakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageVersionArn", GoMethod: "ResetSageMakerImageVersionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrainingPlanArn", GoMethod: "ResetTrainingPlanArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArn", GoGetter: "SageMakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArnInput", GoGetter: "SageMakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArn", GoGetter: "SageMakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArnInput", GoGetter: "SageMakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArn", GoGetter: "TrainingPlanArn"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArnInput", GoGetter: "TrainingPlanArnInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "appLifecycleManagement", GoGetter: "AppLifecycleManagement"},
			_jsii_.MemberProperty{JsiiProperty: "appLifecycleManagementInput", GoGetter: "AppLifecycleManagementInput"},
			_jsii_.MemberProperty{JsiiProperty: "builtInLifecycleConfigArn", GoGetter: "BuiltInLifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "builtInLifecycleConfigArnInput", GoGetter: "BuiltInLifecycleConfigArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "codeRepositories", GoGetter: "CodeRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "codeRepositoriesInput", GoGetter: "CodeRepositoriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customImages", GoGetter: "CustomImages"},
			_jsii_.MemberProperty{JsiiProperty: "customImagesInput", GoGetter: "CustomImagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArns", GoGetter: "LifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnsInput", GoGetter: "LifecycleConfigArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAppLifecycleManagement", GoMethod: "PutAppLifecycleManagement"},
			_jsii_.MemberMethod{JsiiMethod: "putCodeRepositories", GoMethod: "PutCodeRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomImages", GoMethod: "PutCustomImages"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppLifecycleManagement", GoMethod: "ResetAppLifecycleManagement"},
			_jsii_.MemberMethod{JsiiMethod: "resetBuiltInLifecycleConfigArn", GoMethod: "ResetBuiltInLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetCodeRepositories", GoMethod: "ResetCodeRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomImages", GoMethod: "ResetCustomImages"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultResourceSpec", GoMethod: "ResetDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArns", GoMethod: "ResetLifecycleConfigArns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettings",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageArn", GoMethod: "ResetSageMakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageVersionArn", GoMethod: "ResetSageMakerImageVersionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrainingPlanArn", GoMethod: "ResetTrainingPlanArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArn", GoGetter: "SageMakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArnInput", GoGetter: "SageMakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArn", GoGetter: "SageMakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArnInput", GoGetter: "SageMakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArn", GoGetter: "TrainingPlanArn"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArnInput", GoGetter: "TrainingPlanArnInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArns", GoGetter: "LifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnsInput", GoGetter: "LifecycleConfigArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultResourceSpec", GoMethod: "ResetDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArns", GoMethod: "ResetLifecycleConfigArns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettings",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImages",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImages)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImagesList",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImagesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImagesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImagesOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImagesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigName", GoGetter: "AppImageConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigNameInput", GoGetter: "AppImageConfigNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "imageName", GoGetter: "ImageName"},
			_jsii_.MemberProperty{JsiiProperty: "imageNameInput", GoGetter: "ImageNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionNumber", GoGetter: "ImageVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionNumberInput", GoGetter: "ImageVersionNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppImageConfigName", GoMethod: "ResetAppImageConfigName"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageName", GoMethod: "ResetImageName"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageVersionNumber", GoMethod: "ResetImageVersionNumber"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImagesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageArn", GoMethod: "ResetSageMakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageVersionArn", GoMethod: "ResetSageMakerImageVersionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrainingPlanArn", GoMethod: "ResetTrainingPlanArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArn", GoGetter: "SageMakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArnInput", GoGetter: "SageMakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArn", GoGetter: "SageMakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArnInput", GoGetter: "SageMakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArn", GoGetter: "TrainingPlanArn"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArnInput", GoGetter: "TrainingPlanArnInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customImages", GoGetter: "CustomImages"},
			_jsii_.MemberProperty{JsiiProperty: "customImagesInput", GoGetter: "CustomImagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArns", GoGetter: "LifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnsInput", GoGetter: "LifecycleConfigArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomImages", GoMethod: "PutCustomImages"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomImages", GoMethod: "ResetCustomImages"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultResourceSpec", GoMethod: "ResetDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArns", GoMethod: "ResetLifecycleConfigArns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customFileSystemConfigs", GoGetter: "CustomFileSystemConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "customFileSystemConfigsInput", GoGetter: "CustomFileSystemConfigsInput"},
			_jsii_.MemberProperty{JsiiProperty: "customPosixUserConfig", GoGetter: "CustomPosixUserConfig"},
			_jsii_.MemberProperty{JsiiProperty: "customPosixUserConfigInput", GoGetter: "CustomPosixUserConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleInput", GoGetter: "ExecutionRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "jupyterLabAppSettings", GoGetter: "JupyterLabAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "jupyterLabAppSettingsInput", GoGetter: "JupyterLabAppSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "jupyterServerAppSettings", GoGetter: "JupyterServerAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "jupyterServerAppSettingsInput", GoGetter: "JupyterServerAppSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "kernelGatewayAppSettings", GoGetter: "KernelGatewayAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "kernelGatewayAppSettingsInput", GoGetter: "KernelGatewayAppSettingsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomFileSystemConfigs", GoMethod: "PutCustomFileSystemConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomPosixUserConfig", GoMethod: "PutCustomPosixUserConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putJupyterLabAppSettings", GoMethod: "PutJupyterLabAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putJupyterServerAppSettings", GoMethod: "PutJupyterServerAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putKernelGatewayAppSettings", GoMethod: "PutKernelGatewayAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putSpaceStorageSettings", GoMethod: "PutSpaceStorageSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomFileSystemConfigs", GoMethod: "ResetCustomFileSystemConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomPosixUserConfig", GoMethod: "ResetCustomPosixUserConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetExecutionRole", GoMethod: "ResetExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetJupyterLabAppSettings", GoMethod: "ResetJupyterLabAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetJupyterServerAppSettings", GoMethod: "ResetJupyterServerAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetKernelGatewayAppSettings", GoMethod: "ResetKernelGatewayAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroups", GoMethod: "ResetSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpaceStorageSettings", GoMethod: "ResetSpaceStorageSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupsInput", GoGetter: "SecurityGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "spaceStorageSettings", GoGetter: "SpaceStorageSettings"},
			_jsii_.MemberProperty{JsiiProperty: "spaceStorageSettingsInput", GoGetter: "SpaceStorageSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsSpaceStorageSettings",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsSpaceStorageSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettings",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultEbsVolumeSizeInGb", GoGetter: "DefaultEbsVolumeSizeInGb"},
			_jsii_.MemberProperty{JsiiProperty: "defaultEbsVolumeSizeInGbInput", GoGetter: "DefaultEbsVolumeSizeInGbInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maximumEbsVolumeSizeInGb", GoGetter: "MaximumEbsVolumeSizeInGb"},
			_jsii_.MemberProperty{JsiiProperty: "maximumEbsVolumeSizeInGbInput", GoGetter: "MaximumEbsVolumeSizeInGbInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultEbsVolumeSizeInGb", GoMethod: "ResetDefaultEbsVolumeSizeInGb"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaximumEbsVolumeSizeInGb", GoMethod: "ResetMaximumEbsVolumeSizeInGb"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultEbsStorageSettings", GoGetter: "DefaultEbsStorageSettings"},
			_jsii_.MemberProperty{JsiiProperty: "defaultEbsStorageSettingsInput", GoGetter: "DefaultEbsStorageSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultEbsStorageSettings", GoMethod: "PutDefaultEbsStorageSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultEbsStorageSettings", GoMethod: "ResetDefaultEbsStorageSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCodeEditorAppSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCodeEditorAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagement",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagement)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idleTimeoutInMinutes", GoGetter: "IdleTimeoutInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "idleTimeoutInMinutesInput", GoGetter: "IdleTimeoutInMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleManagement", GoGetter: "LifecycleManagement"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleManagementInput", GoGetter: "LifecycleManagementInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxIdleTimeoutInMinutes", GoGetter: "MaxIdleTimeoutInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "maxIdleTimeoutInMinutesInput", GoGetter: "MaxIdleTimeoutInMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "minIdleTimeoutInMinutes", GoGetter: "MinIdleTimeoutInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "minIdleTimeoutInMinutesInput", GoGetter: "MinIdleTimeoutInMinutesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdleTimeoutInMinutes", GoMethod: "ResetIdleTimeoutInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleManagement", GoMethod: "ResetLifecycleManagement"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxIdleTimeoutInMinutes", GoMethod: "ResetMaxIdleTimeoutInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinIdleTimeoutInMinutes", GoMethod: "ResetMinIdleTimeoutInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idleSettings", GoGetter: "IdleSettings"},
			_jsii_.MemberProperty{JsiiProperty: "idleSettingsInput", GoGetter: "IdleSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdleSettings", GoMethod: "PutIdleSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdleSettings", GoMethod: "ResetIdleSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsCustomImages",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsCustomImages)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsCustomImagesList",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsCustomImagesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsCustomImagesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsCustomImagesOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsCustomImagesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigName", GoGetter: "AppImageConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigNameInput", GoGetter: "AppImageConfigNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "imageName", GoGetter: "ImageName"},
			_jsii_.MemberProperty{JsiiProperty: "imageNameInput", GoGetter: "ImageNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionNumber", GoGetter: "ImageVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionNumberInput", GoGetter: "ImageVersionNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppImageConfigName", GoMethod: "ResetAppImageConfigName"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageName", GoMethod: "ResetImageName"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageVersionNumber", GoMethod: "ResetImageVersionNumber"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsCustomImagesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageArn", GoMethod: "ResetSageMakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageVersionArn", GoMethod: "ResetSageMakerImageVersionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrainingPlanArn", GoMethod: "ResetTrainingPlanArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArn", GoGetter: "SageMakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArnInput", GoGetter: "SageMakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArn", GoGetter: "SageMakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArnInput", GoGetter: "SageMakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArn", GoGetter: "TrainingPlanArn"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArnInput", GoGetter: "TrainingPlanArnInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "appLifecycleManagement", GoGetter: "AppLifecycleManagement"},
			_jsii_.MemberProperty{JsiiProperty: "appLifecycleManagementInput", GoGetter: "AppLifecycleManagementInput"},
			_jsii_.MemberProperty{JsiiProperty: "builtInLifecycleConfigArn", GoGetter: "BuiltInLifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "builtInLifecycleConfigArnInput", GoGetter: "BuiltInLifecycleConfigArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customImages", GoGetter: "CustomImages"},
			_jsii_.MemberProperty{JsiiProperty: "customImagesInput", GoGetter: "CustomImagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArns", GoGetter: "LifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnsInput", GoGetter: "LifecycleConfigArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAppLifecycleManagement", GoMethod: "PutAppLifecycleManagement"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomImages", GoMethod: "PutCustomImages"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppLifecycleManagement", GoMethod: "ResetAppLifecycleManagement"},
			_jsii_.MemberMethod{JsiiMethod: "resetBuiltInLifecycleConfigArn", GoMethod: "ResetBuiltInLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomImages", GoMethod: "ResetCustomImages"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultResourceSpec", GoMethod: "ResetDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArns", GoMethod: "ResetLifecycleConfigArns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCustomFileSystemConfigs",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCustomFileSystemConfigs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsEfsFileSystemConfig",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsEfsFileSystemConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsEfsFileSystemConfigOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsEfsFileSystemConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemId", GoGetter: "FileSystemId"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemIdInput", GoGetter: "FileSystemIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemPath", GoGetter: "FileSystemPath"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemPathInput", GoGetter: "FileSystemPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileSystemId", GoMethod: "ResetFileSystemId"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileSystemPath", GoMethod: "ResetFileSystemPath"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsEfsFileSystemConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsFsxLustreFileSystemConfig",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsFsxLustreFileSystemConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsFsxLustreFileSystemConfigOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsFsxLustreFileSystemConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemId", GoGetter: "FileSystemId"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemIdInput", GoGetter: "FileSystemIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemPath", GoGetter: "FileSystemPath"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemPathInput", GoGetter: "FileSystemPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileSystemId", GoMethod: "ResetFileSystemId"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileSystemPath", GoMethod: "ResetFileSystemPath"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsFsxLustreFileSystemConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsList",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "efsFileSystemConfig", GoGetter: "EfsFileSystemConfig"},
			_jsii_.MemberProperty{JsiiProperty: "efsFileSystemConfigInput", GoGetter: "EfsFileSystemConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "fsxLustreFileSystemConfig", GoGetter: "FsxLustreFileSystemConfig"},
			_jsii_.MemberProperty{JsiiProperty: "fsxLustreFileSystemConfigInput", GoGetter: "FsxLustreFileSystemConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putEfsFileSystemConfig", GoMethod: "PutEfsFileSystemConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putFsxLustreFileSystemConfig", GoMethod: "PutFsxLustreFileSystemConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putS3FileSystemConfig", GoMethod: "PutS3FileSystemConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetEfsFileSystemConfig", GoMethod: "ResetEfsFileSystemConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetFsxLustreFileSystemConfig", GoMethod: "ResetFsxLustreFileSystemConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3FileSystemConfig", GoMethod: "ResetS3FileSystemConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3FileSystemConfig", GoGetter: "S3FileSystemConfig"},
			_jsii_.MemberProperty{JsiiProperty: "s3FileSystemConfigInput", GoGetter: "S3FileSystemConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsS3FileSystemConfig",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsS3FileSystemConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsS3FileSystemConfigOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsS3FileSystemConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mountPath", GoGetter: "MountPath"},
			_jsii_.MemberProperty{JsiiProperty: "mountPathInput", GoGetter: "MountPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMountPath", GoMethod: "ResetMountPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3Uri", GoMethod: "ResetS3Uri"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3Uri", GoGetter: "S3Uri"},
			_jsii_.MemberProperty{JsiiProperty: "s3UriInput", GoGetter: "S3UriInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsS3FileSystemConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCustomPosixUserConfig",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCustomPosixUserConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsCustomPosixUserConfigOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsCustomPosixUserConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "gid", GoGetter: "Gid"},
			_jsii_.MemberProperty{JsiiProperty: "gidInput", GoGetter: "GidInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetGid", GoMethod: "ResetGid"},
			_jsii_.MemberMethod{JsiiMethod: "resetUid", GoMethod: "ResetUid"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "uid", GoGetter: "Uid"},
			_jsii_.MemberProperty{JsiiProperty: "uidInput", GoGetter: "UidInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsCustomPosixUserConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagement",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagement)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idleTimeoutInMinutes", GoGetter: "IdleTimeoutInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "idleTimeoutInMinutesInput", GoGetter: "IdleTimeoutInMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleManagement", GoGetter: "LifecycleManagement"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleManagementInput", GoGetter: "LifecycleManagementInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxIdleTimeoutInMinutes", GoGetter: "MaxIdleTimeoutInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "maxIdleTimeoutInMinutesInput", GoGetter: "MaxIdleTimeoutInMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "minIdleTimeoutInMinutes", GoGetter: "MinIdleTimeoutInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "minIdleTimeoutInMinutesInput", GoGetter: "MinIdleTimeoutInMinutesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdleTimeoutInMinutes", GoMethod: "ResetIdleTimeoutInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleManagement", GoMethod: "ResetLifecycleManagement"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxIdleTimeoutInMinutes", GoMethod: "ResetMaxIdleTimeoutInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinIdleTimeoutInMinutes", GoMethod: "ResetMinIdleTimeoutInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idleSettings", GoGetter: "IdleSettings"},
			_jsii_.MemberProperty{JsiiProperty: "idleSettingsInput", GoGetter: "IdleSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdleSettings", GoMethod: "PutIdleSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdleSettings", GoMethod: "ResetIdleSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCodeRepositories",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCodeRepositories)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCodeRepositoriesList",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCodeRepositoriesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCodeRepositoriesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCodeRepositoriesOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCodeRepositoriesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryUrl", GoGetter: "RepositoryUrl"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryUrlInput", GoGetter: "RepositoryUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRepositoryUrl", GoMethod: "ResetRepositoryUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCodeRepositoriesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCustomImages",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCustomImages)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCustomImagesList",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCustomImagesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCustomImagesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCustomImagesOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCustomImagesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigName", GoGetter: "AppImageConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigNameInput", GoGetter: "AppImageConfigNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "imageName", GoGetter: "ImageName"},
			_jsii_.MemberProperty{JsiiProperty: "imageNameInput", GoGetter: "ImageNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionNumber", GoGetter: "ImageVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionNumberInput", GoGetter: "ImageVersionNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppImageConfigName", GoMethod: "ResetAppImageConfigName"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageName", GoMethod: "ResetImageName"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageVersionNumber", GoMethod: "ResetImageVersionNumber"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCustomImagesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageArn", GoMethod: "ResetSageMakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageVersionArn", GoMethod: "ResetSageMakerImageVersionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrainingPlanArn", GoMethod: "ResetTrainingPlanArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArn", GoGetter: "SageMakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArnInput", GoGetter: "SageMakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArn", GoGetter: "SageMakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArnInput", GoGetter: "SageMakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArn", GoGetter: "TrainingPlanArn"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArnInput", GoGetter: "TrainingPlanArnInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "appLifecycleManagement", GoGetter: "AppLifecycleManagement"},
			_jsii_.MemberProperty{JsiiProperty: "appLifecycleManagementInput", GoGetter: "AppLifecycleManagementInput"},
			_jsii_.MemberProperty{JsiiProperty: "builtInLifecycleConfigArn", GoGetter: "BuiltInLifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "builtInLifecycleConfigArnInput", GoGetter: "BuiltInLifecycleConfigArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "codeRepositories", GoGetter: "CodeRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "codeRepositoriesInput", GoGetter: "CodeRepositoriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customImages", GoGetter: "CustomImages"},
			_jsii_.MemberProperty{JsiiProperty: "customImagesInput", GoGetter: "CustomImagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArns", GoGetter: "LifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnsInput", GoGetter: "LifecycleConfigArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAppLifecycleManagement", GoMethod: "PutAppLifecycleManagement"},
			_jsii_.MemberMethod{JsiiMethod: "putCodeRepositories", GoMethod: "PutCodeRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomImages", GoMethod: "PutCustomImages"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppLifecycleManagement", GoMethod: "ResetAppLifecycleManagement"},
			_jsii_.MemberMethod{JsiiMethod: "resetBuiltInLifecycleConfigArn", GoMethod: "ResetBuiltInLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetCodeRepositories", GoMethod: "ResetCodeRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomImages", GoMethod: "ResetCustomImages"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultResourceSpec", GoMethod: "ResetDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArns", GoMethod: "ResetLifecycleConfigArns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterServerAppSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterServerAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageArn", GoMethod: "ResetSageMakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageVersionArn", GoMethod: "ResetSageMakerImageVersionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrainingPlanArn", GoMethod: "ResetTrainingPlanArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArn", GoGetter: "SageMakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArnInput", GoGetter: "SageMakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArn", GoGetter: "SageMakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArnInput", GoGetter: "SageMakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArn", GoGetter: "TrainingPlanArn"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArnInput", GoGetter: "TrainingPlanArnInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArns", GoGetter: "LifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnsInput", GoGetter: "LifecycleConfigArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultResourceSpec", GoMethod: "ResetDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArns", GoMethod: "ResetLifecycleConfigArns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsKernelGatewayAppSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImages",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImages)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImagesList",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImagesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImagesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImagesOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImagesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigName", GoGetter: "AppImageConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigNameInput", GoGetter: "AppImageConfigNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "imageName", GoGetter: "ImageName"},
			_jsii_.MemberProperty{JsiiProperty: "imageNameInput", GoGetter: "ImageNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionNumber", GoGetter: "ImageVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionNumberInput", GoGetter: "ImageVersionNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppImageConfigName", GoMethod: "ResetAppImageConfigName"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageName", GoMethod: "ResetImageName"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageVersionNumber", GoMethod: "ResetImageVersionNumber"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImagesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageArn", GoMethod: "ResetSageMakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageVersionArn", GoMethod: "ResetSageMakerImageVersionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrainingPlanArn", GoMethod: "ResetTrainingPlanArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArn", GoGetter: "SageMakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArnInput", GoGetter: "SageMakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArn", GoGetter: "SageMakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArnInput", GoGetter: "SageMakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArn", GoGetter: "TrainingPlanArn"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArnInput", GoGetter: "TrainingPlanArnInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customImages", GoGetter: "CustomImages"},
			_jsii_.MemberProperty{JsiiProperty: "customImagesInput", GoGetter: "CustomImagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArns", GoGetter: "LifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnsInput", GoGetter: "LifecycleConfigArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomImages", GoMethod: "PutCustomImages"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomImages", GoMethod: "ResetCustomImages"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultResourceSpec", GoMethod: "ResetDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArns", GoMethod: "ResetLifecycleConfigArns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoMountHomeEfs", GoGetter: "AutoMountHomeEfs"},
			_jsii_.MemberProperty{JsiiProperty: "autoMountHomeEfsInput", GoGetter: "AutoMountHomeEfsInput"},
			_jsii_.MemberProperty{JsiiProperty: "codeEditorAppSettings", GoGetter: "CodeEditorAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "codeEditorAppSettingsInput", GoGetter: "CodeEditorAppSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customFileSystemConfigs", GoGetter: "CustomFileSystemConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "customFileSystemConfigsInput", GoGetter: "CustomFileSystemConfigsInput"},
			_jsii_.MemberProperty{JsiiProperty: "customPosixUserConfig", GoGetter: "CustomPosixUserConfig"},
			_jsii_.MemberProperty{JsiiProperty: "customPosixUserConfigInput", GoGetter: "CustomPosixUserConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLandingUri", GoGetter: "DefaultLandingUri"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLandingUriInput", GoGetter: "DefaultLandingUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleInput", GoGetter: "ExecutionRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "jupyterLabAppSettings", GoGetter: "JupyterLabAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "jupyterLabAppSettingsInput", GoGetter: "JupyterLabAppSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "jupyterServerAppSettings", GoGetter: "JupyterServerAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "jupyterServerAppSettingsInput", GoGetter: "JupyterServerAppSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "kernelGatewayAppSettings", GoGetter: "KernelGatewayAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "kernelGatewayAppSettingsInput", GoGetter: "KernelGatewayAppSettingsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCodeEditorAppSettings", GoMethod: "PutCodeEditorAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomFileSystemConfigs", GoMethod: "PutCustomFileSystemConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomPosixUserConfig", GoMethod: "PutCustomPosixUserConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putJupyterLabAppSettings", GoMethod: "PutJupyterLabAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putJupyterServerAppSettings", GoMethod: "PutJupyterServerAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putKernelGatewayAppSettings", GoMethod: "PutKernelGatewayAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putRSessionAppSettings", GoMethod: "PutRSessionAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putRStudioServerProAppSettings", GoMethod: "PutRStudioServerProAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putSharingSettings", GoMethod: "PutSharingSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putSpaceStorageSettings", GoMethod: "PutSpaceStorageSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putStudioWebPortalSettings", GoMethod: "PutStudioWebPortalSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoMountHomeEfs", GoMethod: "ResetAutoMountHomeEfs"},
			_jsii_.MemberMethod{JsiiMethod: "resetCodeEditorAppSettings", GoMethod: "ResetCodeEditorAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomFileSystemConfigs", GoMethod: "ResetCustomFileSystemConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomPosixUserConfig", GoMethod: "ResetCustomPosixUserConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultLandingUri", GoMethod: "ResetDefaultLandingUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetJupyterLabAppSettings", GoMethod: "ResetJupyterLabAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetJupyterServerAppSettings", GoMethod: "ResetJupyterServerAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetKernelGatewayAppSettings", GoMethod: "ResetKernelGatewayAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetRSessionAppSettings", GoMethod: "ResetRSessionAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetRStudioServerProAppSettings", GoMethod: "ResetRStudioServerProAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroups", GoMethod: "ResetSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharingSettings", GoMethod: "ResetSharingSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpaceStorageSettings", GoMethod: "ResetSpaceStorageSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetStudioWebPortal", GoMethod: "ResetStudioWebPortal"},
			_jsii_.MemberMethod{JsiiMethod: "resetStudioWebPortalSettings", GoMethod: "ResetStudioWebPortalSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rSessionAppSettings", GoGetter: "RSessionAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "rSessionAppSettingsInput", GoGetter: "RSessionAppSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "rStudioServerProAppSettings", GoGetter: "RStudioServerProAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "rStudioServerProAppSettingsInput", GoGetter: "RStudioServerProAppSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupsInput", GoGetter: "SecurityGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "sharingSettings", GoGetter: "SharingSettings"},
			_jsii_.MemberProperty{JsiiProperty: "sharingSettingsInput", GoGetter: "SharingSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "spaceStorageSettings", GoGetter: "SpaceStorageSettings"},
			_jsii_.MemberProperty{JsiiProperty: "spaceStorageSettingsInput", GoGetter: "SpaceStorageSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "studioWebPortal", GoGetter: "StudioWebPortal"},
			_jsii_.MemberProperty{JsiiProperty: "studioWebPortalInput", GoGetter: "StudioWebPortalInput"},
			_jsii_.MemberProperty{JsiiProperty: "studioWebPortalSettings", GoGetter: "StudioWebPortalSettings"},
			_jsii_.MemberProperty{JsiiProperty: "studioWebPortalSettingsInput", GoGetter: "StudioWebPortalSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsRSessionAppSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsRSessionAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsRSessionAppSettingsCustomImages",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsRSessionAppSettingsCustomImages)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsRSessionAppSettingsCustomImagesList",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsRSessionAppSettingsCustomImagesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsRSessionAppSettingsCustomImagesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsRSessionAppSettingsCustomImagesOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsRSessionAppSettingsCustomImagesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigName", GoGetter: "AppImageConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigNameInput", GoGetter: "AppImageConfigNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "imageName", GoGetter: "ImageName"},
			_jsii_.MemberProperty{JsiiProperty: "imageNameInput", GoGetter: "ImageNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionNumber", GoGetter: "ImageVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionNumberInput", GoGetter: "ImageVersionNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppImageConfigName", GoMethod: "ResetAppImageConfigName"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageName", GoMethod: "ResetImageName"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageVersionNumber", GoMethod: "ResetImageVersionNumber"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsRSessionAppSettingsCustomImagesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsRSessionAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsRSessionAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsRSessionAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsRSessionAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageArn", GoMethod: "ResetSageMakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageVersionArn", GoMethod: "ResetSageMakerImageVersionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrainingPlanArn", GoMethod: "ResetTrainingPlanArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArn", GoGetter: "SageMakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArnInput", GoGetter: "SageMakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArn", GoGetter: "SageMakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArnInput", GoGetter: "SageMakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArn", GoGetter: "TrainingPlanArn"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArnInput", GoGetter: "TrainingPlanArnInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsRSessionAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsRSessionAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsRSessionAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customImages", GoGetter: "CustomImages"},
			_jsii_.MemberProperty{JsiiProperty: "customImagesInput", GoGetter: "CustomImagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomImages", GoMethod: "PutCustomImages"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomImages", GoMethod: "ResetCustomImages"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultResourceSpec", GoMethod: "ResetDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsRSessionAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsRStudioServerProAppSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsRStudioServerProAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsRStudioServerProAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsRStudioServerProAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessStatus", GoGetter: "AccessStatus"},
			_jsii_.MemberProperty{JsiiProperty: "accessStatusInput", GoGetter: "AccessStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessStatus", GoMethod: "ResetAccessStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserGroup", GoMethod: "ResetUserGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userGroup", GoGetter: "UserGroup"},
			_jsii_.MemberProperty{JsiiProperty: "userGroupInput", GoGetter: "UserGroupInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsRStudioServerProAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsSharingSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsSharingSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsSharingSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsSharingSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "notebookOutputOption", GoGetter: "NotebookOutputOption"},
			_jsii_.MemberProperty{JsiiProperty: "notebookOutputOptionInput", GoGetter: "NotebookOutputOptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotebookOutputOption", GoMethod: "ResetNotebookOutputOption"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3KmsKeyId", GoMethod: "ResetS3KmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3OutputPath", GoMethod: "ResetS3OutputPath"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3KmsKeyId", GoGetter: "S3KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "s3KmsKeyIdInput", GoGetter: "S3KmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputPath", GoGetter: "S3OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputPathInput", GoGetter: "S3OutputPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsSharingSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsSpaceStorageSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsSpaceStorageSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultEbsVolumeSizeInGb", GoGetter: "DefaultEbsVolumeSizeInGb"},
			_jsii_.MemberProperty{JsiiProperty: "defaultEbsVolumeSizeInGbInput", GoGetter: "DefaultEbsVolumeSizeInGbInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maximumEbsVolumeSizeInGb", GoGetter: "MaximumEbsVolumeSizeInGb"},
			_jsii_.MemberProperty{JsiiProperty: "maximumEbsVolumeSizeInGbInput", GoGetter: "MaximumEbsVolumeSizeInGbInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultEbsVolumeSizeInGb", GoMethod: "ResetDefaultEbsVolumeSizeInGb"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaximumEbsVolumeSizeInGb", GoMethod: "ResetMaximumEbsVolumeSizeInGb"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsSpaceStorageSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsSpaceStorageSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultEbsStorageSettings", GoGetter: "DefaultEbsStorageSettings"},
			_jsii_.MemberProperty{JsiiProperty: "defaultEbsStorageSettingsInput", GoGetter: "DefaultEbsStorageSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultEbsStorageSettings", GoMethod: "PutDefaultEbsStorageSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultEbsStorageSettings", GoMethod: "ResetDefaultEbsStorageSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsSpaceStorageSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsStudioWebPortalSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsStudioWebPortalSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsStudioWebPortalSettingsHiddenSageMakerImageVersionAliases",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsStudioWebPortalSettingsHiddenSageMakerImageVersionAliases)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsStudioWebPortalSettingsHiddenSageMakerImageVersionAliasesList",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsStudioWebPortalSettingsHiddenSageMakerImageVersionAliasesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsStudioWebPortalSettingsHiddenSageMakerImageVersionAliasesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsStudioWebPortalSettingsHiddenSageMakerImageVersionAliasesOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsStudioWebPortalSettingsHiddenSageMakerImageVersionAliasesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageName", GoMethod: "ResetSageMakerImageName"},
			_jsii_.MemberMethod{JsiiMethod: "resetVersionAliases", GoMethod: "ResetVersionAliases"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageName", GoGetter: "SageMakerImageName"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageNameInput", GoGetter: "SageMakerImageNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "versionAliases", GoGetter: "VersionAliases"},
			_jsii_.MemberProperty{JsiiProperty: "versionAliasesInput", GoGetter: "VersionAliasesInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsStudioWebPortalSettingsHiddenSageMakerImageVersionAliasesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultUserSettingsStudioWebPortalSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsStudioWebPortalSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hiddenAppTypes", GoGetter: "HiddenAppTypes"},
			_jsii_.MemberProperty{JsiiProperty: "hiddenAppTypesInput", GoGetter: "HiddenAppTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "hiddenInstanceTypes", GoGetter: "HiddenInstanceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "hiddenInstanceTypesInput", GoGetter: "HiddenInstanceTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "hiddenMlTools", GoGetter: "HiddenMlTools"},
			_jsii_.MemberProperty{JsiiProperty: "hiddenMlToolsInput", GoGetter: "HiddenMlToolsInput"},
			_jsii_.MemberProperty{JsiiProperty: "hiddenSageMakerImageVersionAliases", GoGetter: "HiddenSageMakerImageVersionAliases"},
			_jsii_.MemberProperty{JsiiProperty: "hiddenSageMakerImageVersionAliasesInput", GoGetter: "HiddenSageMakerImageVersionAliasesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putHiddenSageMakerImageVersionAliases", GoMethod: "PutHiddenSageMakerImageVersionAliases"},
			_jsii_.MemberMethod{JsiiMethod: "resetHiddenAppTypes", GoMethod: "ResetHiddenAppTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetHiddenInstanceTypes", GoMethod: "ResetHiddenInstanceTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetHiddenMlTools", GoMethod: "ResetHiddenMlTools"},
			_jsii_.MemberMethod{JsiiMethod: "resetHiddenSageMakerImageVersionAliases", GoMethod: "ResetHiddenSageMakerImageVersionAliases"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsStudioWebPortalSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDomainSettings",
		reflect.TypeOf((*SagemakerDomainDomainSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDomainSettingsDockerSettings",
		reflect.TypeOf((*SagemakerDomainDomainSettingsDockerSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDomainSettingsDockerSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDomainSettingsDockerSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableDockerAccess", GoGetter: "EnableDockerAccess"},
			_jsii_.MemberProperty{JsiiProperty: "enableDockerAccessInput", GoGetter: "EnableDockerAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableDockerAccess", GoMethod: "ResetEnableDockerAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcOnlyTrustedAccounts", GoMethod: "ResetVpcOnlyTrustedAccounts"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOnlyTrustedAccounts", GoGetter: "VpcOnlyTrustedAccounts"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOnlyTrustedAccountsInput", GoGetter: "VpcOnlyTrustedAccountsInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDomainSettingsDockerSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDomainSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDomainSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dockerSettings", GoGetter: "DockerSettings"},
			_jsii_.MemberProperty{JsiiProperty: "dockerSettingsInput", GoGetter: "DockerSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleIdentityConfig", GoGetter: "ExecutionRoleIdentityConfig"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleIdentityConfigInput", GoGetter: "ExecutionRoleIdentityConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddressType", GoGetter: "IpAddressType"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddressTypeInput", GoGetter: "IpAddressTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDockerSettings", GoMethod: "PutDockerSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putRStudioServerProDomainSettings", GoMethod: "PutRStudioServerProDomainSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putUnifiedStudioSettings", GoMethod: "PutUnifiedStudioSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetDockerSettings", GoMethod: "ResetDockerSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetExecutionRoleIdentityConfig", GoMethod: "ResetExecutionRoleIdentityConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpAddressType", GoMethod: "ResetIpAddressType"},
			_jsii_.MemberMethod{JsiiMethod: "resetRStudioServerProDomainSettings", GoMethod: "ResetRStudioServerProDomainSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroupIds", GoMethod: "ResetSecurityGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetUnifiedStudioSettings", GoMethod: "ResetUnifiedStudioSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rStudioServerProDomainSettings", GoGetter: "RStudioServerProDomainSettings"},
			_jsii_.MemberProperty{JsiiProperty: "rStudioServerProDomainSettingsInput", GoGetter: "RStudioServerProDomainSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdsInput", GoGetter: "SecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unifiedStudioSettings", GoGetter: "UnifiedStudioSettings"},
			_jsii_.MemberProperty{JsiiProperty: "unifiedStudioSettingsInput", GoGetter: "UnifiedStudioSettingsInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDomainSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDomainSettingsRStudioServerProDomainSettings",
		reflect.TypeOf((*SagemakerDomainDomainSettingsRStudioServerProDomainSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageArn", GoMethod: "ResetSageMakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSageMakerImageVersionArn", GoMethod: "ResetSageMakerImageVersionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrainingPlanArn", GoMethod: "ResetTrainingPlanArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArn", GoGetter: "SageMakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageArnInput", GoGetter: "SageMakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArn", GoGetter: "SageMakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sageMakerImageVersionArnInput", GoGetter: "SageMakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArn", GoGetter: "TrainingPlanArn"},
			_jsii_.MemberProperty{JsiiProperty: "trainingPlanArnInput", GoGetter: "TrainingPlanArnInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberProperty{JsiiProperty: "domainExecutionRoleArn", GoGetter: "DomainExecutionRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "domainExecutionRoleArnInput", GoGetter: "DomainExecutionRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultResourceSpec", GoMethod: "ResetDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomainExecutionRoleArn", GoMethod: "ResetDomainExecutionRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetRStudioConnectUrl", GoMethod: "ResetRStudioConnectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetRStudioPackageManagerUrl", GoMethod: "ResetRStudioPackageManagerUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rStudioConnectUrl", GoGetter: "RStudioConnectUrl"},
			_jsii_.MemberProperty{JsiiProperty: "rStudioConnectUrlInput", GoGetter: "RStudioConnectUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "rStudioPackageManagerUrl", GoGetter: "RStudioPackageManagerUrl"},
			_jsii_.MemberProperty{JsiiProperty: "rStudioPackageManagerUrlInput", GoGetter: "RStudioPackageManagerUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDomainSettingsUnifiedStudioSettings",
		reflect.TypeOf((*SagemakerDomainDomainSettingsUnifiedStudioSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domainAccountId", GoGetter: "DomainAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "domainAccountIdInput", GoGetter: "DomainAccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "domainId", GoGetter: "DomainId"},
			_jsii_.MemberProperty{JsiiProperty: "domainIdInput", GoGetter: "DomainIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "domainRegion", GoGetter: "DomainRegion"},
			_jsii_.MemberProperty{JsiiProperty: "domainRegionInput", GoGetter: "DomainRegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "environmentId", GoGetter: "EnvironmentId"},
			_jsii_.MemberProperty{JsiiProperty: "environmentIdInput", GoGetter: "EnvironmentIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "projectS3Path", GoGetter: "ProjectS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "projectS3PathInput", GoGetter: "ProjectS3PathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomainAccountId", GoMethod: "ResetDomainAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomainId", GoMethod: "ResetDomainId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomainRegion", GoMethod: "ResetDomainRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironmentId", GoMethod: "ResetEnvironmentId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProjectId", GoMethod: "ResetProjectId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProjectS3Path", GoMethod: "ResetProjectS3Path"},
			_jsii_.MemberMethod{JsiiMethod: "resetStudioWebPortalAccess", GoMethod: "ResetStudioWebPortalAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "studioWebPortalAccess", GoGetter: "StudioWebPortalAccess"},
			_jsii_.MemberProperty{JsiiProperty: "studioWebPortalAccessInput", GoGetter: "StudioWebPortalAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDomainSettingsUnifiedStudioSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainTags",
		reflect.TypeOf((*SagemakerDomainTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainTagsList",
		reflect.TypeOf((*SagemakerDomainTagsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainTagsOutputReference",
		reflect.TypeOf((*SagemakerDomainTagsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKey", GoMethod: "ResetKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
