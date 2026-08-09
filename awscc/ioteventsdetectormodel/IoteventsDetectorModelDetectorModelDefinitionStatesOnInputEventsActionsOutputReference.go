// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ioteventsdetectormodel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference interface {
	cdktn.ComplexObject
	ClearTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsClearTimerOutputReference
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
	DynamoDb() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDbOutputReference
	DynamoDbInput() interface{}
	DynamoDBv2() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDBv2OutputReference
	DynamoDBv2Input() interface{}
	Firehose() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsFirehoseOutputReference
	FirehoseInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IotEvents() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotEventsOutputReference
	IotEventsInput() interface{}
	IotSiteWise() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotSiteWiseOutputReference
	IotSiteWiseInput() interface{}
	IotTopicPublish() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotTopicPublishOutputReference
	IotTopicPublishInput() interface{}
	Lambda() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsLambdaOutputReference
	LambdaInput() interface{}
	ResetTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsResetTimerOutputReference
	ResetTimerInput() interface{}
	SetTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSetTimerOutputReference
	SetTimerInput() interface{}
	SetVariable() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSetVariableOutputReference
	SetVariableInput() interface{}
	Sns() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSnsOutputReference
	SnsInput() interface{}
	Sqs() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSqsOutputReference
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
	PutClearTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsClearTimer)
	PutDynamoDb(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDb)
	PutDynamoDBv2(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDBv2)
	PutFirehose(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsFirehose)
	PutIotEvents(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotEvents)
	PutIotSiteWise(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotSiteWise)
	PutIotTopicPublish(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotTopicPublish)
	PutLambda(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsLambda)
	PutResetTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsResetTimer)
	PutSetTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSetTimer)
	PutSetVariable(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSetVariable)
	PutSns(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSns)
	PutSqs(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSqs)
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

// The jsii proxy struct for IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference
type jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ClearTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsClearTimerOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsClearTimerOutputReference
	_jsii_.Get(
		j,
		"clearTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ClearTimerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clearTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) DynamoDb() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDbOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDbOutputReference
	_jsii_.Get(
		j,
		"dynamoDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) DynamoDbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) DynamoDBv2() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDBv2OutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDBv2OutputReference
	_jsii_.Get(
		j,
		"dynamoDBv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) DynamoDBv2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDBv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) Firehose() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsFirehoseOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) FirehoseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) IotEvents() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotEventsOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotEventsOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) IotEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) IotSiteWise() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotSiteWiseOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotSiteWiseOutputReference
	_jsii_.Get(
		j,
		"iotSiteWise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) IotSiteWiseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotSiteWiseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) IotTopicPublish() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotTopicPublishOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotTopicPublishOutputReference
	_jsii_.Get(
		j,
		"iotTopicPublish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) IotTopicPublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotTopicPublishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) Lambda() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsLambdaOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsResetTimerOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsResetTimerOutputReference
	_jsii_.Get(
		j,
		"resetTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetTimerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) SetTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSetTimerOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSetTimerOutputReference
	_jsii_.Get(
		j,
		"setTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) SetTimerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) SetVariable() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSetVariableOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSetVariableOutputReference
	_jsii_.Get(
		j,
		"setVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) SetVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) Sns() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSnsOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) SnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) Sqs() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSqsOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) SqsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference {
	_init_.Initialize()

	if err := validateNewIoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference_Override(i IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) PutClearTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsClearTimer) {
	if err := i.validatePutClearTimerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putClearTimer",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) PutDynamoDb(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDb) {
	if err := i.validatePutDynamoDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamoDb",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) PutDynamoDBv2(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDBv2) {
	if err := i.validatePutDynamoDBv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamoDBv2",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) PutFirehose(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsFirehose) {
	if err := i.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFirehose",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) PutIotEvents(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotEvents) {
	if err := i.validatePutIotEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) PutIotSiteWise(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotSiteWise) {
	if err := i.validatePutIotSiteWiseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotSiteWise",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) PutIotTopicPublish(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotTopicPublish) {
	if err := i.validatePutIotTopicPublishParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotTopicPublish",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) PutLambda(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsLambda) {
	if err := i.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambda",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) PutResetTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsResetTimer) {
	if err := i.validatePutResetTimerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putResetTimer",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) PutSetTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSetTimer) {
	if err := i.validatePutSetTimerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSetTimer",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) PutSetVariable(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSetVariable) {
	if err := i.validatePutSetVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSetVariable",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) PutSns(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSns) {
	if err := i.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSns",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) PutSqs(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSqs) {
	if err := i.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSqs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetClearTimer() {
	_jsii_.InvokeVoid(
		i,
		"resetClearTimer",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetDynamoDb() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamoDb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetDynamoDBv2() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamoDBv2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		i,
		"resetFirehose",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetIotEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetIotSiteWise() {
	_jsii_.InvokeVoid(
		i,
		"resetIotSiteWise",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetIotTopicPublish() {
	_jsii_.InvokeVoid(
		i,
		"resetIotTopicPublish",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		i,
		"resetLambda",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetResetTimer() {
	_jsii_.InvokeVoid(
		i,
		"resetResetTimer",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetSetTimer() {
	_jsii_.InvokeVoid(
		i,
		"resetSetTimer",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetSetVariable() {
	_jsii_.InvokeVoid(
		i,
		"resetSetVariable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetSns() {
	_jsii_.InvokeVoid(
		i,
		"resetSns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ResetSqs() {
	_jsii_.InvokeVoid(
		i,
		"resetSqs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

