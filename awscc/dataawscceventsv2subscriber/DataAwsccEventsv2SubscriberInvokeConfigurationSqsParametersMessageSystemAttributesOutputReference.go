// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscceventsv2subscriber

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscceventsv2subscriber/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference interface {
	cdktn.ComplexObject
	BinaryValue() *string
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
	DataType() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributes
	SetInternalValue(val *DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributes)
	StringValue() *string
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

// The jsii proxy struct for DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference
type jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) BinaryValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) DataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) InternalValue() *DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributes {
	var returns *DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectKey); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEventsv2Subscriber.DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		&j,
	)

	return &j
}

func NewDataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference_Override(d DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEventsv2Subscriber.DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference)SetInternalValue(val *DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

