// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsalarmmodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ioteventsalarmmodel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference interface {
	cdktn.ComplexObject
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
	DynamoDb() IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference
	DynamoDbInput() interface{}
	DynamoDBv2() IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDBv2OutputReference
	DynamoDBv2Input() interface{}
	Firehose() IoteventsAlarmModelAlarmEventActionsAlarmActionsFirehoseOutputReference
	FirehoseInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IotEvents() IoteventsAlarmModelAlarmEventActionsAlarmActionsIotEventsOutputReference
	IotEventsInput() interface{}
	IotSiteWise() IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWiseOutputReference
	IotSiteWiseInput() interface{}
	IotTopicPublish() IoteventsAlarmModelAlarmEventActionsAlarmActionsIotTopicPublishOutputReference
	IotTopicPublishInput() interface{}
	Lambda() IoteventsAlarmModelAlarmEventActionsAlarmActionsLambdaOutputReference
	LambdaInput() interface{}
	Sns() IoteventsAlarmModelAlarmEventActionsAlarmActionsSnsOutputReference
	SnsInput() interface{}
	Sqs() IoteventsAlarmModelAlarmEventActionsAlarmActionsSqsOutputReference
	SqsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutDynamoDb(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDb)
	PutDynamoDBv2(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDBv2)
	PutFirehose(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsFirehose)
	PutIotEvents(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotEvents)
	PutIotSiteWise(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWise)
	PutIotTopicPublish(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotTopicPublish)
	PutLambda(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsLambda)
	PutSns(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsSns)
	PutSqs(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsSqs)
	ResetDynamoDb()
	ResetDynamoDBv2()
	ResetFirehose()
	ResetIotEvents()
	ResetIotSiteWise()
	ResetIotTopicPublish()
	ResetLambda()
	ResetSns()
	ResetSqs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference
type jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) DynamoDb() IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference {
	var returns IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference
	_jsii_.Get(
		j,
		"dynamoDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) DynamoDbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) DynamoDBv2() IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDBv2OutputReference {
	var returns IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDBv2OutputReference
	_jsii_.Get(
		j,
		"dynamoDBv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) DynamoDBv2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDBv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) Firehose() IoteventsAlarmModelAlarmEventActionsAlarmActionsFirehoseOutputReference {
	var returns IoteventsAlarmModelAlarmEventActionsAlarmActionsFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) FirehoseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) IotEvents() IoteventsAlarmModelAlarmEventActionsAlarmActionsIotEventsOutputReference {
	var returns IoteventsAlarmModelAlarmEventActionsAlarmActionsIotEventsOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) IotEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) IotSiteWise() IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWiseOutputReference {
	var returns IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWiseOutputReference
	_jsii_.Get(
		j,
		"iotSiteWise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) IotSiteWiseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotSiteWiseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) IotTopicPublish() IoteventsAlarmModelAlarmEventActionsAlarmActionsIotTopicPublishOutputReference {
	var returns IoteventsAlarmModelAlarmEventActionsAlarmActionsIotTopicPublishOutputReference
	_jsii_.Get(
		j,
		"iotTopicPublish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) IotTopicPublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotTopicPublishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) Lambda() IoteventsAlarmModelAlarmEventActionsAlarmActionsLambdaOutputReference {
	var returns IoteventsAlarmModelAlarmEventActionsAlarmActionsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) Sns() IoteventsAlarmModelAlarmEventActionsAlarmActionsSnsOutputReference {
	var returns IoteventsAlarmModelAlarmEventActionsAlarmActionsSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) SnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) Sqs() IoteventsAlarmModelAlarmEventActionsAlarmActionsSqsOutputReference {
	var returns IoteventsAlarmModelAlarmEventActionsAlarmActionsSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) SqsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference {
	_init_.Initialize()

	if err := validateNewIoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ioteventsAlarmModel.IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference_Override(i IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ioteventsAlarmModel.IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) PutDynamoDb(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDb) {
	if err := i.validatePutDynamoDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamoDb",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) PutDynamoDBv2(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDBv2) {
	if err := i.validatePutDynamoDBv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamoDBv2",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) PutFirehose(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsFirehose) {
	if err := i.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFirehose",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) PutIotEvents(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotEvents) {
	if err := i.validatePutIotEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) PutIotSiteWise(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWise) {
	if err := i.validatePutIotSiteWiseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotSiteWise",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) PutIotTopicPublish(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotTopicPublish) {
	if err := i.validatePutIotTopicPublishParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotTopicPublish",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) PutLambda(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsLambda) {
	if err := i.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambda",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) PutSns(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsSns) {
	if err := i.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSns",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) PutSqs(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsSqs) {
	if err := i.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSqs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) ResetDynamoDb() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamoDb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) ResetDynamoDBv2() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamoDBv2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		i,
		"resetFirehose",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) ResetIotEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) ResetIotSiteWise() {
	_jsii_.InvokeVoid(
		i,
		"resetIotSiteWise",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) ResetIotTopicPublish() {
	_jsii_.InvokeVoid(
		i,
		"resetIotTopicPublish",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		i,
		"resetLambda",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) ResetSns() {
	_jsii_.InvokeVoid(
		i,
		"resetSns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) ResetSqs() {
	_jsii_.InvokeVoid(
		i,
		"resetSqs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

