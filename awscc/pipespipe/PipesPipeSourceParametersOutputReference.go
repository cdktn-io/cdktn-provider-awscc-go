// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/pipespipe/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipesPipeSourceParametersOutputReference interface {
	cdktn.ComplexObject
	ActiveMqBrokerParameters() PipesPipeSourceParametersActiveMqBrokerParametersOutputReference
	ActiveMqBrokerParametersInput() interface{}
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
	DynamoDbStreamParameters() PipesPipeSourceParametersDynamoDbStreamParametersOutputReference
	DynamoDbStreamParametersInput() interface{}
	FilterCriteria() PipesPipeSourceParametersFilterCriteriaOutputReference
	FilterCriteriaInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KinesisStreamParameters() PipesPipeSourceParametersKinesisStreamParametersOutputReference
	KinesisStreamParametersInput() interface{}
	ManagedStreamingKafkaParameters() PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference
	ManagedStreamingKafkaParametersInput() interface{}
	RabbitMqBrokerParameters() PipesPipeSourceParametersRabbitMqBrokerParametersOutputReference
	RabbitMqBrokerParametersInput() interface{}
	SelfManagedKafkaParameters() PipesPipeSourceParametersSelfManagedKafkaParametersOutputReference
	SelfManagedKafkaParametersInput() interface{}
	SqsQueueParameters() PipesPipeSourceParametersSqsQueueParametersOutputReference
	SqsQueueParametersInput() interface{}
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
	PutActiveMqBrokerParameters(value *PipesPipeSourceParametersActiveMqBrokerParameters)
	PutDynamoDbStreamParameters(value *PipesPipeSourceParametersDynamoDbStreamParameters)
	PutFilterCriteria(value *PipesPipeSourceParametersFilterCriteria)
	PutKinesisStreamParameters(value *PipesPipeSourceParametersKinesisStreamParameters)
	PutManagedStreamingKafkaParameters(value *PipesPipeSourceParametersManagedStreamingKafkaParameters)
	PutRabbitMqBrokerParameters(value *PipesPipeSourceParametersRabbitMqBrokerParameters)
	PutSelfManagedKafkaParameters(value *PipesPipeSourceParametersSelfManagedKafkaParameters)
	PutSqsQueueParameters(value *PipesPipeSourceParametersSqsQueueParameters)
	ResetActiveMqBrokerParameters()
	ResetDynamoDbStreamParameters()
	ResetFilterCriteria()
	ResetKinesisStreamParameters()
	ResetManagedStreamingKafkaParameters()
	ResetRabbitMqBrokerParameters()
	ResetSelfManagedKafkaParameters()
	ResetSqsQueueParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipesPipeSourceParametersOutputReference
type jsiiProxy_PipesPipeSourceParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) ActiveMqBrokerParameters() PipesPipeSourceParametersActiveMqBrokerParametersOutputReference {
	var returns PipesPipeSourceParametersActiveMqBrokerParametersOutputReference
	_jsii_.Get(
		j,
		"activeMqBrokerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) ActiveMqBrokerParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeMqBrokerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) DynamoDbStreamParameters() PipesPipeSourceParametersDynamoDbStreamParametersOutputReference {
	var returns PipesPipeSourceParametersDynamoDbStreamParametersOutputReference
	_jsii_.Get(
		j,
		"dynamoDbStreamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) DynamoDbStreamParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbStreamParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) FilterCriteria() PipesPipeSourceParametersFilterCriteriaOutputReference {
	var returns PipesPipeSourceParametersFilterCriteriaOutputReference
	_jsii_.Get(
		j,
		"filterCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) FilterCriteriaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) KinesisStreamParameters() PipesPipeSourceParametersKinesisStreamParametersOutputReference {
	var returns PipesPipeSourceParametersKinesisStreamParametersOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) KinesisStreamParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisStreamParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) ManagedStreamingKafkaParameters() PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference {
	var returns PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference
	_jsii_.Get(
		j,
		"managedStreamingKafkaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) ManagedStreamingKafkaParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedStreamingKafkaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) RabbitMqBrokerParameters() PipesPipeSourceParametersRabbitMqBrokerParametersOutputReference {
	var returns PipesPipeSourceParametersRabbitMqBrokerParametersOutputReference
	_jsii_.Get(
		j,
		"rabbitMqBrokerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) RabbitMqBrokerParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rabbitMqBrokerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) SelfManagedKafkaParameters() PipesPipeSourceParametersSelfManagedKafkaParametersOutputReference {
	var returns PipesPipeSourceParametersSelfManagedKafkaParametersOutputReference
	_jsii_.Get(
		j,
		"selfManagedKafkaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) SelfManagedKafkaParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedKafkaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) SqsQueueParameters() PipesPipeSourceParametersSqsQueueParametersOutputReference {
	var returns PipesPipeSourceParametersSqsQueueParametersOutputReference
	_jsii_.Get(
		j,
		"sqsQueueParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) SqsQueueParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsQueueParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipesPipeSourceParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipesPipeSourceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewPipesPipeSourceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipesPipeSourceParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.pipesPipe.PipesPipeSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipesPipeSourceParametersOutputReference_Override(p PipesPipeSourceParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.pipesPipe.PipesPipeSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) PutActiveMqBrokerParameters(value *PipesPipeSourceParametersActiveMqBrokerParameters) {
	if err := p.validatePutActiveMqBrokerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putActiveMqBrokerParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) PutDynamoDbStreamParameters(value *PipesPipeSourceParametersDynamoDbStreamParameters) {
	if err := p.validatePutDynamoDbStreamParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDynamoDbStreamParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) PutFilterCriteria(value *PipesPipeSourceParametersFilterCriteria) {
	if err := p.validatePutFilterCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFilterCriteria",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) PutKinesisStreamParameters(value *PipesPipeSourceParametersKinesisStreamParameters) {
	if err := p.validatePutKinesisStreamParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putKinesisStreamParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) PutManagedStreamingKafkaParameters(value *PipesPipeSourceParametersManagedStreamingKafkaParameters) {
	if err := p.validatePutManagedStreamingKafkaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putManagedStreamingKafkaParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) PutRabbitMqBrokerParameters(value *PipesPipeSourceParametersRabbitMqBrokerParameters) {
	if err := p.validatePutRabbitMqBrokerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRabbitMqBrokerParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) PutSelfManagedKafkaParameters(value *PipesPipeSourceParametersSelfManagedKafkaParameters) {
	if err := p.validatePutSelfManagedKafkaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSelfManagedKafkaParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) PutSqsQueueParameters(value *PipesPipeSourceParametersSqsQueueParameters) {
	if err := p.validatePutSqsQueueParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSqsQueueParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) ResetActiveMqBrokerParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetActiveMqBrokerParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) ResetDynamoDbStreamParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetDynamoDbStreamParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) ResetFilterCriteria() {
	_jsii_.InvokeVoid(
		p,
		"resetFilterCriteria",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) ResetKinesisStreamParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetKinesisStreamParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) ResetManagedStreamingKafkaParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetManagedStreamingKafkaParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) ResetRabbitMqBrokerParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetRabbitMqBrokerParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) ResetSelfManagedKafkaParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetSelfManagedKafkaParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) ResetSqsQueueParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetSqsQueueParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

