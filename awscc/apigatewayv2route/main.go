// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2route

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.apigatewayv2Route.Apigatewayv2Route",
		reflect.TypeOf((*Apigatewayv2Route)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberProperty{JsiiProperty: "apiIdInput", GoGetter: "ApiIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyRequired", GoGetter: "ApiKeyRequired"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyRequiredInput", GoGetter: "ApiKeyRequiredInput"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationScopes", GoGetter: "AuthorizationScopes"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationScopesInput", GoGetter: "AuthorizationScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationType", GoGetter: "AuthorizationType"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationTypeInput", GoGetter: "AuthorizationTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerId", GoGetter: "AuthorizerId"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerIdInput", GoGetter: "AuthorizerIdInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "modelSelectionExpression", GoGetter: "ModelSelectionExpression"},
			_jsii_.MemberProperty{JsiiProperty: "modelSelectionExpressionInput", GoGetter: "ModelSelectionExpressionInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "operationName", GoGetter: "OperationName"},
			_jsii_.MemberProperty{JsiiProperty: "operationNameInput", GoGetter: "OperationNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberProperty{JsiiProperty: "requestModels", GoGetter: "RequestModels"},
			_jsii_.MemberProperty{JsiiProperty: "requestModelsInput", GoGetter: "RequestModelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestParameters", GoGetter: "RequestParameters"},
			_jsii_.MemberProperty{JsiiProperty: "requestParametersInput", GoGetter: "RequestParametersInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiKeyRequired", GoMethod: "ResetApiKeyRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorizationScopes", GoMethod: "ResetAuthorizationScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorizationType", GoMethod: "ResetAuthorizationType"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorizerId", GoMethod: "ResetAuthorizerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetModelSelectionExpression", GoMethod: "ResetModelSelectionExpression"},
			_jsii_.MemberMethod{JsiiMethod: "resetOperationName", GoMethod: "ResetOperationName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestModels", GoMethod: "ResetRequestModels"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestParameters", GoMethod: "ResetRequestParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetRouteResponseSelectionExpression", GoMethod: "ResetRouteResponseSelectionExpression"},
			_jsii_.MemberMethod{JsiiMethod: "resetTarget", GoMethod: "ResetTarget"},
			_jsii_.MemberProperty{JsiiProperty: "routeId", GoGetter: "RouteId"},
			_jsii_.MemberProperty{JsiiProperty: "routeKey", GoGetter: "RouteKey"},
			_jsii_.MemberProperty{JsiiProperty: "routeKeyInput", GoGetter: "RouteKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "routeResponseSelectionExpression", GoGetter: "RouteResponseSelectionExpression"},
			_jsii_.MemberProperty{JsiiProperty: "routeResponseSelectionExpressionInput", GoGetter: "RouteResponseSelectionExpressionInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberProperty{JsiiProperty: "targetInput", GoGetter: "TargetInput"},
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
			j := jsiiProxy_Apigatewayv2Route{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.apigatewayv2Route.Apigatewayv2RouteConfig",
		reflect.TypeOf((*Apigatewayv2RouteConfig)(nil)).Elem(),
	)
}
