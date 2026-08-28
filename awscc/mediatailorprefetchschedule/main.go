// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchSchedule",
		reflect.TypeOf((*MediatailorPrefetchSchedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "consumption", GoGetter: "Consumption"},
			_jsii_.MemberProperty{JsiiProperty: "consumptionInput", GoGetter: "ConsumptionInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "playbackConfigurationName", GoGetter: "PlaybackConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "playbackConfigurationNameInput", GoGetter: "PlaybackConfigurationNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putConsumption", GoMethod: "PutConsumption"},
			_jsii_.MemberMethod{JsiiMethod: "putRecurringPrefetchConfiguration", GoMethod: "PutRecurringPrefetchConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putRetrieval", GoMethod: "PutRetrieval"},
			_jsii_.MemberMethod{JsiiMethod: "putTags", GoMethod: "PutTags"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "recurringPrefetchConfiguration", GoGetter: "RecurringPrefetchConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "recurringPrefetchConfigurationInput", GoGetter: "RecurringPrefetchConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetConsumption", GoMethod: "ResetConsumption"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecurringPrefetchConfiguration", GoMethod: "ResetRecurringPrefetchConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetrieval", GoMethod: "ResetRetrieval"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduleType", GoMethod: "ResetScheduleType"},
			_jsii_.MemberMethod{JsiiMethod: "resetStreamId", GoMethod: "ResetStreamId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberProperty{JsiiProperty: "retrieval", GoGetter: "Retrieval"},
			_jsii_.MemberProperty{JsiiProperty: "retrievalInput", GoGetter: "RetrievalInput"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleType", GoGetter: "ScheduleType"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleTypeInput", GoGetter: "ScheduleTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "streamId", GoGetter: "StreamId"},
			_jsii_.MemberProperty{JsiiProperty: "streamIdInput", GoGetter: "StreamIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
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
			j := jsiiProxy_MediatailorPrefetchSchedule{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleConfig",
		reflect.TypeOf((*MediatailorPrefetchScheduleConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleConsumption",
		reflect.TypeOf((*MediatailorPrefetchScheduleConsumption)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleConsumptionAvailMatchingCriteria",
		reflect.TypeOf((*MediatailorPrefetchScheduleConsumptionAvailMatchingCriteria)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleConsumptionAvailMatchingCriteriaList",
		reflect.TypeOf((*MediatailorPrefetchScheduleConsumptionAvailMatchingCriteriaList)(nil)).Elem(),
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
			j := jsiiProxy_MediatailorPrefetchScheduleConsumptionAvailMatchingCriteriaList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleConsumptionAvailMatchingCriteriaOutputReference",
		reflect.TypeOf((*MediatailorPrefetchScheduleConsumptionAvailMatchingCriteriaOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicVariable", GoGetter: "DynamicVariable"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicVariableInput", GoGetter: "DynamicVariableInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "operatorInput", GoGetter: "OperatorInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDynamicVariable", GoMethod: "ResetDynamicVariable"},
			_jsii_.MemberMethod{JsiiMethod: "resetOperator", GoMethod: "ResetOperator"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MediatailorPrefetchScheduleConsumptionAvailMatchingCriteriaOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleConsumptionOutputReference",
		reflect.TypeOf((*MediatailorPrefetchScheduleConsumptionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availMatchingCriteria", GoGetter: "AvailMatchingCriteria"},
			_jsii_.MemberProperty{JsiiProperty: "availMatchingCriteriaInput", GoGetter: "AvailMatchingCriteriaInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endTime", GoGetter: "EndTime"},
			_jsii_.MemberProperty{JsiiProperty: "endTimeInput", GoGetter: "EndTimeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putAvailMatchingCriteria", GoMethod: "PutAvailMatchingCriteria"},
			_jsii_.MemberMethod{JsiiMethod: "resetAvailMatchingCriteria", GoMethod: "ResetAvailMatchingCriteria"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndTime", GoMethod: "ResetEndTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartTime", GoMethod: "ResetStartTime"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "startTime", GoGetter: "StartTime"},
			_jsii_.MemberProperty{JsiiProperty: "startTimeInput", GoGetter: "StartTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MediatailorPrefetchScheduleConsumptionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfiguration",
		reflect.TypeOf((*MediatailorPrefetchScheduleRecurringPrefetchConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference",
		reflect.TypeOf((*MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endTime", GoGetter: "EndTime"},
			_jsii_.MemberProperty{JsiiProperty: "endTimeInput", GoGetter: "EndTimeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putRecurringConsumption", GoMethod: "PutRecurringConsumption"},
			_jsii_.MemberMethod{JsiiMethod: "putRecurringRetrieval", GoMethod: "PutRecurringRetrieval"},
			_jsii_.MemberProperty{JsiiProperty: "recurringConsumption", GoGetter: "RecurringConsumption"},
			_jsii_.MemberProperty{JsiiProperty: "recurringConsumptionInput", GoGetter: "RecurringConsumptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "recurringRetrieval", GoGetter: "RecurringRetrieval"},
			_jsii_.MemberProperty{JsiiProperty: "recurringRetrievalInput", GoGetter: "RecurringRetrievalInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndTime", GoMethod: "ResetEndTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecurringConsumption", GoMethod: "ResetRecurringConsumption"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecurringRetrieval", GoMethod: "ResetRecurringRetrieval"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartTime", GoMethod: "ResetStartTime"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "startTime", GoGetter: "StartTime"},
			_jsii_.MemberProperty{JsiiProperty: "startTimeInput", GoGetter: "StartTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumption",
		reflect.TypeOf((*MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumption)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionAvailMatchingCriteria",
		reflect.TypeOf((*MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionAvailMatchingCriteria)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionAvailMatchingCriteriaList",
		reflect.TypeOf((*MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionAvailMatchingCriteriaList)(nil)).Elem(),
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
			j := jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionAvailMatchingCriteriaList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionAvailMatchingCriteriaOutputReference",
		reflect.TypeOf((*MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionAvailMatchingCriteriaOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicVariable", GoGetter: "DynamicVariable"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicVariableInput", GoGetter: "DynamicVariableInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "operatorInput", GoGetter: "OperatorInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDynamicVariable", GoMethod: "ResetDynamicVariable"},
			_jsii_.MemberMethod{JsiiMethod: "resetOperator", GoMethod: "ResetOperator"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionAvailMatchingCriteriaOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionOutputReference",
		reflect.TypeOf((*MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availMatchingCriteria", GoGetter: "AvailMatchingCriteria"},
			_jsii_.MemberProperty{JsiiProperty: "availMatchingCriteriaInput", GoGetter: "AvailMatchingCriteriaInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putAvailMatchingCriteria", GoMethod: "PutAvailMatchingCriteria"},
			_jsii_.MemberMethod{JsiiMethod: "resetAvailMatchingCriteria", GoMethod: "ResetAvailMatchingCriteria"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetrievedAdExpirationSeconds", GoMethod: "ResetRetrievedAdExpirationSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retrievedAdExpirationSeconds", GoGetter: "RetrievedAdExpirationSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "retrievedAdExpirationSecondsInput", GoGetter: "RetrievedAdExpirationSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrieval",
		reflect.TypeOf((*MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrieval)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference",
		reflect.TypeOf((*MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delayAfterAvailEndSeconds", GoGetter: "DelayAfterAvailEndSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "delayAfterAvailEndSecondsInput", GoGetter: "DelayAfterAvailEndSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicVariables", GoGetter: "DynamicVariables"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicVariablesInput", GoGetter: "DynamicVariablesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putTrafficShapingRetrievalWindow", GoMethod: "PutTrafficShapingRetrievalWindow"},
			_jsii_.MemberMethod{JsiiMethod: "putTrafficShapingTpsConfiguration", GoMethod: "PutTrafficShapingTpsConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelayAfterAvailEndSeconds", GoMethod: "ResetDelayAfterAvailEndSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetDynamicVariables", GoMethod: "ResetDynamicVariables"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrafficShapingRetrievalWindow", GoMethod: "ResetTrafficShapingRetrievalWindow"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrafficShapingTpsConfiguration", GoMethod: "ResetTrafficShapingTpsConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrafficShapingType", GoMethod: "ResetTrafficShapingType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trafficShapingRetrievalWindow", GoGetter: "TrafficShapingRetrievalWindow"},
			_jsii_.MemberProperty{JsiiProperty: "trafficShapingRetrievalWindowInput", GoGetter: "TrafficShapingRetrievalWindowInput"},
			_jsii_.MemberProperty{JsiiProperty: "trafficShapingTpsConfiguration", GoGetter: "TrafficShapingTpsConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "trafficShapingTpsConfigurationInput", GoGetter: "TrafficShapingTpsConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "trafficShapingType", GoGetter: "TrafficShapingType"},
			_jsii_.MemberProperty{JsiiProperty: "trafficShapingTypeInput", GoGetter: "TrafficShapingTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindow",
		reflect.TypeOf((*MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindow)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference",
		reflect.TypeOf((*MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetRetrievalWindowDurationSeconds", GoMethod: "ResetRetrievalWindowDurationSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retrievalWindowDurationSeconds", GoGetter: "RetrievalWindowDurationSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "retrievalWindowDurationSecondsInput", GoGetter: "RetrievalWindowDurationSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingTpsConfiguration",
		reflect.TypeOf((*MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingTpsConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingTpsConfigurationOutputReference",
		reflect.TypeOf((*MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingTpsConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "peakConcurrentUsers", GoGetter: "PeakConcurrentUsers"},
			_jsii_.MemberProperty{JsiiProperty: "peakConcurrentUsersInput", GoGetter: "PeakConcurrentUsersInput"},
			_jsii_.MemberProperty{JsiiProperty: "peakTps", GoGetter: "PeakTps"},
			_jsii_.MemberProperty{JsiiProperty: "peakTpsInput", GoGetter: "PeakTpsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPeakConcurrentUsers", GoMethod: "ResetPeakConcurrentUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resetPeakTps", GoMethod: "ResetPeakTps"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingTpsConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRetrieval",
		reflect.TypeOf((*MediatailorPrefetchScheduleRetrieval)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRetrievalOutputReference",
		reflect.TypeOf((*MediatailorPrefetchScheduleRetrievalOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicVariables", GoGetter: "DynamicVariables"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicVariablesInput", GoGetter: "DynamicVariablesInput"},
			_jsii_.MemberProperty{JsiiProperty: "endTime", GoGetter: "EndTime"},
			_jsii_.MemberProperty{JsiiProperty: "endTimeInput", GoGetter: "EndTimeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putTrafficShapingRetrievalWindow", GoMethod: "PutTrafficShapingRetrievalWindow"},
			_jsii_.MemberMethod{JsiiMethod: "putTrafficShapingTpsConfiguration", GoMethod: "PutTrafficShapingTpsConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetDynamicVariables", GoMethod: "ResetDynamicVariables"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndTime", GoMethod: "ResetEndTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartTime", GoMethod: "ResetStartTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrafficShapingRetrievalWindow", GoMethod: "ResetTrafficShapingRetrievalWindow"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrafficShapingTpsConfiguration", GoMethod: "ResetTrafficShapingTpsConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrafficShapingType", GoMethod: "ResetTrafficShapingType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "startTime", GoGetter: "StartTime"},
			_jsii_.MemberProperty{JsiiProperty: "startTimeInput", GoGetter: "StartTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trafficShapingRetrievalWindow", GoGetter: "TrafficShapingRetrievalWindow"},
			_jsii_.MemberProperty{JsiiProperty: "trafficShapingRetrievalWindowInput", GoGetter: "TrafficShapingRetrievalWindowInput"},
			_jsii_.MemberProperty{JsiiProperty: "trafficShapingTpsConfiguration", GoGetter: "TrafficShapingTpsConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "trafficShapingTpsConfigurationInput", GoGetter: "TrafficShapingTpsConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "trafficShapingType", GoGetter: "TrafficShapingType"},
			_jsii_.MemberProperty{JsiiProperty: "trafficShapingTypeInput", GoGetter: "TrafficShapingTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRetrievalTrafficShapingRetrievalWindow",
		reflect.TypeOf((*MediatailorPrefetchScheduleRetrievalTrafficShapingRetrievalWindow)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRetrievalTrafficShapingRetrievalWindowOutputReference",
		reflect.TypeOf((*MediatailorPrefetchScheduleRetrievalTrafficShapingRetrievalWindowOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetRetrievalWindowDurationSeconds", GoMethod: "ResetRetrievalWindowDurationSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retrievalWindowDurationSeconds", GoGetter: "RetrievalWindowDurationSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "retrievalWindowDurationSecondsInput", GoGetter: "RetrievalWindowDurationSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingRetrievalWindowOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfiguration",
		reflect.TypeOf((*MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference",
		reflect.TypeOf((*MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "peakConcurrentUsers", GoGetter: "PeakConcurrentUsers"},
			_jsii_.MemberProperty{JsiiProperty: "peakConcurrentUsersInput", GoGetter: "PeakConcurrentUsersInput"},
			_jsii_.MemberProperty{JsiiProperty: "peakTps", GoGetter: "PeakTps"},
			_jsii_.MemberProperty{JsiiProperty: "peakTpsInput", GoGetter: "PeakTpsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPeakConcurrentUsers", GoMethod: "ResetPeakConcurrentUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resetPeakTps", GoMethod: "ResetPeakTps"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleTags",
		reflect.TypeOf((*MediatailorPrefetchScheduleTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleTagsList",
		reflect.TypeOf((*MediatailorPrefetchScheduleTagsList)(nil)).Elem(),
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
			j := jsiiProxy_MediatailorPrefetchScheduleTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleTagsOutputReference",
		reflect.TypeOf((*MediatailorPrefetchScheduleTagsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_MediatailorPrefetchScheduleTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
