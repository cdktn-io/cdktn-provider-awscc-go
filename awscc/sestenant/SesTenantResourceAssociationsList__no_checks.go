// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sestenant

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SesTenantResourceAssociationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SesTenantResourceAssociationsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SesTenantResourceAssociationsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SesTenantResourceAssociationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SesTenantResourceAssociationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SesTenantResourceAssociationsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SesTenantResourceAssociationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSesTenantResourceAssociationsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

