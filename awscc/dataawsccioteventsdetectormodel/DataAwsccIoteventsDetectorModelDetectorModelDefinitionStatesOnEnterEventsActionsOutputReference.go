// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccioteventsdetectormodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccioteventsdetectormodel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference interface {
	cdktn.ComplexObject
	ClearTimer() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsClearTimerOutputReference
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
	DynamoDb() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDbOutputReference
	DynamoDBv2() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDBv2OutputReference
	Firehose() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsFirehoseOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActions
	SetInternalValue(val *DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActions)
	IotEvents() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotEventsOutputReference
	IotSiteWise() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWiseOutputReference
	IotTopicPublish() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishOutputReference
	Lambda() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsLambdaOutputReference
	ResetTimer() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsResetTimerOutputReference
	SetTimer() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetTimerOutputReference
	SetVariable() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetVariableOutputReference
	Sns() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSnsOutputReference
	Sqs() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSqsOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference
type jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ClearTimer() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsClearTimerOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsClearTimerOutputReference
	_jsii_.Get(
		j,
		"clearTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) DynamoDb() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDbOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDbOutputReference
	_jsii_.Get(
		j,
		"dynamoDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) DynamoDBv2() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDBv2OutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsDynamoDBv2OutputReference
	_jsii_.Get(
		j,
		"dynamoDBv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) Firehose() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsFirehoseOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) InternalValue() *DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActions {
	var returns *DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) IotEvents() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotEventsOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotEventsOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) IotSiteWise() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWiseOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWiseOutputReference
	_jsii_.Get(
		j,
		"iotSiteWise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) IotTopicPublish() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishOutputReference
	_jsii_.Get(
		j,
		"iotTopicPublish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) Lambda() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsLambdaOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ResetTimer() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsResetTimerOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsResetTimerOutputReference
	_jsii_.Get(
		j,
		"resetTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) SetTimer() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetTimerOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetTimerOutputReference
	_jsii_.Get(
		j,
		"setTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) SetVariable() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetVariableOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSetVariableOutputReference
	_jsii_.Get(
		j,
		"setVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) Sns() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSnsOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) Sqs() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSqsOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccIoteventsDetectorModel.DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference_Override(d DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccIoteventsDetectorModel.DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference)SetInternalValue(val *DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

