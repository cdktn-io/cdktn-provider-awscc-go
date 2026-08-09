// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sqsqueue

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sqsQueue.SqsQueue",
		reflect.TypeOf((*SqsQueue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "contentBasedDeduplication", GoGetter: "ContentBasedDeduplication"},
			_jsii_.MemberProperty{JsiiProperty: "contentBasedDeduplicationInput", GoGetter: "ContentBasedDeduplicationInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "deduplicationScope", GoGetter: "DeduplicationScope"},
			_jsii_.MemberProperty{JsiiProperty: "deduplicationScopeInput", GoGetter: "DeduplicationScopeInput"},
			_jsii_.MemberProperty{JsiiProperty: "delaySeconds", GoGetter: "DelaySeconds"},
			_jsii_.MemberProperty{JsiiProperty: "delaySecondsInput", GoGetter: "DelaySecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fifoQueue", GoGetter: "FifoQueue"},
			_jsii_.MemberProperty{JsiiProperty: "fifoQueueInput", GoGetter: "FifoQueueInput"},
			_jsii_.MemberProperty{JsiiProperty: "fifoThroughputLimit", GoGetter: "FifoThroughputLimit"},
			_jsii_.MemberProperty{JsiiProperty: "fifoThroughputLimitInput", GoGetter: "FifoThroughputLimitInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "kmsDataKeyReusePeriodSeconds", GoGetter: "KmsDataKeyReusePeriodSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "kmsDataKeyReusePeriodSecondsInput", GoGetter: "KmsDataKeyReusePeriodSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "kmsMasterKeyId", GoGetter: "KmsMasterKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsMasterKeyIdInput", GoGetter: "KmsMasterKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maximumMessageSize", GoGetter: "MaximumMessageSize"},
			_jsii_.MemberProperty{JsiiProperty: "maximumMessageSizeInput", GoGetter: "MaximumMessageSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "messageRetentionPeriod", GoGetter: "MessageRetentionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "messageRetentionPeriodInput", GoGetter: "MessageRetentionPeriodInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putTags", GoMethod: "PutTags"},
			_jsii_.MemberProperty{JsiiProperty: "queueName", GoGetter: "QueueName"},
			_jsii_.MemberProperty{JsiiProperty: "queueNameInput", GoGetter: "QueueNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "queueUrl", GoGetter: "QueueUrl"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "receiveMessageWaitTimeSeconds", GoGetter: "ReceiveMessageWaitTimeSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "receiveMessageWaitTimeSecondsInput", GoGetter: "ReceiveMessageWaitTimeSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "redriveAllowPolicy", GoGetter: "RedriveAllowPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "redriveAllowPolicyInput", GoGetter: "RedriveAllowPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "redrivePolicy", GoGetter: "RedrivePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "redrivePolicyInput", GoGetter: "RedrivePolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentBasedDeduplication", GoMethod: "ResetContentBasedDeduplication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeduplicationScope", GoMethod: "ResetDeduplicationScope"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelaySeconds", GoMethod: "ResetDelaySeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetFifoQueue", GoMethod: "ResetFifoQueue"},
			_jsii_.MemberMethod{JsiiMethod: "resetFifoThroughputLimit", GoMethod: "ResetFifoThroughputLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsDataKeyReusePeriodSeconds", GoMethod: "ResetKmsDataKeyReusePeriodSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsMasterKeyId", GoMethod: "ResetKmsMasterKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaximumMessageSize", GoMethod: "ResetMaximumMessageSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessageRetentionPeriod", GoMethod: "ResetMessageRetentionPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueueName", GoMethod: "ResetQueueName"},
			_jsii_.MemberMethod{JsiiMethod: "resetReceiveMessageWaitTimeSeconds", GoMethod: "ResetReceiveMessageWaitTimeSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedriveAllowPolicy", GoMethod: "ResetRedriveAllowPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedrivePolicy", GoMethod: "ResetRedrivePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetSqsManagedSseEnabled", GoMethod: "ResetSqsManagedSseEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetVisibilityTimeout", GoMethod: "ResetVisibilityTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "sqsManagedSseEnabled", GoGetter: "SqsManagedSseEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "sqsManagedSseEnabledInput", GoGetter: "SqsManagedSseEnabledInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "visibilityTimeout", GoGetter: "VisibilityTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityTimeoutInput", GoGetter: "VisibilityTimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_SqsQueue{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sqsQueue.SqsQueueConfig",
		reflect.TypeOf((*SqsQueueConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.sqsQueue.SqsQueueTags",
		reflect.TypeOf((*SqsQueueTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sqsQueue.SqsQueueTagsList",
		reflect.TypeOf((*SqsQueueTagsList)(nil)).Elem(),
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
			j := jsiiProxy_SqsQueueTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.sqsQueue.SqsQueueTagsOutputReference",
		reflect.TypeOf((*SqsQueueTagsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_SqsQueueTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
