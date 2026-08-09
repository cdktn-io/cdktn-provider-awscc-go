// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iottopicrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotTopicRuleTopicRulePayloadErrorActionOutputReference interface {
	cdktn.ComplexObject
	CloudwatchAlarm() IotTopicRuleTopicRulePayloadErrorActionCloudwatchAlarmOutputReference
	CloudwatchAlarmInput() interface{}
	CloudwatchLogs() IotTopicRuleTopicRulePayloadErrorActionCloudwatchLogsOutputReference
	CloudwatchLogsInput() interface{}
	CloudwatchMetric() IotTopicRuleTopicRulePayloadErrorActionCloudwatchMetricOutputReference
	CloudwatchMetricInput() interface{}
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
	DynamoDb() IotTopicRuleTopicRulePayloadErrorActionDynamoDbOutputReference
	DynamoDbInput() interface{}
	DynamoDBv2() IotTopicRuleTopicRulePayloadErrorActionDynamoDBv2OutputReference
	DynamoDBv2Input() interface{}
	Elasticsearch() IotTopicRuleTopicRulePayloadErrorActionElasticsearchOutputReference
	ElasticsearchInput() interface{}
	Firehose() IotTopicRuleTopicRulePayloadErrorActionFirehoseOutputReference
	FirehoseInput() interface{}
	// Experimental.
	Fqn() *string
	Http() IotTopicRuleTopicRulePayloadErrorActionHttpOutputReference
	HttpInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IotAnalytics() IotTopicRuleTopicRulePayloadErrorActionIotAnalyticsOutputReference
	IotAnalyticsInput() interface{}
	IotEvents() IotTopicRuleTopicRulePayloadErrorActionIotEventsOutputReference
	IotEventsInput() interface{}
	IotSiteWise() IotTopicRuleTopicRulePayloadErrorActionIotSiteWiseOutputReference
	IotSiteWiseInput() interface{}
	Kafka() IotTopicRuleTopicRulePayloadErrorActionKafkaOutputReference
	KafkaInput() interface{}
	Kinesis() IotTopicRuleTopicRulePayloadErrorActionKinesisOutputReference
	KinesisInput() interface{}
	Lambda() IotTopicRuleTopicRulePayloadErrorActionLambdaOutputReference
	LambdaInput() interface{}
	Location() IotTopicRuleTopicRulePayloadErrorActionLocationOutputReference
	LocationInput() interface{}
	OpenSearch() IotTopicRuleTopicRulePayloadErrorActionOpenSearchOutputReference
	OpenSearchInput() interface{}
	Republish() IotTopicRuleTopicRulePayloadErrorActionRepublishOutputReference
	RepublishInput() interface{}
	S3() IotTopicRuleTopicRulePayloadErrorActionS3OutputReference
	S3Input() interface{}
	Sns() IotTopicRuleTopicRulePayloadErrorActionSnsOutputReference
	SnsInput() interface{}
	Sqs() IotTopicRuleTopicRulePayloadErrorActionSqsOutputReference
	SqsInput() interface{}
	StepFunctions() IotTopicRuleTopicRulePayloadErrorActionStepFunctionsOutputReference
	StepFunctionsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Timestream() IotTopicRuleTopicRulePayloadErrorActionTimestreamOutputReference
	TimestreamInput() interface{}
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
	PutCloudwatchAlarm(value *IotTopicRuleTopicRulePayloadErrorActionCloudwatchAlarm)
	PutCloudwatchLogs(value *IotTopicRuleTopicRulePayloadErrorActionCloudwatchLogs)
	PutCloudwatchMetric(value *IotTopicRuleTopicRulePayloadErrorActionCloudwatchMetric)
	PutDynamoDb(value *IotTopicRuleTopicRulePayloadErrorActionDynamoDb)
	PutDynamoDBv2(value *IotTopicRuleTopicRulePayloadErrorActionDynamoDBv2)
	PutElasticsearch(value *IotTopicRuleTopicRulePayloadErrorActionElasticsearch)
	PutFirehose(value *IotTopicRuleTopicRulePayloadErrorActionFirehose)
	PutHttp(value *IotTopicRuleTopicRulePayloadErrorActionHttp)
	PutIotAnalytics(value *IotTopicRuleTopicRulePayloadErrorActionIotAnalytics)
	PutIotEvents(value *IotTopicRuleTopicRulePayloadErrorActionIotEvents)
	PutIotSiteWise(value *IotTopicRuleTopicRulePayloadErrorActionIotSiteWise)
	PutKafka(value *IotTopicRuleTopicRulePayloadErrorActionKafka)
	PutKinesis(value *IotTopicRuleTopicRulePayloadErrorActionKinesis)
	PutLambda(value *IotTopicRuleTopicRulePayloadErrorActionLambda)
	PutLocation(value *IotTopicRuleTopicRulePayloadErrorActionLocation)
	PutOpenSearch(value *IotTopicRuleTopicRulePayloadErrorActionOpenSearch)
	PutRepublish(value *IotTopicRuleTopicRulePayloadErrorActionRepublish)
	PutS3(value *IotTopicRuleTopicRulePayloadErrorActionS3)
	PutSns(value *IotTopicRuleTopicRulePayloadErrorActionSns)
	PutSqs(value *IotTopicRuleTopicRulePayloadErrorActionSqs)
	PutStepFunctions(value *IotTopicRuleTopicRulePayloadErrorActionStepFunctions)
	PutTimestream(value *IotTopicRuleTopicRulePayloadErrorActionTimestream)
	ResetCloudwatchAlarm()
	ResetCloudwatchLogs()
	ResetCloudwatchMetric()
	ResetDynamoDb()
	ResetDynamoDBv2()
	ResetElasticsearch()
	ResetFirehose()
	ResetHttp()
	ResetIotAnalytics()
	ResetIotEvents()
	ResetIotSiteWise()
	ResetKafka()
	ResetKinesis()
	ResetLambda()
	ResetLocation()
	ResetOpenSearch()
	ResetRepublish()
	ResetS3()
	ResetSns()
	ResetSqs()
	ResetStepFunctions()
	ResetTimestream()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotTopicRuleTopicRulePayloadErrorActionOutputReference
type jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) CloudwatchAlarm() IotTopicRuleTopicRulePayloadErrorActionCloudwatchAlarmOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionCloudwatchAlarmOutputReference
	_jsii_.Get(
		j,
		"cloudwatchAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) CloudwatchAlarmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) CloudwatchLogs() IotTopicRuleTopicRulePayloadErrorActionCloudwatchLogsOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionCloudwatchLogsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) CloudwatchLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) CloudwatchMetric() IotTopicRuleTopicRulePayloadErrorActionCloudwatchMetricOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionCloudwatchMetricOutputReference
	_jsii_.Get(
		j,
		"cloudwatchMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) CloudwatchMetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) DynamoDb() IotTopicRuleTopicRulePayloadErrorActionDynamoDbOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionDynamoDbOutputReference
	_jsii_.Get(
		j,
		"dynamoDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) DynamoDbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) DynamoDBv2() IotTopicRuleTopicRulePayloadErrorActionDynamoDBv2OutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionDynamoDBv2OutputReference
	_jsii_.Get(
		j,
		"dynamoDBv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) DynamoDBv2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDBv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) Elasticsearch() IotTopicRuleTopicRulePayloadErrorActionElasticsearchOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionElasticsearchOutputReference
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ElasticsearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) Firehose() IotTopicRuleTopicRulePayloadErrorActionFirehoseOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) FirehoseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) Http() IotTopicRuleTopicRulePayloadErrorActionHttpOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionHttpOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) HttpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) IotAnalytics() IotTopicRuleTopicRulePayloadErrorActionIotAnalyticsOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionIotAnalyticsOutputReference
	_jsii_.Get(
		j,
		"iotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) IotAnalyticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) IotEvents() IotTopicRuleTopicRulePayloadErrorActionIotEventsOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionIotEventsOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) IotEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) IotSiteWise() IotTopicRuleTopicRulePayloadErrorActionIotSiteWiseOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionIotSiteWiseOutputReference
	_jsii_.Get(
		j,
		"iotSiteWise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) IotSiteWiseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotSiteWiseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) Kafka() IotTopicRuleTopicRulePayloadErrorActionKafkaOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionKafkaOutputReference
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) KafkaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) Kinesis() IotTopicRuleTopicRulePayloadErrorActionKinesisOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionKinesisOutputReference
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) KinesisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) Lambda() IotTopicRuleTopicRulePayloadErrorActionLambdaOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) Location() IotTopicRuleTopicRulePayloadErrorActionLocationOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionLocationOutputReference
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) LocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) OpenSearch() IotTopicRuleTopicRulePayloadErrorActionOpenSearchOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionOpenSearchOutputReference
	_jsii_.Get(
		j,
		"openSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) OpenSearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) Republish() IotTopicRuleTopicRulePayloadErrorActionRepublishOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionRepublishOutputReference
	_jsii_.Get(
		j,
		"republish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) RepublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"republishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) S3() IotTopicRuleTopicRulePayloadErrorActionS3OutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) S3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) Sns() IotTopicRuleTopicRulePayloadErrorActionSnsOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) SnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) Sqs() IotTopicRuleTopicRulePayloadErrorActionSqsOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) SqsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) StepFunctions() IotTopicRuleTopicRulePayloadErrorActionStepFunctionsOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionStepFunctionsOutputReference
	_jsii_.Get(
		j,
		"stepFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) StepFunctionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) Timestream() IotTopicRuleTopicRulePayloadErrorActionTimestreamOutputReference {
	var returns IotTopicRuleTopicRulePayloadErrorActionTimestreamOutputReference
	_jsii_.Get(
		j,
		"timestream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) TimestreamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestreamInput",
		&returns,
	)
	return returns
}


func NewIotTopicRuleTopicRulePayloadErrorActionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotTopicRuleTopicRulePayloadErrorActionOutputReference {
	_init_.Initialize()

	if err := validateNewIotTopicRuleTopicRulePayloadErrorActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotTopicRule.IotTopicRuleTopicRulePayloadErrorActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotTopicRuleTopicRulePayloadErrorActionOutputReference_Override(i IotTopicRuleTopicRulePayloadErrorActionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotTopicRule.IotTopicRuleTopicRulePayloadErrorActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutCloudwatchAlarm(value *IotTopicRuleTopicRulePayloadErrorActionCloudwatchAlarm) {
	if err := i.validatePutCloudwatchAlarmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchAlarm",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutCloudwatchLogs(value *IotTopicRuleTopicRulePayloadErrorActionCloudwatchLogs) {
	if err := i.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutCloudwatchMetric(value *IotTopicRuleTopicRulePayloadErrorActionCloudwatchMetric) {
	if err := i.validatePutCloudwatchMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchMetric",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutDynamoDb(value *IotTopicRuleTopicRulePayloadErrorActionDynamoDb) {
	if err := i.validatePutDynamoDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamoDb",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutDynamoDBv2(value *IotTopicRuleTopicRulePayloadErrorActionDynamoDBv2) {
	if err := i.validatePutDynamoDBv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamoDBv2",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutElasticsearch(value *IotTopicRuleTopicRulePayloadErrorActionElasticsearch) {
	if err := i.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutFirehose(value *IotTopicRuleTopicRulePayloadErrorActionFirehose) {
	if err := i.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFirehose",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutHttp(value *IotTopicRuleTopicRulePayloadErrorActionHttp) {
	if err := i.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putHttp",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutIotAnalytics(value *IotTopicRuleTopicRulePayloadErrorActionIotAnalytics) {
	if err := i.validatePutIotAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotAnalytics",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutIotEvents(value *IotTopicRuleTopicRulePayloadErrorActionIotEvents) {
	if err := i.validatePutIotEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutIotSiteWise(value *IotTopicRuleTopicRulePayloadErrorActionIotSiteWise) {
	if err := i.validatePutIotSiteWiseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotSiteWise",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutKafka(value *IotTopicRuleTopicRulePayloadErrorActionKafka) {
	if err := i.validatePutKafkaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putKafka",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutKinesis(value *IotTopicRuleTopicRulePayloadErrorActionKinesis) {
	if err := i.validatePutKinesisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putKinesis",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutLambda(value *IotTopicRuleTopicRulePayloadErrorActionLambda) {
	if err := i.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambda",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutLocation(value *IotTopicRuleTopicRulePayloadErrorActionLocation) {
	if err := i.validatePutLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLocation",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutOpenSearch(value *IotTopicRuleTopicRulePayloadErrorActionOpenSearch) {
	if err := i.validatePutOpenSearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOpenSearch",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutRepublish(value *IotTopicRuleTopicRulePayloadErrorActionRepublish) {
	if err := i.validatePutRepublishParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRepublish",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutS3(value *IotTopicRuleTopicRulePayloadErrorActionS3) {
	if err := i.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putS3",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutSns(value *IotTopicRuleTopicRulePayloadErrorActionSns) {
	if err := i.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSns",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutSqs(value *IotTopicRuleTopicRulePayloadErrorActionSqs) {
	if err := i.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSqs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutStepFunctions(value *IotTopicRuleTopicRulePayloadErrorActionStepFunctions) {
	if err := i.validatePutStepFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putStepFunctions",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) PutTimestream(value *IotTopicRuleTopicRulePayloadErrorActionTimestream) {
	if err := i.validatePutTimestreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimestream",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetCloudwatchAlarm() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchAlarm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetCloudwatchMetric() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchMetric",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetDynamoDb() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamoDb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetDynamoDBv2() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamoDBv2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		i,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		i,
		"resetFirehose",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetHttp() {
	_jsii_.InvokeVoid(
		i,
		"resetHttp",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetIotAnalytics() {
	_jsii_.InvokeVoid(
		i,
		"resetIotAnalytics",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetIotEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetIotSiteWise() {
	_jsii_.InvokeVoid(
		i,
		"resetIotSiteWise",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetKafka() {
	_jsii_.InvokeVoid(
		i,
		"resetKafka",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetKinesis() {
	_jsii_.InvokeVoid(
		i,
		"resetKinesis",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		i,
		"resetLambda",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		i,
		"resetLocation",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetOpenSearch() {
	_jsii_.InvokeVoid(
		i,
		"resetOpenSearch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetRepublish() {
	_jsii_.InvokeVoid(
		i,
		"resetRepublish",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		i,
		"resetS3",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetSns() {
	_jsii_.InvokeVoid(
		i,
		"resetSns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetSqs() {
	_jsii_.InvokeVoid(
		i,
		"resetSqs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetStepFunctions() {
	_jsii_.InvokeVoid(
		i,
		"resetStepFunctions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ResetTimestream() {
	_jsii_.InvokeVoid(
		i,
		"resetTimestream",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadErrorActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

