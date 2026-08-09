// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeoptimizerautomationrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/computeoptimizerautomationrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComputeoptimizerAutomationRuleCriteriaOutputReference interface {
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
	EbsVolumeSizeInGib() ComputeoptimizerAutomationRuleCriteriaEbsVolumeSizeInGibList
	EbsVolumeSizeInGibInput() interface{}
	EbsVolumeType() ComputeoptimizerAutomationRuleCriteriaEbsVolumeTypeList
	EbsVolumeTypeInput() interface{}
	EstimatedMonthlySavings() ComputeoptimizerAutomationRuleCriteriaEstimatedMonthlySavingsList
	EstimatedMonthlySavingsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LookBackPeriodInDays() ComputeoptimizerAutomationRuleCriteriaLookBackPeriodInDaysList
	LookBackPeriodInDaysInput() interface{}
	Region() ComputeoptimizerAutomationRuleCriteriaRegionList
	RegionInput() interface{}
	ResourceArn() ComputeoptimizerAutomationRuleCriteriaResourceArnList
	ResourceArnInput() interface{}
	ResourceTag() ComputeoptimizerAutomationRuleCriteriaResourceTagList
	ResourceTagInput() interface{}
	RestartNeeded() ComputeoptimizerAutomationRuleCriteriaRestartNeededList
	RestartNeededInput() interface{}
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
	PutEbsVolumeSizeInGib(value interface{})
	PutEbsVolumeType(value interface{})
	PutEstimatedMonthlySavings(value interface{})
	PutLookBackPeriodInDays(value interface{})
	PutRegion(value interface{})
	PutResourceArn(value interface{})
	PutResourceTag(value interface{})
	PutRestartNeeded(value interface{})
	ResetEbsVolumeSizeInGib()
	ResetEbsVolumeType()
	ResetEstimatedMonthlySavings()
	ResetLookBackPeriodInDays()
	ResetRegion()
	ResetResourceArn()
	ResetResourceTag()
	ResetRestartNeeded()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeoptimizerAutomationRuleCriteriaOutputReference
type jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) EbsVolumeSizeInGib() ComputeoptimizerAutomationRuleCriteriaEbsVolumeSizeInGibList {
	var returns ComputeoptimizerAutomationRuleCriteriaEbsVolumeSizeInGibList
	_jsii_.Get(
		j,
		"ebsVolumeSizeInGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) EbsVolumeSizeInGibInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolumeSizeInGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) EbsVolumeType() ComputeoptimizerAutomationRuleCriteriaEbsVolumeTypeList {
	var returns ComputeoptimizerAutomationRuleCriteriaEbsVolumeTypeList
	_jsii_.Get(
		j,
		"ebsVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) EbsVolumeTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) EstimatedMonthlySavings() ComputeoptimizerAutomationRuleCriteriaEstimatedMonthlySavingsList {
	var returns ComputeoptimizerAutomationRuleCriteriaEstimatedMonthlySavingsList
	_jsii_.Get(
		j,
		"estimatedMonthlySavings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) EstimatedMonthlySavingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"estimatedMonthlySavingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) LookBackPeriodInDays() ComputeoptimizerAutomationRuleCriteriaLookBackPeriodInDaysList {
	var returns ComputeoptimizerAutomationRuleCriteriaLookBackPeriodInDaysList
	_jsii_.Get(
		j,
		"lookBackPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) LookBackPeriodInDaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lookBackPeriodInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) Region() ComputeoptimizerAutomationRuleCriteriaRegionList {
	var returns ComputeoptimizerAutomationRuleCriteriaRegionList
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) RegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ResourceArn() ComputeoptimizerAutomationRuleCriteriaResourceArnList {
	var returns ComputeoptimizerAutomationRuleCriteriaResourceArnList
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ResourceArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ResourceTag() ComputeoptimizerAutomationRuleCriteriaResourceTagList {
	var returns ComputeoptimizerAutomationRuleCriteriaResourceTagList
	_jsii_.Get(
		j,
		"resourceTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ResourceTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) RestartNeeded() ComputeoptimizerAutomationRuleCriteriaRestartNeededList {
	var returns ComputeoptimizerAutomationRuleCriteriaRestartNeededList
	_jsii_.Get(
		j,
		"restartNeeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) RestartNeededInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartNeededInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeoptimizerAutomationRuleCriteriaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ComputeoptimizerAutomationRuleCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewComputeoptimizerAutomationRuleCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.computeoptimizerAutomationRule.ComputeoptimizerAutomationRuleCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeoptimizerAutomationRuleCriteriaOutputReference_Override(c ComputeoptimizerAutomationRuleCriteriaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.computeoptimizerAutomationRule.ComputeoptimizerAutomationRuleCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) PutEbsVolumeSizeInGib(value interface{}) {
	if err := c.validatePutEbsVolumeSizeInGibParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEbsVolumeSizeInGib",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) PutEbsVolumeType(value interface{}) {
	if err := c.validatePutEbsVolumeTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEbsVolumeType",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) PutEstimatedMonthlySavings(value interface{}) {
	if err := c.validatePutEstimatedMonthlySavingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEstimatedMonthlySavings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) PutLookBackPeriodInDays(value interface{}) {
	if err := c.validatePutLookBackPeriodInDaysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLookBackPeriodInDays",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) PutRegion(value interface{}) {
	if err := c.validatePutRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRegion",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) PutResourceArn(value interface{}) {
	if err := c.validatePutResourceArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putResourceArn",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) PutResourceTag(value interface{}) {
	if err := c.validatePutResourceTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putResourceTag",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) PutRestartNeeded(value interface{}) {
	if err := c.validatePutRestartNeededParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRestartNeeded",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ResetEbsVolumeSizeInGib() {
	_jsii_.InvokeVoid(
		c,
		"resetEbsVolumeSizeInGib",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ResetEbsVolumeType() {
	_jsii_.InvokeVoid(
		c,
		"resetEbsVolumeType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ResetEstimatedMonthlySavings() {
	_jsii_.InvokeVoid(
		c,
		"resetEstimatedMonthlySavings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ResetLookBackPeriodInDays() {
	_jsii_.InvokeVoid(
		c,
		"resetLookBackPeriodInDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ResetResourceArn() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ResetResourceTag() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceTag",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ResetRestartNeeded() {
	_jsii_.InvokeVoid(
		c,
		"resetRestartNeeded",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeoptimizerAutomationRuleCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

