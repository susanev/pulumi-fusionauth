// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # Apple Identity Provider Resource
 *
 * The Apple identity provider type will use the Sign in with Apple APIs and will provide a Sign with Apple button on FusionAuth’s login page that will either redirect to an Apple sign in page or leverage native controls when using Safari on macOS or iOS. Additionally, this identity provider will call Apples’s /auth/token API to load additional details about the user and store them in FusionAuth.
 *
 * FusionAuth will also store the Apple refreshToken that is returned from the /auth/token endpoint in the UserRegistration object inside the tokens Map. This Map stores the tokens from the various identity providers so that you can use them in your application to call their APIs.
 *
 * [Apple Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/apple/#create-the-apple-identity-provider)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "@pulumi/fusionauth";
 *
 * const apple = new fusionauth.FusionAuthIdpApple("apple", {
 *     applicationConfigurations: [{
 *         applicationId: "1c212e59-0d0e-6b1a-ad48-f4f92793be32",
 *         createRegistration: true,
 *         enabled: true,
 *     }],
 *     buttonText: "Sign in with Apple",
 *     debug: false,
 *     enabled: true,
 *     keyId: "2f81529c-4d39-4ce2-982e-cf5fbb1325f6",
 *     scope: "email name",
 *     servicesId: "com.piedpiper.webapp",
 *     teamId: "R4NQ1P4UEB",
 * });
 * ```
 */
export class FusionAuthIdpApple extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthIdpApple resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthIdpAppleState, opts?: pulumi.CustomResourceOptions): FusionAuthIdpApple {
        return new FusionAuthIdpApple(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthIdpApple:FusionAuthIdpApple';

    /**
     * Returns true if the given object is an instance of FusionAuthIdpApple.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthIdpApple {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthIdpApple.__pulumiType;
    }

    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    public readonly applicationConfigurations!: pulumi.Output<outputs.FusionAuthIdpAppleApplicationConfiguration[] | undefined>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    public readonly buttonText!: pulumi.Output<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
     */
    public readonly debug!: pulumi.Output<boolean | undefined>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
     */
    public readonly keyId!: pulumi.Output<string>;
    /**
     * The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
     */
    public readonly lambdaReconcileId!: pulumi.Output<string | undefined>;
    /**
     * The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
     */
    public readonly linkingStrategy!: pulumi.Output<string>;
    /**
     * The top-level space separated scope that you are requesting from Apple.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
     */
    public readonly servicesId!: pulumi.Output<string>;
    /**
     * The Apple App ID Prefix, or Team ID found in your Apple Developer Account which has been configured for Sign in with Apple.
     */
    public readonly teamId!: pulumi.Output<string>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    public readonly tenantConfigurations!: pulumi.Output<outputs.FusionAuthIdpAppleTenantConfiguration[] | undefined>;

    /**
     * Create a FusionAuthIdpApple resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthIdpAppleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthIdpAppleArgs | FusionAuthIdpAppleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthIdpAppleState | undefined;
            resourceInputs["applicationConfigurations"] = state ? state.applicationConfigurations : undefined;
            resourceInputs["buttonText"] = state ? state.buttonText : undefined;
            resourceInputs["debug"] = state ? state.debug : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["lambdaReconcileId"] = state ? state.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = state ? state.linkingStrategy : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["servicesId"] = state ? state.servicesId : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
            resourceInputs["tenantConfigurations"] = state ? state.tenantConfigurations : undefined;
        } else {
            const args = argsOrState as FusionAuthIdpAppleArgs | undefined;
            if ((!args || args.buttonText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'buttonText'");
            }
            if ((!args || args.keyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyId'");
            }
            if ((!args || args.servicesId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'servicesId'");
            }
            if ((!args || args.teamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamId'");
            }
            resourceInputs["applicationConfigurations"] = args ? args.applicationConfigurations : undefined;
            resourceInputs["buttonText"] = args ? args.buttonText : undefined;
            resourceInputs["debug"] = args ? args.debug : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["keyId"] = args ? args.keyId : undefined;
            resourceInputs["lambdaReconcileId"] = args ? args.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = args ? args.linkingStrategy : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["servicesId"] = args ? args.servicesId : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["tenantConfigurations"] = args ? args.tenantConfigurations : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FusionAuthIdpApple.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthIdpApple resources.
 */
export interface FusionAuthIdpAppleState {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpAppleApplicationConfiguration>[]>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    buttonText?: pulumi.Input<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
     */
    debug?: pulumi.Input<boolean>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
     */
    lambdaReconcileId?: pulumi.Input<string>;
    /**
     * The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
     */
    linkingStrategy?: pulumi.Input<string>;
    /**
     * The top-level space separated scope that you are requesting from Apple.
     */
    scope?: pulumi.Input<string>;
    /**
     * The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
     */
    servicesId?: pulumi.Input<string>;
    /**
     * The Apple App ID Prefix, or Team ID found in your Apple Developer Account which has been configured for Sign in with Apple.
     */
    teamId?: pulumi.Input<string>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpAppleTenantConfiguration>[]>;
}

/**
 * The set of arguments for constructing a FusionAuthIdpApple resource.
 */
export interface FusionAuthIdpAppleArgs {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpAppleApplicationConfiguration>[]>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    buttonText: pulumi.Input<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
     */
    debug?: pulumi.Input<boolean>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
     */
    keyId: pulumi.Input<string>;
    /**
     * The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
     */
    lambdaReconcileId?: pulumi.Input<string>;
    /**
     * The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
     */
    linkingStrategy?: pulumi.Input<string>;
    /**
     * The top-level space separated scope that you are requesting from Apple.
     */
    scope?: pulumi.Input<string>;
    /**
     * The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
     */
    servicesId: pulumi.Input<string>;
    /**
     * The Apple App ID Prefix, or Team ID found in your Apple Developer Account which has been configured for Sign in with Apple.
     */
    teamId: pulumi.Input<string>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpAppleTenantConfiguration>[]>;
}
