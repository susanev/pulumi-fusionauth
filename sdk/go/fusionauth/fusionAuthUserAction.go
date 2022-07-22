// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # User Action Resource
//
// [User Actions API](https://fusionauth.io/docs/v1/tech/apis/user-actions/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fusionauth.NewFusionAuthUserAction(ctx, "example", &fusionauth.FusionAuthUserActionArgs{
// 			PreventLogin: pulumi.Bool(true),
// 			Temporal:     pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FusionAuthUserAction struct {
	pulumi.CustomResourceState

	// The Id of the Email Template that is used when User Actions are canceled.
	CancelEmailTemplateId pulumi.StringPtrOutput `pulumi:"cancelEmailTemplateId"`
	// The Id of the Email Template that is used when User Actions expired automatically (end).
	EndEmailTemplateId pulumi.StringPtrOutput `pulumi:"endEmailTemplateId"`
	// Whether to include the email information in the JSON that is sent to the Webhook when a user action is taken.
	IncludeEmailInEventJson pulumi.BoolPtrOutput `pulumi:"includeEmailInEventJson"`
	// A mapping of localized names for this User Action Option. The key is the Locale and the value is the name of the User Action Option for that language.
	LocalizedNames pulumi.MapOutput `pulumi:"localizedNames"`
	// The Id of the Email Template that is used when User Actions are modified.
	ModifyEmailTemplateId pulumi.StringPtrOutput `pulumi:"modifyEmailTemplateId"`
	// The name of this User Action Option.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of User Action Options.
	Options FusionAuthUserActionOptionArrayOutput `pulumi:"options"`
	// Whether or not this User Action will prevent user login. When this value is set to true the action must also be marked as a time based action. See `temporal`.
	PreventLogin pulumi.BoolPtrOutput `pulumi:"preventLogin"`
	// Whether or not FusionAuth will send events to any registered Webhooks when this User Action expires.
	SendEndEvent pulumi.BoolPtrOutput `pulumi:"sendEndEvent"`
	// The Id of the Email Template that is used when User Actions are started (created).
	StartEmailTemplateId pulumi.StringPtrOutput `pulumi:"startEmailTemplateId"`
	// Whether or not this User Action is time-based (temporal).
	Temporal pulumi.BoolPtrOutput `pulumi:"temporal"`
	// Whether or not email is enabled for this User Action.
	UserEmailingEnabled pulumi.BoolPtrOutput `pulumi:"userEmailingEnabled"`
	// Whether or not user notifications are enabled for this User Action.
	UserNotificationsEnabled pulumi.BoolPtrOutput `pulumi:"userNotificationsEnabled"`
}

// NewFusionAuthUserAction registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthUserAction(ctx *pulumi.Context,
	name string, args *FusionAuthUserActionArgs, opts ...pulumi.ResourceOption) (*FusionAuthUserAction, error) {
	if args == nil {
		args = &FusionAuthUserActionArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FusionAuthUserAction
	err := ctx.RegisterResource("fusionauth:index/fusionAuthUserAction:FusionAuthUserAction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthUserAction gets an existing FusionAuthUserAction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthUserAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthUserActionState, opts ...pulumi.ResourceOption) (*FusionAuthUserAction, error) {
	var resource FusionAuthUserAction
	err := ctx.ReadResource("fusionauth:index/fusionAuthUserAction:FusionAuthUserAction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthUserAction resources.
type fusionAuthUserActionState struct {
	// The Id of the Email Template that is used when User Actions are canceled.
	CancelEmailTemplateId *string `pulumi:"cancelEmailTemplateId"`
	// The Id of the Email Template that is used when User Actions expired automatically (end).
	EndEmailTemplateId *string `pulumi:"endEmailTemplateId"`
	// Whether to include the email information in the JSON that is sent to the Webhook when a user action is taken.
	IncludeEmailInEventJson *bool `pulumi:"includeEmailInEventJson"`
	// A mapping of localized names for this User Action Option. The key is the Locale and the value is the name of the User Action Option for that language.
	LocalizedNames map[string]interface{} `pulumi:"localizedNames"`
	// The Id of the Email Template that is used when User Actions are modified.
	ModifyEmailTemplateId *string `pulumi:"modifyEmailTemplateId"`
	// The name of this User Action Option.
	Name *string `pulumi:"name"`
	// The list of User Action Options.
	Options []FusionAuthUserActionOption `pulumi:"options"`
	// Whether or not this User Action will prevent user login. When this value is set to true the action must also be marked as a time based action. See `temporal`.
	PreventLogin *bool `pulumi:"preventLogin"`
	// Whether or not FusionAuth will send events to any registered Webhooks when this User Action expires.
	SendEndEvent *bool `pulumi:"sendEndEvent"`
	// The Id of the Email Template that is used when User Actions are started (created).
	StartEmailTemplateId *string `pulumi:"startEmailTemplateId"`
	// Whether or not this User Action is time-based (temporal).
	Temporal *bool `pulumi:"temporal"`
	// Whether or not email is enabled for this User Action.
	UserEmailingEnabled *bool `pulumi:"userEmailingEnabled"`
	// Whether or not user notifications are enabled for this User Action.
	UserNotificationsEnabled *bool `pulumi:"userNotificationsEnabled"`
}

type FusionAuthUserActionState struct {
	// The Id of the Email Template that is used when User Actions are canceled.
	CancelEmailTemplateId pulumi.StringPtrInput
	// The Id of the Email Template that is used when User Actions expired automatically (end).
	EndEmailTemplateId pulumi.StringPtrInput
	// Whether to include the email information in the JSON that is sent to the Webhook when a user action is taken.
	IncludeEmailInEventJson pulumi.BoolPtrInput
	// A mapping of localized names for this User Action Option. The key is the Locale and the value is the name of the User Action Option for that language.
	LocalizedNames pulumi.MapInput
	// The Id of the Email Template that is used when User Actions are modified.
	ModifyEmailTemplateId pulumi.StringPtrInput
	// The name of this User Action Option.
	Name pulumi.StringPtrInput
	// The list of User Action Options.
	Options FusionAuthUserActionOptionArrayInput
	// Whether or not this User Action will prevent user login. When this value is set to true the action must also be marked as a time based action. See `temporal`.
	PreventLogin pulumi.BoolPtrInput
	// Whether or not FusionAuth will send events to any registered Webhooks when this User Action expires.
	SendEndEvent pulumi.BoolPtrInput
	// The Id of the Email Template that is used when User Actions are started (created).
	StartEmailTemplateId pulumi.StringPtrInput
	// Whether or not this User Action is time-based (temporal).
	Temporal pulumi.BoolPtrInput
	// Whether or not email is enabled for this User Action.
	UserEmailingEnabled pulumi.BoolPtrInput
	// Whether or not user notifications are enabled for this User Action.
	UserNotificationsEnabled pulumi.BoolPtrInput
}

func (FusionAuthUserActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthUserActionState)(nil)).Elem()
}

type fusionAuthUserActionArgs struct {
	// The Id of the Email Template that is used when User Actions are canceled.
	CancelEmailTemplateId *string `pulumi:"cancelEmailTemplateId"`
	// The Id of the Email Template that is used when User Actions expired automatically (end).
	EndEmailTemplateId *string `pulumi:"endEmailTemplateId"`
	// Whether to include the email information in the JSON that is sent to the Webhook when a user action is taken.
	IncludeEmailInEventJson *bool `pulumi:"includeEmailInEventJson"`
	// A mapping of localized names for this User Action Option. The key is the Locale and the value is the name of the User Action Option for that language.
	LocalizedNames map[string]interface{} `pulumi:"localizedNames"`
	// The Id of the Email Template that is used when User Actions are modified.
	ModifyEmailTemplateId *string `pulumi:"modifyEmailTemplateId"`
	// The name of this User Action Option.
	Name *string `pulumi:"name"`
	// The list of User Action Options.
	Options []FusionAuthUserActionOption `pulumi:"options"`
	// Whether or not this User Action will prevent user login. When this value is set to true the action must also be marked as a time based action. See `temporal`.
	PreventLogin *bool `pulumi:"preventLogin"`
	// Whether or not FusionAuth will send events to any registered Webhooks when this User Action expires.
	SendEndEvent *bool `pulumi:"sendEndEvent"`
	// The Id of the Email Template that is used when User Actions are started (created).
	StartEmailTemplateId *string `pulumi:"startEmailTemplateId"`
	// Whether or not this User Action is time-based (temporal).
	Temporal *bool `pulumi:"temporal"`
	// Whether or not email is enabled for this User Action.
	UserEmailingEnabled *bool `pulumi:"userEmailingEnabled"`
	// Whether or not user notifications are enabled for this User Action.
	UserNotificationsEnabled *bool `pulumi:"userNotificationsEnabled"`
}

// The set of arguments for constructing a FusionAuthUserAction resource.
type FusionAuthUserActionArgs struct {
	// The Id of the Email Template that is used when User Actions are canceled.
	CancelEmailTemplateId pulumi.StringPtrInput
	// The Id of the Email Template that is used when User Actions expired automatically (end).
	EndEmailTemplateId pulumi.StringPtrInput
	// Whether to include the email information in the JSON that is sent to the Webhook when a user action is taken.
	IncludeEmailInEventJson pulumi.BoolPtrInput
	// A mapping of localized names for this User Action Option. The key is the Locale and the value is the name of the User Action Option for that language.
	LocalizedNames pulumi.MapInput
	// The Id of the Email Template that is used when User Actions are modified.
	ModifyEmailTemplateId pulumi.StringPtrInput
	// The name of this User Action Option.
	Name pulumi.StringPtrInput
	// The list of User Action Options.
	Options FusionAuthUserActionOptionArrayInput
	// Whether or not this User Action will prevent user login. When this value is set to true the action must also be marked as a time based action. See `temporal`.
	PreventLogin pulumi.BoolPtrInput
	// Whether or not FusionAuth will send events to any registered Webhooks when this User Action expires.
	SendEndEvent pulumi.BoolPtrInput
	// The Id of the Email Template that is used when User Actions are started (created).
	StartEmailTemplateId pulumi.StringPtrInput
	// Whether or not this User Action is time-based (temporal).
	Temporal pulumi.BoolPtrInput
	// Whether or not email is enabled for this User Action.
	UserEmailingEnabled pulumi.BoolPtrInput
	// Whether or not user notifications are enabled for this User Action.
	UserNotificationsEnabled pulumi.BoolPtrInput
}

func (FusionAuthUserActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthUserActionArgs)(nil)).Elem()
}

type FusionAuthUserActionInput interface {
	pulumi.Input

	ToFusionAuthUserActionOutput() FusionAuthUserActionOutput
	ToFusionAuthUserActionOutputWithContext(ctx context.Context) FusionAuthUserActionOutput
}

func (*FusionAuthUserAction) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthUserAction)(nil)).Elem()
}

func (i *FusionAuthUserAction) ToFusionAuthUserActionOutput() FusionAuthUserActionOutput {
	return i.ToFusionAuthUserActionOutputWithContext(context.Background())
}

func (i *FusionAuthUserAction) ToFusionAuthUserActionOutputWithContext(ctx context.Context) FusionAuthUserActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthUserActionOutput)
}

// FusionAuthUserActionArrayInput is an input type that accepts FusionAuthUserActionArray and FusionAuthUserActionArrayOutput values.
// You can construct a concrete instance of `FusionAuthUserActionArrayInput` via:
//
//          FusionAuthUserActionArray{ FusionAuthUserActionArgs{...} }
type FusionAuthUserActionArrayInput interface {
	pulumi.Input

	ToFusionAuthUserActionArrayOutput() FusionAuthUserActionArrayOutput
	ToFusionAuthUserActionArrayOutputWithContext(context.Context) FusionAuthUserActionArrayOutput
}

type FusionAuthUserActionArray []FusionAuthUserActionInput

func (FusionAuthUserActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthUserAction)(nil)).Elem()
}

func (i FusionAuthUserActionArray) ToFusionAuthUserActionArrayOutput() FusionAuthUserActionArrayOutput {
	return i.ToFusionAuthUserActionArrayOutputWithContext(context.Background())
}

func (i FusionAuthUserActionArray) ToFusionAuthUserActionArrayOutputWithContext(ctx context.Context) FusionAuthUserActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthUserActionArrayOutput)
}

// FusionAuthUserActionMapInput is an input type that accepts FusionAuthUserActionMap and FusionAuthUserActionMapOutput values.
// You can construct a concrete instance of `FusionAuthUserActionMapInput` via:
//
//          FusionAuthUserActionMap{ "key": FusionAuthUserActionArgs{...} }
type FusionAuthUserActionMapInput interface {
	pulumi.Input

	ToFusionAuthUserActionMapOutput() FusionAuthUserActionMapOutput
	ToFusionAuthUserActionMapOutputWithContext(context.Context) FusionAuthUserActionMapOutput
}

type FusionAuthUserActionMap map[string]FusionAuthUserActionInput

func (FusionAuthUserActionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthUserAction)(nil)).Elem()
}

func (i FusionAuthUserActionMap) ToFusionAuthUserActionMapOutput() FusionAuthUserActionMapOutput {
	return i.ToFusionAuthUserActionMapOutputWithContext(context.Background())
}

func (i FusionAuthUserActionMap) ToFusionAuthUserActionMapOutputWithContext(ctx context.Context) FusionAuthUserActionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthUserActionMapOutput)
}

type FusionAuthUserActionOutput struct{ *pulumi.OutputState }

func (FusionAuthUserActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthUserAction)(nil)).Elem()
}

func (o FusionAuthUserActionOutput) ToFusionAuthUserActionOutput() FusionAuthUserActionOutput {
	return o
}

func (o FusionAuthUserActionOutput) ToFusionAuthUserActionOutputWithContext(ctx context.Context) FusionAuthUserActionOutput {
	return o
}

// The Id of the Email Template that is used when User Actions are canceled.
func (o FusionAuthUserActionOutput) CancelEmailTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthUserAction) pulumi.StringPtrOutput { return v.CancelEmailTemplateId }).(pulumi.StringPtrOutput)
}

// The Id of the Email Template that is used when User Actions expired automatically (end).
func (o FusionAuthUserActionOutput) EndEmailTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthUserAction) pulumi.StringPtrOutput { return v.EndEmailTemplateId }).(pulumi.StringPtrOutput)
}

// Whether to include the email information in the JSON that is sent to the Webhook when a user action is taken.
func (o FusionAuthUserActionOutput) IncludeEmailInEventJson() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthUserAction) pulumi.BoolPtrOutput { return v.IncludeEmailInEventJson }).(pulumi.BoolPtrOutput)
}

// A mapping of localized names for this User Action Option. The key is the Locale and the value is the name of the User Action Option for that language.
func (o FusionAuthUserActionOutput) LocalizedNames() pulumi.MapOutput {
	return o.ApplyT(func(v *FusionAuthUserAction) pulumi.MapOutput { return v.LocalizedNames }).(pulumi.MapOutput)
}

// The Id of the Email Template that is used when User Actions are modified.
func (o FusionAuthUserActionOutput) ModifyEmailTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthUserAction) pulumi.StringPtrOutput { return v.ModifyEmailTemplateId }).(pulumi.StringPtrOutput)
}

// The name of this User Action Option.
func (o FusionAuthUserActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthUserAction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of User Action Options.
func (o FusionAuthUserActionOutput) Options() FusionAuthUserActionOptionArrayOutput {
	return o.ApplyT(func(v *FusionAuthUserAction) FusionAuthUserActionOptionArrayOutput { return v.Options }).(FusionAuthUserActionOptionArrayOutput)
}

// Whether or not this User Action will prevent user login. When this value is set to true the action must also be marked as a time based action. See `temporal`.
func (o FusionAuthUserActionOutput) PreventLogin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthUserAction) pulumi.BoolPtrOutput { return v.PreventLogin }).(pulumi.BoolPtrOutput)
}

// Whether or not FusionAuth will send events to any registered Webhooks when this User Action expires.
func (o FusionAuthUserActionOutput) SendEndEvent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthUserAction) pulumi.BoolPtrOutput { return v.SendEndEvent }).(pulumi.BoolPtrOutput)
}

// The Id of the Email Template that is used when User Actions are started (created).
func (o FusionAuthUserActionOutput) StartEmailTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthUserAction) pulumi.StringPtrOutput { return v.StartEmailTemplateId }).(pulumi.StringPtrOutput)
}

// Whether or not this User Action is time-based (temporal).
func (o FusionAuthUserActionOutput) Temporal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthUserAction) pulumi.BoolPtrOutput { return v.Temporal }).(pulumi.BoolPtrOutput)
}

// Whether or not email is enabled for this User Action.
func (o FusionAuthUserActionOutput) UserEmailingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthUserAction) pulumi.BoolPtrOutput { return v.UserEmailingEnabled }).(pulumi.BoolPtrOutput)
}

// Whether or not user notifications are enabled for this User Action.
func (o FusionAuthUserActionOutput) UserNotificationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthUserAction) pulumi.BoolPtrOutput { return v.UserNotificationsEnabled }).(pulumi.BoolPtrOutput)
}

type FusionAuthUserActionArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthUserActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthUserAction)(nil)).Elem()
}

func (o FusionAuthUserActionArrayOutput) ToFusionAuthUserActionArrayOutput() FusionAuthUserActionArrayOutput {
	return o
}

func (o FusionAuthUserActionArrayOutput) ToFusionAuthUserActionArrayOutputWithContext(ctx context.Context) FusionAuthUserActionArrayOutput {
	return o
}

func (o FusionAuthUserActionArrayOutput) Index(i pulumi.IntInput) FusionAuthUserActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthUserAction {
		return vs[0].([]*FusionAuthUserAction)[vs[1].(int)]
	}).(FusionAuthUserActionOutput)
}

type FusionAuthUserActionMapOutput struct{ *pulumi.OutputState }

func (FusionAuthUserActionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthUserAction)(nil)).Elem()
}

func (o FusionAuthUserActionMapOutput) ToFusionAuthUserActionMapOutput() FusionAuthUserActionMapOutput {
	return o
}

func (o FusionAuthUserActionMapOutput) ToFusionAuthUserActionMapOutputWithContext(ctx context.Context) FusionAuthUserActionMapOutput {
	return o
}

func (o FusionAuthUserActionMapOutput) MapIndex(k pulumi.StringInput) FusionAuthUserActionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthUserAction {
		return vs[0].(map[string]*FusionAuthUserAction)[vs[1].(string)]
	}).(FusionAuthUserActionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthUserActionInput)(nil)).Elem(), &FusionAuthUserAction{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthUserActionArrayInput)(nil)).Elem(), FusionAuthUserActionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthUserActionMapInput)(nil)).Elem(), FusionAuthUserActionMap{})
	pulumi.RegisterOutputType(FusionAuthUserActionOutput{})
	pulumi.RegisterOutputType(FusionAuthUserActionArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthUserActionMapOutput{})
}
