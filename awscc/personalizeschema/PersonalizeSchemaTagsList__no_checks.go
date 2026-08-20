// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package personalizeschema

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PersonalizeSchemaTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PersonalizeSchemaTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PersonalizeSchemaTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PersonalizeSchemaTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PersonalizeSchemaTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PersonalizeSchemaTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PersonalizeSchemaTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPersonalizeSchemaTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

