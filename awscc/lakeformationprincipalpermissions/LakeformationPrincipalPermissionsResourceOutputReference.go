// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationprincipalpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/lakeformationprincipalpermissions/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LakeformationPrincipalPermissionsResourceOutputReference interface {
	cdktn.ComplexObject
	Catalog() *string
	SetCatalog(val *string)
	CatalogInput() *string
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
	Database() LakeformationPrincipalPermissionsResourceDatabaseOutputReference
	DatabaseInput() interface{}
	DataCellsFilter() LakeformationPrincipalPermissionsResourceDataCellsFilterOutputReference
	DataCellsFilterInput() interface{}
	DataLocation() LakeformationPrincipalPermissionsResourceDataLocationOutputReference
	DataLocationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LfTag() LakeformationPrincipalPermissionsResourceLfTagOutputReference
	LfTagInput() interface{}
	LfTagPolicy() LakeformationPrincipalPermissionsResourceLfTagPolicyOutputReference
	LfTagPolicyInput() interface{}
	Table() LakeformationPrincipalPermissionsResourceTableOutputReference
	TableInput() interface{}
	TableWithColumns() LakeformationPrincipalPermissionsResourceTableWithColumnsOutputReference
	TableWithColumnsInput() interface{}
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
	PutDatabase(value *LakeformationPrincipalPermissionsResourceDatabase)
	PutDataCellsFilter(value *LakeformationPrincipalPermissionsResourceDataCellsFilter)
	PutDataLocation(value *LakeformationPrincipalPermissionsResourceDataLocation)
	PutLfTag(value *LakeformationPrincipalPermissionsResourceLfTag)
	PutLfTagPolicy(value *LakeformationPrincipalPermissionsResourceLfTagPolicy)
	PutTable(value *LakeformationPrincipalPermissionsResourceTable)
	PutTableWithColumns(value *LakeformationPrincipalPermissionsResourceTableWithColumns)
	ResetCatalog()
	ResetDatabase()
	ResetDataCellsFilter()
	ResetDataLocation()
	ResetLfTag()
	ResetLfTagPolicy()
	ResetTable()
	ResetTableWithColumns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LakeformationPrincipalPermissionsResourceOutputReference
type jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) Catalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) CatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) Database() LakeformationPrincipalPermissionsResourceDatabaseOutputReference {
	var returns LakeformationPrincipalPermissionsResourceDatabaseOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) DataCellsFilter() LakeformationPrincipalPermissionsResourceDataCellsFilterOutputReference {
	var returns LakeformationPrincipalPermissionsResourceDataCellsFilterOutputReference
	_jsii_.Get(
		j,
		"dataCellsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) DataCellsFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCellsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) DataLocation() LakeformationPrincipalPermissionsResourceDataLocationOutputReference {
	var returns LakeformationPrincipalPermissionsResourceDataLocationOutputReference
	_jsii_.Get(
		j,
		"dataLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) DataLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) LfTag() LakeformationPrincipalPermissionsResourceLfTagOutputReference {
	var returns LakeformationPrincipalPermissionsResourceLfTagOutputReference
	_jsii_.Get(
		j,
		"lfTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) LfTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) LfTagPolicy() LakeformationPrincipalPermissionsResourceLfTagPolicyOutputReference {
	var returns LakeformationPrincipalPermissionsResourceLfTagPolicyOutputReference
	_jsii_.Get(
		j,
		"lfTagPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) LfTagPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfTagPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) Table() LakeformationPrincipalPermissionsResourceTableOutputReference {
	var returns LakeformationPrincipalPermissionsResourceTableOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) TableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) TableWithColumns() LakeformationPrincipalPermissionsResourceTableWithColumnsOutputReference {
	var returns LakeformationPrincipalPermissionsResourceTableWithColumnsOutputReference
	_jsii_.Get(
		j,
		"tableWithColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) TableWithColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableWithColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLakeformationPrincipalPermissionsResourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LakeformationPrincipalPermissionsResourceOutputReference {
	_init_.Initialize()

	if err := validateNewLakeformationPrincipalPermissionsResourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.lakeformationPrincipalPermissions.LakeformationPrincipalPermissionsResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLakeformationPrincipalPermissionsResourceOutputReference_Override(l LakeformationPrincipalPermissionsResourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.lakeformationPrincipalPermissions.LakeformationPrincipalPermissionsResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference)SetCatalog(val *string) {
	if err := j.validateSetCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalog",
		val,
	)
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) PutDatabase(value *LakeformationPrincipalPermissionsResourceDatabase) {
	if err := l.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDatabase",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) PutDataCellsFilter(value *LakeformationPrincipalPermissionsResourceDataCellsFilter) {
	if err := l.validatePutDataCellsFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDataCellsFilter",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) PutDataLocation(value *LakeformationPrincipalPermissionsResourceDataLocation) {
	if err := l.validatePutDataLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDataLocation",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) PutLfTag(value *LakeformationPrincipalPermissionsResourceLfTag) {
	if err := l.validatePutLfTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLfTag",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) PutLfTagPolicy(value *LakeformationPrincipalPermissionsResourceLfTagPolicy) {
	if err := l.validatePutLfTagPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLfTagPolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) PutTable(value *LakeformationPrincipalPermissionsResourceTable) {
	if err := l.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTable",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) PutTableWithColumns(value *LakeformationPrincipalPermissionsResourceTableWithColumns) {
	if err := l.validatePutTableWithColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTableWithColumns",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) ResetCatalog() {
	_jsii_.InvokeVoid(
		l,
		"resetCatalog",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		l,
		"resetDatabase",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) ResetDataCellsFilter() {
	_jsii_.InvokeVoid(
		l,
		"resetDataCellsFilter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) ResetDataLocation() {
	_jsii_.InvokeVoid(
		l,
		"resetDataLocation",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) ResetLfTag() {
	_jsii_.InvokeVoid(
		l,
		"resetLfTag",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) ResetLfTagPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetLfTagPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) ResetTable() {
	_jsii_.InvokeVoid(
		l,
		"resetTable",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) ResetTableWithColumns() {
	_jsii_.InvokeVoid(
		l,
		"resetTableWithColumns",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPrincipalPermissionsResourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

