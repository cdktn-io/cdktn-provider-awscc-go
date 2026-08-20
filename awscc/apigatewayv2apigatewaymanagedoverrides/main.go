// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2apigatewaymanagedoverrides

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverrides",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverrides)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiGatewayManagedOverridesId", GoGetter: "ApiGatewayManagedOverridesId"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberProperty{JsiiProperty: "apiIdInput", GoGetter: "ApiIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
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
			_jsii_.MemberProperty{JsiiProperty: "integration", GoGetter: "Integration"},
			_jsii_.MemberProperty{JsiiProperty: "integrationInput", GoGetter: "IntegrationInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putIntegration", GoMethod: "PutIntegration"},
			_jsii_.MemberMethod{JsiiMethod: "putRoute", GoMethod: "PutRoute"},
			_jsii_.MemberMethod{JsiiMethod: "putStage", GoMethod: "PutStage"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntegration", GoMethod: "ResetIntegration"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoute", GoMethod: "ResetRoute"},
			_jsii_.MemberMethod{JsiiMethod: "resetStage", GoMethod: "ResetStage"},
			_jsii_.MemberProperty{JsiiProperty: "route", GoGetter: "Route"},
			_jsii_.MemberProperty{JsiiProperty: "routeInput", GoGetter: "RouteInput"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberProperty{JsiiProperty: "stageInput", GoGetter: "StageInput"},
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
			j := jsiiProxy_Apigatewayv2ApiGatewayManagedOverrides{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesConfig",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesIntegration",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesIntegration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "integrationMethod", GoGetter: "IntegrationMethod"},
			_jsii_.MemberProperty{JsiiProperty: "integrationMethodInput", GoGetter: "IntegrationMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "payloadFormatVersion", GoGetter: "PayloadFormatVersion"},
			_jsii_.MemberProperty{JsiiProperty: "payloadFormatVersionInput", GoGetter: "PayloadFormatVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntegrationMethod", GoMethod: "ResetIntegrationMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetPayloadFormatVersion", GoMethod: "ResetPayloadFormatVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeoutInMillis", GoMethod: "ResetTimeoutInMillis"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInMillis", GoGetter: "TimeoutInMillis"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInMillisInput", GoGetter: "TimeoutInMillisInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesIntegrationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesRoute",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesRoute)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesRouteOutputReference",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesRouteOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorizationScopes", GoGetter: "AuthorizationScopes"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationScopesInput", GoGetter: "AuthorizationScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationType", GoGetter: "AuthorizationType"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationTypeInput", GoGetter: "AuthorizationTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerId", GoGetter: "AuthorizerId"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerIdInput", GoGetter: "AuthorizerIdInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "operationName", GoGetter: "OperationName"},
			_jsii_.MemberProperty{JsiiProperty: "operationNameInput", GoGetter: "OperationNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorizationScopes", GoMethod: "ResetAuthorizationScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorizationType", GoMethod: "ResetAuthorizationType"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorizerId", GoMethod: "ResetAuthorizerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOperationName", GoMethod: "ResetOperationName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTarget", GoMethod: "ResetTarget"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberProperty{JsiiProperty: "targetInput", GoGetter: "TargetInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesRouteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesStage",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesStage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesStageAccessLogSettings",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesStageAccessLogSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesStageAccessLogSettingsOutputReference",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesStageAccessLogSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinationArn", GoGetter: "DestinationArn"},
			_jsii_.MemberProperty{JsiiProperty: "destinationArnInput", GoGetter: "DestinationArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberProperty{JsiiProperty: "formatInput", GoGetter: "FormatInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDestinationArn", GoMethod: "ResetDestinationArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetFormat", GoMethod: "ResetFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageAccessLogSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettings",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettingsOutputReference",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "detailedMetricsEnabled", GoGetter: "DetailedMetricsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "detailedMetricsEnabledInput", GoGetter: "DetailedMetricsEnabledInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDetailedMetricsEnabled", GoMethod: "ResetDetailedMetricsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetThrottlingBurstLimit", GoMethod: "ResetThrottlingBurstLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetThrottlingRateLimit", GoMethod: "ResetThrottlingRateLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "throttlingBurstLimit", GoGetter: "ThrottlingBurstLimit"},
			_jsii_.MemberProperty{JsiiProperty: "throttlingBurstLimitInput", GoGetter: "ThrottlingBurstLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "throttlingRateLimit", GoGetter: "ThrottlingRateLimit"},
			_jsii_.MemberProperty{JsiiProperty: "throttlingRateLimitInput", GoGetter: "ThrottlingRateLimitInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesStageOutputReference",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesStageOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessLogSettings", GoGetter: "AccessLogSettings"},
			_jsii_.MemberProperty{JsiiProperty: "accessLogSettingsInput", GoGetter: "AccessLogSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoDeploy", GoGetter: "AutoDeploy"},
			_jsii_.MemberProperty{JsiiProperty: "autoDeployInput", GoGetter: "AutoDeployInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteSettings", GoGetter: "DefaultRouteSettings"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteSettingsInput", GoGetter: "DefaultRouteSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putAccessLogSettings", GoMethod: "PutAccessLogSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultRouteSettings", GoMethod: "PutDefaultRouteSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putRouteSettings", GoMethod: "PutRouteSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessLogSettings", GoMethod: "ResetAccessLogSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoDeploy", GoMethod: "ResetAutoDeploy"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultRouteSettings", GoMethod: "ResetDefaultRouteSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetRouteSettings", GoMethod: "ResetRouteSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetStageVariables", GoMethod: "ResetStageVariables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "routeSettings", GoGetter: "RouteSettings"},
			_jsii_.MemberProperty{JsiiProperty: "routeSettingsInput", GoGetter: "RouteSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "stageVariables", GoGetter: "StageVariables"},
			_jsii_.MemberProperty{JsiiProperty: "stageVariablesInput", GoGetter: "StageVariablesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesStageRouteSettings",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesStageRouteSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsMap{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexMap)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.apigatewayv2ApiGatewayManagedOverrides.Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsOutputReference",
		reflect.TypeOf((*Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "detailedMetricsEnabled", GoGetter: "DetailedMetricsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "detailedMetricsEnabledInput", GoGetter: "DetailedMetricsEnabledInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDetailedMetricsEnabled", GoMethod: "ResetDetailedMetricsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetThrottlingBurstLimit", GoMethod: "ResetThrottlingBurstLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetThrottlingRateLimit", GoMethod: "ResetThrottlingRateLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "throttlingBurstLimit", GoGetter: "ThrottlingBurstLimit"},
			_jsii_.MemberProperty{JsiiProperty: "throttlingBurstLimitInput", GoGetter: "ThrottlingBurstLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "throttlingRateLimit", GoGetter: "ThrottlingRateLimit"},
			_jsii_.MemberProperty{JsiiProperty: "throttlingRateLimitInput", GoGetter: "ThrottlingRateLimitInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Apigatewayv2ApiGatewayManagedOverridesStageRouteSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
