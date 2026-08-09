// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscciottopicrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscciottopicrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference interface {
	cdktn.ComplexObject
	CloudwatchAlarm() DataAwsccIotTopicRuleTopicRulePayloadActionsCloudwatchAlarmOutputReference
	CloudwatchLogs() DataAwsccIotTopicRuleTopicRulePayloadActionsCloudwatchLogsOutputReference
	CloudwatchMetric() DataAwsccIotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference
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
	DynamoDb() DataAwsccIotTopicRuleTopicRulePayloadActionsDynamoDbOutputReference
	DynamoDBv2() DataAwsccIotTopicRuleTopicRulePayloadActionsDynamoDBv2OutputReference
	Elasticsearch() DataAwsccIotTopicRuleTopicRulePayloadActionsElasticsearchOutputReference
	Firehose() DataAwsccIotTopicRuleTopicRulePayloadActionsFirehoseOutputReference
	// Experimental.
	Fqn() *string
	Http() DataAwsccIotTopicRuleTopicRulePayloadActionsHttpOutputReference
	InternalValue() *DataAwsccIotTopicRuleTopicRulePayloadActions
	SetInternalValue(val *DataAwsccIotTopicRuleTopicRulePayloadActions)
	IotAnalytics() DataAwsccIotTopicRuleTopicRulePayloadActionsIotAnalyticsOutputReference
	IotEvents() DataAwsccIotTopicRuleTopicRulePayloadActionsIotEventsOutputReference
	IotSiteWise() DataAwsccIotTopicRuleTopicRulePayloadActionsIotSiteWiseOutputReference
	Kafka() DataAwsccIotTopicRuleTopicRulePayloadActionsKafkaOutputReference
	Kinesis() DataAwsccIotTopicRuleTopicRulePayloadActionsKinesisOutputReference
	Lambda() DataAwsccIotTopicRuleTopicRulePayloadActionsLambdaOutputReference
	Location() DataAwsccIotTopicRuleTopicRulePayloadActionsLocationOutputReference
	OpenSearch() DataAwsccIotTopicRuleTopicRulePayloadActionsOpenSearchOutputReference
	Republish() DataAwsccIotTopicRuleTopicRulePayloadActionsRepublishOutputReference
	S3() DataAwsccIotTopicRuleTopicRulePayloadActionsS3OutputReference
	Sns() DataAwsccIotTopicRuleTopicRulePayloadActionsSnsOutputReference
	Sqs() DataAwsccIotTopicRuleTopicRulePayloadActionsSqsOutputReference
	StepFunctions() DataAwsccIotTopicRuleTopicRulePayloadActionsStepFunctionsOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Timestream() DataAwsccIotTopicRuleTopicRulePayloadActionsTimestreamOutputReference
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

// The jsii proxy struct for DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference
type jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) CloudwatchAlarm() DataAwsccIotTopicRuleTopicRulePayloadActionsCloudwatchAlarmOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsCloudwatchAlarmOutputReference
	_jsii_.Get(
		j,
		"cloudwatchAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) CloudwatchLogs() DataAwsccIotTopicRuleTopicRulePayloadActionsCloudwatchLogsOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsCloudwatchLogsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) CloudwatchMetric() DataAwsccIotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference
	_jsii_.Get(
		j,
		"cloudwatchMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) DynamoDb() DataAwsccIotTopicRuleTopicRulePayloadActionsDynamoDbOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsDynamoDbOutputReference
	_jsii_.Get(
		j,
		"dynamoDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) DynamoDBv2() DataAwsccIotTopicRuleTopicRulePayloadActionsDynamoDBv2OutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsDynamoDBv2OutputReference
	_jsii_.Get(
		j,
		"dynamoDBv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) Elasticsearch() DataAwsccIotTopicRuleTopicRulePayloadActionsElasticsearchOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsElasticsearchOutputReference
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) Firehose() DataAwsccIotTopicRuleTopicRulePayloadActionsFirehoseOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) Http() DataAwsccIotTopicRuleTopicRulePayloadActionsHttpOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsHttpOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) InternalValue() *DataAwsccIotTopicRuleTopicRulePayloadActions {
	var returns *DataAwsccIotTopicRuleTopicRulePayloadActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) IotAnalytics() DataAwsccIotTopicRuleTopicRulePayloadActionsIotAnalyticsOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsIotAnalyticsOutputReference
	_jsii_.Get(
		j,
		"iotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) IotEvents() DataAwsccIotTopicRuleTopicRulePayloadActionsIotEventsOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsIotEventsOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) IotSiteWise() DataAwsccIotTopicRuleTopicRulePayloadActionsIotSiteWiseOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsIotSiteWiseOutputReference
	_jsii_.Get(
		j,
		"iotSiteWise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) Kafka() DataAwsccIotTopicRuleTopicRulePayloadActionsKafkaOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsKafkaOutputReference
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) Kinesis() DataAwsccIotTopicRuleTopicRulePayloadActionsKinesisOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsKinesisOutputReference
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) Lambda() DataAwsccIotTopicRuleTopicRulePayloadActionsLambdaOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) Location() DataAwsccIotTopicRuleTopicRulePayloadActionsLocationOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsLocationOutputReference
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) OpenSearch() DataAwsccIotTopicRuleTopicRulePayloadActionsOpenSearchOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsOpenSearchOutputReference
	_jsii_.Get(
		j,
		"openSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) Republish() DataAwsccIotTopicRuleTopicRulePayloadActionsRepublishOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsRepublishOutputReference
	_jsii_.Get(
		j,
		"republish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) S3() DataAwsccIotTopicRuleTopicRulePayloadActionsS3OutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) Sns() DataAwsccIotTopicRuleTopicRulePayloadActionsSnsOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) Sqs() DataAwsccIotTopicRuleTopicRulePayloadActionsSqsOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) StepFunctions() DataAwsccIotTopicRuleTopicRulePayloadActionsStepFunctionsOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsStepFunctionsOutputReference
	_jsii_.Get(
		j,
		"stepFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) Timestream() DataAwsccIotTopicRuleTopicRulePayloadActionsTimestreamOutputReference {
	var returns DataAwsccIotTopicRuleTopicRulePayloadActionsTimestreamOutputReference
	_jsii_.Get(
		j,
		"timestream",
		&returns,
	)
	return returns
}


func NewDataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccIotTopicRuleTopicRulePayloadActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccIotTopicRule.DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference_Override(d DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccIotTopicRule.DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference)SetInternalValue(val *DataAwsccIotTopicRuleTopicRulePayloadActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccIotTopicRuleTopicRulePayloadActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

