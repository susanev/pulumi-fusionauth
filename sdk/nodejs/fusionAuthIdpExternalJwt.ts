// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # External JWT Identity Provider Resource
 *
 * This is a special type of identity provider that is only used via the JWT Reconcile API. This identity provider defines the claims inside the incoming JWT and how they map to fields in the FusionAuth User object.
 *
 * In order for this identity provider to use the JWT, it also needs the public key or HMAC secret that the JWT was signed with. FusionAuth will verify that the JWT is valid and has not expired. Once the JWT has been validated, FusionAuth will reconcile it to ensure that the User exists and is up-to-date.
 *
 * [External JWT Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/external-jwt/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "@pulumi/fusionauth";
 *
 * const jwt = new fusionauth.FusionAuthIdpExternalJwt("jwt", {
 *     claimMap: {
 *         dept: "RegistrationData",
 *         first_name: "firstName",
 *         last_name: "lastName",
 *     },
 *     debug: false,
 *     enabled: true,
 *     headerKeyParameter: "kid",
 *     oauth2AuthorizationEndpoint: "https://acme.com/adfs/oauth2/authorize?client_id=cf3b00da-9551-460a-ad18-33232e6cbff0&response_type=code&redirect_uri=https://acme.com/oauth2/redirect",
 *     oauth2TokenEndpoint: "https://acme.com/adfs/oauth2/token",
 *     uniqueIdentityClaim: "email",
 * });
 * ```
 */
export class FusionAuthIdpExternalJwt extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthIdpExternalJwt resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthIdpExternalJwtState, opts?: pulumi.CustomResourceOptions): FusionAuthIdpExternalJwt {
        return new FusionAuthIdpExternalJwt(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthIdpExternalJwt:FusionAuthIdpExternalJwt';

    /**
     * Returns true if the given object is an instance of FusionAuthIdpExternalJwt.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthIdpExternalJwt {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthIdpExternalJwt.__pulumiType;
    }

    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    public readonly applicationConfigurations!: pulumi.Output<outputs.FusionAuthIdpExternalJwtApplicationConfiguration[] | undefined>;
    /**
     * A map of incoming claims to User fields, User data or Registration data. The key of the map is the incoming claim name
     * from the configured identity provider.
     */
    public readonly claimMap!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
     */
    public readonly debug!: pulumi.Output<boolean | undefined>;
    /**
     * An array of domains that are managed by this Identity Provider.
     */
    public readonly domains!: pulumi.Output<string[] | undefined>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name header claim that identifies the public key used to verify the signature. In most cases this be kid or x5t.
     */
    public readonly headerKeyParameter!: pulumi.Output<string>;
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
     * The name of the Identity Provider.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The authorization endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to perform the browser redirect to the OAuth2 authorize endpoint.
     */
    public readonly oauth2AuthorizationEndpoint!: pulumi.Output<string | undefined>;
    /**
     * TThe token endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to complete the OAuth2 grant workflow.
     */
    public readonly oauth2TokenEndpoint!: pulumi.Output<string | undefined>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    public readonly tenantConfigurations!: pulumi.Output<outputs.FusionAuthIdpExternalJwtTenantConfiguration[] | undefined>;
    /**
     * The name of the claim that represents the unique identify of the User. This will generally be email or the name of the claim that provides the email address.
     */
    public readonly uniqueIdentityClaim!: pulumi.Output<string>;

    /**
     * Create a FusionAuthIdpExternalJwt resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthIdpExternalJwtArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthIdpExternalJwtArgs | FusionAuthIdpExternalJwtState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthIdpExternalJwtState | undefined;
            resourceInputs["applicationConfigurations"] = state ? state.applicationConfigurations : undefined;
            resourceInputs["claimMap"] = state ? state.claimMap : undefined;
            resourceInputs["debug"] = state ? state.debug : undefined;
            resourceInputs["domains"] = state ? state.domains : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["headerKeyParameter"] = state ? state.headerKeyParameter : undefined;
            resourceInputs["idpId"] = state ? state.idpId : undefined;
            resourceInputs["lambdaReconcileId"] = state ? state.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = state ? state.linkingStrategy : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["oauth2AuthorizationEndpoint"] = state ? state.oauth2AuthorizationEndpoint : undefined;
            resourceInputs["oauth2TokenEndpoint"] = state ? state.oauth2TokenEndpoint : undefined;
            resourceInputs["tenantConfigurations"] = state ? state.tenantConfigurations : undefined;
            resourceInputs["uniqueIdentityClaim"] = state ? state.uniqueIdentityClaim : undefined;
        } else {
            const args = argsOrState as FusionAuthIdpExternalJwtArgs | undefined;
            if ((!args || args.headerKeyParameter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'headerKeyParameter'");
            }
            if ((!args || args.uniqueIdentityClaim === undefined) && !opts.urn) {
                throw new Error("Missing required property 'uniqueIdentityClaim'");
            }
            resourceInputs["applicationConfigurations"] = args ? args.applicationConfigurations : undefined;
            resourceInputs["claimMap"] = args ? args.claimMap : undefined;
            resourceInputs["debug"] = args ? args.debug : undefined;
            resourceInputs["domains"] = args ? args.domains : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["headerKeyParameter"] = args ? args.headerKeyParameter : undefined;
            resourceInputs["idpId"] = args ? args.idpId : undefined;
            resourceInputs["lambdaReconcileId"] = args ? args.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = args ? args.linkingStrategy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["oauth2AuthorizationEndpoint"] = args ? args.oauth2AuthorizationEndpoint : undefined;
            resourceInputs["oauth2TokenEndpoint"] = args ? args.oauth2TokenEndpoint : undefined;
            resourceInputs["tenantConfigurations"] = args ? args.tenantConfigurations : undefined;
            resourceInputs["uniqueIdentityClaim"] = args ? args.uniqueIdentityClaim : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FusionAuthIdpExternalJwt.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthIdpExternalJwt resources.
 */
export interface FusionAuthIdpExternalJwtState {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpExternalJwtApplicationConfiguration>[]>;
    /**
     * A map of incoming claims to User fields, User data or Registration data. The key of the map is the incoming claim name
     * from the configured identity provider.
     */
    claimMap?: pulumi.Input<{[key: string]: any}>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
     */
    debug?: pulumi.Input<boolean>;
    /**
     * An array of domains that are managed by this Identity Provider.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The name header claim that identifies the public key used to verify the signature. In most cases this be kid or x5t.
     */
    headerKeyParameter?: pulumi.Input<string>;
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
     * The name of the Identity Provider.
     */
    name?: pulumi.Input<string>;
    /**
     * The authorization endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to perform the browser redirect to the OAuth2 authorize endpoint.
     */
    oauth2AuthorizationEndpoint?: pulumi.Input<string>;
    /**
     * TThe token endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to complete the OAuth2 grant workflow.
     */
    oauth2TokenEndpoint?: pulumi.Input<string>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpExternalJwtTenantConfiguration>[]>;
    /**
     * The name of the claim that represents the unique identify of the User. This will generally be email or the name of the claim that provides the email address.
     */
    uniqueIdentityClaim?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FusionAuthIdpExternalJwt resource.
 */
export interface FusionAuthIdpExternalJwtArgs {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpExternalJwtApplicationConfiguration>[]>;
    /**
     * A map of incoming claims to User fields, User data or Registration data. The key of the map is the incoming claim name
     * from the configured identity provider.
     */
    claimMap?: pulumi.Input<{[key: string]: any}>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
     */
    debug?: pulumi.Input<boolean>;
    /**
     * An array of domains that are managed by this Identity Provider.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The name header claim that identifies the public key used to verify the signature. In most cases this be kid or x5t.
     */
    headerKeyParameter: pulumi.Input<string>;
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
     * The name of the Identity Provider.
     */
    name?: pulumi.Input<string>;
    /**
     * The authorization endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to perform the browser redirect to the OAuth2 authorize endpoint.
     */
    oauth2AuthorizationEndpoint?: pulumi.Input<string>;
    /**
     * TThe token endpoint for this Identity Provider. This value is not utilized by FusionAuth is only provided to be returned by the Lookup Identity Provider API response. During integration you may then utilize this value to complete the OAuth2 grant workflow.
     */
    oauth2TokenEndpoint?: pulumi.Input<string>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpExternalJwtTenantConfiguration>[]>;
    /**
     * The name of the claim that represents the unique identify of the User. This will generally be email or the name of the claim that provides the email address.
     */
    uniqueIdentityClaim: pulumi.Input<string>;
}
