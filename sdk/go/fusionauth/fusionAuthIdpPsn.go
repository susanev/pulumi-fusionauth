// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Sony Playstation Network Identity Provider Resource
//
// The Sony PlayStation Network identity provider type will use the Sony OAuth v2.0 login API. It will also provide a Login with Sony PlayStation Network button on FusionAuth’s login page that will direct a user to the Sony login page.
//
// This identity provider will call Sony’s API to load the user’s email and onlineId and use those as email and username to lookup or create a user in FusionAuth depending on the linking strategy configured for this identity provider. Additional claims returned by Sony PlayStation Network can be used to reconcile the user to FusionAuth by using a Sony PlayStation Network Reconcile Lambda.
//
// FusionAuth will also store the Sony PlayStation Network accessToken returned from the Sony PlayStation Network API in the link between the user and the identity provider. This token can be used by an application to make further requests to Sony PlayStation Network APIs on behalf of the user.
//
// [Sony PlayStation Network Identity Provider APIs](https://fusionauth.io/docs/v1/tech/apis/identity-providers/sonypsn/)
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
// 		_, err := fusionauth.NewFusionAuthIdpPsn(ctx, "sonyPsn", &fusionauth.FusionAuthIdpPsnArgs{
// 			ApplicationConfigurations: FusionAuthIdpPsnApplicationConfigurationArray{
// 				&FusionAuthIdpPsnApplicationConfigurationArgs{
// 					ApplicationId:      pulumi.Any(fusionauth_application.My_app.Id),
// 					CreateRegistration: pulumi.Bool(true),
// 					Enabled:            pulumi.Bool(true),
// 				},
// 			},
// 			ButtonText:   pulumi.String("Login with Playstation"),
// 			ClientId:     pulumi.String("0eb1ce3c-2fb1-4ae9-b361-d49fc6e764cc"),
// 			ClientSecret: pulumi.String("693s000cbn66k0mxtqzr_c_NfLy3~6_SEA"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FusionAuthIdpPsn struct {
	pulumi.CustomResourceState

	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpPsnApplicationConfigurationArrayOutput `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringOutput `pulumi:"buttonText"`
	// The top-level Sony PlayStation Network client id for your Application. This value is retrieved from the Sony PlayStation Network developer website when you setup your Sony PlayStation Network developer account.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The top-level client secret to use with the Sony PlayStation Network Identity Provider when retrieving the long-lived token. This value is retrieved from the Sony PlayStation Network developer website when you setup your Sony PlayStation Network developer account.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrOutput `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrOutput `pulumi:"idpId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrOutput `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringOutput `pulumi:"linkingStrategy"`
	// The top-level scope that you are requesting from Sony PlayStation Network.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpPsnTenantConfigurationArrayOutput `pulumi:"tenantConfigurations"`
}

// NewFusionAuthIdpPsn registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthIdpPsn(ctx *pulumi.Context,
	name string, args *FusionAuthIdpPsnArgs, opts ...pulumi.ResourceOption) (*FusionAuthIdpPsn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ButtonText == nil {
		return nil, errors.New("invalid value for required argument 'ButtonText'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FusionAuthIdpPsn
	err := ctx.RegisterResource("fusionauth:index/fusionAuthIdpPsn:FusionAuthIdpPsn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthIdpPsn gets an existing FusionAuthIdpPsn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthIdpPsn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthIdpPsnState, opts ...pulumi.ResourceOption) (*FusionAuthIdpPsn, error) {
	var resource FusionAuthIdpPsn
	err := ctx.ReadResource("fusionauth:index/fusionAuthIdpPsn:FusionAuthIdpPsn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthIdpPsn resources.
type fusionAuthIdpPsnState struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpPsnApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText *string `pulumi:"buttonText"`
	// The top-level Sony PlayStation Network client id for your Application. This value is retrieved from the Sony PlayStation Network developer website when you setup your Sony PlayStation Network developer account.
	ClientId *string `pulumi:"clientId"`
	// The top-level client secret to use with the Sony PlayStation Network Identity Provider when retrieving the long-lived token. This value is retrieved from the Sony PlayStation Network developer website when you setup your Sony PlayStation Network developer account.
	ClientSecret *string `pulumi:"clientSecret"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId *string `pulumi:"idpId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The top-level scope that you are requesting from Sony PlayStation Network.
	Scope *string `pulumi:"scope"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpPsnTenantConfiguration `pulumi:"tenantConfigurations"`
}

type FusionAuthIdpPsnState struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpPsnApplicationConfigurationArrayInput
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringPtrInput
	// The top-level Sony PlayStation Network client id for your Application. This value is retrieved from the Sony PlayStation Network developer website when you setup your Sony PlayStation Network developer account.
	ClientId pulumi.StringPtrInput
	// The top-level client secret to use with the Sony PlayStation Network Identity Provider when retrieving the long-lived token. This value is retrieved from the Sony PlayStation Network developer website when you setup your Sony PlayStation Network developer account.
	ClientSecret pulumi.StringPtrInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The top-level scope that you are requesting from Sony PlayStation Network.
	Scope pulumi.StringPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpPsnTenantConfigurationArrayInput
}

func (FusionAuthIdpPsnState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpPsnState)(nil)).Elem()
}

type fusionAuthIdpPsnArgs struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpPsnApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText string `pulumi:"buttonText"`
	// The top-level Sony PlayStation Network client id for your Application. This value is retrieved from the Sony PlayStation Network developer website when you setup your Sony PlayStation Network developer account.
	ClientId string `pulumi:"clientId"`
	// The top-level client secret to use with the Sony PlayStation Network Identity Provider when retrieving the long-lived token. This value is retrieved from the Sony PlayStation Network developer website when you setup your Sony PlayStation Network developer account.
	ClientSecret string `pulumi:"clientSecret"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId *string `pulumi:"idpId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The top-level scope that you are requesting from Sony PlayStation Network.
	Scope *string `pulumi:"scope"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpPsnTenantConfiguration `pulumi:"tenantConfigurations"`
}

// The set of arguments for constructing a FusionAuthIdpPsn resource.
type FusionAuthIdpPsnArgs struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpPsnApplicationConfigurationArrayInput
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringInput
	// The top-level Sony PlayStation Network client id for your Application. This value is retrieved from the Sony PlayStation Network developer website when you setup your Sony PlayStation Network developer account.
	ClientId pulumi.StringInput
	// The top-level client secret to use with the Sony PlayStation Network Identity Provider when retrieving the long-lived token. This value is retrieved from the Sony PlayStation Network developer website when you setup your Sony PlayStation Network developer account.
	ClientSecret pulumi.StringInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The top-level scope that you are requesting from Sony PlayStation Network.
	Scope pulumi.StringPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpPsnTenantConfigurationArrayInput
}

func (FusionAuthIdpPsnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpPsnArgs)(nil)).Elem()
}

type FusionAuthIdpPsnInput interface {
	pulumi.Input

	ToFusionAuthIdpPsnOutput() FusionAuthIdpPsnOutput
	ToFusionAuthIdpPsnOutputWithContext(ctx context.Context) FusionAuthIdpPsnOutput
}

func (*FusionAuthIdpPsn) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpPsn)(nil)).Elem()
}

func (i *FusionAuthIdpPsn) ToFusionAuthIdpPsnOutput() FusionAuthIdpPsnOutput {
	return i.ToFusionAuthIdpPsnOutputWithContext(context.Background())
}

func (i *FusionAuthIdpPsn) ToFusionAuthIdpPsnOutputWithContext(ctx context.Context) FusionAuthIdpPsnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpPsnOutput)
}

// FusionAuthIdpPsnArrayInput is an input type that accepts FusionAuthIdpPsnArray and FusionAuthIdpPsnArrayOutput values.
// You can construct a concrete instance of `FusionAuthIdpPsnArrayInput` via:
//
//          FusionAuthIdpPsnArray{ FusionAuthIdpPsnArgs{...} }
type FusionAuthIdpPsnArrayInput interface {
	pulumi.Input

	ToFusionAuthIdpPsnArrayOutput() FusionAuthIdpPsnArrayOutput
	ToFusionAuthIdpPsnArrayOutputWithContext(context.Context) FusionAuthIdpPsnArrayOutput
}

type FusionAuthIdpPsnArray []FusionAuthIdpPsnInput

func (FusionAuthIdpPsnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpPsn)(nil)).Elem()
}

func (i FusionAuthIdpPsnArray) ToFusionAuthIdpPsnArrayOutput() FusionAuthIdpPsnArrayOutput {
	return i.ToFusionAuthIdpPsnArrayOutputWithContext(context.Background())
}

func (i FusionAuthIdpPsnArray) ToFusionAuthIdpPsnArrayOutputWithContext(ctx context.Context) FusionAuthIdpPsnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpPsnArrayOutput)
}

// FusionAuthIdpPsnMapInput is an input type that accepts FusionAuthIdpPsnMap and FusionAuthIdpPsnMapOutput values.
// You can construct a concrete instance of `FusionAuthIdpPsnMapInput` via:
//
//          FusionAuthIdpPsnMap{ "key": FusionAuthIdpPsnArgs{...} }
type FusionAuthIdpPsnMapInput interface {
	pulumi.Input

	ToFusionAuthIdpPsnMapOutput() FusionAuthIdpPsnMapOutput
	ToFusionAuthIdpPsnMapOutputWithContext(context.Context) FusionAuthIdpPsnMapOutput
}

type FusionAuthIdpPsnMap map[string]FusionAuthIdpPsnInput

func (FusionAuthIdpPsnMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpPsn)(nil)).Elem()
}

func (i FusionAuthIdpPsnMap) ToFusionAuthIdpPsnMapOutput() FusionAuthIdpPsnMapOutput {
	return i.ToFusionAuthIdpPsnMapOutputWithContext(context.Background())
}

func (i FusionAuthIdpPsnMap) ToFusionAuthIdpPsnMapOutputWithContext(ctx context.Context) FusionAuthIdpPsnMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpPsnMapOutput)
}

type FusionAuthIdpPsnOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpPsnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpPsn)(nil)).Elem()
}

func (o FusionAuthIdpPsnOutput) ToFusionAuthIdpPsnOutput() FusionAuthIdpPsnOutput {
	return o
}

func (o FusionAuthIdpPsnOutput) ToFusionAuthIdpPsnOutputWithContext(ctx context.Context) FusionAuthIdpPsnOutput {
	return o
}

// The configuration for each Application that the identity provider is enabled for.
func (o FusionAuthIdpPsnOutput) ApplicationConfigurations() FusionAuthIdpPsnApplicationConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpPsn) FusionAuthIdpPsnApplicationConfigurationArrayOutput {
		return v.ApplicationConfigurations
	}).(FusionAuthIdpPsnApplicationConfigurationArrayOutput)
}

// The top-level button text to use on the FusionAuth login page for this Identity Provider.
func (o FusionAuthIdpPsnOutput) ButtonText() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpPsn) pulumi.StringOutput { return v.ButtonText }).(pulumi.StringOutput)
}

// The top-level Sony PlayStation Network client id for your Application. This value is retrieved from the Sony PlayStation Network developer website when you setup your Sony PlayStation Network developer account.
func (o FusionAuthIdpPsnOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpPsn) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The top-level client secret to use with the Sony PlayStation Network Identity Provider when retrieving the long-lived token. This value is retrieved from the Sony PlayStation Network developer website when you setup your Sony PlayStation Network developer account.
func (o FusionAuthIdpPsnOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpPsn) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
func (o FusionAuthIdpPsnOutput) Debug() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpPsn) pulumi.BoolPtrOutput { return v.Debug }).(pulumi.BoolPtrOutput)
}

// Determines if this provider is enabled. If it is false then it will be disabled globally.
func (o FusionAuthIdpPsnOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpPsn) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
func (o FusionAuthIdpPsnOutput) IdpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpPsn) pulumi.StringPtrOutput { return v.IdpId }).(pulumi.StringPtrOutput)
}

// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
func (o FusionAuthIdpPsnOutput) LambdaReconcileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpPsn) pulumi.StringPtrOutput { return v.LambdaReconcileId }).(pulumi.StringPtrOutput)
}

// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
func (o FusionAuthIdpPsnOutput) LinkingStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpPsn) pulumi.StringOutput { return v.LinkingStrategy }).(pulumi.StringOutput)
}

// The top-level scope that you are requesting from Sony PlayStation Network.
func (o FusionAuthIdpPsnOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpPsn) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
func (o FusionAuthIdpPsnOutput) TenantConfigurations() FusionAuthIdpPsnTenantConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpPsn) FusionAuthIdpPsnTenantConfigurationArrayOutput {
		return v.TenantConfigurations
	}).(FusionAuthIdpPsnTenantConfigurationArrayOutput)
}

type FusionAuthIdpPsnArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpPsnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpPsn)(nil)).Elem()
}

func (o FusionAuthIdpPsnArrayOutput) ToFusionAuthIdpPsnArrayOutput() FusionAuthIdpPsnArrayOutput {
	return o
}

func (o FusionAuthIdpPsnArrayOutput) ToFusionAuthIdpPsnArrayOutputWithContext(ctx context.Context) FusionAuthIdpPsnArrayOutput {
	return o
}

func (o FusionAuthIdpPsnArrayOutput) Index(i pulumi.IntInput) FusionAuthIdpPsnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthIdpPsn {
		return vs[0].([]*FusionAuthIdpPsn)[vs[1].(int)]
	}).(FusionAuthIdpPsnOutput)
}

type FusionAuthIdpPsnMapOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpPsnMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpPsn)(nil)).Elem()
}

func (o FusionAuthIdpPsnMapOutput) ToFusionAuthIdpPsnMapOutput() FusionAuthIdpPsnMapOutput {
	return o
}

func (o FusionAuthIdpPsnMapOutput) ToFusionAuthIdpPsnMapOutputWithContext(ctx context.Context) FusionAuthIdpPsnMapOutput {
	return o
}

func (o FusionAuthIdpPsnMapOutput) MapIndex(k pulumi.StringInput) FusionAuthIdpPsnOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthIdpPsn {
		return vs[0].(map[string]*FusionAuthIdpPsn)[vs[1].(string)]
	}).(FusionAuthIdpPsnOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpPsnInput)(nil)).Elem(), &FusionAuthIdpPsn{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpPsnArrayInput)(nil)).Elem(), FusionAuthIdpPsnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpPsnMapInput)(nil)).Elem(), FusionAuthIdpPsnMap{})
	pulumi.RegisterOutputType(FusionAuthIdpPsnOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpPsnArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpPsnMapOutput{})
}
