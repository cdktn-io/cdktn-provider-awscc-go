// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ivsstreamkey

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IvsStreamKeyTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IvsStreamKeyTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IvsStreamKeyTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IvsStreamKeyTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IvsStreamKeyTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IvsStreamKeyTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IvsStreamKeyTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIvsStreamKeyTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

