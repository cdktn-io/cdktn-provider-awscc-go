// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbproxytargetgroup

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.rdsDbProxyTargetGroup.RdsDbProxyTargetGroup",
		reflect.TypeOf((*RdsDbProxyTargetGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "connectionPoolConfigurationInfo", GoGetter: "ConnectionPoolConfigurationInfo"},
			_jsii_.MemberProperty{JsiiProperty: "connectionPoolConfigurationInfoInput", GoGetter: "ConnectionPoolConfigurationInfoInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dbClusterIdentifiers", GoGetter: "DbClusterIdentifiers"},
			_jsii_.MemberProperty{JsiiProperty: "dbClusterIdentifiersInput", GoGetter: "DbClusterIdentifiersInput"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceIdentifiers", GoGetter: "DbInstanceIdentifiers"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceIdentifiersInput", GoGetter: "DbInstanceIdentifiersInput"},
			_jsii_.MemberProperty{JsiiProperty: "dbProxyName", GoGetter: "DbProxyName"},
			_jsii_.MemberProperty{JsiiProperty: "dbProxyNameInput", GoGetter: "DbProxyNameInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putConnectionPoolConfigurationInfo", GoMethod: "PutConnectionPoolConfigurationInfo"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionPoolConfigurationInfo", GoMethod: "ResetConnectionPoolConfigurationInfo"},
			_jsii_.MemberMethod{JsiiMethod: "resetDbClusterIdentifiers", GoMethod: "ResetDbClusterIdentifiers"},
			_jsii_.MemberMethod{JsiiMethod: "resetDbInstanceIdentifiers", GoMethod: "ResetDbInstanceIdentifiers"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupArn", GoGetter: "TargetGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupName", GoGetter: "TargetGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupNameInput", GoGetter: "TargetGroupNameInput"},
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
			j := jsiiProxy_RdsDbProxyTargetGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.rdsDbProxyTargetGroup.RdsDbProxyTargetGroupConfig",
		reflect.TypeOf((*RdsDbProxyTargetGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.rdsDbProxyTargetGroup.RdsDbProxyTargetGroupConnectionPoolConfigurationInfo",
		reflect.TypeOf((*RdsDbProxyTargetGroupConnectionPoolConfigurationInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.rdsDbProxyTargetGroup.RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference",
		reflect.TypeOf((*RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "connectionBorrowTimeout", GoGetter: "ConnectionBorrowTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "connectionBorrowTimeoutInput", GoGetter: "ConnectionBorrowTimeoutInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "initQuery", GoGetter: "InitQuery"},
			_jsii_.MemberProperty{JsiiProperty: "initQueryInput", GoGetter: "InitQueryInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxConnectionsPercent", GoGetter: "MaxConnectionsPercent"},
			_jsii_.MemberProperty{JsiiProperty: "maxConnectionsPercentInput", GoGetter: "MaxConnectionsPercentInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxIdleConnectionsPercent", GoGetter: "MaxIdleConnectionsPercent"},
			_jsii_.MemberProperty{JsiiProperty: "maxIdleConnectionsPercentInput", GoGetter: "MaxIdleConnectionsPercentInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionBorrowTimeout", GoMethod: "ResetConnectionBorrowTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetInitQuery", GoMethod: "ResetInitQuery"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxConnectionsPercent", GoMethod: "ResetMaxConnectionsPercent"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxIdleConnectionsPercent", GoMethod: "ResetMaxIdleConnectionsPercent"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionPinningFilters", GoMethod: "ResetSessionPinningFilters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sessionPinningFilters", GoGetter: "SessionPinningFilters"},
			_jsii_.MemberProperty{JsiiProperty: "sessionPinningFiltersInput", GoGetter: "SessionPinningFiltersInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
