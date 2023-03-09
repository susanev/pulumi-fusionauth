// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "fusionauth:index/fusionAuthApiKey:FusionAuthApiKey":
		r = &FusionAuthApiKey{}
	case "fusionauth:index/fusionAuthApplication:FusionAuthApplication":
		r = &FusionAuthApplication{}
	case "fusionauth:index/fusionAuthApplicationRole:FusionAuthApplicationRole":
		r = &FusionAuthApplicationRole{}
	case "fusionauth:index/fusionAuthEMail:FusionAuthEMail":
		r = &FusionAuthEMail{}
	case "fusionauth:index/fusionAuthEntity:FusionAuthEntity":
		r = &FusionAuthEntity{}
	case "fusionauth:index/fusionAuthEntityGrant:FusionAuthEntityGrant":
		r = &FusionAuthEntityGrant{}
	case "fusionauth:index/fusionAuthEntityType:FusionAuthEntityType":
		r = &FusionAuthEntityType{}
	case "fusionauth:index/fusionAuthEntityTypePermission:FusionAuthEntityTypePermission":
		r = &FusionAuthEntityTypePermission{}
	case "fusionauth:index/fusionAuthForm:FusionAuthForm":
		r = &FusionAuthForm{}
	case "fusionauth:index/fusionAuthFormField:FusionAuthFormField":
		r = &FusionAuthFormField{}
	case "fusionauth:index/fusionAuthGenericConnector:FusionAuthGenericConnector":
		r = &FusionAuthGenericConnector{}
	case "fusionauth:index/fusionAuthGroup:FusionAuthGroup":
		r = &FusionAuthGroup{}
	case "fusionauth:index/fusionAuthIdpApple:FusionAuthIdpApple":
		r = &FusionAuthIdpApple{}
	case "fusionauth:index/fusionAuthIdpExternalJwt:FusionAuthIdpExternalJwt":
		r = &FusionAuthIdpExternalJwt{}
	case "fusionauth:index/fusionAuthIdpFacebook:FusionAuthIdpFacebook":
		r = &FusionAuthIdpFacebook{}
	case "fusionauth:index/fusionAuthIdpGoogle:FusionAuthIdpGoogle":
		r = &FusionAuthIdpGoogle{}
	case "fusionauth:index/fusionAuthIdpLinkedIn:FusionAuthIdpLinkedIn":
		r = &FusionAuthIdpLinkedIn{}
	case "fusionauth:index/fusionAuthIdpOpenIdConnect:FusionAuthIdpOpenIdConnect":
		r = &FusionAuthIdpOpenIdConnect{}
	case "fusionauth:index/fusionAuthIdpPsn:FusionAuthIdpPsn":
		r = &FusionAuthIdpPsn{}
	case "fusionauth:index/fusionAuthIdpSamlV2IdpInitiated:FusionAuthIdpSamlV2IdpInitiated":
		r = &FusionAuthIdpSamlV2IdpInitiated{}
	case "fusionauth:index/fusionAuthIdpSamlv2:FusionAuthIdpSamlv2":
		r = &FusionAuthIdpSamlv2{}
	case "fusionauth:index/fusionAuthIdpSteam:FusionAuthIdpSteam":
		r = &FusionAuthIdpSteam{}
	case "fusionauth:index/fusionAuthIdpTwitch:FusionAuthIdpTwitch":
		r = &FusionAuthIdpTwitch{}
	case "fusionauth:index/fusionAuthIdpXBox:FusionAuthIdpXBox":
		r = &FusionAuthIdpXBox{}
	case "fusionauth:index/fusionAuthImportedKey:FusionAuthImportedKey":
		r = &FusionAuthImportedKey{}
	case "fusionauth:index/fusionAuthKey:FusionAuthKey":
		r = &FusionAuthKey{}
	case "fusionauth:index/fusionAuthLambda:FusionAuthLambda":
		r = &FusionAuthLambda{}
	case "fusionauth:index/fusionAuthReactor:FusionAuthReactor":
		r = &FusionAuthReactor{}
	case "fusionauth:index/fusionAuthRegistration:FusionAuthRegistration":
		r = &FusionAuthRegistration{}
	case "fusionauth:index/fusionAuthSystemConfiguration:FusionAuthSystemConfiguration":
		r = &FusionAuthSystemConfiguration{}
	case "fusionauth:index/fusionAuthTenant:FusionAuthTenant":
		r = &FusionAuthTenant{}
	case "fusionauth:index/fusionAuthTheme:FusionAuthTheme":
		r = &FusionAuthTheme{}
	case "fusionauth:index/fusionAuthUser:FusionAuthUser":
		r = &FusionAuthUser{}
	case "fusionauth:index/fusionAuthUserAction:FusionAuthUserAction":
		r = &FusionAuthUserAction{}
	case "fusionauth:index/fusionAuthWebhook:FusionAuthWebhook":
		r = &FusionAuthWebhook{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:fusionauth" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthApiKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthApplication",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthApplicationRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthEMail",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthEntity",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthEntityGrant",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthEntityType",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthEntityTypePermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthForm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthFormField",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthGenericConnector",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthIdpApple",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthIdpExternalJwt",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthIdpFacebook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthIdpGoogle",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthIdpLinkedIn",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthIdpOpenIdConnect",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthIdpPsn",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthIdpSamlV2IdpInitiated",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthIdpSamlv2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthIdpSteam",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthIdpTwitch",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthIdpXBox",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthImportedKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthLambda",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthReactor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthRegistration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthSystemConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthTenant",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthTheme",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthUserAction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fusionauth",
		"index/fusionAuthWebhook",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"fusionauth",
		&pkg{version},
	)
}
