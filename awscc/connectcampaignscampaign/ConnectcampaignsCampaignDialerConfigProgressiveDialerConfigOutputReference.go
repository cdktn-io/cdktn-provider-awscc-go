// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignscampaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectcampaignscampaign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference interface {
	cdktn.ComplexObject
	BandwidthAllocation() *float64
	SetBandwidthAllocation(val *float64)
	BandwidthAllocationInput() *float64
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
	DialingCapacity() *float64
	SetDialingCapacity(val *float64)
	DialingCapacityInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetBandwidthAllocation()
	ResetDialingCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference
type jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) BandwidthAllocation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthAllocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) BandwidthAllocationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthAllocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) DialingCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dialingCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) DialingCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dialingCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsCampaign.ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference_Override(c ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsCampaign.ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference)SetBandwidthAllocation(val *float64) {
	if err := j.validateSetBandwidthAllocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthAllocation",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference)SetDialingCapacity(val *float64) {
	if err := j.validateSetDialingCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dialingCapacity",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) ResetBandwidthAllocation() {
	_jsii_.InvokeVoid(
		c,
		"resetBandwidthAllocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) ResetDialingCapacity() {
	_jsii_.InvokeVoid(
		c,
		"resetDialingCapacity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectcampaignsCampaignDialerConfigProgressiveDialerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

