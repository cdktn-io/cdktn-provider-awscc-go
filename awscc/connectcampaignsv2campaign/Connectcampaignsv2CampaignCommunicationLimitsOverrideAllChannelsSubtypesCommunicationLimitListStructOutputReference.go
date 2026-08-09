// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectcampaignsv2campaign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference interface {
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
	// Experimental.
	Fqn() *string
	Frequency() *float64
	SetFrequency(val *float64)
	FrequencyInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxCountPerRecipient() *float64
	SetMaxCountPerRecipient(val *float64)
	MaxCountPerRecipientInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
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
	ResetFrequency()
	ResetMaxCountPerRecipient()
	ResetUnit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference
type jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) Frequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) FrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) MaxCountPerRecipient() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCountPerRecipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) MaxCountPerRecipientInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCountPerRecipientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}


func NewConnectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference {
	_init_.Initialize()

	if err := validateNewConnectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewConnectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference_Override(c Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference)SetFrequency(val *float64) {
	if err := j.validateSetFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference)SetMaxCountPerRecipient(val *float64) {
	if err := j.validateSetMaxCountPerRecipientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCountPerRecipient",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) ResetFrequency() {
	_jsii_.InvokeVoid(
		c,
		"resetFrequency",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) ResetMaxCountPerRecipient() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxCountPerRecipient",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) ResetUnit() {
	_jsii_.InvokeVoid(
		c,
		"resetUnit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

