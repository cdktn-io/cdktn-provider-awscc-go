// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ioteventsdetectormodel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference interface {
	cdktn.ComplexObject
	ClearTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsClearTimerOutputReference
	ClearTimerInput() interface{}
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
	DynamoDb() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDbOutputReference
	DynamoDbInput() interface{}
	DynamoDBv2() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDBv2OutputReference
	DynamoDBv2Input() interface{}
	Firehose() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsFirehoseOutputReference
	FirehoseInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IotEvents() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotEventsOutputReference
	IotEventsInput() interface{}
	IotSiteWise() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWiseOutputReference
	IotSiteWiseInput() interface{}
	IotTopicPublish() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishOutputReference
	IotTopicPublishInput() interface{}
	Lambda() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsLambdaOutputReference
	LambdaInput() interface{}
	ResetTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsResetTimerOutputReference
	ResetTimerInput() interface{}
	SetTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetTimerOutputReference
	SetTimerInput() interface{}
	SetVariable() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetVariableOutputReference
	SetVariableInput() interface{}
	Sns() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSnsOutputReference
	SnsInput() interface{}
	Sqs() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSqsOutputReference
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
	PutClearTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsClearTimer)
	PutDynamoDb(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDb)
	PutDynamoDBv2(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDBv2)
	PutFirehose(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsFirehose)
	PutIotEvents(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotEvents)
	PutIotSiteWise(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWise)
	PutIotTopicPublish(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublish)
	PutLambda(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsLambda)
	PutResetTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsResetTimer)
	PutSetTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetTimer)
	PutSetVariable(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetVariable)
	PutSns(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSns)
	PutSqs(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSqs)
	ResetClearTimer()
	ResetDynamoDb()
	ResetDynamoDBv2()
	ResetFirehose()
	ResetIotEvents()
	ResetIotSiteWise()
	ResetIotTopicPublish()
	ResetLambda()
	ResetResetTimer()
	ResetSetTimer()
	ResetSetVariable()
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

// The jsii proxy struct for IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference
type jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ClearTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsClearTimerOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsClearTimerOutputReference
	_jsii_.Get(
		j,
		"clearTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ClearTimerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clearTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) DynamoDb() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDbOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDbOutputReference
	_jsii_.Get(
		j,
		"dynamoDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) DynamoDbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) DynamoDBv2() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDBv2OutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDBv2OutputReference
	_jsii_.Get(
		j,
		"dynamoDBv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) DynamoDBv2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDBv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) Firehose() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsFirehoseOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) FirehoseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) IotEvents() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotEventsOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotEventsOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) IotEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) IotSiteWise() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWiseOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWiseOutputReference
	_jsii_.Get(
		j,
		"iotSiteWise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) IotSiteWiseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotSiteWiseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) IotTopicPublish() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishOutputReference
	_jsii_.Get(
		j,
		"iotTopicPublish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) IotTopicPublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotTopicPublishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) Lambda() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsLambdaOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsResetTimerOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsResetTimerOutputReference
	_jsii_.Get(
		j,
		"resetTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetTimerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) SetTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetTimerOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetTimerOutputReference
	_jsii_.Get(
		j,
		"setTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) SetTimerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) SetVariable() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetVariableOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetVariableOutputReference
	_jsii_.Get(
		j,
		"setVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) SetVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) Sns() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSnsOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) SnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) Sqs() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSqsOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) SqsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference {
	_init_.Initialize()

	if err := validateNewIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference_Override(i IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) PutClearTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsClearTimer) {
	if err := i.validatePutClearTimerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putClearTimer",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) PutDynamoDb(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDb) {
	if err := i.validatePutDynamoDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamoDb",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) PutDynamoDBv2(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDBv2) {
	if err := i.validatePutDynamoDBv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamoDBv2",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) PutFirehose(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsFirehose) {
	if err := i.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFirehose",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) PutIotEvents(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotEvents) {
	if err := i.validatePutIotEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) PutIotSiteWise(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWise) {
	if err := i.validatePutIotSiteWiseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotSiteWise",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) PutIotTopicPublish(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublish) {
	if err := i.validatePutIotTopicPublishParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotTopicPublish",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) PutLambda(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsLambda) {
	if err := i.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambda",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) PutResetTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsResetTimer) {
	if err := i.validatePutResetTimerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putResetTimer",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) PutSetTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetTimer) {
	if err := i.validatePutSetTimerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSetTimer",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) PutSetVariable(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetVariable) {
	if err := i.validatePutSetVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSetVariable",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) PutSns(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSns) {
	if err := i.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSns",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) PutSqs(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSqs) {
	if err := i.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSqs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetClearTimer() {
	_jsii_.InvokeVoid(
		i,
		"resetClearTimer",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetDynamoDb() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamoDb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetDynamoDBv2() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamoDBv2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		i,
		"resetFirehose",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetIotEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetIotSiteWise() {
	_jsii_.InvokeVoid(
		i,
		"resetIotSiteWise",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetIotTopicPublish() {
	_jsii_.InvokeVoid(
		i,
		"resetIotTopicPublish",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		i,
		"resetLambda",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetResetTimer() {
	_jsii_.InvokeVoid(
		i,
		"resetResetTimer",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetSetTimer() {
	_jsii_.InvokeVoid(
		i,
		"resetSetTimer",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetSetVariable() {
	_jsii_.InvokeVoid(
		i,
		"resetSetVariable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetSns() {
	_jsii_.InvokeVoid(
		i,
		"resetSns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetSqs() {
	_jsii_.InvokeVoid(
		i,
		"resetSqs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

