// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # LinkedIn Identity Provider Resource
//
// The LinkedIn identity provider type will use OAuth 2.0 to authenticate a user with LinkedIn. It will also provide a
// `Login with LinkedIn` button on FusionAuth’s login page that will direct a user to the LinkedIn login page.
// Additionally, after successful user authentication, this identity provider will call LinkedIn’s `/v2/me` and
// `/v2/emailAddress` APIs to load additional details about the user and store them in FusionAuth.
//
// The email address returned by the LinkedIn `/v2/emailAddress` API will be used to create or look up the existing user.
// Additional claims returned by LinkedIn can be used to reconcile the User to FusionAuth by using a LinkedIn Reconcile
// lambda. Unless you assign a reconcile lambda to this provider, only the email address will be used from the available
// claims returned by LinkedIn.
//
// FusionAuth will also store the LinkedIn `accessToken` returned from the login endpoint in the `identityProviderLink`
// object. This object is accessible using the Link API.
//
// The `identityProviderLink` object stores the token so that you can use it in your application to call LinkedIn APIs on
// behalf of the user if desired.
//
// [LinkedIn Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/linkedin)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fusionauth/sdk/v2/go/fusionauth"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/theogravity/pulumi-fusionauth/sdk/v2/go/fusionauth"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fusionauth.NewFusionAuthIdpLinkedIn(ctx, "linkedin", &fusionauth.FusionAuthIdpLinkedInArgs{
//				ApplicationConfigurations: FusionAuthIdpLinkedInApplicationConfigurationArray{
//					&FusionAuthIdpLinkedInApplicationConfigurationArgs{
//						ApplicationId:      pulumi.Any(fusionauth_application.Myapp.Id),
//						CreateRegistration: pulumi.Bool(true),
//						Enabled:            pulumi.Bool(true),
//					},
//				},
//				ButtonText:   pulumi.String("Login with LinkedIn"),
//				Debug:        pulumi.Bool(false),
//				Enabled:      pulumi.Bool(true),
//				ClientId:     pulumi.String("9876543210"),
//				ClientSecret: pulumi.String("716a572f917640698cdb99e9d7e64115"),
//				Scope:        pulumi.String("r_emailaddress r_liteprofile"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FusionAuthIdpLinkedIn struct {
	pulumi.CustomResourceState

	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpLinkedInApplicationConfigurationArrayOutput `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringOutput `pulumi:"buttonText"`
	// The top-level LinkedIn client id for your Application. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The top-level client secret to use with the LinkedIn Identity Provider when retrieving the long-lived token. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// Determines if debug is enabled for this provider. When enabled, an Event Log is created each time this provider is invoked to reconcile a login.
	Debug pulumi.BoolPtrOutput `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringOutput `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the Facebook Identity Provider and the user.
	// The valid values are:
	LinkingStrategy pulumi.StringOutput `pulumi:"linkingStrategy"`
	// The top-level scope that you are requesting from LinkedIn.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpLinkedInTenantConfigurationArrayOutput `pulumi:"tenantConfigurations"`
}

// NewFusionAuthIdpLinkedIn registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthIdpLinkedIn(ctx *pulumi.Context,
	name string, args *FusionAuthIdpLinkedInArgs, opts ...pulumi.ResourceOption) (*FusionAuthIdpLinkedIn, error) {
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
	var resource FusionAuthIdpLinkedIn
	err := ctx.RegisterResource("fusionauth:index/fusionAuthIdpLinkedIn:FusionAuthIdpLinkedIn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthIdpLinkedIn gets an existing FusionAuthIdpLinkedIn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthIdpLinkedIn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthIdpLinkedInState, opts ...pulumi.ResourceOption) (*FusionAuthIdpLinkedIn, error) {
	var resource FusionAuthIdpLinkedIn
	err := ctx.ReadResource("fusionauth:index/fusionAuthIdpLinkedIn:FusionAuthIdpLinkedIn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthIdpLinkedIn resources.
type fusionAuthIdpLinkedInState struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpLinkedInApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText *string `pulumi:"buttonText"`
	// The top-level LinkedIn client id for your Application. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
	ClientId *string `pulumi:"clientId"`
	// The top-level client secret to use with the LinkedIn Identity Provider when retrieving the long-lived token. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
	ClientSecret *string `pulumi:"clientSecret"`
	// Determines if debug is enabled for this provider. When enabled, an Event Log is created each time this provider is invoked to reconcile a login.
	Debug *bool `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the Facebook Identity Provider and the user.
	// The valid values are:
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The top-level scope that you are requesting from LinkedIn.
	Scope *string `pulumi:"scope"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpLinkedInTenantConfiguration `pulumi:"tenantConfigurations"`
}

type FusionAuthIdpLinkedInState struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpLinkedInApplicationConfigurationArrayInput
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringPtrInput
	// The top-level LinkedIn client id for your Application. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
	ClientId pulumi.StringPtrInput
	// The top-level client secret to use with the LinkedIn Identity Provider when retrieving the long-lived token. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
	ClientSecret pulumi.StringPtrInput
	// Determines if debug is enabled for this provider. When enabled, an Event Log is created each time this provider is invoked to reconcile a login.
	Debug pulumi.BoolPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the Facebook Identity Provider and the user.
	// The valid values are:
	LinkingStrategy pulumi.StringPtrInput
	// The top-level scope that you are requesting from LinkedIn.
	Scope pulumi.StringPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpLinkedInTenantConfigurationArrayInput
}

func (FusionAuthIdpLinkedInState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpLinkedInState)(nil)).Elem()
}

type fusionAuthIdpLinkedInArgs struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpLinkedInApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText string `pulumi:"buttonText"`
	// The top-level LinkedIn client id for your Application. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
	ClientId string `pulumi:"clientId"`
	// The top-level client secret to use with the LinkedIn Identity Provider when retrieving the long-lived token. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
	ClientSecret string `pulumi:"clientSecret"`
	// Determines if debug is enabled for this provider. When enabled, an Event Log is created each time this provider is invoked to reconcile a login.
	Debug *bool `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the Facebook Identity Provider and the user.
	// The valid values are:
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The top-level scope that you are requesting from LinkedIn.
	Scope *string `pulumi:"scope"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpLinkedInTenantConfiguration `pulumi:"tenantConfigurations"`
}

// The set of arguments for constructing a FusionAuthIdpLinkedIn resource.
type FusionAuthIdpLinkedInArgs struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpLinkedInApplicationConfigurationArrayInput
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringInput
	// The top-level LinkedIn client id for your Application. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
	ClientId pulumi.StringInput
	// The top-level client secret to use with the LinkedIn Identity Provider when retrieving the long-lived token. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
	ClientSecret pulumi.StringInput
	// Determines if debug is enabled for this provider. When enabled, an Event Log is created each time this provider is invoked to reconcile a login.
	Debug pulumi.BoolPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the Facebook Identity Provider and the user.
	// The valid values are:
	LinkingStrategy pulumi.StringPtrInput
	// The top-level scope that you are requesting from LinkedIn.
	Scope pulumi.StringPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpLinkedInTenantConfigurationArrayInput
}

func (FusionAuthIdpLinkedInArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpLinkedInArgs)(nil)).Elem()
}

type FusionAuthIdpLinkedInInput interface {
	pulumi.Input

	ToFusionAuthIdpLinkedInOutput() FusionAuthIdpLinkedInOutput
	ToFusionAuthIdpLinkedInOutputWithContext(ctx context.Context) FusionAuthIdpLinkedInOutput
}

func (*FusionAuthIdpLinkedIn) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpLinkedIn)(nil)).Elem()
}

func (i *FusionAuthIdpLinkedIn) ToFusionAuthIdpLinkedInOutput() FusionAuthIdpLinkedInOutput {
	return i.ToFusionAuthIdpLinkedInOutputWithContext(context.Background())
}

func (i *FusionAuthIdpLinkedIn) ToFusionAuthIdpLinkedInOutputWithContext(ctx context.Context) FusionAuthIdpLinkedInOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpLinkedInOutput)
}

// FusionAuthIdpLinkedInArrayInput is an input type that accepts FusionAuthIdpLinkedInArray and FusionAuthIdpLinkedInArrayOutput values.
// You can construct a concrete instance of `FusionAuthIdpLinkedInArrayInput` via:
//
//	FusionAuthIdpLinkedInArray{ FusionAuthIdpLinkedInArgs{...} }
type FusionAuthIdpLinkedInArrayInput interface {
	pulumi.Input

	ToFusionAuthIdpLinkedInArrayOutput() FusionAuthIdpLinkedInArrayOutput
	ToFusionAuthIdpLinkedInArrayOutputWithContext(context.Context) FusionAuthIdpLinkedInArrayOutput
}

type FusionAuthIdpLinkedInArray []FusionAuthIdpLinkedInInput

func (FusionAuthIdpLinkedInArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpLinkedIn)(nil)).Elem()
}

func (i FusionAuthIdpLinkedInArray) ToFusionAuthIdpLinkedInArrayOutput() FusionAuthIdpLinkedInArrayOutput {
	return i.ToFusionAuthIdpLinkedInArrayOutputWithContext(context.Background())
}

func (i FusionAuthIdpLinkedInArray) ToFusionAuthIdpLinkedInArrayOutputWithContext(ctx context.Context) FusionAuthIdpLinkedInArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpLinkedInArrayOutput)
}

// FusionAuthIdpLinkedInMapInput is an input type that accepts FusionAuthIdpLinkedInMap and FusionAuthIdpLinkedInMapOutput values.
// You can construct a concrete instance of `FusionAuthIdpLinkedInMapInput` via:
//
//	FusionAuthIdpLinkedInMap{ "key": FusionAuthIdpLinkedInArgs{...} }
type FusionAuthIdpLinkedInMapInput interface {
	pulumi.Input

	ToFusionAuthIdpLinkedInMapOutput() FusionAuthIdpLinkedInMapOutput
	ToFusionAuthIdpLinkedInMapOutputWithContext(context.Context) FusionAuthIdpLinkedInMapOutput
}

type FusionAuthIdpLinkedInMap map[string]FusionAuthIdpLinkedInInput

func (FusionAuthIdpLinkedInMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpLinkedIn)(nil)).Elem()
}

func (i FusionAuthIdpLinkedInMap) ToFusionAuthIdpLinkedInMapOutput() FusionAuthIdpLinkedInMapOutput {
	return i.ToFusionAuthIdpLinkedInMapOutputWithContext(context.Background())
}

func (i FusionAuthIdpLinkedInMap) ToFusionAuthIdpLinkedInMapOutputWithContext(ctx context.Context) FusionAuthIdpLinkedInMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpLinkedInMapOutput)
}

type FusionAuthIdpLinkedInOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpLinkedInOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpLinkedIn)(nil)).Elem()
}

func (o FusionAuthIdpLinkedInOutput) ToFusionAuthIdpLinkedInOutput() FusionAuthIdpLinkedInOutput {
	return o
}

func (o FusionAuthIdpLinkedInOutput) ToFusionAuthIdpLinkedInOutputWithContext(ctx context.Context) FusionAuthIdpLinkedInOutput {
	return o
}

// The configuration for each Application that the identity provider is enabled for.
func (o FusionAuthIdpLinkedInOutput) ApplicationConfigurations() FusionAuthIdpLinkedInApplicationConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpLinkedIn) FusionAuthIdpLinkedInApplicationConfigurationArrayOutput {
		return v.ApplicationConfigurations
	}).(FusionAuthIdpLinkedInApplicationConfigurationArrayOutput)
}

// The top-level button text to use on the FusionAuth login page for this Identity Provider.
func (o FusionAuthIdpLinkedInOutput) ButtonText() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpLinkedIn) pulumi.StringOutput { return v.ButtonText }).(pulumi.StringOutput)
}

// The top-level LinkedIn client id for your Application. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
func (o FusionAuthIdpLinkedInOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpLinkedIn) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The top-level client secret to use with the LinkedIn Identity Provider when retrieving the long-lived token. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
func (o FusionAuthIdpLinkedInOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpLinkedIn) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// Determines if debug is enabled for this provider. When enabled, an Event Log is created each time this provider is invoked to reconcile a login.
func (o FusionAuthIdpLinkedInOutput) Debug() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpLinkedIn) pulumi.BoolPtrOutput { return v.Debug }).(pulumi.BoolPtrOutput)
}

// Determines if this provider is enabled. If it is false then it will be disabled globally.
func (o FusionAuthIdpLinkedInOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpLinkedIn) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
func (o FusionAuthIdpLinkedInOutput) LambdaReconcileId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpLinkedIn) pulumi.StringOutput { return v.LambdaReconcileId }).(pulumi.StringOutput)
}

// The linking strategy to use when creating the link between the Facebook Identity Provider and the user.
// The valid values are:
func (o FusionAuthIdpLinkedInOutput) LinkingStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpLinkedIn) pulumi.StringOutput { return v.LinkingStrategy }).(pulumi.StringOutput)
}

// The top-level scope that you are requesting from LinkedIn.
func (o FusionAuthIdpLinkedInOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpLinkedIn) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
func (o FusionAuthIdpLinkedInOutput) TenantConfigurations() FusionAuthIdpLinkedInTenantConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpLinkedIn) FusionAuthIdpLinkedInTenantConfigurationArrayOutput {
		return v.TenantConfigurations
	}).(FusionAuthIdpLinkedInTenantConfigurationArrayOutput)
}

type FusionAuthIdpLinkedInArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpLinkedInArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpLinkedIn)(nil)).Elem()
}

func (o FusionAuthIdpLinkedInArrayOutput) ToFusionAuthIdpLinkedInArrayOutput() FusionAuthIdpLinkedInArrayOutput {
	return o
}

func (o FusionAuthIdpLinkedInArrayOutput) ToFusionAuthIdpLinkedInArrayOutputWithContext(ctx context.Context) FusionAuthIdpLinkedInArrayOutput {
	return o
}

func (o FusionAuthIdpLinkedInArrayOutput) Index(i pulumi.IntInput) FusionAuthIdpLinkedInOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthIdpLinkedIn {
		return vs[0].([]*FusionAuthIdpLinkedIn)[vs[1].(int)]
	}).(FusionAuthIdpLinkedInOutput)
}

type FusionAuthIdpLinkedInMapOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpLinkedInMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpLinkedIn)(nil)).Elem()
}

func (o FusionAuthIdpLinkedInMapOutput) ToFusionAuthIdpLinkedInMapOutput() FusionAuthIdpLinkedInMapOutput {
	return o
}

func (o FusionAuthIdpLinkedInMapOutput) ToFusionAuthIdpLinkedInMapOutputWithContext(ctx context.Context) FusionAuthIdpLinkedInMapOutput {
	return o
}

func (o FusionAuthIdpLinkedInMapOutput) MapIndex(k pulumi.StringInput) FusionAuthIdpLinkedInOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthIdpLinkedIn {
		return vs[0].(map[string]*FusionAuthIdpLinkedIn)[vs[1].(string)]
	}).(FusionAuthIdpLinkedInOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpLinkedInInput)(nil)).Elem(), &FusionAuthIdpLinkedIn{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpLinkedInArrayInput)(nil)).Elem(), FusionAuthIdpLinkedInArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpLinkedInMapInput)(nil)).Elem(), FusionAuthIdpLinkedInMap{})
	pulumi.RegisterOutputType(FusionAuthIdpLinkedInOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpLinkedInArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpLinkedInMapOutput{})
}