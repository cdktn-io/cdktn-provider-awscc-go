// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectcampaignsv2campaign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference interface {
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
	EndDate() *string
	SetEndDate(val *string)
	EndDateInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	StartDate() *string
	SetStartDate(val *string)
	StartDateInput() *string
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
	ResetEndDate()
	ResetName()
	ResetStartDate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference
type jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) EndDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) EndDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) StartDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) StartDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference {
	_init_.Initialize()

	if err := validateNewConnectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewConnectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference_Override(c Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference)SetEndDate(val *string) {
	if err := j.validateSetEndDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endDate",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference)SetStartDate(val *string) {
	if err := j.validateSetStartDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startDate",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) ResetEndDate() {
	_jsii_.InvokeVoid(
		c,
		"resetEndDate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) ResetStartDate() {
	_jsii_.InvokeVoid(
		c,
		"resetStartDate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigEmailRestrictedPeriodsRestrictedPeriodListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

