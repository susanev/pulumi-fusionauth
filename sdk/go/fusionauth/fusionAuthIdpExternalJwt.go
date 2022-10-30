// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # External JWT Identity Provider Resource
//
// This is a special type of identity provider that is only used via the JWT Reconcile API. This identity provider defines the claims inside the incoming JWT and how they map to fields in the FusionAuth User object.
//
// In order for this identity provider to use the JWT, it also needs the public key or HMAC secret that the JWT was signed with. FusionAuth will verify that the JWT is valid and has not expired. Once the JWT has been validated, FusionAuth will reconcile it to ensure that the User exists and is up-to-date.
//
// [External JWT Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/external-jwt/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/theogravity/pulumi-fusionauth/sdk/v2/go/fusionauth"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fusionauth.NewFusionAuthIdpExternalJwt(ctx, "jwt", &fusionauth.FusionAuthIdpExternalJwtArgs{
// 			ClaimMap: pulumi.AnyMap{
// 				"dept":       pulumi.Any("RegistrationData"),
// 				"first_name": pulumi.Any("firstName"),
// 				"last_name":  pulumi.Any("lastName"),
// 			},
// 			Debug:                       pulumi.Bool(false),
// 			Enabled:                     pulumi.Bool(true),
// 			HeaderKeyParameter:          pulumi.String("kid"),
// 			Oauth2AuthorizationEndpoint: pulumi.String("https://acme.com/adfs/oauth2/authorize?client_id=cf3b00da-9551-460a-ad18-33232e6cbff0&response_type=code&redirect_uri=https://acme.com/oauth2/redirect"),
// 			Oauth2TokenEndpoint:         pulumi.String("https://acme.com/adfs/oauth2/token"),
// 			UniqueIdentityClaim:         pulumi.String("email"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FusionAuthIdpExternalJwt struct {
	pulumi.CustomResourceState

	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpExternalJwtApplicationConfigurationArrayOutput `pulumi:"applicationConfigurations"`
	// A map of incoming claims to User fields, User data or Registration data. The key of the map is the incoming claim name
	// from the configured identity provider.
	ClaimMap pulumi.MapOutput `pulumi:"claimMap"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrOutput `pulumi:"debug"`
	// An array of domains that are managed by this Identity Provider.
	Domains pulumi.StringArrayOutput `pulumi:"domains"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The name header claim that identifies the public key used to verify the signature. In most cases this be kid or x5t.
	HeaderKeyParameter pulumi.StringOutput `pulumi:"headerKeyParameter"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrOutput `pulumi:"idpId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrOutput `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringOutput `pulumi:"linkingStrategy"`
	// The name of the Identity Provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// The authorization endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to perform the browser redirect to the OAuth2 authorize endpoint.
	Oauth2AuthorizationEndpoint pulumi.StringPtrOutput `pulumi:"oauth2AuthorizationEndpoint"`
	// TThe token endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to complete the OAuth2 grant workflow.
	Oauth2TokenEndpoint pulumi.StringPtrOutput `pulumi:"oauth2TokenEndpoint"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpExternalJwtTenantConfigurationArrayOutput `pulumi:"tenantConfigurations"`
	// The name of the claim that represents the unique identify of the User. This will generally be email or the name of the claim that provides the email address.
	UniqueIdentityClaim pulumi.StringOutput `pulumi:"uniqueIdentityClaim"`
}

// NewFusionAuthIdpExternalJwt registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthIdpExternalJwt(ctx *pulumi.Context,
	name string, args *FusionAuthIdpExternalJwtArgs, opts ...pulumi.ResourceOption) (*FusionAuthIdpExternalJwt, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HeaderKeyParameter == nil {
		return nil, errors.New("invalid value for required argument 'HeaderKeyParameter'")
	}
	if args.UniqueIdentityClaim == nil {
		return nil, errors.New("invalid value for required argument 'UniqueIdentityClaim'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FusionAuthIdpExternalJwt
	err := ctx.RegisterResource("fusionauth:index/fusionAuthIdpExternalJwt:FusionAuthIdpExternalJwt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthIdpExternalJwt gets an existing FusionAuthIdpExternalJwt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthIdpExternalJwt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthIdpExternalJwtState, opts ...pulumi.ResourceOption) (*FusionAuthIdpExternalJwt, error) {
	var resource FusionAuthIdpExternalJwt
	err := ctx.ReadResource("fusionauth:index/fusionAuthIdpExternalJwt:FusionAuthIdpExternalJwt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthIdpExternalJwt resources.
type fusionAuthIdpExternalJwtState struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpExternalJwtApplicationConfiguration `pulumi:"applicationConfigurations"`
	// A map of incoming claims to User fields, User data or Registration data. The key of the map is the incoming claim name
	// from the configured identity provider.
	ClaimMap map[string]interface{} `pulumi:"claimMap"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// An array of domains that are managed by this Identity Provider.
	Domains []string `pulumi:"domains"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The name header claim that identifies the public key used to verify the signature. In most cases this be kid or x5t.
	HeaderKeyParameter *string `pulumi:"headerKeyParameter"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId *string `pulumi:"idpId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The name of the Identity Provider.
	Name *string `pulumi:"name"`
	// The authorization endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to perform the browser redirect to the OAuth2 authorize endpoint.
	Oauth2AuthorizationEndpoint *string `pulumi:"oauth2AuthorizationEndpoint"`
	// TThe token endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to complete the OAuth2 grant workflow.
	Oauth2TokenEndpoint *string `pulumi:"oauth2TokenEndpoint"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpExternalJwtTenantConfiguration `pulumi:"tenantConfigurations"`
	// The name of the claim that represents the unique identify of the User. This will generally be email or the name of the claim that provides the email address.
	UniqueIdentityClaim *string `pulumi:"uniqueIdentityClaim"`
}

type FusionAuthIdpExternalJwtState struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpExternalJwtApplicationConfigurationArrayInput
	// A map of incoming claims to User fields, User data or Registration data. The key of the map is the incoming claim name
	// from the configured identity provider.
	ClaimMap pulumi.MapInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// An array of domains that are managed by this Identity Provider.
	Domains pulumi.StringArrayInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The name header claim that identifies the public key used to verify the signature. In most cases this be kid or x5t.
	HeaderKeyParameter pulumi.StringPtrInput
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The name of the Identity Provider.
	Name pulumi.StringPtrInput
	// The authorization endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to perform the browser redirect to the OAuth2 authorize endpoint.
	Oauth2AuthorizationEndpoint pulumi.StringPtrInput
	// TThe token endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to complete the OAuth2 grant workflow.
	Oauth2TokenEndpoint pulumi.StringPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpExternalJwtTenantConfigurationArrayInput
	// The name of the claim that represents the unique identify of the User. This will generally be email or the name of the claim that provides the email address.
	UniqueIdentityClaim pulumi.StringPtrInput
}

func (FusionAuthIdpExternalJwtState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpExternalJwtState)(nil)).Elem()
}

type fusionAuthIdpExternalJwtArgs struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpExternalJwtApplicationConfiguration `pulumi:"applicationConfigurations"`
	// A map of incoming claims to User fields, User data or Registration data. The key of the map is the incoming claim name
	// from the configured identity provider.
	ClaimMap map[string]interface{} `pulumi:"claimMap"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// An array of domains that are managed by this Identity Provider.
	Domains []string `pulumi:"domains"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The name header claim that identifies the public key used to verify the signature. In most cases this be kid or x5t.
	HeaderKeyParameter string `pulumi:"headerKeyParameter"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId *string `pulumi:"idpId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The name of the Identity Provider.
	Name *string `pulumi:"name"`
	// The authorization endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to perform the browser redirect to the OAuth2 authorize endpoint.
	Oauth2AuthorizationEndpoint *string `pulumi:"oauth2AuthorizationEndpoint"`
	// TThe token endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to complete the OAuth2 grant workflow.
	Oauth2TokenEndpoint *string `pulumi:"oauth2TokenEndpoint"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpExternalJwtTenantConfiguration `pulumi:"tenantConfigurations"`
	// The name of the claim that represents the unique identify of the User. This will generally be email or the name of the claim that provides the email address.
	UniqueIdentityClaim string `pulumi:"uniqueIdentityClaim"`
}

// The set of arguments for constructing a FusionAuthIdpExternalJwt resource.
type FusionAuthIdpExternalJwtArgs struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpExternalJwtApplicationConfigurationArrayInput
	// A map of incoming claims to User fields, User data or Registration data. The key of the map is the incoming claim name
	// from the configured identity provider.
	ClaimMap pulumi.MapInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// An array of domains that are managed by this Identity Provider.
	Domains pulumi.StringArrayInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The name header claim that identifies the public key used to verify the signature. In most cases this be kid or x5t.
	HeaderKeyParameter pulumi.StringInput
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The name of the Identity Provider.
	Name pulumi.StringPtrInput
	// The authorization endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to perform the browser redirect to the OAuth2 authorize endpoint.
	Oauth2AuthorizationEndpoint pulumi.StringPtrInput
	// TThe token endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to complete the OAuth2 grant workflow.
	Oauth2TokenEndpoint pulumi.StringPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpExternalJwtTenantConfigurationArrayInput
	// The name of the claim that represents the unique identify of the User. This will generally be email or the name of the claim that provides the email address.
	UniqueIdentityClaim pulumi.StringInput
}

func (FusionAuthIdpExternalJwtArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpExternalJwtArgs)(nil)).Elem()
}

type FusionAuthIdpExternalJwtInput interface {
	pulumi.Input

	ToFusionAuthIdpExternalJwtOutput() FusionAuthIdpExternalJwtOutput
	ToFusionAuthIdpExternalJwtOutputWithContext(ctx context.Context) FusionAuthIdpExternalJwtOutput
}

func (*FusionAuthIdpExternalJwt) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpExternalJwt)(nil)).Elem()
}

func (i *FusionAuthIdpExternalJwt) ToFusionAuthIdpExternalJwtOutput() FusionAuthIdpExternalJwtOutput {
	return i.ToFusionAuthIdpExternalJwtOutputWithContext(context.Background())
}

func (i *FusionAuthIdpExternalJwt) ToFusionAuthIdpExternalJwtOutputWithContext(ctx context.Context) FusionAuthIdpExternalJwtOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpExternalJwtOutput)
}

// FusionAuthIdpExternalJwtArrayInput is an input type that accepts FusionAuthIdpExternalJwtArray and FusionAuthIdpExternalJwtArrayOutput values.
// You can construct a concrete instance of `FusionAuthIdpExternalJwtArrayInput` via:
//
//          FusionAuthIdpExternalJwtArray{ FusionAuthIdpExternalJwtArgs{...} }
type FusionAuthIdpExternalJwtArrayInput interface {
	pulumi.Input

	ToFusionAuthIdpExternalJwtArrayOutput() FusionAuthIdpExternalJwtArrayOutput
	ToFusionAuthIdpExternalJwtArrayOutputWithContext(context.Context) FusionAuthIdpExternalJwtArrayOutput
}

type FusionAuthIdpExternalJwtArray []FusionAuthIdpExternalJwtInput

func (FusionAuthIdpExternalJwtArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpExternalJwt)(nil)).Elem()
}

func (i FusionAuthIdpExternalJwtArray) ToFusionAuthIdpExternalJwtArrayOutput() FusionAuthIdpExternalJwtArrayOutput {
	return i.ToFusionAuthIdpExternalJwtArrayOutputWithContext(context.Background())
}

func (i FusionAuthIdpExternalJwtArray) ToFusionAuthIdpExternalJwtArrayOutputWithContext(ctx context.Context) FusionAuthIdpExternalJwtArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpExternalJwtArrayOutput)
}

// FusionAuthIdpExternalJwtMapInput is an input type that accepts FusionAuthIdpExternalJwtMap and FusionAuthIdpExternalJwtMapOutput values.
// You can construct a concrete instance of `FusionAuthIdpExternalJwtMapInput` via:
//
//          FusionAuthIdpExternalJwtMap{ "key": FusionAuthIdpExternalJwtArgs{...} }
type FusionAuthIdpExternalJwtMapInput interface {
	pulumi.Input

	ToFusionAuthIdpExternalJwtMapOutput() FusionAuthIdpExternalJwtMapOutput
	ToFusionAuthIdpExternalJwtMapOutputWithContext(context.Context) FusionAuthIdpExternalJwtMapOutput
}

type FusionAuthIdpExternalJwtMap map[string]FusionAuthIdpExternalJwtInput

func (FusionAuthIdpExternalJwtMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpExternalJwt)(nil)).Elem()
}

func (i FusionAuthIdpExternalJwtMap) ToFusionAuthIdpExternalJwtMapOutput() FusionAuthIdpExternalJwtMapOutput {
	return i.ToFusionAuthIdpExternalJwtMapOutputWithContext(context.Background())
}

func (i FusionAuthIdpExternalJwtMap) ToFusionAuthIdpExternalJwtMapOutputWithContext(ctx context.Context) FusionAuthIdpExternalJwtMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpExternalJwtMapOutput)
}

type FusionAuthIdpExternalJwtOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpExternalJwtOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpExternalJwt)(nil)).Elem()
}

func (o FusionAuthIdpExternalJwtOutput) ToFusionAuthIdpExternalJwtOutput() FusionAuthIdpExternalJwtOutput {
	return o
}

func (o FusionAuthIdpExternalJwtOutput) ToFusionAuthIdpExternalJwtOutputWithContext(ctx context.Context) FusionAuthIdpExternalJwtOutput {
	return o
}

// The configuration for each Application that the identity provider is enabled for.
func (o FusionAuthIdpExternalJwtOutput) ApplicationConfigurations() FusionAuthIdpExternalJwtApplicationConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) FusionAuthIdpExternalJwtApplicationConfigurationArrayOutput {
		return v.ApplicationConfigurations
	}).(FusionAuthIdpExternalJwtApplicationConfigurationArrayOutput)
}

// A map of incoming claims to User fields, User data or Registration data. The key of the map is the incoming claim name
// from the configured identity provider.
func (o FusionAuthIdpExternalJwtOutput) ClaimMap() pulumi.MapOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) pulumi.MapOutput { return v.ClaimMap }).(pulumi.MapOutput)
}

// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
func (o FusionAuthIdpExternalJwtOutput) Debug() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) pulumi.BoolPtrOutput { return v.Debug }).(pulumi.BoolPtrOutput)
}

// An array of domains that are managed by this Identity Provider.
func (o FusionAuthIdpExternalJwtOutput) Domains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) pulumi.StringArrayOutput { return v.Domains }).(pulumi.StringArrayOutput)
}

// Determines if this provider is enabled. If it is false then it will be disabled globally.
func (o FusionAuthIdpExternalJwtOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The name header claim that identifies the public key used to verify the signature. In most cases this be kid or x5t.
func (o FusionAuthIdpExternalJwtOutput) HeaderKeyParameter() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) pulumi.StringOutput { return v.HeaderKeyParameter }).(pulumi.StringOutput)
}

// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
func (o FusionAuthIdpExternalJwtOutput) IdpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) pulumi.StringPtrOutput { return v.IdpId }).(pulumi.StringPtrOutput)
}

// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
func (o FusionAuthIdpExternalJwtOutput) LambdaReconcileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) pulumi.StringPtrOutput { return v.LambdaReconcileId }).(pulumi.StringPtrOutput)
}

// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
func (o FusionAuthIdpExternalJwtOutput) LinkingStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) pulumi.StringOutput { return v.LinkingStrategy }).(pulumi.StringOutput)
}

// The name of the Identity Provider.
func (o FusionAuthIdpExternalJwtOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The authorization endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to perform the browser redirect to the OAuth2 authorize endpoint.
func (o FusionAuthIdpExternalJwtOutput) Oauth2AuthorizationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) pulumi.StringPtrOutput { return v.Oauth2AuthorizationEndpoint }).(pulumi.StringPtrOutput)
}

// TThe token endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to complete the OAuth2 grant workflow.
func (o FusionAuthIdpExternalJwtOutput) Oauth2TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) pulumi.StringPtrOutput { return v.Oauth2TokenEndpoint }).(pulumi.StringPtrOutput)
}

// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
func (o FusionAuthIdpExternalJwtOutput) TenantConfigurations() FusionAuthIdpExternalJwtTenantConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) FusionAuthIdpExternalJwtTenantConfigurationArrayOutput {
		return v.TenantConfigurations
	}).(FusionAuthIdpExternalJwtTenantConfigurationArrayOutput)
}

// The name of the claim that represents the unique identify of the User. This will generally be email or the name of the claim that provides the email address.
func (o FusionAuthIdpExternalJwtOutput) UniqueIdentityClaim() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpExternalJwt) pulumi.StringOutput { return v.UniqueIdentityClaim }).(pulumi.StringOutput)
}

type FusionAuthIdpExternalJwtArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpExternalJwtArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpExternalJwt)(nil)).Elem()
}

func (o FusionAuthIdpExternalJwtArrayOutput) ToFusionAuthIdpExternalJwtArrayOutput() FusionAuthIdpExternalJwtArrayOutput {
	return o
}

func (o FusionAuthIdpExternalJwtArrayOutput) ToFusionAuthIdpExternalJwtArrayOutputWithContext(ctx context.Context) FusionAuthIdpExternalJwtArrayOutput {
	return o
}

func (o FusionAuthIdpExternalJwtArrayOutput) Index(i pulumi.IntInput) FusionAuthIdpExternalJwtOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthIdpExternalJwt {
		return vs[0].([]*FusionAuthIdpExternalJwt)[vs[1].(int)]
	}).(FusionAuthIdpExternalJwtOutput)
}

type FusionAuthIdpExternalJwtMapOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpExternalJwtMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpExternalJwt)(nil)).Elem()
}

func (o FusionAuthIdpExternalJwtMapOutput) ToFusionAuthIdpExternalJwtMapOutput() FusionAuthIdpExternalJwtMapOutput {
	return o
}

func (o FusionAuthIdpExternalJwtMapOutput) ToFusionAuthIdpExternalJwtMapOutputWithContext(ctx context.Context) FusionAuthIdpExternalJwtMapOutput {
	return o
}

func (o FusionAuthIdpExternalJwtMapOutput) MapIndex(k pulumi.StringInput) FusionAuthIdpExternalJwtOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthIdpExternalJwt {
		return vs[0].(map[string]*FusionAuthIdpExternalJwt)[vs[1].(string)]
	}).(FusionAuthIdpExternalJwtOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpExternalJwtInput)(nil)).Elem(), &FusionAuthIdpExternalJwt{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpExternalJwtArrayInput)(nil)).Elem(), FusionAuthIdpExternalJwtArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpExternalJwtMapInput)(nil)).Elem(), FusionAuthIdpExternalJwtMap{})
	pulumi.RegisterOutputType(FusionAuthIdpExternalJwtOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpExternalJwtArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpExternalJwtMapOutput{})
}
