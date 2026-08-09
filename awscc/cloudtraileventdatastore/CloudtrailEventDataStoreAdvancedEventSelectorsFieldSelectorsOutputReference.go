// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtraileventdatastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cloudtraileventdatastore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference interface {
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
	EndsWith() *[]*string
	SetEndsWith(val *[]*string)
	EndsWithInput() *[]*string
	EqualTo() *[]*string
	SetEqualTo(val *[]*string)
	EqualToInput() *[]*string
	Field() *string
	SetField(val *string)
	FieldInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NotEndsWith() *[]*string
	SetNotEndsWith(val *[]*string)
	NotEndsWithInput() *[]*string
	NotEquals() *[]*string
	SetNotEquals(val *[]*string)
	NotEqualsInput() *[]*string
	NotStartsWith() *[]*string
	SetNotStartsWith(val *[]*string)
	NotStartsWithInput() *[]*string
	StartsWith() *[]*string
	SetStartsWith(val *[]*string)
	StartsWithInput() *[]*string
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
	ResetEndsWith()
	ResetEqualTo()
	ResetField()
	ResetNotEndsWith()
	ResetNotEquals()
	ResetNotStartsWith()
	ResetStartsWith()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference
type jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) EndsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) EndsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) EqualTo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) EqualToInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) NotEndsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEndsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) NotEndsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEndsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) NotEquals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) NotEqualsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) NotStartsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notStartsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) NotStartsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notStartsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) StartsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"startsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) StartsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"startsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference {
	_init_.Initialize()

	if err := validateNewCloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cloudtrailEventDataStore.CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference_Override(c CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cloudtrailEventDataStore.CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference)SetEndsWith(val *[]*string) {
	if err := j.validateSetEndsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endsWith",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference)SetEqualTo(val *[]*string) {
	if err := j.validateSetEqualToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"equalTo",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference)SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference)SetNotEndsWith(val *[]*string) {
	if err := j.validateSetNotEndsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notEndsWith",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference)SetNotEquals(val *[]*string) {
	if err := j.validateSetNotEqualsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notEquals",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference)SetNotStartsWith(val *[]*string) {
	if err := j.validateSetNotStartsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notStartsWith",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference)SetStartsWith(val *[]*string) {
	if err := j.validateSetStartsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startsWith",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) ResetEndsWith() {
	_jsii_.InvokeVoid(
		c,
		"resetEndsWith",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) ResetEqualTo() {
	_jsii_.InvokeVoid(
		c,
		"resetEqualTo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) ResetField() {
	_jsii_.InvokeVoid(
		c,
		"resetField",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) ResetNotEndsWith() {
	_jsii_.InvokeVoid(
		c,
		"resetNotEndsWith",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) ResetNotEquals() {
	_jsii_.InvokeVoid(
		c,
		"resetNotEquals",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) ResetNotStartsWith() {
	_jsii_.InvokeVoid(
		c,
		"resetNotStartsWith",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) ResetStartsWith() {
	_jsii_.InvokeVoid(
		c,
		"resetStartsWith",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

