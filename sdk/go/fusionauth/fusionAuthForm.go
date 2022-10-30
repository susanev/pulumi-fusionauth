// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Form Resource
//
// A FusionAuth Form is a customizable object that contains one-to-many ordered steps. Each step is comprised of one or more Form Fields.
//
// [Form API](https://fusionauth.io/docs/v1/tech/apis/forms/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fusionauth/sdk/v2/go/fusionauth"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/theogravity/pulumi-fusionauth/sdk/v2/go/fusionauth"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fusionauth.NewFusionAuthForm(ctx, "form", &fusionauth.FusionAuthFormArgs{
// 			Data: pulumi.AnyMap{
// 				"description": pulumi.Any("This form customizes the registration experience."),
// 			},
// 			Steps: FusionAuthFormStepArray{
// 				&FusionAuthFormStepArgs{
// 					Fields: pulumi.StringArray{
// 						pulumi.String("91909721-7d4f-b110-8f21-cfdee2a1edb8"),
// 					},
// 				},
// 				&FusionAuthFormStepArgs{
// 					Fields: pulumi.StringArray{
// 						pulumi.String("8ed89a31-c325-3156-72ed-6e89183af917"),
// 						pulumi.String("a977cfd4-a9ed-c4cf-650f-f4539268ac38"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FusionAuthForm struct {
	pulumi.CustomResourceState

	// An object that can hold any information about the Form Field that should be persisted.
	Data pulumi.MapOutput `pulumi:"data"`
	// The Id to use for the new Form. If not specified a secure random UUID will be generated.
	FormId pulumi.StringPtrOutput `pulumi:"formId"`
	// The unique name of the Form Field.
	Name pulumi.StringOutput `pulumi:"name"`
	// An ordered list of objects containing one or more Form Fields. A Form must have at least one step defined.
	Steps FusionAuthFormStepArrayOutput `pulumi:"steps"`
	// The type of form being created, a form type cannot be changed after the form has been created.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewFusionAuthForm registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthForm(ctx *pulumi.Context,
	name string, args *FusionAuthFormArgs, opts ...pulumi.ResourceOption) (*FusionAuthForm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Steps == nil {
		return nil, errors.New("invalid value for required argument 'Steps'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FusionAuthForm
	err := ctx.RegisterResource("fusionauth:index/fusionAuthForm:FusionAuthForm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthForm gets an existing FusionAuthForm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthForm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthFormState, opts ...pulumi.ResourceOption) (*FusionAuthForm, error) {
	var resource FusionAuthForm
	err := ctx.ReadResource("fusionauth:index/fusionAuthForm:FusionAuthForm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthForm resources.
type fusionAuthFormState struct {
	// An object that can hold any information about the Form Field that should be persisted.
	Data map[string]interface{} `pulumi:"data"`
	// The Id to use for the new Form. If not specified a secure random UUID will be generated.
	FormId *string `pulumi:"formId"`
	// The unique name of the Form Field.
	Name *string `pulumi:"name"`
	// An ordered list of objects containing one or more Form Fields. A Form must have at least one step defined.
	Steps []FusionAuthFormStep `pulumi:"steps"`
	// The type of form being created, a form type cannot be changed after the form has been created.
	Type *string `pulumi:"type"`
}

type FusionAuthFormState struct {
	// An object that can hold any information about the Form Field that should be persisted.
	Data pulumi.MapInput
	// The Id to use for the new Form. If not specified a secure random UUID will be generated.
	FormId pulumi.StringPtrInput
	// The unique name of the Form Field.
	Name pulumi.StringPtrInput
	// An ordered list of objects containing one or more Form Fields. A Form must have at least one step defined.
	Steps FusionAuthFormStepArrayInput
	// The type of form being created, a form type cannot be changed after the form has been created.
	Type pulumi.StringPtrInput
}

func (FusionAuthFormState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthFormState)(nil)).Elem()
}

type fusionAuthFormArgs struct {
	// An object that can hold any information about the Form Field that should be persisted.
	Data map[string]interface{} `pulumi:"data"`
	// The Id to use for the new Form. If not specified a secure random UUID will be generated.
	FormId *string `pulumi:"formId"`
	// The unique name of the Form Field.
	Name *string `pulumi:"name"`
	// An ordered list of objects containing one or more Form Fields. A Form must have at least one step defined.
	Steps []FusionAuthFormStep `pulumi:"steps"`
	// The type of form being created, a form type cannot be changed after the form has been created.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a FusionAuthForm resource.
type FusionAuthFormArgs struct {
	// An object that can hold any information about the Form Field that should be persisted.
	Data pulumi.MapInput
	// The Id to use for the new Form. If not specified a secure random UUID will be generated.
	FormId pulumi.StringPtrInput
	// The unique name of the Form Field.
	Name pulumi.StringPtrInput
	// An ordered list of objects containing one or more Form Fields. A Form must have at least one step defined.
	Steps FusionAuthFormStepArrayInput
	// The type of form being created, a form type cannot be changed after the form has been created.
	Type pulumi.StringPtrInput
}

func (FusionAuthFormArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthFormArgs)(nil)).Elem()
}

type FusionAuthFormInput interface {
	pulumi.Input

	ToFusionAuthFormOutput() FusionAuthFormOutput
	ToFusionAuthFormOutputWithContext(ctx context.Context) FusionAuthFormOutput
}

func (*FusionAuthForm) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthForm)(nil)).Elem()
}

func (i *FusionAuthForm) ToFusionAuthFormOutput() FusionAuthFormOutput {
	return i.ToFusionAuthFormOutputWithContext(context.Background())
}

func (i *FusionAuthForm) ToFusionAuthFormOutputWithContext(ctx context.Context) FusionAuthFormOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthFormOutput)
}

// FusionAuthFormArrayInput is an input type that accepts FusionAuthFormArray and FusionAuthFormArrayOutput values.
// You can construct a concrete instance of `FusionAuthFormArrayInput` via:
//
//          FusionAuthFormArray{ FusionAuthFormArgs{...} }
type FusionAuthFormArrayInput interface {
	pulumi.Input

	ToFusionAuthFormArrayOutput() FusionAuthFormArrayOutput
	ToFusionAuthFormArrayOutputWithContext(context.Context) FusionAuthFormArrayOutput
}

type FusionAuthFormArray []FusionAuthFormInput

func (FusionAuthFormArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthForm)(nil)).Elem()
}

func (i FusionAuthFormArray) ToFusionAuthFormArrayOutput() FusionAuthFormArrayOutput {
	return i.ToFusionAuthFormArrayOutputWithContext(context.Background())
}

func (i FusionAuthFormArray) ToFusionAuthFormArrayOutputWithContext(ctx context.Context) FusionAuthFormArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthFormArrayOutput)
}

// FusionAuthFormMapInput is an input type that accepts FusionAuthFormMap and FusionAuthFormMapOutput values.
// You can construct a concrete instance of `FusionAuthFormMapInput` via:
//
//          FusionAuthFormMap{ "key": FusionAuthFormArgs{...} }
type FusionAuthFormMapInput interface {
	pulumi.Input

	ToFusionAuthFormMapOutput() FusionAuthFormMapOutput
	ToFusionAuthFormMapOutputWithContext(context.Context) FusionAuthFormMapOutput
}

type FusionAuthFormMap map[string]FusionAuthFormInput

func (FusionAuthFormMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthForm)(nil)).Elem()
}

func (i FusionAuthFormMap) ToFusionAuthFormMapOutput() FusionAuthFormMapOutput {
	return i.ToFusionAuthFormMapOutputWithContext(context.Background())
}

func (i FusionAuthFormMap) ToFusionAuthFormMapOutputWithContext(ctx context.Context) FusionAuthFormMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthFormMapOutput)
}

type FusionAuthFormOutput struct{ *pulumi.OutputState }

func (FusionAuthFormOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthForm)(nil)).Elem()
}

func (o FusionAuthFormOutput) ToFusionAuthFormOutput() FusionAuthFormOutput {
	return o
}

func (o FusionAuthFormOutput) ToFusionAuthFormOutputWithContext(ctx context.Context) FusionAuthFormOutput {
	return o
}

// An object that can hold any information about the Form Field that should be persisted.
func (o FusionAuthFormOutput) Data() pulumi.MapOutput {
	return o.ApplyT(func(v *FusionAuthForm) pulumi.MapOutput { return v.Data }).(pulumi.MapOutput)
}

// The Id to use for the new Form. If not specified a secure random UUID will be generated.
func (o FusionAuthFormOutput) FormId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthForm) pulumi.StringPtrOutput { return v.FormId }).(pulumi.StringPtrOutput)
}

// The unique name of the Form Field.
func (o FusionAuthFormOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthForm) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An ordered list of objects containing one or more Form Fields. A Form must have at least one step defined.
func (o FusionAuthFormOutput) Steps() FusionAuthFormStepArrayOutput {
	return o.ApplyT(func(v *FusionAuthForm) FusionAuthFormStepArrayOutput { return v.Steps }).(FusionAuthFormStepArrayOutput)
}

// The type of form being created, a form type cannot be changed after the form has been created.
func (o FusionAuthFormOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthForm) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type FusionAuthFormArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthFormArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthForm)(nil)).Elem()
}

func (o FusionAuthFormArrayOutput) ToFusionAuthFormArrayOutput() FusionAuthFormArrayOutput {
	return o
}

func (o FusionAuthFormArrayOutput) ToFusionAuthFormArrayOutputWithContext(ctx context.Context) FusionAuthFormArrayOutput {
	return o
}

func (o FusionAuthFormArrayOutput) Index(i pulumi.IntInput) FusionAuthFormOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthForm {
		return vs[0].([]*FusionAuthForm)[vs[1].(int)]
	}).(FusionAuthFormOutput)
}

type FusionAuthFormMapOutput struct{ *pulumi.OutputState }

func (FusionAuthFormMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthForm)(nil)).Elem()
}

func (o FusionAuthFormMapOutput) ToFusionAuthFormMapOutput() FusionAuthFormMapOutput {
	return o
}

func (o FusionAuthFormMapOutput) ToFusionAuthFormMapOutputWithContext(ctx context.Context) FusionAuthFormMapOutput {
	return o
}

func (o FusionAuthFormMapOutput) MapIndex(k pulumi.StringInput) FusionAuthFormOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthForm {
		return vs[0].(map[string]*FusionAuthForm)[vs[1].(string)]
	}).(FusionAuthFormOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthFormInput)(nil)).Elem(), &FusionAuthForm{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthFormArrayInput)(nil)).Elem(), FusionAuthFormArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthFormMapInput)(nil)).Elem(), FusionAuthFormMap{})
	pulumi.RegisterOutputType(FusionAuthFormOutput{})
	pulumi.RegisterOutputType(FusionAuthFormArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthFormMapOutput{})
}
