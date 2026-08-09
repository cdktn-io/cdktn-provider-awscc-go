// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ssmdocument

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmDocumentRequiresList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SsmDocumentRequiresList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsmDocumentRequiresList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentRequiresList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentRequiresList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentRequiresList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentRequiresList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsmDocumentRequiresListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

