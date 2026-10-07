// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscceventsv2subscriber

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscceventsv2subscriber/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference interface {
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
	EventBusV2Parameters() DataAwsccEventsv2SubscriberInvokeConfigurationEventBusV2ParametersOutputReference
	// Experimental.
	Fqn() *string
	HttpParameters() DataAwsccEventsv2SubscriberInvokeConfigurationHttpParametersOutputReference
	InternalValue() *DataAwsccEventsv2SubscriberInvokeConfiguration
	SetInternalValue(val *DataAwsccEventsv2SubscriberInvokeConfiguration)
	KinesisParameters() DataAwsccEventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference
	LambdaParameters() DataAwsccEventsv2SubscriberInvokeConfigurationLambdaParametersOutputReference
	RoleArn() *string
	SnsParameters() DataAwsccEventsv2SubscriberInvokeConfigurationSnsParametersOutputReference
	SqsParameters() DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersOutputReference
	StepFunctionsParameters() DataAwsccEventsv2SubscriberInvokeConfigurationStepFunctionsParametersOutputReference
	TargetArn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UniversalTargetParameters() DataAwsccEventsv2SubscriberInvokeConfigurationUniversalTargetParametersOutputReference
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

// The jsii proxy struct for DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference
type jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) EventBusV2Parameters() DataAwsccEventsv2SubscriberInvokeConfigurationEventBusV2ParametersOutputReference {
	var returns DataAwsccEventsv2SubscriberInvokeConfigurationEventBusV2ParametersOutputReference
	_jsii_.Get(
		j,
		"eventBusV2Parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) HttpParameters() DataAwsccEventsv2SubscriberInvokeConfigurationHttpParametersOutputReference {
	var returns DataAwsccEventsv2SubscriberInvokeConfigurationHttpParametersOutputReference
	_jsii_.Get(
		j,
		"httpParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) InternalValue() *DataAwsccEventsv2SubscriberInvokeConfiguration {
	var returns *DataAwsccEventsv2SubscriberInvokeConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) KinesisParameters() DataAwsccEventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference {
	var returns DataAwsccEventsv2SubscriberInvokeConfigurationKinesisParametersOutputReference
	_jsii_.Get(
		j,
		"kinesisParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) LambdaParameters() DataAwsccEventsv2SubscriberInvokeConfigurationLambdaParametersOutputReference {
	var returns DataAwsccEventsv2SubscriberInvokeConfigurationLambdaParametersOutputReference
	_jsii_.Get(
		j,
		"lambdaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) SnsParameters() DataAwsccEventsv2SubscriberInvokeConfigurationSnsParametersOutputReference {
	var returns DataAwsccEventsv2SubscriberInvokeConfigurationSnsParametersOutputReference
	_jsii_.Get(
		j,
		"snsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) SqsParameters() DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersOutputReference {
	var returns DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersOutputReference
	_jsii_.Get(
		j,
		"sqsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) StepFunctionsParameters() DataAwsccEventsv2SubscriberInvokeConfigurationStepFunctionsParametersOutputReference {
	var returns DataAwsccEventsv2SubscriberInvokeConfigurationStepFunctionsParametersOutputReference
	_jsii_.Get(
		j,
		"stepFunctionsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) UniversalTargetParameters() DataAwsccEventsv2SubscriberInvokeConfigurationUniversalTargetParametersOutputReference {
	var returns DataAwsccEventsv2SubscriberInvokeConfigurationUniversalTargetParametersOutputReference
	_jsii_.Get(
		j,
		"universalTargetParameters",
		&returns,
	)
	return returns
}


func NewDataAwsccEventsv2SubscriberInvokeConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEventsv2SubscriberInvokeConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEventsv2Subscriber.DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccEventsv2SubscriberInvokeConfigurationOutputReference_Override(d DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEventsv2Subscriber.DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference)SetInternalValue(val *DataAwsccEventsv2SubscriberInvokeConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

