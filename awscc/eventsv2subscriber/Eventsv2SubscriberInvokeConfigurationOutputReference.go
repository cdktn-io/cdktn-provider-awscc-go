// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/eventsv2subscriber/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Eventsv2SubscriberInvokeConfigurationOutputReference interface {
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
	EventBusV2Parameters() Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersOutputReference
	EventBusV2ParametersInput() interface{}
	// Experimental.
	Fqn() *string
	HttpParameters() Eventsv2SubscriberInvokeConfigurationHttpParametersOutputReference
	HttpParametersInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KinesisParameters() Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference
	KinesisParametersInput() interface{}
	LambdaParameters() Eventsv2SubscriberInvokeConfigurationLambdaParametersOutputReference
	LambdaParametersInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	SnsParameters() Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference
	SnsParametersInput() interface{}
	SqsParameters() Eventsv2SubscriberInvokeConfigurationSqsParametersOutputReference
	SqsParametersInput() interface{}
	StepFunctionsParameters() Eventsv2SubscriberInvokeConfigurationStepFunctionsParametersOutputReference
	StepFunctionsParametersInput() interface{}
	TargetArn() *string
	SetTargetArn(val *string)
	TargetArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UniversalTargetParameters() Eventsv2SubscriberInvokeConfigurationUniversalTargetParametersOutputReference
	UniversalTargetParametersInput() interface{}
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
	PutEventBusV2Parameters(value *Eventsv2SubscriberInvokeConfigurationEventBusV2Parameters)
	PutHttpParameters(value *Eventsv2SubscriberInvokeConfigurationHttpParameters)
	PutKinesisParameters(value *Eventsv2SubscriberInvokeConfigurationKinesisParameters)
	PutLambdaParameters(value *Eventsv2SubscriberInvokeConfigurationLambdaParameters)
	PutSnsParameters(value *Eventsv2SubscriberInvokeConfigurationSnsParameters)
	PutSqsParameters(value *Eventsv2SubscriberInvokeConfigurationSqsParameters)
	PutStepFunctionsParameters(value *Eventsv2SubscriberInvokeConfigurationStepFunctionsParameters)
	PutUniversalTargetParameters(value *Eventsv2SubscriberInvokeConfigurationUniversalTargetParameters)
	ResetEventBusV2Parameters()
	ResetHttpParameters()
	ResetKinesisParameters()
	ResetLambdaParameters()
	ResetSnsParameters()
	ResetSqsParameters()
	ResetStepFunctionsParameters()
	ResetUniversalTargetParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Eventsv2SubscriberInvokeConfigurationOutputReference
type jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) EventBusV2Parameters() Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersOutputReference {
	var returns Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersOutputReference
	_jsii_.Get(
		j,
		"eventBusV2Parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) EventBusV2ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventBusV2ParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) HttpParameters() Eventsv2SubscriberInvokeConfigurationHttpParametersOutputReference {
	var returns Eventsv2SubscriberInvokeConfigurationHttpParametersOutputReference
	_jsii_.Get(
		j,
		"httpParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) HttpParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) KinesisParameters() Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference {
	var returns Eventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference
	_jsii_.Get(
		j,
		"kinesisParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) KinesisParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) LambdaParameters() Eventsv2SubscriberInvokeConfigurationLambdaParametersOutputReference {
	var returns Eventsv2SubscriberInvokeConfigurationLambdaParametersOutputReference
	_jsii_.Get(
		j,
		"lambdaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) LambdaParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) SnsParameters() Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference {
	var returns Eventsv2SubscriberInvokeConfigurationSnsParametersOutputReference
	_jsii_.Get(
		j,
		"snsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) SnsParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) SqsParameters() Eventsv2SubscriberInvokeConfigurationSqsParametersOutputReference {
	var returns Eventsv2SubscriberInvokeConfigurationSqsParametersOutputReference
	_jsii_.Get(
		j,
		"sqsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) SqsParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) StepFunctionsParameters() Eventsv2SubscriberInvokeConfigurationStepFunctionsParametersOutputReference {
	var returns Eventsv2SubscriberInvokeConfigurationStepFunctionsParametersOutputReference
	_jsii_.Get(
		j,
		"stepFunctionsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) StepFunctionsParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepFunctionsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) TargetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) UniversalTargetParameters() Eventsv2SubscriberInvokeConfigurationUniversalTargetParametersOutputReference {
	var returns Eventsv2SubscriberInvokeConfigurationUniversalTargetParametersOutputReference
	_jsii_.Get(
		j,
		"universalTargetParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) UniversalTargetParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"universalTargetParametersInput",
		&returns,
	)
	return returns
}


func NewEventsv2SubscriberInvokeConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Eventsv2SubscriberInvokeConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEventsv2SubscriberInvokeConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2SubscriberInvokeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventsv2SubscriberInvokeConfigurationOutputReference_Override(e Eventsv2SubscriberInvokeConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2SubscriberInvokeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference)SetTargetArn(val *string) {
	if err := j.validateSetTargetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetArn",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) PutEventBusV2Parameters(value *Eventsv2SubscriberInvokeConfigurationEventBusV2Parameters) {
	if err := e.validatePutEventBusV2ParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEventBusV2Parameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) PutHttpParameters(value *Eventsv2SubscriberInvokeConfigurationHttpParameters) {
	if err := e.validatePutHttpParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHttpParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) PutKinesisParameters(value *Eventsv2SubscriberInvokeConfigurationKinesisParameters) {
	if err := e.validatePutKinesisParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putKinesisParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) PutLambdaParameters(value *Eventsv2SubscriberInvokeConfigurationLambdaParameters) {
	if err := e.validatePutLambdaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLambdaParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) PutSnsParameters(value *Eventsv2SubscriberInvokeConfigurationSnsParameters) {
	if err := e.validatePutSnsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSnsParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) PutSqsParameters(value *Eventsv2SubscriberInvokeConfigurationSqsParameters) {
	if err := e.validatePutSqsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSqsParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) PutStepFunctionsParameters(value *Eventsv2SubscriberInvokeConfigurationStepFunctionsParameters) {
	if err := e.validatePutStepFunctionsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStepFunctionsParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) PutUniversalTargetParameters(value *Eventsv2SubscriberInvokeConfigurationUniversalTargetParameters) {
	if err := e.validatePutUniversalTargetParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putUniversalTargetParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) ResetEventBusV2Parameters() {
	_jsii_.InvokeVoid(
		e,
		"resetEventBusV2Parameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) ResetHttpParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetHttpParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) ResetKinesisParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetKinesisParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) ResetLambdaParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetLambdaParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) ResetSnsParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetSnsParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) ResetSqsParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetSqsParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) ResetStepFunctionsParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetStepFunctionsParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) ResetUniversalTargetParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetUniversalTargetParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2SubscriberInvokeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

