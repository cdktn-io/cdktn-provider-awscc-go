// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package quicksightdatasource

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateSetCatalogParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateSetHostParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParameters:
		val := val.(*QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParameters)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParameters:
		val_ := val.(QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParameters)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParameters; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateSetPortParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersTrinoParametersOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

