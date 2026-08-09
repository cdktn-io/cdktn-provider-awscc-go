// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package scnnamespace

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScnNamespaceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ScnNamespaceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ScnNamespaceTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ScnNamespaceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ScnNamespaceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ScnNamespaceTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ScnNamespaceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewScnNamespaceTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

