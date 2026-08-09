// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/greengrassv2deployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference interface {
	cdktn.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	FailureType() *string
	SetFailureType(val *string)
	FailureTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinNumberOfExecutedThings() *float64
	SetMinNumberOfExecutedThings(val *float64)
	MinNumberOfExecutedThingsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThresholdPercentage() *float64
	SetThresholdPercentage(val *float64)
	ThresholdPercentageInput() *float64
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
	ResetAction()
	ResetFailureType()
	ResetMinNumberOfExecutedThings()
	ResetThresholdPercentage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference
type jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) FailureType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) FailureTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) MinNumberOfExecutedThings() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumberOfExecutedThings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) MinNumberOfExecutedThingsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumberOfExecutedThingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) ThresholdPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) ThresholdPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdPercentageInput",
		&returns,
	)
	return returns
}


func NewGreengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference {
	_init_.Initialize()

	if err := validateNewGreengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.greengrassv2Deployment.Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGreengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference_Override(g Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.greengrassv2Deployment.Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference)SetFailureType(val *string) {
	if err := j.validateSetFailureTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureType",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference)SetMinNumberOfExecutedThings(val *float64) {
	if err := j.validateSetMinNumberOfExecutedThingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNumberOfExecutedThings",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference)SetThresholdPercentage(val *float64) {
	if err := j.validateSetThresholdPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdPercentage",
		val,
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		g,
		"resetAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) ResetFailureType() {
	_jsii_.InvokeVoid(
		g,
		"resetFailureType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) ResetMinNumberOfExecutedThings() {
	_jsii_.InvokeVoid(
		g,
		"resetMinNumberOfExecutedThings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) ResetThresholdPercentage() {
	_jsii_.InvokeVoid(
		g,
		"resetThresholdPercentage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

