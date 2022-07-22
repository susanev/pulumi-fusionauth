// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # SAML v2 Identity Provider Resource
//
// SAML v2 identity providers connect to external SAML v2 login systems. This type of login will optionally provide a Login with …​ button on FusionAuth’s login page. This button is customizable by using different properties of the identity provider.
//
// Optionally, this identity provider can define one or more domains it is associated with. This is useful for allowing employees to log in with their corporate credentials. As long as the company has an identity solution that provides SAML v2, you can leverage this feature. This is referred to as a Domain Based Identity Provider. If you enable domains for an identity provider, the Login with …​ button will not be displayed. Instead, only the email form field will be displayed initially on the FusionAuth login page. Once the user types in their email address, FusionAuth will determine if the user is logging in locally or if they should be redirected to this identity provider. This is determined by extracting the domain from their email address and comparing it to the domains associated with the identity provider.
//
// FusionAuth will locate the user’s email address in the SAML assertion which will be used to create or lookup the existing user. Additional claims from the SAML response can be used to reconcile the User to FusionAuth by using a SAML v2 Reconcile Lambda. Unless you assign a reconcile lambda to this provider, on the email address will be used from the available assertions returned by the SAML v2 identity provider.
//
// [SAML v2 Connect Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/samlv2/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fusionauth/sdk/go/fusionauth"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fusionauth.NewFusionAuthIdpSamlv2(ctx, "saml", &fusionauth.FusionAuthIdpSamlv2Args{
// 			ApplicationConfigurations: FusionAuthIdpSamlv2ApplicationConfigurationArray{
// 				&FusionAuthIdpSamlv2ApplicationConfigurationArgs{
// 					ApplicationId:      pulumi.Any(fusionauth_application.Myapp.Id),
// 					ButtonText:         pulumi.String("Login with SAML (app text)"),
// 					CreateRegistration: pulumi.Bool(true),
// 					Enabled:            pulumi.Bool(true),
// 				},
// 			},
// 			ButtonText:        pulumi.String("Login with SAML"),
// 			Debug:             pulumi.Bool(false),
// 			EmailClaim:        pulumi.String("email"),
// 			IdpEndpoint:       pulumi.String("https://www.example.com/login"),
// 			PostRequest:       pulumi.Bool(true),
// 			RequestSigningKey: pulumi.String("3168129b-91fa-46f4-9676-947f5509fdce"),
// 			SignRequest:       pulumi.Bool(true),
// 			UseNameForEmail:   pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FusionAuthIdpSamlv2 struct {
	pulumi.CustomResourceState

	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpSamlv2ApplicationConfigurationArrayOutput `pulumi:"applicationConfigurations"`
	// The top-level button image (URL) to use on the FusionAuth login page for this Identity Provider.
	ButtonImageUrl pulumi.StringPtrOutput `pulumi:"buttonImageUrl"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringOutput `pulumi:"buttonText"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrOutput `pulumi:"debug"`
	// This is an optional list of domains that this OpenID Connect provider should be used for. This converts the FusionAuth login form to a domain-based login form. This type of form first asks the user for their email. FusionAuth then uses their email to determine if an OpenID Connect identity provider should be used. If an OpenID Connect provider should be used, the browser is redirected to the authorization endpoint of that identity provider. Otherwise, the password field is revealed on the form so that the user can login using FusionAuth.
	Domains pulumi.StringArrayOutput `pulumi:"domains"`
	// The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely identity the user. If this is not set, the `useNameForEmail` flag must be true.
	EmailClaim pulumi.StringPtrOutput `pulumi:"emailClaim"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The SAML v2 login page of the identity provider.
	IdpEndpoint pulumi.StringPtrOutput `pulumi:"idpEndpoint"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrOutput `pulumi:"idpId"`
	// The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the identity provider. This key must be a verification only key or certificate (meaning that it only has a public key component).
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrOutput `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringOutput `pulumi:"linkingStrategy"`
	// The name of this OpenID Connect identity provider. This is only used for display purposes.
	Name pulumi.StringOutput `pulumi:"name"`
	// Either urn:oasis:names:tc:SAML:2.0:nameid-format:persistent or urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress depending on which NameId format you wish to use.
	NameIdFormat pulumi.StringOutput `pulumi:"nameIdFormat"`
	// Set this value equal to true if you wish to use POST bindings with this OpenID Connect identity provider. The default value of false means that a redirect binding which uses a GET request will be used.
	PostRequest pulumi.BoolPtrOutput `pulumi:"postRequest"`
	// TThe key pair Id to use to sign the SAML request. Required when `signRequest` is true.
	RequestSigningKey pulumi.StringPtrOutput `pulumi:"requestSigningKey"`
	// When true authentication requests sent to the identity provider will be signed.
	SignRequest pulumi.BoolPtrOutput `pulumi:"signRequest"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpSamlv2TenantConfigurationArrayOutput `pulumi:"tenantConfigurations"`
	// Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation processing. If this is false, then the `emailClaim` property must be set.
	UseNameForEmail pulumi.BoolPtrOutput `pulumi:"useNameForEmail"`
	// The XML signature canonicalization method used when digesting and signing the SAML request.
	XmlSignatureCanonicalizationMethod pulumi.StringPtrOutput `pulumi:"xmlSignatureCanonicalizationMethod"`
}

// NewFusionAuthIdpSamlv2 registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthIdpSamlv2(ctx *pulumi.Context,
	name string, args *FusionAuthIdpSamlv2Args, opts ...pulumi.ResourceOption) (*FusionAuthIdpSamlv2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ButtonText == nil {
		return nil, errors.New("invalid value for required argument 'ButtonText'")
	}
	if args.KeyId == nil {
		return nil, errors.New("invalid value for required argument 'KeyId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FusionAuthIdpSamlv2
	err := ctx.RegisterResource("fusionauth:index/fusionAuthIdpSamlv2:FusionAuthIdpSamlv2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthIdpSamlv2 gets an existing FusionAuthIdpSamlv2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthIdpSamlv2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthIdpSamlv2State, opts ...pulumi.ResourceOption) (*FusionAuthIdpSamlv2, error) {
	var resource FusionAuthIdpSamlv2
	err := ctx.ReadResource("fusionauth:index/fusionAuthIdpSamlv2:FusionAuthIdpSamlv2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthIdpSamlv2 resources.
type fusionAuthIdpSamlv2State struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpSamlv2ApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The top-level button image (URL) to use on the FusionAuth login page for this Identity Provider.
	ButtonImageUrl *string `pulumi:"buttonImageUrl"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText *string `pulumi:"buttonText"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// This is an optional list of domains that this OpenID Connect provider should be used for. This converts the FusionAuth login form to a domain-based login form. This type of form first asks the user for their email. FusionAuth then uses their email to determine if an OpenID Connect identity provider should be used. If an OpenID Connect provider should be used, the browser is redirected to the authorization endpoint of that identity provider. Otherwise, the password field is revealed on the form so that the user can login using FusionAuth.
	Domains []string `pulumi:"domains"`
	// The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely identity the user. If this is not set, the `useNameForEmail` flag must be true.
	EmailClaim *string `pulumi:"emailClaim"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The SAML v2 login page of the identity provider.
	IdpEndpoint *string `pulumi:"idpEndpoint"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId *string `pulumi:"idpId"`
	// The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the identity provider. This key must be a verification only key or certificate (meaning that it only has a public key component).
	KeyId *string `pulumi:"keyId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The name of this OpenID Connect identity provider. This is only used for display purposes.
	Name *string `pulumi:"name"`
	// Either urn:oasis:names:tc:SAML:2.0:nameid-format:persistent or urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress depending on which NameId format you wish to use.
	NameIdFormat *string `pulumi:"nameIdFormat"`
	// Set this value equal to true if you wish to use POST bindings with this OpenID Connect identity provider. The default value of false means that a redirect binding which uses a GET request will be used.
	PostRequest *bool `pulumi:"postRequest"`
	// TThe key pair Id to use to sign the SAML request. Required when `signRequest` is true.
	RequestSigningKey *string `pulumi:"requestSigningKey"`
	// When true authentication requests sent to the identity provider will be signed.
	SignRequest *bool `pulumi:"signRequest"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpSamlv2TenantConfiguration `pulumi:"tenantConfigurations"`
	// Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation processing. If this is false, then the `emailClaim` property must be set.
	UseNameForEmail *bool `pulumi:"useNameForEmail"`
	// The XML signature canonicalization method used when digesting and signing the SAML request.
	XmlSignatureCanonicalizationMethod *string `pulumi:"xmlSignatureCanonicalizationMethod"`
}

type FusionAuthIdpSamlv2State struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpSamlv2ApplicationConfigurationArrayInput
	// The top-level button image (URL) to use on the FusionAuth login page for this Identity Provider.
	ButtonImageUrl pulumi.StringPtrInput
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringPtrInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// This is an optional list of domains that this OpenID Connect provider should be used for. This converts the FusionAuth login form to a domain-based login form. This type of form first asks the user for their email. FusionAuth then uses their email to determine if an OpenID Connect identity provider should be used. If an OpenID Connect provider should be used, the browser is redirected to the authorization endpoint of that identity provider. Otherwise, the password field is revealed on the form so that the user can login using FusionAuth.
	Domains pulumi.StringArrayInput
	// The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely identity the user. If this is not set, the `useNameForEmail` flag must be true.
	EmailClaim pulumi.StringPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The SAML v2 login page of the identity provider.
	IdpEndpoint pulumi.StringPtrInput
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrInput
	// The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the identity provider. This key must be a verification only key or certificate (meaning that it only has a public key component).
	KeyId pulumi.StringPtrInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The name of this OpenID Connect identity provider. This is only used for display purposes.
	Name pulumi.StringPtrInput
	// Either urn:oasis:names:tc:SAML:2.0:nameid-format:persistent or urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress depending on which NameId format you wish to use.
	NameIdFormat pulumi.StringPtrInput
	// Set this value equal to true if you wish to use POST bindings with this OpenID Connect identity provider. The default value of false means that a redirect binding which uses a GET request will be used.
	PostRequest pulumi.BoolPtrInput
	// TThe key pair Id to use to sign the SAML request. Required when `signRequest` is true.
	RequestSigningKey pulumi.StringPtrInput
	// When true authentication requests sent to the identity provider will be signed.
	SignRequest pulumi.BoolPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpSamlv2TenantConfigurationArrayInput
	// Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation processing. If this is false, then the `emailClaim` property must be set.
	UseNameForEmail pulumi.BoolPtrInput
	// The XML signature canonicalization method used when digesting and signing the SAML request.
	XmlSignatureCanonicalizationMethod pulumi.StringPtrInput
}

func (FusionAuthIdpSamlv2State) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpSamlv2State)(nil)).Elem()
}

type fusionAuthIdpSamlv2Args struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpSamlv2ApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The top-level button image (URL) to use on the FusionAuth login page for this Identity Provider.
	ButtonImageUrl *string `pulumi:"buttonImageUrl"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText string `pulumi:"buttonText"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// This is an optional list of domains that this OpenID Connect provider should be used for. This converts the FusionAuth login form to a domain-based login form. This type of form first asks the user for their email. FusionAuth then uses their email to determine if an OpenID Connect identity provider should be used. If an OpenID Connect provider should be used, the browser is redirected to the authorization endpoint of that identity provider. Otherwise, the password field is revealed on the form so that the user can login using FusionAuth.
	Domains []string `pulumi:"domains"`
	// The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely identity the user. If this is not set, the `useNameForEmail` flag must be true.
	EmailClaim *string `pulumi:"emailClaim"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The SAML v2 login page of the identity provider.
	IdpEndpoint *string `pulumi:"idpEndpoint"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId *string `pulumi:"idpId"`
	// The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the identity provider. This key must be a verification only key or certificate (meaning that it only has a public key component).
	KeyId string `pulumi:"keyId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The name of this OpenID Connect identity provider. This is only used for display purposes.
	Name *string `pulumi:"name"`
	// Either urn:oasis:names:tc:SAML:2.0:nameid-format:persistent or urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress depending on which NameId format you wish to use.
	NameIdFormat *string `pulumi:"nameIdFormat"`
	// Set this value equal to true if you wish to use POST bindings with this OpenID Connect identity provider. The default value of false means that a redirect binding which uses a GET request will be used.
	PostRequest *bool `pulumi:"postRequest"`
	// TThe key pair Id to use to sign the SAML request. Required when `signRequest` is true.
	RequestSigningKey *string `pulumi:"requestSigningKey"`
	// When true authentication requests sent to the identity provider will be signed.
	SignRequest *bool `pulumi:"signRequest"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpSamlv2TenantConfiguration `pulumi:"tenantConfigurations"`
	// Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation processing. If this is false, then the `emailClaim` property must be set.
	UseNameForEmail *bool `pulumi:"useNameForEmail"`
	// The XML signature canonicalization method used when digesting and signing the SAML request.
	XmlSignatureCanonicalizationMethod *string `pulumi:"xmlSignatureCanonicalizationMethod"`
}

// The set of arguments for constructing a FusionAuthIdpSamlv2 resource.
type FusionAuthIdpSamlv2Args struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpSamlv2ApplicationConfigurationArrayInput
	// The top-level button image (URL) to use on the FusionAuth login page for this Identity Provider.
	ButtonImageUrl pulumi.StringPtrInput
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// This is an optional list of domains that this OpenID Connect provider should be used for. This converts the FusionAuth login form to a domain-based login form. This type of form first asks the user for their email. FusionAuth then uses their email to determine if an OpenID Connect identity provider should be used. If an OpenID Connect provider should be used, the browser is redirected to the authorization endpoint of that identity provider. Otherwise, the password field is revealed on the form so that the user can login using FusionAuth.
	Domains pulumi.StringArrayInput
	// The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely identity the user. If this is not set, the `useNameForEmail` flag must be true.
	EmailClaim pulumi.StringPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The SAML v2 login page of the identity provider.
	IdpEndpoint pulumi.StringPtrInput
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrInput
	// The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the identity provider. This key must be a verification only key or certificate (meaning that it only has a public key component).
	KeyId pulumi.StringInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The name of this OpenID Connect identity provider. This is only used for display purposes.
	Name pulumi.StringPtrInput
	// Either urn:oasis:names:tc:SAML:2.0:nameid-format:persistent or urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress depending on which NameId format you wish to use.
	NameIdFormat pulumi.StringPtrInput
	// Set this value equal to true if you wish to use POST bindings with this OpenID Connect identity provider. The default value of false means that a redirect binding which uses a GET request will be used.
	PostRequest pulumi.BoolPtrInput
	// TThe key pair Id to use to sign the SAML request. Required when `signRequest` is true.
	RequestSigningKey pulumi.StringPtrInput
	// When true authentication requests sent to the identity provider will be signed.
	SignRequest pulumi.BoolPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpSamlv2TenantConfigurationArrayInput
	// Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation processing. If this is false, then the `emailClaim` property must be set.
	UseNameForEmail pulumi.BoolPtrInput
	// The XML signature canonicalization method used when digesting and signing the SAML request.
	XmlSignatureCanonicalizationMethod pulumi.StringPtrInput
}

func (FusionAuthIdpSamlv2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpSamlv2Args)(nil)).Elem()
}

type FusionAuthIdpSamlv2Input interface {
	pulumi.Input

	ToFusionAuthIdpSamlv2Output() FusionAuthIdpSamlv2Output
	ToFusionAuthIdpSamlv2OutputWithContext(ctx context.Context) FusionAuthIdpSamlv2Output
}

func (*FusionAuthIdpSamlv2) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpSamlv2)(nil)).Elem()
}

func (i *FusionAuthIdpSamlv2) ToFusionAuthIdpSamlv2Output() FusionAuthIdpSamlv2Output {
	return i.ToFusionAuthIdpSamlv2OutputWithContext(context.Background())
}

func (i *FusionAuthIdpSamlv2) ToFusionAuthIdpSamlv2OutputWithContext(ctx context.Context) FusionAuthIdpSamlv2Output {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpSamlv2Output)
}

// FusionAuthIdpSamlv2ArrayInput is an input type that accepts FusionAuthIdpSamlv2Array and FusionAuthIdpSamlv2ArrayOutput values.
// You can construct a concrete instance of `FusionAuthIdpSamlv2ArrayInput` via:
//
//          FusionAuthIdpSamlv2Array{ FusionAuthIdpSamlv2Args{...} }
type FusionAuthIdpSamlv2ArrayInput interface {
	pulumi.Input

	ToFusionAuthIdpSamlv2ArrayOutput() FusionAuthIdpSamlv2ArrayOutput
	ToFusionAuthIdpSamlv2ArrayOutputWithContext(context.Context) FusionAuthIdpSamlv2ArrayOutput
}

type FusionAuthIdpSamlv2Array []FusionAuthIdpSamlv2Input

func (FusionAuthIdpSamlv2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpSamlv2)(nil)).Elem()
}

func (i FusionAuthIdpSamlv2Array) ToFusionAuthIdpSamlv2ArrayOutput() FusionAuthIdpSamlv2ArrayOutput {
	return i.ToFusionAuthIdpSamlv2ArrayOutputWithContext(context.Background())
}

func (i FusionAuthIdpSamlv2Array) ToFusionAuthIdpSamlv2ArrayOutputWithContext(ctx context.Context) FusionAuthIdpSamlv2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpSamlv2ArrayOutput)
}

// FusionAuthIdpSamlv2MapInput is an input type that accepts FusionAuthIdpSamlv2Map and FusionAuthIdpSamlv2MapOutput values.
// You can construct a concrete instance of `FusionAuthIdpSamlv2MapInput` via:
//
//          FusionAuthIdpSamlv2Map{ "key": FusionAuthIdpSamlv2Args{...} }
type FusionAuthIdpSamlv2MapInput interface {
	pulumi.Input

	ToFusionAuthIdpSamlv2MapOutput() FusionAuthIdpSamlv2MapOutput
	ToFusionAuthIdpSamlv2MapOutputWithContext(context.Context) FusionAuthIdpSamlv2MapOutput
}

type FusionAuthIdpSamlv2Map map[string]FusionAuthIdpSamlv2Input

func (FusionAuthIdpSamlv2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpSamlv2)(nil)).Elem()
}

func (i FusionAuthIdpSamlv2Map) ToFusionAuthIdpSamlv2MapOutput() FusionAuthIdpSamlv2MapOutput {
	return i.ToFusionAuthIdpSamlv2MapOutputWithContext(context.Background())
}

func (i FusionAuthIdpSamlv2Map) ToFusionAuthIdpSamlv2MapOutputWithContext(ctx context.Context) FusionAuthIdpSamlv2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpSamlv2MapOutput)
}

type FusionAuthIdpSamlv2Output struct{ *pulumi.OutputState }

func (FusionAuthIdpSamlv2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpSamlv2)(nil)).Elem()
}

func (o FusionAuthIdpSamlv2Output) ToFusionAuthIdpSamlv2Output() FusionAuthIdpSamlv2Output {
	return o
}

func (o FusionAuthIdpSamlv2Output) ToFusionAuthIdpSamlv2OutputWithContext(ctx context.Context) FusionAuthIdpSamlv2Output {
	return o
}

// The configuration for each Application that the identity provider is enabled for.
func (o FusionAuthIdpSamlv2Output) ApplicationConfigurations() FusionAuthIdpSamlv2ApplicationConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) FusionAuthIdpSamlv2ApplicationConfigurationArrayOutput {
		return v.ApplicationConfigurations
	}).(FusionAuthIdpSamlv2ApplicationConfigurationArrayOutput)
}

// The top-level button image (URL) to use on the FusionAuth login page for this Identity Provider.
func (o FusionAuthIdpSamlv2Output) ButtonImageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.StringPtrOutput { return v.ButtonImageUrl }).(pulumi.StringPtrOutput)
}

// The top-level button text to use on the FusionAuth login page for this Identity Provider.
func (o FusionAuthIdpSamlv2Output) ButtonText() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.StringOutput { return v.ButtonText }).(pulumi.StringOutput)
}

// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
func (o FusionAuthIdpSamlv2Output) Debug() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.BoolPtrOutput { return v.Debug }).(pulumi.BoolPtrOutput)
}

// This is an optional list of domains that this OpenID Connect provider should be used for. This converts the FusionAuth login form to a domain-based login form. This type of form first asks the user for their email. FusionAuth then uses their email to determine if an OpenID Connect identity provider should be used. If an OpenID Connect provider should be used, the browser is redirected to the authorization endpoint of that identity provider. Otherwise, the password field is revealed on the form so that the user can login using FusionAuth.
func (o FusionAuthIdpSamlv2Output) Domains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.StringArrayOutput { return v.Domains }).(pulumi.StringArrayOutput)
}

// The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely identity the user. If this is not set, the `useNameForEmail` flag must be true.
func (o FusionAuthIdpSamlv2Output) EmailClaim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.StringPtrOutput { return v.EmailClaim }).(pulumi.StringPtrOutput)
}

// Determines if this provider is enabled. If it is false then it will be disabled globally.
func (o FusionAuthIdpSamlv2Output) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The SAML v2 login page of the identity provider.
func (o FusionAuthIdpSamlv2Output) IdpEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.StringPtrOutput { return v.IdpEndpoint }).(pulumi.StringPtrOutput)
}

// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
func (o FusionAuthIdpSamlv2Output) IdpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.StringPtrOutput { return v.IdpId }).(pulumi.StringPtrOutput)
}

// The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the identity provider. This key must be a verification only key or certificate (meaning that it only has a public key component).
func (o FusionAuthIdpSamlv2Output) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
func (o FusionAuthIdpSamlv2Output) LambdaReconcileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.StringPtrOutput { return v.LambdaReconcileId }).(pulumi.StringPtrOutput)
}

// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
func (o FusionAuthIdpSamlv2Output) LinkingStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.StringOutput { return v.LinkingStrategy }).(pulumi.StringOutput)
}

// The name of this OpenID Connect identity provider. This is only used for display purposes.
func (o FusionAuthIdpSamlv2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Either urn:oasis:names:tc:SAML:2.0:nameid-format:persistent or urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress depending on which NameId format you wish to use.
func (o FusionAuthIdpSamlv2Output) NameIdFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.StringOutput { return v.NameIdFormat }).(pulumi.StringOutput)
}

// Set this value equal to true if you wish to use POST bindings with this OpenID Connect identity provider. The default value of false means that a redirect binding which uses a GET request will be used.
func (o FusionAuthIdpSamlv2Output) PostRequest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.BoolPtrOutput { return v.PostRequest }).(pulumi.BoolPtrOutput)
}

// TThe key pair Id to use to sign the SAML request. Required when `signRequest` is true.
func (o FusionAuthIdpSamlv2Output) RequestSigningKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.StringPtrOutput { return v.RequestSigningKey }).(pulumi.StringPtrOutput)
}

// When true authentication requests sent to the identity provider will be signed.
func (o FusionAuthIdpSamlv2Output) SignRequest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.BoolPtrOutput { return v.SignRequest }).(pulumi.BoolPtrOutput)
}

// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
func (o FusionAuthIdpSamlv2Output) TenantConfigurations() FusionAuthIdpSamlv2TenantConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) FusionAuthIdpSamlv2TenantConfigurationArrayOutput {
		return v.TenantConfigurations
	}).(FusionAuthIdpSamlv2TenantConfigurationArrayOutput)
}

// Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation processing. If this is false, then the `emailClaim` property must be set.
func (o FusionAuthIdpSamlv2Output) UseNameForEmail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.BoolPtrOutput { return v.UseNameForEmail }).(pulumi.BoolPtrOutput)
}

// The XML signature canonicalization method used when digesting and signing the SAML request.
func (o FusionAuthIdpSamlv2Output) XmlSignatureCanonicalizationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlv2) pulumi.StringPtrOutput { return v.XmlSignatureCanonicalizationMethod }).(pulumi.StringPtrOutput)
}

type FusionAuthIdpSamlv2ArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpSamlv2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpSamlv2)(nil)).Elem()
}

func (o FusionAuthIdpSamlv2ArrayOutput) ToFusionAuthIdpSamlv2ArrayOutput() FusionAuthIdpSamlv2ArrayOutput {
	return o
}

func (o FusionAuthIdpSamlv2ArrayOutput) ToFusionAuthIdpSamlv2ArrayOutputWithContext(ctx context.Context) FusionAuthIdpSamlv2ArrayOutput {
	return o
}

func (o FusionAuthIdpSamlv2ArrayOutput) Index(i pulumi.IntInput) FusionAuthIdpSamlv2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthIdpSamlv2 {
		return vs[0].([]*FusionAuthIdpSamlv2)[vs[1].(int)]
	}).(FusionAuthIdpSamlv2Output)
}

type FusionAuthIdpSamlv2MapOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpSamlv2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpSamlv2)(nil)).Elem()
}

func (o FusionAuthIdpSamlv2MapOutput) ToFusionAuthIdpSamlv2MapOutput() FusionAuthIdpSamlv2MapOutput {
	return o
}

func (o FusionAuthIdpSamlv2MapOutput) ToFusionAuthIdpSamlv2MapOutputWithContext(ctx context.Context) FusionAuthIdpSamlv2MapOutput {
	return o
}

func (o FusionAuthIdpSamlv2MapOutput) MapIndex(k pulumi.StringInput) FusionAuthIdpSamlv2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthIdpSamlv2 {
		return vs[0].(map[string]*FusionAuthIdpSamlv2)[vs[1].(string)]
	}).(FusionAuthIdpSamlv2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpSamlv2Input)(nil)).Elem(), &FusionAuthIdpSamlv2{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpSamlv2ArrayInput)(nil)).Elem(), FusionAuthIdpSamlv2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpSamlv2MapInput)(nil)).Elem(), FusionAuthIdpSamlv2Map{})
	pulumi.RegisterOutputType(FusionAuthIdpSamlv2Output{})
	pulumi.RegisterOutputType(FusionAuthIdpSamlv2ArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpSamlv2MapOutput{})
}