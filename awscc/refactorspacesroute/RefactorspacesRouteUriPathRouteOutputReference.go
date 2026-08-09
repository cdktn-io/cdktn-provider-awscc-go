// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package refactorspacesroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/refactorspacesroute/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RefactorspacesRouteUriPathRouteOutputReference interface {
	cdktn.ComplexObject
	ActivationState() *string
	SetActivationState(val *string)
	ActivationStateInput() *string
	AppendSourcePath() interface{}
	SetAppendSourcePath(val interface{})
	AppendSourcePathInput() interface{}
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
	IncludeChildPaths() interface{}
	SetIncludeChildPaths(val interface{})
	IncludeChildPathsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Methods() *[]*string
	SetMethods(val *[]*string)
	MethodsInput() *[]*string
	SourcePath() *string
	SetSourcePath(val *string)
	SourcePathInput() *string
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
	ResetActivationState()
	ResetAppendSourcePath()
	ResetIncludeChildPaths()
	ResetMethods()
	ResetSourcePath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RefactorspacesRouteUriPathRouteOutputReference
type jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) ActivationState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) ActivationStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) AppendSourcePath() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendSourcePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) AppendSourcePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendSourcePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) IncludeChildPaths() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeChildPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) IncludeChildPathsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeChildPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) Methods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) MethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) SourcePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) SourcePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRefactorspacesRouteUriPathRouteOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) RefactorspacesRouteUriPathRouteOutputReference {
	_init_.Initialize()

	if err := validateNewRefactorspacesRouteUriPathRouteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.refactorspacesRoute.RefactorspacesRouteUriPathRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRefactorspacesRouteUriPathRouteOutputReference_Override(r RefactorspacesRouteUriPathRouteOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.refactorspacesRoute.RefactorspacesRouteUriPathRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference)SetActivationState(val *string) {
	if err := j.validateSetActivationStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activationState",
		val,
	)
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference)SetAppendSourcePath(val interface{}) {
	if err := j.validateSetAppendSourcePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appendSourcePath",
		val,
	)
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference)SetIncludeChildPaths(val interface{}) {
	if err := j.validateSetIncludeChildPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeChildPaths",
		val,
	)
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference)SetMethods(val *[]*string) {
	if err := j.validateSetMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"methods",
		val,
	)
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference)SetSourcePath(val *string) {
	if err := j.validateSetSourcePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePath",
		val,
	)
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) ResetActivationState() {
	_jsii_.InvokeVoid(
		r,
		"resetActivationState",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) ResetAppendSourcePath() {
	_jsii_.InvokeVoid(
		r,
		"resetAppendSourcePath",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) ResetIncludeChildPaths() {
	_jsii_.InvokeVoid(
		r,
		"resetIncludeChildPaths",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) ResetMethods() {
	_jsii_.InvokeVoid(
		r,
		"resetMethods",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) ResetSourcePath() {
	_jsii_.InvokeVoid(
		r,
		"resetSourcePath",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RefactorspacesRouteUriPathRouteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

