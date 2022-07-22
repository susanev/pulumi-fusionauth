// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # OpenID Connect Identity Provider Resource
 *
 * OpenID Connect identity providers connect to external OpenID Connect login systems. This type of login will optionally provide a Login with …​ button on FusionAuth’s login page. This button is customizable by using different properties of the identity provider.
 *
 * Optionally, this identity provider can define one or more domains it is associated with. This is useful for allowing employees to log in with their corporate credentials. As long as the company has an identity solution that provides OpenID Connect, you can leverage this feature. This is referred to as a Domain Based Identity Provider. If you enable domains for an identity provider, the Login with …​ button will not be displayed. Instead, only the email form field will be displayed initially on the FusionAuth login page. Once the user types in their email address, FusionAuth will determine if the user is logging in locally or if they should be redirected to this identity provider. This is determined by extracting the domain from their email address and comparing it to the domains associated with the identity provider.
 *
 * FusionAuth will also leverage the /userinfo API that is part of the OpenID Connect specification. The email address returned from the Userinfo response will be used to create or lookup the existing user. Additional claims from the Userinfo response can be used to reconcile the User in FusionAuth by using an OpenID Connect Reconcile Lambda. Unless you assign a reconcile lambda to this provider, on the email address will be used from the available claims returned by the OpenID Connect identity provider.
 *
 * If the external OpenID Connect identity provider returns a refresh token, it will be stored in the UserRegistration object inside the tokens Map. This Map stores the tokens from the various identity providers so that you can use them in your application to call their APIs.
 *
 * [OpenID Connect Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/openid-connect)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi_fusionauth from "@theogravity/pulumi-fusionauth";
 *
 * const openID = new fusionauth.FusionAuthIdpOpenIdConnect("openID", {
 *     applicationConfigurations: [{
 *         applicationId: fusionauth_application.myapp.id,
 *         createRegistration: true,
 *         enabled: true,
 *     }],
 *     buttonText: "Login with OpenID Connect",
 *     debug: false,
 *     oauth2AuthorizationEndpoint: "https://acme.com/oauth2/authorization",
 *     oauth2ClientId: "191c23dc-b772-4558-bd21-dc1cbf74ae21",
 *     oauth2ClientSecret: "SUsnoP0pWUYfXvWbSe5pvj8Di5nAxOvO",
 *     oauth2ClientAuthenticationMethod: "client_secret_basic",
 *     oauth2Scope: "openid offline_access",
 *     oauth2TokenEndpoint: "https://acme.com/oauth2/token",
 *     oauth2UserInfoEndpoint: "https://acme.com/oauth2/userinfo",
 * });
 * ```
 */
export class FusionAuthIdpOpenIdConnect extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthIdpOpenIdConnect resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthIdpOpenIdConnectState, opts?: pulumi.CustomResourceOptions): FusionAuthIdpOpenIdConnect {
        return new FusionAuthIdpOpenIdConnect(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthIdpOpenIdConnect:FusionAuthIdpOpenIdConnect';

    /**
     * Returns true if the given object is an instance of FusionAuthIdpOpenIdConnect.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthIdpOpenIdConnect {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthIdpOpenIdConnect.__pulumiType;
    }

    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    public readonly applicationConfigurations!: pulumi.Output<outputs.FusionAuthIdpOpenIdConnectApplicationConfiguration[] | undefined>;
    /**
     * The top-level button image (URL) to use on the FusionAuth login page for this Identity Provider.
     */
    public readonly buttonImageUrl!: pulumi.Output<string | undefined>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    public readonly buttonText!: pulumi.Output<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
     */
    public readonly debug!: pulumi.Output<boolean | undefined>;
    /**
     * This is an optional list of domains that this OpenID Connect provider should be used for. This converts the FusionAuth login form to a domain-based login form. This type of form first asks the user for their email. FusionAuth then uses their email to determine if an OpenID Connect identity provider should be used. If an OpenID Connect provider should be used, the browser is redirected to the authorization endpoint of that identity provider. Otherwise, the password field is revealed on the form so that the user can login using FusionAuth.
     */
    public readonly domains!: pulumi.Output<string[] | undefined>;
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
     * The name of this OpenID Connect identity provider. This is only used for display purposes.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The top-level authorization endpoint for the OpenID Connect identity provider. You can leave this blank if you provide the issuer field, which will be used to make a request to the OpenID Connect .well-known endpoint in order to dynamically resolve the authorization endpoint. If you provide an issuer then this field will be ignored.
     */
    public readonly oauth2AuthorizationEndpoint!: pulumi.Output<string | undefined>;
    /**
     * The client authentication method to use with the OpenID Connect identity provider.
     */
    public readonly oauth2ClientAuthenticationMethod!: pulumi.Output<string | undefined>;
    /**
     * The top-level client id for your Application.
     */
    public readonly oauth2ClientId!: pulumi.Output<string>;
    /**
     * The top-level client secret to use with the OpenID Connect identity provider.
     */
    public readonly oauth2ClientSecret!: pulumi.Output<string | undefined>;
    /**
     * An optional configuration to modify the expected name of the claim returned by the IdP that contains the email address.
     */
    public readonly oauth2EmailClaim!: pulumi.Output<string | undefined>;
    /**
     * The top-level issuer URI for the OpenID Connect identity provider. If this is provided, the authorization endpoint, token endpoint and userinfo endpoint will all be resolved using the issuer URI plus /.well-known/openid-configuration.
     */
    public readonly oauth2Issuer!: pulumi.Output<string | undefined>;
    /**
     * The top-level scope that you are requesting from the OpenID Connect identity provider.
     */
    public readonly oauth2Scope!: pulumi.Output<string | undefined>;
    /**
     * The top-level token endpoint for the OpenID Connect identity provider. You can leave this blank if you provide the issuer field, which will be used to make a request to the OpenID Connect .well-known endpoint in order to dynamically resolve the token endpoint. If you provide an issuer then this field will be ignored.
     */
    public readonly oauth2TokenEndpoint!: pulumi.Output<string | undefined>;
    /**
     * The top-level userinfo endpoint for the OpenID Connect identity provider. You can leave this blank if you provide the issuer field, which will be used to make a request to the OpenID Connect .well-known endpoint in order to dynamically resolve the userinfo endpoint. If you provide an issuer then this field will be ignored.
     */
    public readonly oauth2UserInfoEndpoint!: pulumi.Output<string | undefined>;
    /**
     * Set this value equal to true if you wish to use POST bindings with this OpenID Connect identity provider. The default
     * value of false means that a redirect binding which uses a GET request will be used.
     */
    public readonly postRequest!: pulumi.Output<boolean | undefined>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    public readonly tenantConfigurations!: pulumi.Output<outputs.FusionAuthIdpOpenIdConnectTenantConfiguration[] | undefined>;

    /**
     * Create a FusionAuthIdpOpenIdConnect resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthIdpOpenIdConnectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthIdpOpenIdConnectArgs | FusionAuthIdpOpenIdConnectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthIdpOpenIdConnectState | undefined;
            resourceInputs["applicationConfigurations"] = state ? state.applicationConfigurations : undefined;
            resourceInputs["buttonImageUrl"] = state ? state.buttonImageUrl : undefined;
            resourceInputs["buttonText"] = state ? state.buttonText : undefined;
            resourceInputs["debug"] = state ? state.debug : undefined;
            resourceInputs["domains"] = state ? state.domains : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["idpId"] = state ? state.idpId : undefined;
            resourceInputs["lambdaReconcileId"] = state ? state.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = state ? state.linkingStrategy : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["oauth2AuthorizationEndpoint"] = state ? state.oauth2AuthorizationEndpoint : undefined;
            resourceInputs["oauth2ClientAuthenticationMethod"] = state ? state.oauth2ClientAuthenticationMethod : undefined;
            resourceInputs["oauth2ClientId"] = state ? state.oauth2ClientId : undefined;
            resourceInputs["oauth2ClientSecret"] = state ? state.oauth2ClientSecret : undefined;
            resourceInputs["oauth2EmailClaim"] = state ? state.oauth2EmailClaim : undefined;
            resourceInputs["oauth2Issuer"] = state ? state.oauth2Issuer : undefined;
            resourceInputs["oauth2Scope"] = state ? state.oauth2Scope : undefined;
            resourceInputs["oauth2TokenEndpoint"] = state ? state.oauth2TokenEndpoint : undefined;
            resourceInputs["oauth2UserInfoEndpoint"] = state ? state.oauth2UserInfoEndpoint : undefined;
            resourceInputs["postRequest"] = state ? state.postRequest : undefined;
            resourceInputs["tenantConfigurations"] = state ? state.tenantConfigurations : undefined;
        } else {
            const args = argsOrState as FusionAuthIdpOpenIdConnectArgs | undefined;
            if ((!args || args.buttonText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'buttonText'");
            }
            if ((!args || args.oauth2ClientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oauth2ClientId'");
            }
            resourceInputs["applicationConfigurations"] = args ? args.applicationConfigurations : undefined;
            resourceInputs["buttonImageUrl"] = args ? args.buttonImageUrl : undefined;
            resourceInputs["buttonText"] = args ? args.buttonText : undefined;
            resourceInputs["debug"] = args ? args.debug : undefined;
            resourceInputs["domains"] = args ? args.domains : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["idpId"] = args ? args.idpId : undefined;
            resourceInputs["lambdaReconcileId"] = args ? args.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = args ? args.linkingStrategy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["oauth2AuthorizationEndpoint"] = args ? args.oauth2AuthorizationEndpoint : undefined;
            resourceInputs["oauth2ClientAuthenticationMethod"] = args ? args.oauth2ClientAuthenticationMethod : undefined;
            resourceInputs["oauth2ClientId"] = args ? args.oauth2ClientId : undefined;
            resourceInputs["oauth2ClientSecret"] = args ? args.oauth2ClientSecret : undefined;
            resourceInputs["oauth2EmailClaim"] = args ? args.oauth2EmailClaim : undefined;
            resourceInputs["oauth2Issuer"] = args ? args.oauth2Issuer : undefined;
            resourceInputs["oauth2Scope"] = args ? args.oauth2Scope : undefined;
            resourceInputs["oauth2TokenEndpoint"] = args ? args.oauth2TokenEndpoint : undefined;
            resourceInputs["oauth2UserInfoEndpoint"] = args ? args.oauth2UserInfoEndpoint : undefined;
            resourceInputs["postRequest"] = args ? args.postRequest : undefined;
            resourceInputs["tenantConfigurations"] = args ? args.tenantConfigurations : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FusionAuthIdpOpenIdConnect.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthIdpOpenIdConnect resources.
 */
export interface FusionAuthIdpOpenIdConnectState {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpOpenIdConnectApplicationConfiguration>[]>;
    /**
     * The top-level button image (URL) to use on the FusionAuth login page for this Identity Provider.
     */
    buttonImageUrl?: pulumi.Input<string>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    buttonText?: pulumi.Input<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
     */
    debug?: pulumi.Input<boolean>;
    /**
     * This is an optional list of domains that this OpenID Connect provider should be used for. This converts the FusionAuth login form to a domain-based login form. This type of form first asks the user for their email. FusionAuth then uses their email to determine if an OpenID Connect identity provider should be used. If an OpenID Connect provider should be used, the browser is redirected to the authorization endpoint of that identity provider. Otherwise, the password field is revealed on the form so that the user can login using FusionAuth.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
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
     * The name of this OpenID Connect identity provider. This is only used for display purposes.
     */
    name?: pulumi.Input<string>;
    /**
     * The top-level authorization endpoint for the OpenID Connect identity provider. You can leave this blank if you provide the issuer field, which will be used to make a request to the OpenID Connect .well-known endpoint in order to dynamically resolve the authorization endpoint. If you provide an issuer then this field will be ignored.
     */
    oauth2AuthorizationEndpoint?: pulumi.Input<string>;
    /**
     * The client authentication method to use with the OpenID Connect identity provider.
     */
    oauth2ClientAuthenticationMethod?: pulumi.Input<string>;
    /**
     * The top-level client id for your Application.
     */
    oauth2ClientId?: pulumi.Input<string>;
    /**
     * The top-level client secret to use with the OpenID Connect identity provider.
     */
    oauth2ClientSecret?: pulumi.Input<string>;
    /**
     * An optional configuration to modify the expected name of the claim returned by the IdP that contains the email address.
     */
    oauth2EmailClaim?: pulumi.Input<string>;
    /**
     * The top-level issuer URI for the OpenID Connect identity provider. If this is provided, the authorization endpoint, token endpoint and userinfo endpoint will all be resolved using the issuer URI plus /.well-known/openid-configuration.
     */
    oauth2Issuer?: pulumi.Input<string>;
    /**
     * The top-level scope that you are requesting from the OpenID Connect identity provider.
     */
    oauth2Scope?: pulumi.Input<string>;
    /**
     * The top-level token endpoint for the OpenID Connect identity provider. You can leave this blank if you provide the issuer field, which will be used to make a request to the OpenID Connect .well-known endpoint in order to dynamically resolve the token endpoint. If you provide an issuer then this field will be ignored.
     */
    oauth2TokenEndpoint?: pulumi.Input<string>;
    /**
     * The top-level userinfo endpoint for the OpenID Connect identity provider. You can leave this blank if you provide the issuer field, which will be used to make a request to the OpenID Connect .well-known endpoint in order to dynamically resolve the userinfo endpoint. If you provide an issuer then this field will be ignored.
     */
    oauth2UserInfoEndpoint?: pulumi.Input<string>;
    /**
     * Set this value equal to true if you wish to use POST bindings with this OpenID Connect identity provider. The default
     * value of false means that a redirect binding which uses a GET request will be used.
     */
    postRequest?: pulumi.Input<boolean>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpOpenIdConnectTenantConfiguration>[]>;
}

/**
 * The set of arguments for constructing a FusionAuthIdpOpenIdConnect resource.
 */
export interface FusionAuthIdpOpenIdConnectArgs {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpOpenIdConnectApplicationConfiguration>[]>;
    /**
     * The top-level button image (URL) to use on the FusionAuth login page for this Identity Provider.
     */
    buttonImageUrl?: pulumi.Input<string>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    buttonText: pulumi.Input<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
     */
    debug?: pulumi.Input<boolean>;
    /**
     * This is an optional list of domains that this OpenID Connect provider should be used for. This converts the FusionAuth login form to a domain-based login form. This type of form first asks the user for their email. FusionAuth then uses their email to determine if an OpenID Connect identity provider should be used. If an OpenID Connect provider should be used, the browser is redirected to the authorization endpoint of that identity provider. Otherwise, the password field is revealed on the form so that the user can login using FusionAuth.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
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
     * The name of this OpenID Connect identity provider. This is only used for display purposes.
     */
    name?: pulumi.Input<string>;
    /**
     * The top-level authorization endpoint for the OpenID Connect identity provider. You can leave this blank if you provide the issuer field, which will be used to make a request to the OpenID Connect .well-known endpoint in order to dynamically resolve the authorization endpoint. If you provide an issuer then this field will be ignored.
     */
    oauth2AuthorizationEndpoint?: pulumi.Input<string>;
    /**
     * The client authentication method to use with the OpenID Connect identity provider.
     */
    oauth2ClientAuthenticationMethod?: pulumi.Input<string>;
    /**
     * The top-level client id for your Application.
     */
    oauth2ClientId: pulumi.Input<string>;
    /**
     * The top-level client secret to use with the OpenID Connect identity provider.
     */
    oauth2ClientSecret?: pulumi.Input<string>;
    /**
     * An optional configuration to modify the expected name of the claim returned by the IdP that contains the email address.
     */
    oauth2EmailClaim?: pulumi.Input<string>;
    /**
     * The top-level issuer URI for the OpenID Connect identity provider. If this is provided, the authorization endpoint, token endpoint and userinfo endpoint will all be resolved using the issuer URI plus /.well-known/openid-configuration.
     */
    oauth2Issuer?: pulumi.Input<string>;
    /**
     * The top-level scope that you are requesting from the OpenID Connect identity provider.
     */
    oauth2Scope?: pulumi.Input<string>;
    /**
     * The top-level token endpoint for the OpenID Connect identity provider. You can leave this blank if you provide the issuer field, which will be used to make a request to the OpenID Connect .well-known endpoint in order to dynamically resolve the token endpoint. If you provide an issuer then this field will be ignored.
     */
    oauth2TokenEndpoint?: pulumi.Input<string>;
    /**
     * The top-level userinfo endpoint for the OpenID Connect identity provider. You can leave this blank if you provide the issuer field, which will be used to make a request to the OpenID Connect .well-known endpoint in order to dynamically resolve the userinfo endpoint. If you provide an issuer then this field will be ignored.
     */
    oauth2UserInfoEndpoint?: pulumi.Input<string>;
    /**
     * Set this value equal to true if you wish to use POST bindings with this OpenID Connect identity provider. The default
     * value of false means that a redirect binding which uses a GET request will be used.
     */
    postRequest?: pulumi.Input<boolean>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpOpenIdConnectTenantConfiguration>[]>;
}