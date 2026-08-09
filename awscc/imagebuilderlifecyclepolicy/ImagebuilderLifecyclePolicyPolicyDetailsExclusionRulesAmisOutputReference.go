// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderlifecyclepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/imagebuilderlifecyclepolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsPublic() interface{}
	SetIsPublic(val interface{})
	IsPublicInput() interface{}
	LastLaunched() ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisLastLaunchedOutputReference
	LastLaunchedInput() interface{}
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
	SharedAccounts() *[]*string
	SetSharedAccounts(val *[]*string)
	SharedAccountsInput() *[]*string
	TagMap() *map[string]*string
	SetTagMap(val *map[string]*string)
	TagMapInput() *map[string]*string
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
	PutLastLaunched(value *ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisLastLaunched)
	ResetIsPublic()
	ResetLastLaunched()
	ResetRegions()
	ResetSharedAccounts()
	ResetTagMap()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference
type jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) IsPublic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) IsPublicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPublicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) LastLaunched() ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisLastLaunchedOutputReference {
	var returns ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisLastLaunchedOutputReference
	_jsii_.Get(
		j,
		"lastLaunched",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) LastLaunchedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastLaunchedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) SharedAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) SharedAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) TagMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) TagMapInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference {
	_init_.Initialize()

	if err := validateNewImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.imagebuilderLifecyclePolicy.ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference_Override(i ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.imagebuilderLifecyclePolicy.ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference)SetIsPublic(val interface{}) {
	if err := j.validateSetIsPublicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPublic",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference)SetSharedAccounts(val *[]*string) {
	if err := j.validateSetSharedAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAccounts",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference)SetTagMap(val *map[string]*string) {
	if err := j.validateSetTagMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagMap",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) PutLastLaunched(value *ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisLastLaunched) {
	if err := i.validatePutLastLaunchedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLastLaunched",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) ResetIsPublic() {
	_jsii_.InvokeVoid(
		i,
		"resetIsPublic",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) ResetLastLaunched() {
	_jsii_.InvokeVoid(
		i,
		"resetLastLaunched",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		i,
		"resetRegions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) ResetSharedAccounts() {
	_jsii_.InvokeVoid(
		i,
		"resetSharedAccounts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) ResetTagMap() {
	_jsii_.InvokeVoid(
		i,
		"resetTagMap",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

