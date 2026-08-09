// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kendradatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference interface {
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
	CrawlArchivedSpaces() interface{}
	SetCrawlArchivedSpaces(val interface{})
	CrawlArchivedSpacesInput() interface{}
	CrawlPersonalSpaces() interface{}
	SetCrawlPersonalSpaces(val interface{})
	CrawlPersonalSpacesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludeSpaces() *[]*string
	SetExcludeSpaces(val *[]*string)
	ExcludeSpacesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeSpaces() *[]*string
	SetIncludeSpaces(val *[]*string)
	IncludeSpacesInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SpaceFieldMappings() KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationSpaceFieldMappingsList
	SpaceFieldMappingsInput() interface{}
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
	PutSpaceFieldMappings(value interface{})
	ResetCrawlArchivedSpaces()
	ResetCrawlPersonalSpaces()
	ResetExcludeSpaces()
	ResetIncludeSpaces()
	ResetSpaceFieldMappings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference
type jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) CrawlArchivedSpaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlArchivedSpaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) CrawlArchivedSpacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlArchivedSpacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) CrawlPersonalSpaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlPersonalSpaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) CrawlPersonalSpacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlPersonalSpacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) ExcludeSpaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeSpaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) ExcludeSpacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeSpacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) IncludeSpaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeSpaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) IncludeSpacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeSpacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) SpaceFieldMappings() KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationSpaceFieldMappingsList {
	var returns KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationSpaceFieldMappingsList
	_jsii_.Get(
		j,
		"spaceFieldMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) SpaceFieldMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spaceFieldMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference_Override(k KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference)SetCrawlArchivedSpaces(val interface{}) {
	if err := j.validateSetCrawlArchivedSpacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlArchivedSpaces",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference)SetCrawlPersonalSpaces(val interface{}) {
	if err := j.validateSetCrawlPersonalSpacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlPersonalSpaces",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference)SetExcludeSpaces(val *[]*string) {
	if err := j.validateSetExcludeSpacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeSpaces",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference)SetIncludeSpaces(val *[]*string) {
	if err := j.validateSetIncludeSpacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSpaces",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) PutSpaceFieldMappings(value interface{}) {
	if err := k.validatePutSpaceFieldMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSpaceFieldMappings",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) ResetCrawlArchivedSpaces() {
	_jsii_.InvokeVoid(
		k,
		"resetCrawlArchivedSpaces",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) ResetCrawlPersonalSpaces() {
	_jsii_.InvokeVoid(
		k,
		"resetCrawlPersonalSpaces",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) ResetExcludeSpaces() {
	_jsii_.InvokeVoid(
		k,
		"resetExcludeSpaces",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) ResetIncludeSpaces() {
	_jsii_.InvokeVoid(
		k,
		"resetIncludeSpaces",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) ResetSpaceFieldMappings() {
	_jsii_.InvokeVoid(
		k,
		"resetSpaceFieldMappings",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

