// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ioteventsdetectormodel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference interface {
	cdktn.ComplexObject
	ClearTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsClearTimerOutputReference
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
	DynamoDb() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDbOutputReference
	DynamoDbInput() interface{}
	DynamoDBv2() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDBv2OutputReference
	DynamoDBv2Input() interface{}
	Firehose() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsFirehoseOutputReference
	FirehoseInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IotEvents() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotEventsOutputReference
	IotEventsInput() interface{}
	IotSiteWise() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotSiteWiseOutputReference
	IotSiteWiseInput() interface{}
	IotTopicPublish() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotTopicPublishOutputReference
	IotTopicPublishInput() interface{}
	Lambda() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsLambdaOutputReference
	LambdaInput() interface{}
	ResetTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsResetTimerOutputReference
	ResetTimerInput() interface{}
	SetTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetTimerOutputReference
	SetTimerInput() interface{}
	SetVariable() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetVariableOutputReference
	SetVariableInput() interface{}
	Sns() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSnsOutputReference
	SnsInput() interface{}
	Sqs() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSqsOutputReference
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
	PutClearTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsClearTimer)
	PutDynamoDb(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDb)
	PutDynamoDBv2(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDBv2)
	PutFirehose(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsFirehose)
	PutIotEvents(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotEvents)
	PutIotSiteWise(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotSiteWise)
	PutIotTopicPublish(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotTopicPublish)
	PutLambda(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsLambda)
	PutResetTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsResetTimer)
	PutSetTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetTimer)
	PutSetVariable(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetVariable)
	PutSns(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSns)
	PutSqs(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSqs)
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

// The jsii proxy struct for IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference
type jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ClearTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsClearTimerOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsClearTimerOutputReference
	_jsii_.Get(
		j,
		"clearTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ClearTimerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clearTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) DynamoDb() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDbOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDbOutputReference
	_jsii_.Get(
		j,
		"dynamoDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) DynamoDbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) DynamoDBv2() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDBv2OutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDBv2OutputReference
	_jsii_.Get(
		j,
		"dynamoDBv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) DynamoDBv2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDBv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) Firehose() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsFirehoseOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) FirehoseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) IotEvents() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotEventsOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotEventsOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) IotEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) IotSiteWise() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotSiteWiseOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotSiteWiseOutputReference
	_jsii_.Get(
		j,
		"iotSiteWise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) IotSiteWiseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotSiteWiseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) IotTopicPublish() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotTopicPublishOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotTopicPublishOutputReference
	_jsii_.Get(
		j,
		"iotTopicPublish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) IotTopicPublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotTopicPublishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) Lambda() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsLambdaOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsResetTimerOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsResetTimerOutputReference
	_jsii_.Get(
		j,
		"resetTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetTimerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) SetTimer() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetTimerOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetTimerOutputReference
	_jsii_.Get(
		j,
		"setTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) SetTimerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) SetVariable() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetVariableOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetVariableOutputReference
	_jsii_.Get(
		j,
		"setVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) SetVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) Sns() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSnsOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) SnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) Sqs() IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSqsOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) SqsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference {
	_init_.Initialize()

	if err := validateNewIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference_Override(i IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) PutClearTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsClearTimer) {
	if err := i.validatePutClearTimerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putClearTimer",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) PutDynamoDb(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDb) {
	if err := i.validatePutDynamoDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamoDb",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) PutDynamoDBv2(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDBv2) {
	if err := i.validatePutDynamoDBv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamoDBv2",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) PutFirehose(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsFirehose) {
	if err := i.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFirehose",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) PutIotEvents(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotEvents) {
	if err := i.validatePutIotEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) PutIotSiteWise(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotSiteWise) {
	if err := i.validatePutIotSiteWiseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotSiteWise",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) PutIotTopicPublish(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotTopicPublish) {
	if err := i.validatePutIotTopicPublishParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotTopicPublish",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) PutLambda(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsLambda) {
	if err := i.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambda",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) PutResetTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsResetTimer) {
	if err := i.validatePutResetTimerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putResetTimer",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) PutSetTimer(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetTimer) {
	if err := i.validatePutSetTimerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSetTimer",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) PutSetVariable(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetVariable) {
	if err := i.validatePutSetVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSetVariable",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) PutSns(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSns) {
	if err := i.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSns",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) PutSqs(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSqs) {
	if err := i.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSqs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetClearTimer() {
	_jsii_.InvokeVoid(
		i,
		"resetClearTimer",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetDynamoDb() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamoDb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetDynamoDBv2() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamoDBv2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		i,
		"resetFirehose",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetIotEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetIotSiteWise() {
	_jsii_.InvokeVoid(
		i,
		"resetIotSiteWise",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetIotTopicPublish() {
	_jsii_.InvokeVoid(
		i,
		"resetIotTopicPublish",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		i,
		"resetLambda",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetResetTimer() {
	_jsii_.InvokeVoid(
		i,
		"resetResetTimer",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetSetTimer() {
	_jsii_.InvokeVoid(
		i,
		"resetSetTimer",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetSetVariable() {
	_jsii_.InvokeVoid(
		i,
		"resetSetVariable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetSns() {
	_jsii_.InvokeVoid(
		i,
		"resetSns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetSqs() {
	_jsii_.InvokeVoid(
		i,
		"resetSqs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

