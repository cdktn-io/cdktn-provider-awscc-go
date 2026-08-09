// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscciottopicrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscciottopicrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference interface {
	cdktn.ComplexObject
	CloudwatchAlarm() DataAwsccIotTopicRuleTopicRulePayloadErrorActionCloudwatchAlarmOutputReference
	CloudwatchLogs() DataAwsccIotTopicRuleTopicRulePayloadErrorActionCloudwatchLogsOutputReference
	CloudwatchMetric() DataAwsccIotTopicRuleTopicRulePayloadErrorActionCloudwatchMetricOutputReference
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DynamoDb() DataAwsccIotTopicRuleTopicRulePayloadErrorActionDynamoDbOutputReference
	DynamoDBv2() DataAwsccIotTopicRuleTopicRulePayloadErrorActionDynamoDBv2OutputReference
	Elasticsearch() DataAwsccIotTopicRuleTopicRulePayloadErrorActionElasticsearchOutputReference
	Firehose() DataAwsccIotTopicRuleTopicRulePayloadErrorActionFirehoseOutputReference
	// Experimental.
	Fqn() *string
	Http() DataAwsccIotTopicRuleTopicRulePayloadErrorActionHttpOutputReference
	InternalValue() *DataAwsccIotTopicRuleTopicRulePayloadErrorAction
	SetInternalValue(val *DataAwsccIotTopicRuleTopicRulePayloadErrorAction)
	IotAnalytics() DataAwsccIotTopicRuleTopicRulePayloadErrorActionIotAnalyticsOutputReference
	IotEvents() DataAwsccIotTopicRuleTopicRulePayloadErrorActionIotEventsOutputReference
	IotSiteWise() DataAwsccIotTopicRuleTopicRulePayloadErrorActionIotSiteWiseOutputReference
	Kafka() DataAwsccIotTopicRuleTopicRulePayloadErrorActionKafkaOutputReference
	Kinesis() DataAwsccIotTopicRuleTopicRulePayloadErrorActionKinesisOutputReference
	Lambda() DataAwsccIotTopicRuleTopicRulePayloadErrorActionLambdaOutputReference
	Location() DataAwsccIotTopicRuleTopicRulePayloadErrorActionLocationOutputReference
	OpenSearch() DataAwsccIotTopicRuleTopicRulePayloadErrorActionOpenSearchOutputReference
	Republish() DataAwsccIotTopicRuleTopicRulePayloadErrorActionRepublishOutputReference
	S3() DataAwsccIotTopicRuleTopicRulePayloadErrorActionS3OutputReference
	Sns() DataAwsccIotTopicRuleTopicRulePayloadErrorActionSnsOutputReference
	Sqs() DataAwsccIotTopicRuleTopicRulePayloadErrorActionSqsOutputReference
	StepFunctions() DataAwsccIotTopicRuleTopicRulePayloadErrorActionStepFunctionsOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Timestream() DataAwsccIotTopicRuleTopicRulePayloadErrorActionTimestreamOutputReference
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference
type jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) CloudwatchAlarm() DataAwsccIotTopicRuleTopicRulePayloadErrorActionCloudwatchAlarmOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionCloudwatchAlarmOutputReference
	_jsii_.Get(
		j,
		"cloudwatchAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) CloudwatchLogs() DataAwsccIotTopicRuleTopicRulePayloadErrorActionCloudwatchLogsOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionCloudwatchLogsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) CloudwatchMetric() DataAwsccIotTopicRuleTopicRulePayloadErrorActionCloudwatchMetricOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionCloudwatchMetricOutputReference
	_jsii_.Get(
		j,
		"cloudwatchMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) DynamoDb() DataAwsccIotTopicRuleTopicRulePayloadErrorActionDynamoDbOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionDynamoDbOutputReference
	_jsii_.Get(
		j,
		"dynamoDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) DynamoDBv2() DataAwsccIotTopicRuleTopicRulePayloadErrorActionDynamoDBv2OutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionDynamoDBv2OutputReference
	_jsii_.Get(
		j,
		"dynamoDBv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) Elasticsearch() DataAwsccIotTopicRuleTopicRulePayloadErrorActionElasticsearchOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionElasticsearchOutputReference
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) Firehose() DataAwsccIotTopicRuleTopicRulePayloadErrorActionFirehoseOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) Http() DataAwsccIotTopicRuleTopicRulePayloadErrorActionHttpOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionHttpOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) InternalValue() *DataAwsccIotTopicRuleTopicRulePayloadErrorAction {
	var returns *DataAwsccIotTopicRuleTopicRulePayloadErrorAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) IotAnalytics() DataAwsccIotTopicRuleTopicRulePayloadErrorActionIotAnalyticsOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionIotAnalyticsOutputReference
	_jsii_.Get(
		j,
		"iotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) IotEvents() DataAwsccIotTopicRuleTopicRulePayloadErrorActionIotEventsOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionIotEventsOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) IotSiteWise() DataAwsccIotTopicRuleTopicRulePayloadErrorActionIotSiteWiseOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionIotSiteWiseOutputReference
	_jsii_.Get(
		j,
		"iotSiteWise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) Kafka() DataAwsccIotTopicRuleTopicRulePayloadErrorActionKafkaOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionKafkaOutputReference
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) Kinesis() DataAwsccIotTopicRuleTopicRulePayloadErrorActionKinesisOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionKinesisOutputReference
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) Lambda() DataAwsccIotTopicRuleTopicRulePayloadErrorActionLambdaOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) Location() DataAwsccIotTopicRuleTopicRulePayloadErrorActionLocationOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionLocationOutputReference
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) OpenSearch() DataAwsccIotTopicRuleTopicRulePayloadErrorActionOpenSearchOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionOpenSearchOutputReference
	_jsii_.Get(
		j,
		"openSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) Republish() DataAwsccIotTopicRuleTopicRulePayloadErrorActionRepublishOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionRepublishOutputReference
	_jsii_.Get(
		j,
		"republish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) S3() DataAwsccIotTopicRuleTopicRulePayloadErrorActionS3OutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) Sns() DataAwsccIotTopicRuleTopicRulePayloadErrorActionSnsOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) Sqs() DataAwsccIotTopicRuleTopicRulePayloadErrorActionSqsOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) StepFunctions() DataAwsccIotTopicRuleTopicRulePayloadErrorActionStepFunctionsOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionStepFunctionsOutputReference
	_jsii_.Get(
		j,
		"stepFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) Timestream() DataAwsccIotTopicRuleTopicRulePayloadErrorActionTimestreamOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadErrorActionTimestreamOutputReference
	_jsii_.Get(
		j,
		"timestream",
		&returns,
	)
	return returns
}


func NewDataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccIotTopicRule.DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference_Override(d DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccIotTopicRule.DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference)SetInternalValue(val *DataAwsccIotTopicRuleTopicRulePayloadErrorAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadErrorActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

