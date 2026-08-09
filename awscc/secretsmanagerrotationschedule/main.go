// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretsmanagerrotationschedule

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.secretsmanagerRotationSchedule.SecretsmanagerRotationSchedule",
		reflect.TypeOf((*SecretsmanagerRotationSchedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "externalSecretRotationMetadata", GoGetter: "ExternalSecretRotationMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "externalSecretRotationMetadataInput", GoGetter: "ExternalSecretRotationMetadataInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalSecretRotationRoleArn", GoGetter: "ExternalSecretRotationRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "externalSecretRotationRoleArnInput", GoGetter: "ExternalSecretRotationRoleArnInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hostedRotationLambda", GoGetter: "HostedRotationLambda"},
			_jsii_.MemberProperty{JsiiProperty: "hostedRotationLambdaInput", GoGetter: "HostedRotationLambdaInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putExternalSecretRotationMetadata", GoMethod: "PutExternalSecretRotationMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "putHostedRotationLambda", GoMethod: "PutHostedRotationLambda"},
			_jsii_.MemberMethod{JsiiMethod: "putRotationRules", GoMethod: "PutRotationRules"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalSecretRotationMetadata", GoMethod: "ResetExternalSecretRotationMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalSecretRotationRoleArn", GoMethod: "ResetExternalSecretRotationRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostedRotationLambda", GoMethod: "ResetHostedRotationLambda"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotateImmediatelyOnUpdate", GoMethod: "ResetRotateImmediatelyOnUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationLambdaArn", GoMethod: "ResetRotationLambdaArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationRules", GoMethod: "ResetRotationRules"},
			_jsii_.MemberProperty{JsiiProperty: "rotateImmediatelyOnUpdate", GoGetter: "RotateImmediatelyOnUpdate"},
			_jsii_.MemberProperty{JsiiProperty: "rotateImmediatelyOnUpdateInput", GoGetter: "RotateImmediatelyOnUpdateInput"},
			_jsii_.MemberProperty{JsiiProperty: "rotationLambdaArn", GoGetter: "RotationLambdaArn"},
			_jsii_.MemberProperty{JsiiProperty: "rotationLambdaArnInput", GoGetter: "RotationLambdaArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "rotationRules", GoGetter: "RotationRules"},
			_jsii_.MemberProperty{JsiiProperty: "rotationRulesInput", GoGetter: "RotationRulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "rotationScheduleId", GoGetter: "RotationScheduleId"},
			_jsii_.MemberProperty{JsiiProperty: "secretId", GoGetter: "SecretId"},
			_jsii_.MemberProperty{JsiiProperty: "secretIdInput", GoGetter: "SecretIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_SecretsmanagerRotationSchedule{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.secretsmanagerRotationSchedule.SecretsmanagerRotationScheduleConfig",
		reflect.TypeOf((*SecretsmanagerRotationScheduleConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.secretsmanagerRotationSchedule.SecretsmanagerRotationScheduleExternalSecretRotationMetadata",
		reflect.TypeOf((*SecretsmanagerRotationScheduleExternalSecretRotationMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.secretsmanagerRotationSchedule.SecretsmanagerRotationScheduleExternalSecretRotationMetadataList",
		reflect.TypeOf((*SecretsmanagerRotationScheduleExternalSecretRotationMetadataList)(nil)).Elem(),
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
			j := jsiiProxy_SecretsmanagerRotationScheduleExternalSecretRotationMetadataList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.secretsmanagerRotationSchedule.SecretsmanagerRotationScheduleExternalSecretRotationMetadataOutputReference",
		reflect.TypeOf((*SecretsmanagerRotationScheduleExternalSecretRotationMetadataOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_SecretsmanagerRotationScheduleExternalSecretRotationMetadataOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.secretsmanagerRotationSchedule.SecretsmanagerRotationScheduleHostedRotationLambda",
		reflect.TypeOf((*SecretsmanagerRotationScheduleHostedRotationLambda)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.secretsmanagerRotationSchedule.SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference",
		reflect.TypeOf((*SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCharacters", GoGetter: "ExcludeCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCharactersInput", GoGetter: "ExcludeCharactersInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyArn", GoGetter: "KmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyArnInput", GoGetter: "KmsKeyArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "masterSecretArn", GoGetter: "MasterSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "masterSecretArnInput", GoGetter: "MasterSecretArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "masterSecretKmsKeyArn", GoGetter: "MasterSecretKmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "masterSecretKmsKeyArnInput", GoGetter: "MasterSecretKmsKeyArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludeCharacters", GoMethod: "ResetExcludeCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyArn", GoMethod: "ResetKmsKeyArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetMasterSecretArn", GoMethod: "ResetMasterSecretArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetMasterSecretKmsKeyArn", GoMethod: "ResetMasterSecretKmsKeyArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationLambdaName", GoMethod: "ResetRotationLambdaName"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationType", GoMethod: "ResetRotationType"},
			_jsii_.MemberMethod{JsiiMethod: "resetRuntime", GoMethod: "ResetRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuperuserSecretArn", GoMethod: "ResetSuperuserSecretArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuperuserSecretKmsKeyArn", GoMethod: "ResetSuperuserSecretKmsKeyArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcSecurityGroupIds", GoMethod: "ResetVpcSecurityGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcSubnetIds", GoMethod: "ResetVpcSubnetIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rotationLambdaName", GoGetter: "RotationLambdaName"},
			_jsii_.MemberProperty{JsiiProperty: "rotationLambdaNameInput", GoGetter: "RotationLambdaNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "rotationType", GoGetter: "RotationType"},
			_jsii_.MemberProperty{JsiiProperty: "rotationTypeInput", GoGetter: "RotationTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "runtime", GoGetter: "Runtime"},
			_jsii_.MemberProperty{JsiiProperty: "runtimeInput", GoGetter: "RuntimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "superuserSecretArn", GoGetter: "SuperuserSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "superuserSecretArnInput", GoGetter: "SuperuserSecretArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "superuserSecretKmsKeyArn", GoGetter: "SuperuserSecretKmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "superuserSecretKmsKeyArnInput", GoGetter: "SuperuserSecretKmsKeyArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSecurityGroupIds", GoGetter: "VpcSecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSecurityGroupIdsInput", GoGetter: "VpcSecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnetIds", GoGetter: "VpcSubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnetIdsInput", GoGetter: "VpcSubnetIdsInput"},
		},
		func() interface{} {
			j := jsiiProxy_SecretsmanagerRotationScheduleHostedRotationLambdaOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.secretsmanagerRotationSchedule.SecretsmanagerRotationScheduleRotationRules",
		reflect.TypeOf((*SecretsmanagerRotationScheduleRotationRules)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.secretsmanagerRotationSchedule.SecretsmanagerRotationScheduleRotationRulesOutputReference",
		reflect.TypeOf((*SecretsmanagerRotationScheduleRotationRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "automaticallyAfterDays", GoGetter: "AutomaticallyAfterDays"},
			_jsii_.MemberProperty{JsiiProperty: "automaticallyAfterDaysInput", GoGetter: "AutomaticallyAfterDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "duration", GoGetter: "Duration"},
			_jsii_.MemberProperty{JsiiProperty: "durationInput", GoGetter: "DurationInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAutomaticallyAfterDays", GoMethod: "ResetAutomaticallyAfterDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetDuration", GoMethod: "ResetDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduleExpression", GoMethod: "ResetScheduleExpression"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleExpression", GoGetter: "ScheduleExpression"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleExpressionInput", GoGetter: "ScheduleExpressionInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SecretsmanagerRotationScheduleRotationRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
