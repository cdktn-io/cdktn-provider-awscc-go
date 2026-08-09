// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ssmassociation

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmAssociationTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SsmAssociationTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsmAssociationTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsmAssociationTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SsmAssociationTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsmAssociationTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsmAssociationTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsmAssociationTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

