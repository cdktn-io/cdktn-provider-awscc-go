// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscomputenodegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/pcscomputenodegroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference interface {
	cdktn.ComplexObject
	Checksum() *string
	SetChecksum(val *string)
	ChecksumInput() *string
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
	S3VersionId() *string
	SetS3VersionId(val *string)
	S3VersionIdInput() *string
	ScriptLocation() *string
	SetScriptLocation(val *string)
	ScriptLocationInput() *string
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
	ResetChecksum()
	ResetS3VersionId()
	ResetScriptLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference
type jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) ChecksumInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) S3VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3VersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) S3VersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3VersionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) ScriptLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) ScriptLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference {
	_init_.Initialize()

	if err := validateNewPcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.pcsComputeNodeGroup.PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference_Override(p PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.pcsComputeNodeGroup.PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference)SetChecksum(val *string) {
	if err := j.validateSetChecksumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checksum",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference)SetS3VersionId(val *string) {
	if err := j.validateSetS3VersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3VersionId",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference)SetScriptLocation(val *string) {
	if err := j.validateSetScriptLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptLocation",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) ResetChecksum() {
	_jsii_.InvokeVoid(
		p,
		"resetChecksum",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) ResetS3VersionId() {
	_jsii_.InvokeVoid(
		p,
		"resetS3VersionId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) ResetScriptLocation() {
	_jsii_.InvokeVoid(
		p,
		"resetScriptLocation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

