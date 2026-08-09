// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package acmpcacertificateauthority

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validatePutDirectoryNameParameters(value *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationDirectoryName) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validatePutEdiPartyNameParameters(value *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationEdiPartyName) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validatePutOtherNameParameters(value *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOtherName) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateSetDnsNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocation:
		val := val.(*AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocation)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocation:
		val_ := val.(AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocation)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocation; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateSetIpAddressParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateSetRegisteredIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateSetRfc822NameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference) validateSetUniformResourceIdentifierParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

