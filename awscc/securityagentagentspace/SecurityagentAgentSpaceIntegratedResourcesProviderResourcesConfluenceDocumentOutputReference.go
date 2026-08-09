// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/securityagentagentspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference interface {
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	PageId() *string
	SetPageId(val *string)
	PageIdInput() *string
	SpaceKey() *string
	SetSpaceKey(val *string)
	SpaceKeyInput() *string
	SpaceTitle() *string
	SetSpaceTitle(val *string)
	SpaceTitleInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
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
	ResetName()
	ResetPageId()
	ResetSpaceKey()
	ResetSpaceTitle()
	ResetTitle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference
type jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) PageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) PageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) SpaceKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spaceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) SpaceKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spaceKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) SpaceTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spaceTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) SpaceTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spaceTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}


func NewSecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.securityagentAgentSpace.SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference_Override(s SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.securityagentAgentSpace.SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference)SetPageId(val *string) {
	if err := j.validateSetPageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pageId",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference)SetSpaceKey(val *string) {
	if err := j.validateSetSpaceKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spaceKey",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference)SetSpaceTitle(val *string) {
	if err := j.validateSetSpaceTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spaceTitle",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) ResetPageId() {
	_jsii_.InvokeVoid(
		s,
		"resetPageId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) ResetSpaceKey() {
	_jsii_.InvokeVoid(
		s,
		"resetSpaceKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) ResetSpaceTitle() {
	_jsii_.InvokeVoid(
		s,
		"resetSpaceTitle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		s,
		"resetTitle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocumentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

