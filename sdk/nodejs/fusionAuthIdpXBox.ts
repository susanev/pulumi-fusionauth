// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # Xbox Identity Provider Resource
 *
 * The Xbox identity provider type will use the Xbox OAuth v2.0 login API. It will also provide a Login with Xbox button on FusionAuth’s login page that will direct a user to the Xbox login page.
 *
 * This identity provider will call Xbox’s API to load the user’s email and gtg (Gamer Tag) and use those as email and username to lookup or create a user in FusionAuth depending on the linking strategy configured for this identity provider. Additional claims returned by Xbox can be used to reconcile the user to FusionAuth by using an Xbox Reconcile Lambda.
 *
 * FusionAuth will also store the Xbox refreshToken returned from the Xbox API in the link between the user and the identity provider. This token can be used by an application to make further requests to Xbox APIs on behalf of the user.
 *
 * [Xbox Identity Provider APIs](https://fusionauth.io/docs/v1/tech/apis/identity-providers/xbox/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi_fusionauth from "pulumi-fusionauth";
 *
 * const xbox = new fusionauth.FusionAuthIdpXBox("xbox", {
 *     applicationConfigurations: [{
 *         applicationId: fusionauth_application.GPS_Insight.id,
 *         createRegistration: true,
 *         enabled: true,
 *     }],
 *     buttonText: "Login with Xbox",
 *     clientId: "0eb1ce3c-2fb1-4ae9-b361-d49fc6e764cc",
 *     clientSecret: "693s000cbn66k0mxtqzr_c_NfLy3~6_SEA",
 *     scope: "Xboxlive.signin Xboxlive.offline_access",
 * });
 * ```
 */
export class FusionAuthIdpXBox extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthIdpXBox resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthIdpXBoxState, opts?: pulumi.CustomResourceOptions): FusionAuthIdpXBox {
        return new FusionAuthIdpXBox(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthIdpXBox:FusionAuthIdpXBox';

    /**
     * Returns true if the given object is an instance of FusionAuthIdpXBox.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthIdpXBox {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthIdpXBox.__pulumiType;
    }

    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    public readonly applicationConfigurations!: pulumi.Output<outputs.FusionAuthIdpXBoxApplicationConfiguration[] | undefined>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    public readonly buttonText!: pulumi.Output<string>;
    /**
     * TThe top-level Xbox client id for your Application. This value is retrieved from the Xbox developer website when you setup your Xbox developer account.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The top-level client secret to use with the Xbox Identity Provider when retrieving the long-lived token. This value is retrieved from the Xbox developer website when you setup your Xbox developer account.
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
     */
    public readonly debug!: pulumi.Output<boolean | undefined>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
     */
    public readonly idpId!: pulumi.Output<string | undefined>;
    /**
     * The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
     */
    public readonly lambdaReconcileId!: pulumi.Output<string | undefined>;
    /**
     * The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
     */
    public readonly linkingStrategy!: pulumi.Output<string>;
    /**
     * The top-level scope that you are requesting from Xbox.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    public readonly tenantConfigurations!: pulumi.Output<outputs.FusionAuthIdpXBoxTenantConfiguration[] | undefined>;

    /**
     * Create a FusionAuthIdpXBox resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthIdpXBoxArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthIdpXBoxArgs | FusionAuthIdpXBoxState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthIdpXBoxState | undefined;
            resourceInputs["applicationConfigurations"] = state ? state.applicationConfigurations : undefined;
            resourceInputs["buttonText"] = state ? state.buttonText : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["debug"] = state ? state.debug : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["idpId"] = state ? state.idpId : undefined;
            resourceInputs["lambdaReconcileId"] = state ? state.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = state ? state.linkingStrategy : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["tenantConfigurations"] = state ? state.tenantConfigurations : undefined;
        } else {
            const args = argsOrState as FusionAuthIdpXBoxArgs | undefined;
            if ((!args || args.buttonText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'buttonText'");
            }
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.clientSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientSecret'");
            }
            resourceInputs["applicationConfigurations"] = args ? args.applicationConfigurations : undefined;
            resourceInputs["buttonText"] = args ? args.buttonText : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args ? args.clientSecret : undefined;
            resourceInputs["debug"] = args ? args.debug : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["idpId"] = args ? args.idpId : undefined;
            resourceInputs["lambdaReconcileId"] = args ? args.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = args ? args.linkingStrategy : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["tenantConfigurations"] = args ? args.tenantConfigurations : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FusionAuthIdpXBox.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthIdpXBox resources.
 */
export interface FusionAuthIdpXBoxState {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpXBoxApplicationConfiguration>[]>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    buttonText?: pulumi.Input<string>;
    /**
     * TThe top-level Xbox client id for your Application. This value is retrieved from the Xbox developer website when you setup your Xbox developer account.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The top-level client secret to use with the Xbox Identity Provider when retrieving the long-lived token. This value is retrieved from the Xbox developer website when you setup your Xbox developer account.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
     */
    debug?: pulumi.Input<boolean>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
     */
    idpId?: pulumi.Input<string>;
    /**
     * The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
     */
    lambdaReconcileId?: pulumi.Input<string>;
    /**
     * The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
     */
    linkingStrategy?: pulumi.Input<string>;
    /**
     * The top-level scope that you are requesting from Xbox.
     */
    scope?: pulumi.Input<string>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpXBoxTenantConfiguration>[]>;
}

/**
 * The set of arguments for constructing a FusionAuthIdpXBox resource.
 */
export interface FusionAuthIdpXBoxArgs {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpXBoxApplicationConfiguration>[]>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    buttonText: pulumi.Input<string>;
    /**
     * TThe top-level Xbox client id for your Application. This value is retrieved from the Xbox developer website when you setup your Xbox developer account.
     */
    clientId: pulumi.Input<string>;
    /**
     * The top-level client secret to use with the Xbox Identity Provider when retrieving the long-lived token. This value is retrieved from the Xbox developer website when you setup your Xbox developer account.
     */
    clientSecret: pulumi.Input<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
     */
    debug?: pulumi.Input<boolean>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
     */
    idpId?: pulumi.Input<string>;
    /**
     * The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
     */
    lambdaReconcileId?: pulumi.Input<string>;
    /**
     * The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
     */
    linkingStrategy?: pulumi.Input<string>;
    /**
     * The top-level scope that you are requesting from Xbox.
     */
    scope?: pulumi.Input<string>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpXBoxTenantConfiguration>[]>;
}
