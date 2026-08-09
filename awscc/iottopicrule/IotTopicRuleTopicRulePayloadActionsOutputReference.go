// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iottopicrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotTopicRuleTopicRulePayloadActionsOutputReference interface {
	cdktn.ComplexObject
	CloudwatchAlarm() IotTopicRuleTopicRulePayloadActionsCloudwatchAlarmOutputReference
	CloudwatchAlarmInput() interface{}
	CloudwatchLogs() IotTopicRuleTopicRulePayloadActionsCloudwatchLogsOutputReference
	CloudwatchLogsInput() interface{}
	CloudwatchMetric() IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference
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
	DynamoDb() IotTopicRuleTopicRulePayloadActionsDynamoDbOutputReference
	DynamoDbInput() interface{}
	DynamoDBv2() IotTopicRuleTopicRulePayloadActionsDynamoDBv2OutputReference
	DynamoDBv2Input() interface{}
	Elasticsearch() IotTopicRuleTopicRulePayloadActionsElasticsearchOutputReference
	ElasticsearchInput() interface{}
	Firehose() IotTopicRuleTopicRulePayloadActionsFirehoseOutputReference
	FirehoseInput() interface{}
	// Experimental.
	Fqn() *string
	Http() IotTopicRuleTopicRulePayloadActionsHttpOutputReference
	HttpInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IotAnalytics() IotTopicRuleTopicRulePayloadActionsIotAnalyticsOutputReference
	IotAnalyticsInput() interface{}
	IotEvents() IotTopicRuleTopicRulePayloadActionsIotEventsOutputReference
	IotEventsInput() interface{}
	IotSiteWise() IotTopicRuleTopicRulePayloadActionsIotSiteWiseOutputReference
	IotSiteWiseInput() interface{}
	Kafka() IotTopicRuleTopicRulePayloadActionsKafkaOutputReference
	KafkaInput() interface{}
	Kinesis() IotTopicRuleTopicRulePayloadActionsKinesisOutputReference
	KinesisInput() interface{}
	Lambda() IotTopicRuleTopicRulePayloadActionsLambdaOutputReference
	LambdaInput() interface{}
	Location() IotTopicRuleTopicRulePayloadActionsLocationOutputReference
	LocationInput() interface{}
	OpenSearch() IotTopicRuleTopicRulePayloadActionsOpenSearchOutputReference
	OpenSearchInput() interface{}
	Republish() IotTopicRuleTopicRulePayloadActionsRepublishOutputReference
	RepublishInput() interface{}
	S3() IotTopicRuleTopicRulePayloadActionsS3OutputReference
	S3Input() interface{}
	Sns() IotTopicRuleTopicRulePayloadActionsSnsOutputReference
	SnsInput() interface{}
	Sqs() IotTopicRuleTopicRulePayloadActionsSqsOutputReference
	SqsInput() interface{}
	StepFunctions() IotTopicRuleTopicRulePayloadActionsStepFunctionsOutputReference
	StepFunctionsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Timestream() IotTopicRuleTopicRulePayloadActionsTimestreamOutputReference
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
	PutCloudwatchAlarm(value *IotTopicRuleTopicRulePayloadActionsCloudwatchAlarm)
	PutCloudwatchLogs(value *IotTopicRuleTopicRulePayloadActionsCloudwatchLogs)
	PutCloudwatchMetric(value *IotTopicRuleTopicRulePayloadActionsCloudwatchMetric)
	PutDynamoDb(value *IotTopicRuleTopicRulePayloadActionsDynamoDb)
	PutDynamoDBv2(value *IotTopicRuleTopicRulePayloadActionsDynamoDBv2)
	PutElasticsearch(value *IotTopicRuleTopicRulePayloadActionsElasticsearch)
	PutFirehose(value *IotTopicRuleTopicRulePayloadActionsFirehose)
	PutHttp(value *IotTopicRuleTopicRulePayloadActionsHttp)
	PutIotAnalytics(value *IotTopicRuleTopicRulePayloadActionsIotAnalytics)
	PutIotEvents(value *IotTopicRuleTopicRulePayloadActionsIotEvents)
	PutIotSiteWise(value *IotTopicRuleTopicRulePayloadActionsIotSiteWise)
	PutKafka(value *IotTopicRuleTopicRulePayloadActionsKafka)
	PutKinesis(value *IotTopicRuleTopicRulePayloadActionsKinesis)
	PutLambda(value *IotTopicRuleTopicRulePayloadActionsLambda)
	PutLocation(value *IotTopicRuleTopicRulePayloadActionsLocation)
	PutOpenSearch(value *IotTopicRuleTopicRulePayloadActionsOpenSearch)
	PutRepublish(value *IotTopicRuleTopicRulePayloadActionsRepublish)
	PutS3(value *IotTopicRuleTopicRulePayloadActionsS3)
	PutSns(value *IotTopicRuleTopicRulePayloadActionsSns)
	PutSqs(value *IotTopicRuleTopicRulePayloadActionsSqs)
	PutStepFunctions(value *IotTopicRuleTopicRulePayloadActionsStepFunctions)
	PutTimestream(value *IotTopicRuleTopicRulePayloadActionsTimestream)
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

// The jsii proxy struct for IotTopicRuleTopicRulePayloadActionsOutputReference
type jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) CloudwatchAlarm() IotTopicRuleTopicRulePayloadActionsCloudwatchAlarmOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsCloudwatchAlarmOutputReference
	_jsii_.Get(
		j,
		"cloudwatchAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) CloudwatchAlarmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) CloudwatchLogs() IotTopicRuleTopicRulePayloadActionsCloudwatchLogsOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsCloudwatchLogsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) CloudwatchLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) CloudwatchMetric() IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsCloudwatchMetricOutputReference
	_jsii_.Get(
		j,
		"cloudwatchMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) CloudwatchMetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) DynamoDb() IotTopicRuleTopicRulePayloadActionsDynamoDbOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsDynamoDbOutputReference
	_jsii_.Get(
		j,
		"dynamoDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) DynamoDbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) DynamoDBv2() IotTopicRuleTopicRulePayloadActionsDynamoDBv2OutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsDynamoDBv2OutputReference
	_jsii_.Get(
		j,
		"dynamoDBv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) DynamoDBv2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDBv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) Elasticsearch() IotTopicRuleTopicRulePayloadActionsElasticsearchOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsElasticsearchOutputReference
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ElasticsearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) Firehose() IotTopicRuleTopicRulePayloadActionsFirehoseOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) FirehoseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) Http() IotTopicRuleTopicRulePayloadActionsHttpOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsHttpOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) HttpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) IotAnalytics() IotTopicRuleTopicRulePayloadActionsIotAnalyticsOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsIotAnalyticsOutputReference
	_jsii_.Get(
		j,
		"iotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) IotAnalyticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) IotEvents() IotTopicRuleTopicRulePayloadActionsIotEventsOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsIotEventsOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) IotEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) IotSiteWise() IotTopicRuleTopicRulePayloadActionsIotSiteWiseOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsIotSiteWiseOutputReference
	_jsii_.Get(
		j,
		"iotSiteWise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) IotSiteWiseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotSiteWiseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) Kafka() IotTopicRuleTopicRulePayloadActionsKafkaOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsKafkaOutputReference
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) KafkaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) Kinesis() IotTopicRuleTopicRulePayloadActionsKinesisOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsKinesisOutputReference
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) KinesisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) Lambda() IotTopicRuleTopicRulePayloadActionsLambdaOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) Location() IotTopicRuleTopicRulePayloadActionsLocationOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsLocationOutputReference
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) LocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) OpenSearch() IotTopicRuleTopicRulePayloadActionsOpenSearchOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsOpenSearchOutputReference
	_jsii_.Get(
		j,
		"openSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) OpenSearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) Republish() IotTopicRuleTopicRulePayloadActionsRepublishOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsRepublishOutputReference
	_jsii_.Get(
		j,
		"republish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) RepublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"republishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) S3() IotTopicRuleTopicRulePayloadActionsS3OutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) S3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) Sns() IotTopicRuleTopicRulePayloadActionsSnsOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) SnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) Sqs() IotTopicRuleTopicRulePayloadActionsSqsOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) SqsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) StepFunctions() IotTopicRuleTopicRulePayloadActionsStepFunctionsOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsStepFunctionsOutputReference
	_jsii_.Get(
		j,
		"stepFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) StepFunctionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) Timestream() IotTopicRuleTopicRulePayloadActionsTimestreamOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsTimestreamOutputReference
	_jsii_.Get(
		j,
		"timestream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) TimestreamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestreamInput",
		&returns,
	)
	return returns
}


func NewIotTopicRuleTopicRulePayloadActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IotTopicRuleTopicRulePayloadActionsOutputReference {
	_init_.Initialize()

	if err := validateNewIotTopicRuleTopicRulePayloadActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotTopicRule.IotTopicRuleTopicRulePayloadActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIotTopicRuleTopicRulePayloadActionsOutputReference_Override(i IotTopicRuleTopicRulePayloadActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotTopicRule.IotTopicRuleTopicRulePayloadActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutCloudwatchAlarm(value *IotTopicRuleTopicRulePayloadActionsCloudwatchAlarm) {
	if err := i.validatePutCloudwatchAlarmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchAlarm",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutCloudwatchLogs(value *IotTopicRuleTopicRulePayloadActionsCloudwatchLogs) {
	if err := i.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutCloudwatchMetric(value *IotTopicRuleTopicRulePayloadActionsCloudwatchMetric) {
	if err := i.validatePutCloudwatchMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchMetric",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutDynamoDb(value *IotTopicRuleTopicRulePayloadActionsDynamoDb) {
	if err := i.validatePutDynamoDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamoDb",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutDynamoDBv2(value *IotTopicRuleTopicRulePayloadActionsDynamoDBv2) {
	if err := i.validatePutDynamoDBv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamoDBv2",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutElasticsearch(value *IotTopicRuleTopicRulePayloadActionsElasticsearch) {
	if err := i.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutFirehose(value *IotTopicRuleTopicRulePayloadActionsFirehose) {
	if err := i.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFirehose",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutHttp(value *IotTopicRuleTopicRulePayloadActionsHttp) {
	if err := i.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putHttp",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutIotAnalytics(value *IotTopicRuleTopicRulePayloadActionsIotAnalytics) {
	if err := i.validatePutIotAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotAnalytics",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutIotEvents(value *IotTopicRuleTopicRulePayloadActionsIotEvents) {
	if err := i.validatePutIotEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutIotSiteWise(value *IotTopicRuleTopicRulePayloadActionsIotSiteWise) {
	if err := i.validatePutIotSiteWiseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotSiteWise",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutKafka(value *IotTopicRuleTopicRulePayloadActionsKafka) {
	if err := i.validatePutKafkaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putKafka",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutKinesis(value *IotTopicRuleTopicRulePayloadActionsKinesis) {
	if err := i.validatePutKinesisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putKinesis",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutLambda(value *IotTopicRuleTopicRulePayloadActionsLambda) {
	if err := i.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambda",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutLocation(value *IotTopicRuleTopicRulePayloadActionsLocation) {
	if err := i.validatePutLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLocation",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutOpenSearch(value *IotTopicRuleTopicRulePayloadActionsOpenSearch) {
	if err := i.validatePutOpenSearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOpenSearch",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutRepublish(value *IotTopicRuleTopicRulePayloadActionsRepublish) {
	if err := i.validatePutRepublishParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRepublish",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutS3(value *IotTopicRuleTopicRulePayloadActionsS3) {
	if err := i.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putS3",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutSns(value *IotTopicRuleTopicRulePayloadActionsSns) {
	if err := i.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSns",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutSqs(value *IotTopicRuleTopicRulePayloadActionsSqs) {
	if err := i.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSqs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutStepFunctions(value *IotTopicRuleTopicRulePayloadActionsStepFunctions) {
	if err := i.validatePutStepFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putStepFunctions",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) PutTimestream(value *IotTopicRuleTopicRulePayloadActionsTimestream) {
	if err := i.validatePutTimestreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimestream",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetCloudwatchAlarm() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchAlarm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetCloudwatchMetric() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchMetric",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetDynamoDb() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamoDb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetDynamoDBv2() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamoDBv2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		i,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		i,
		"resetFirehose",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetHttp() {
	_jsii_.InvokeVoid(
		i,
		"resetHttp",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetIotAnalytics() {
	_jsii_.InvokeVoid(
		i,
		"resetIotAnalytics",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetIotEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetIotSiteWise() {
	_jsii_.InvokeVoid(
		i,
		"resetIotSiteWise",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetKafka() {
	_jsii_.InvokeVoid(
		i,
		"resetKafka",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetKinesis() {
	_jsii_.InvokeVoid(
		i,
		"resetKinesis",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		i,
		"resetLambda",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		i,
		"resetLocation",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetOpenSearch() {
	_jsii_.InvokeVoid(
		i,
		"resetOpenSearch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetRepublish() {
	_jsii_.InvokeVoid(
		i,
		"resetRepublish",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		i,
		"resetS3",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetSns() {
	_jsii_.InvokeVoid(
		i,
		"resetSns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetSqs() {
	_jsii_.InvokeVoid(
		i,
		"resetSqs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetStepFunctions() {
	_jsii_.InvokeVoid(
		i,
		"resetStepFunctions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ResetTimestream() {
	_jsii_.InvokeVoid(
		i,
		"resetTimestream",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

