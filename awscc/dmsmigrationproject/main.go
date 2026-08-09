// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsmigrationproject

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProject",
		reflect.TypeOf((*DmsMigrationProject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfileArn", GoGetter: "InstanceProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfileArnInput", GoGetter: "InstanceProfileArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfileIdentifier", GoGetter: "InstanceProfileIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfileIdentifierInput", GoGetter: "InstanceProfileIdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfileName", GoGetter: "InstanceProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfileNameInput", GoGetter: "InstanceProfileNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "migrationProjectArn", GoGetter: "MigrationProjectArn"},
			_jsii_.MemberProperty{JsiiProperty: "migrationProjectCreationTime", GoGetter: "MigrationProjectCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "migrationProjectCreationTimeInput", GoGetter: "MigrationProjectCreationTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "migrationProjectIdentifier", GoGetter: "MigrationProjectIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "migrationProjectIdentifierInput", GoGetter: "MigrationProjectIdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "migrationProjectName", GoGetter: "MigrationProjectName"},
			_jsii_.MemberProperty{JsiiProperty: "migrationProjectNameInput", GoGetter: "MigrationProjectNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putSchemaConversionApplicationAttributes", GoMethod: "PutSchemaConversionApplicationAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "putSourceDataProviderDescriptors", GoMethod: "PutSourceDataProviderDescriptors"},
			_jsii_.MemberMethod{JsiiMethod: "putTags", GoMethod: "PutTags"},
			_jsii_.MemberMethod{JsiiMethod: "putTargetDataProviderDescriptors", GoMethod: "PutTargetDataProviderDescriptors"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceProfileArn", GoMethod: "ResetInstanceProfileArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceProfileIdentifier", GoMethod: "ResetInstanceProfileIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceProfileName", GoMethod: "ResetInstanceProfileName"},
			_jsii_.MemberMethod{JsiiMethod: "resetMigrationProjectCreationTime", GoMethod: "ResetMigrationProjectCreationTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetMigrationProjectIdentifier", GoMethod: "ResetMigrationProjectIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetMigrationProjectName", GoMethod: "ResetMigrationProjectName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchemaConversionApplicationAttributes", GoMethod: "ResetSchemaConversionApplicationAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceDataProviderDescriptors", GoMethod: "ResetSourceDataProviderDescriptors"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetDataProviderDescriptors", GoMethod: "ResetTargetDataProviderDescriptors"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransformationRules", GoMethod: "ResetTransformationRules"},
			_jsii_.MemberProperty{JsiiProperty: "schemaConversionApplicationAttributes", GoGetter: "SchemaConversionApplicationAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "schemaConversionApplicationAttributesInput", GoGetter: "SchemaConversionApplicationAttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceDataProviderDescriptors", GoGetter: "SourceDataProviderDescriptors"},
			_jsii_.MemberProperty{JsiiProperty: "sourceDataProviderDescriptorsInput", GoGetter: "SourceDataProviderDescriptorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetDataProviderDescriptors", GoGetter: "TargetDataProviderDescriptors"},
			_jsii_.MemberProperty{JsiiProperty: "targetDataProviderDescriptorsInput", GoGetter: "TargetDataProviderDescriptorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "transformationRules", GoGetter: "TransformationRules"},
			_jsii_.MemberProperty{JsiiProperty: "transformationRulesInput", GoGetter: "TransformationRulesInput"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_DmsMigrationProject{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectConfig",
		reflect.TypeOf((*DmsMigrationProjectConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectSchemaConversionApplicationAttributes",
		reflect.TypeOf((*DmsMigrationProjectSchemaConversionApplicationAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectSchemaConversionApplicationAttributesOutputReference",
		reflect.TypeOf((*DmsMigrationProjectSchemaConversionApplicationAttributesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetS3BucketPath", GoMethod: "ResetS3BucketPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3BucketRoleArn", GoMethod: "ResetS3BucketRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3BucketPath", GoGetter: "S3BucketPath"},
			_jsii_.MemberProperty{JsiiProperty: "s3BucketPathInput", GoGetter: "S3BucketPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3BucketRoleArn", GoGetter: "S3BucketRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "s3BucketRoleArnInput", GoGetter: "S3BucketRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DmsMigrationProjectSchemaConversionApplicationAttributesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectSourceDataProviderDescriptors",
		reflect.TypeOf((*DmsMigrationProjectSourceDataProviderDescriptors)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectSourceDataProviderDescriptorsList",
		reflect.TypeOf((*DmsMigrationProjectSourceDataProviderDescriptorsList)(nil)).Elem(),
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
			j := jsiiProxy_DmsMigrationProjectSourceDataProviderDescriptorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectSourceDataProviderDescriptorsOutputReference",
		reflect.TypeOf((*DmsMigrationProjectSourceDataProviderDescriptorsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataProviderArn", GoGetter: "DataProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "dataProviderArnInput", GoGetter: "DataProviderArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataProviderIdentifier", GoGetter: "DataProviderIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "dataProviderIdentifierInput", GoGetter: "DataProviderIdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataProviderName", GoGetter: "DataProviderName"},
			_jsii_.MemberProperty{JsiiProperty: "dataProviderNameInput", GoGetter: "DataProviderNameInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDataProviderArn", GoMethod: "ResetDataProviderArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataProviderIdentifier", GoMethod: "ResetDataProviderIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataProviderName", GoMethod: "ResetDataProviderName"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretsManagerAccessRoleArn", GoMethod: "ResetSecretsManagerAccessRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretsManagerSecretId", GoMethod: "ResetSecretsManagerSecretId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secretsManagerAccessRoleArn", GoGetter: "SecretsManagerAccessRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "secretsManagerAccessRoleArnInput", GoGetter: "SecretsManagerAccessRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "secretsManagerSecretId", GoGetter: "SecretsManagerSecretId"},
			_jsii_.MemberProperty{JsiiProperty: "secretsManagerSecretIdInput", GoGetter: "SecretsManagerSecretIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DmsMigrationProjectSourceDataProviderDescriptorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectTags",
		reflect.TypeOf((*DmsMigrationProjectTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectTagsList",
		reflect.TypeOf((*DmsMigrationProjectTagsList)(nil)).Elem(),
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
			j := jsiiProxy_DmsMigrationProjectTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectTagsOutputReference",
		reflect.TypeOf((*DmsMigrationProjectTagsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_DmsMigrationProjectTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectTargetDataProviderDescriptors",
		reflect.TypeOf((*DmsMigrationProjectTargetDataProviderDescriptors)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectTargetDataProviderDescriptorsList",
		reflect.TypeOf((*DmsMigrationProjectTargetDataProviderDescriptorsList)(nil)).Elem(),
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
			j := jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectTargetDataProviderDescriptorsOutputReference",
		reflect.TypeOf((*DmsMigrationProjectTargetDataProviderDescriptorsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataProviderArn", GoGetter: "DataProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "dataProviderArnInput", GoGetter: "DataProviderArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataProviderIdentifier", GoGetter: "DataProviderIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "dataProviderIdentifierInput", GoGetter: "DataProviderIdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataProviderName", GoGetter: "DataProviderName"},
			_jsii_.MemberProperty{JsiiProperty: "dataProviderNameInput", GoGetter: "DataProviderNameInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDataProviderArn", GoMethod: "ResetDataProviderArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataProviderIdentifier", GoMethod: "ResetDataProviderIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataProviderName", GoMethod: "ResetDataProviderName"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretsManagerAccessRoleArn", GoMethod: "ResetSecretsManagerAccessRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretsManagerSecretId", GoMethod: "ResetSecretsManagerSecretId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secretsManagerAccessRoleArn", GoGetter: "SecretsManagerAccessRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "secretsManagerAccessRoleArnInput", GoGetter: "SecretsManagerAccessRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "secretsManagerSecretId", GoGetter: "SecretsManagerSecretId"},
			_jsii_.MemberProperty{JsiiProperty: "secretsManagerSecretIdInput", GoGetter: "SecretsManagerSecretIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
