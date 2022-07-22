// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # SAML v2 Identity Provider Resource
 *
 * SAML v2 identity providers connect to external SAML v2 login systems. This type of login will optionally provide a Login with …​ button on FusionAuth’s login page. This button is customizable by using different properties of the identity provider.
 *
 * Optionally, this identity provider can define one or more domains it is associated with. This is useful for allowing employees to log in with their corporate credentials. As long as the company has an identity solution that provides SAML v2, you can leverage this feature. This is referred to as a Domain Based Identity Provider. If you enable domains for an identity provider, the Login with …​ button will not be displayed. Instead, only the email form field will be displayed initially on the FusionAuth login page. Once the user types in their email address, FusionAuth will determine if the user is logging in locally or if they should be redirected to this identity provider. This is determined by extracting the domain from their email address and comparing it to the domains associated with the identity provider.
 *
 * FusionAuth will locate the user’s email address in the SAML assertion which will be used to create or lookup the existing user. Additional claims from the SAML response can be used to reconcile the User to FusionAuth by using a SAML v2 Reconcile Lambda. Unless you assign a reconcile lambda to this provider, on the email address will be used from the available assertions returned by the SAML v2 identity provider.
 *
 * [SAML v2 Connect Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/samlv2/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi_fusionauth from "@theogravity/pulumi-fusionauth";
 *
 * const saml = new fusionauth.FusionAuthIdpSamlv2("saml", {
 *     applicationConfigurations: [{
 *         applicationId: fusionauth_application.myapp.id,
 *         buttonText: "Login with SAML (app text)",
 *         createRegistration: true,
 *         enabled: true,
 *     }],
 *     buttonText: "Login with SAML",
 *     debug: false,
 *     emailClaim: "email",
 *     idpEndpoint: "https://www.example.com/login",
 *     postRequest: true,
 *     requestSigningKey: "3168129b-91fa-46f4-9676-947f5509fdce",
 *     signRequest: true,
 *     useNameForEmail: true,
 * });
 * ```
 */
export class FusionAuthIdpSamlv2 extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthIdpSamlv2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthIdpSamlv2State, opts?: pulumi.CustomResourceOptions): FusionAuthIdpSamlv2 {
        return new FusionAuthIdpSamlv2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthIdpSamlv2:FusionAuthIdpSamlv2';

    /**
     * Returns true if the given object is an instance of FusionAuthIdpSamlv2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthIdpSamlv2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthIdpSamlv2.__pulumiType;
    }

    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    public readonly applicationConfigurations!: pulumi.Output<outputs.FusionAuthIdpSamlv2ApplicationConfiguration[] | undefined>;
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
     * The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely identity the user. If this is not set, the `useNameForEmail` flag must be true.
     */
    public readonly emailClaim!: pulumi.Output<string | undefined>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The SAML v2 login page of the identity provider.
     */
    public readonly idpEndpoint!: pulumi.Output<string | undefined>;
    /**
     * The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
     */
    public readonly idpId!: pulumi.Output<string | undefined>;
    /**
     * The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the identity provider. This key must be a verification only key or certificate (meaning that it only has a public key component).
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
     * The name of this OpenID Connect identity provider. This is only used for display purposes.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Either urn:oasis:names:tc:SAML:2.0:nameid-format:persistent or urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress depending on which NameId format you wish to use.
     */
    public readonly nameIdFormat!: pulumi.Output<string>;
    /**
     * Set this value equal to true if you wish to use POST bindings with this OpenID Connect identity provider. The default value of false means that a redirect binding which uses a GET request will be used.
     */
    public readonly postRequest!: pulumi.Output<boolean | undefined>;
    /**
     * TThe key pair Id to use to sign the SAML request. Required when `signRequest` is true.
     */
    public readonly requestSigningKey!: pulumi.Output<string | undefined>;
    /**
     * When true authentication requests sent to the identity provider will be signed.
     */
    public readonly signRequest!: pulumi.Output<boolean | undefined>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    public readonly tenantConfigurations!: pulumi.Output<outputs.FusionAuthIdpSamlv2TenantConfiguration[] | undefined>;
    /**
     * Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation processing. If this is false, then the `emailClaim` property must be set.
     */
    public readonly useNameForEmail!: pulumi.Output<boolean | undefined>;
    /**
     * The XML signature canonicalization method used when digesting and signing the SAML request.
     */
    public readonly xmlSignatureCanonicalizationMethod!: pulumi.Output<string | undefined>;

    /**
     * Create a FusionAuthIdpSamlv2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthIdpSamlv2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthIdpSamlv2Args | FusionAuthIdpSamlv2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthIdpSamlv2State | undefined;
            resourceInputs["applicationConfigurations"] = state ? state.applicationConfigurations : undefined;
            resourceInputs["buttonImageUrl"] = state ? state.buttonImageUrl : undefined;
            resourceInputs["buttonText"] = state ? state.buttonText : undefined;
            resourceInputs["debug"] = state ? state.debug : undefined;
            resourceInputs["domains"] = state ? state.domains : undefined;
            resourceInputs["emailClaim"] = state ? state.emailClaim : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["idpEndpoint"] = state ? state.idpEndpoint : undefined;
            resourceInputs["idpId"] = state ? state.idpId : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["lambdaReconcileId"] = state ? state.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = state ? state.linkingStrategy : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nameIdFormat"] = state ? state.nameIdFormat : undefined;
            resourceInputs["postRequest"] = state ? state.postRequest : undefined;
            resourceInputs["requestSigningKey"] = state ? state.requestSigningKey : undefined;
            resourceInputs["signRequest"] = state ? state.signRequest : undefined;
            resourceInputs["tenantConfigurations"] = state ? state.tenantConfigurations : undefined;
            resourceInputs["useNameForEmail"] = state ? state.useNameForEmail : undefined;
            resourceInputs["xmlSignatureCanonicalizationMethod"] = state ? state.xmlSignatureCanonicalizationMethod : undefined;
        } else {
            const args = argsOrState as FusionAuthIdpSamlv2Args | undefined;
            if ((!args || args.buttonText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'buttonText'");
            }
            if ((!args || args.keyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyId'");
            }
            resourceInputs["applicationConfigurations"] = args ? args.applicationConfigurations : undefined;
            resourceInputs["buttonImageUrl"] = args ? args.buttonImageUrl : undefined;
            resourceInputs["buttonText"] = args ? args.buttonText : undefined;
            resourceInputs["debug"] = args ? args.debug : undefined;
            resourceInputs["domains"] = args ? args.domains : undefined;
            resourceInputs["emailClaim"] = args ? args.emailClaim : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["idpEndpoint"] = args ? args.idpEndpoint : undefined;
            resourceInputs["idpId"] = args ? args.idpId : undefined;
            resourceInputs["keyId"] = args ? args.keyId : undefined;
            resourceInputs["lambdaReconcileId"] = args ? args.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = args ? args.linkingStrategy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nameIdFormat"] = args ? args.nameIdFormat : undefined;
            resourceInputs["postRequest"] = args ? args.postRequest : undefined;
            resourceInputs["requestSigningKey"] = args ? args.requestSigningKey : undefined;
            resourceInputs["signRequest"] = args ? args.signRequest : undefined;
            resourceInputs["tenantConfigurations"] = args ? args.tenantConfigurations : undefined;
            resourceInputs["useNameForEmail"] = args ? args.useNameForEmail : undefined;
            resourceInputs["xmlSignatureCanonicalizationMethod"] = args ? args.xmlSignatureCanonicalizationMethod : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FusionAuthIdpSamlv2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthIdpSamlv2 resources.
 */
export interface FusionAuthIdpSamlv2State {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpSamlv2ApplicationConfiguration>[]>;
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
     * The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely identity the user. If this is not set, the `useNameForEmail` flag must be true.
     */
    emailClaim?: pulumi.Input<string>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The SAML v2 login page of the identity provider.
     */
    idpEndpoint?: pulumi.Input<string>;
    /**
     * The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
     */
    idpId?: pulumi.Input<string>;
    /**
     * The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the identity provider. This key must be a verification only key or certificate (meaning that it only has a public key component).
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
     * The name of this OpenID Connect identity provider. This is only used for display purposes.
     */
    name?: pulumi.Input<string>;
    /**
     * Either urn:oasis:names:tc:SAML:2.0:nameid-format:persistent or urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress depending on which NameId format you wish to use.
     */
    nameIdFormat?: pulumi.Input<string>;
    /**
     * Set this value equal to true if you wish to use POST bindings with this OpenID Connect identity provider. The default value of false means that a redirect binding which uses a GET request will be used.
     */
    postRequest?: pulumi.Input<boolean>;
    /**
     * TThe key pair Id to use to sign the SAML request. Required when `signRequest` is true.
     */
    requestSigningKey?: pulumi.Input<string>;
    /**
     * When true authentication requests sent to the identity provider will be signed.
     */
    signRequest?: pulumi.Input<boolean>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpSamlv2TenantConfiguration>[]>;
    /**
     * Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation processing. If this is false, then the `emailClaim` property must be set.
     */
    useNameForEmail?: pulumi.Input<boolean>;
    /**
     * The XML signature canonicalization method used when digesting and signing the SAML request.
     */
    xmlSignatureCanonicalizationMethod?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FusionAuthIdpSamlv2 resource.
 */
export interface FusionAuthIdpSamlv2Args {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpSamlv2ApplicationConfiguration>[]>;
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
     * The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely identity the user. If this is not set, the `useNameForEmail` flag must be true.
     */
    emailClaim?: pulumi.Input<string>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The SAML v2 login page of the identity provider.
     */
    idpEndpoint?: pulumi.Input<string>;
    /**
     * The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
     */
    idpId?: pulumi.Input<string>;
    /**
     * The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the identity provider. This key must be a verification only key or certificate (meaning that it only has a public key component).
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
     * The name of this OpenID Connect identity provider. This is only used for display purposes.
     */
    name?: pulumi.Input<string>;
    /**
     * Either urn:oasis:names:tc:SAML:2.0:nameid-format:persistent or urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress depending on which NameId format you wish to use.
     */
    nameIdFormat?: pulumi.Input<string>;
    /**
     * Set this value equal to true if you wish to use POST bindings with this OpenID Connect identity provider. The default value of false means that a redirect binding which uses a GET request will be used.
     */
    postRequest?: pulumi.Input<boolean>;
    /**
     * TThe key pair Id to use to sign the SAML request. Required when `signRequest` is true.
     */
    requestSigningKey?: pulumi.Input<string>;
    /**
     * When true authentication requests sent to the identity provider will be signed.
     */
    signRequest?: pulumi.Input<boolean>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpSamlv2TenantConfiguration>[]>;
    /**
     * Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation processing. If this is false, then the `emailClaim` property must be set.
     */
    useNameForEmail?: pulumi.Input<boolean>;
    /**
     * The XML signature canonicalization method used when digesting and signing the SAML request.
     */
    xmlSignatureCanonicalizationMethod?: pulumi.Input<string>;
}
