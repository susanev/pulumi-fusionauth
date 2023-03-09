// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth.Inputs
{

    public sealed class FusionAuthApplicationSamlv2ConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The audience for the SAML response sent to back to the service provider from FusionAuth. Some service providers require different audience values than the issuer and this configuration option lets you change the audience in the response.
        /// </summary>
        [Input("audience")]
        public Input<string>? Audience { get; set; }

        [Input("authorizedRedirectUrls", required: true)]
        private InputList<string>? _authorizedRedirectUrls;

        /// <summary>
        /// An array of URLs that are the authorized redirect URLs for FusionAuth OAuth.
        /// </summary>
        public InputList<string> AuthorizedRedirectUrls
        {
            get => _authorizedRedirectUrls ?? (_authorizedRedirectUrls = new InputList<string>());
            set => _authorizedRedirectUrls = value;
        }

        /// <summary>
        /// The URL of the callback (sometimes called the Assertion Consumer Service or ACS). This is where FusionAuth sends the browser after the user logs in via SAML.
        /// </summary>
        [Input("callbackUrl")]
        public Input<string>? CallbackUrl { get; set; }

        /// <summary>
        /// Whether or not FusionAuth will log SAML debug messages to the event log. This is useful for debugging purposes.
        /// </summary>
        [Input("debug")]
        public Input<bool>? Debug { get; set; }

        /// <summary>
        /// The unique Id of the Key used to verify the signature if the public key cannot be determined by the KeyInfo element when using POST bindings, or the key used to verify the signature when using HTTP Redirect bindings.
        /// </summary>
        [Input("defaultVerificationKeyId")]
        public Input<string>? DefaultVerificationKeyId { get; set; }

        /// <summary>
        /// Whether or not SAML Single Logout for this SAML IdP is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The issuer that identifies the service provider and allows FusionAuth to load the correct Application and SAML configuration. If you don’t know the issuer, you can often times put in anything here and FusionAuth will display an error message with the issuer from the service provider when you test the SAML login.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        /// <summary>
        /// The unique Id of the Key used to sign the SAML Single Logout response.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        [Input("logout")]
        public Input<Inputs.FusionAuthApplicationSamlv2ConfigurationLogoutArgs>? Logout { get; set; }

        /// <summary>
        /// The URL that the browser is taken to after the user logs out of the SAML service provider. Often service providers need this URL in order to correctly hook up single-logout. Note that FusionAuth does not support the SAML single-logout profile because most service providers to not support it properly.
        /// </summary>
        [Input("logoutUrl")]
        public Input<string>? LogoutUrl { get; set; }

        /// <summary>
        /// If set to true, will force verification through the key store.
        /// </summary>
        [Input("requiredSignedRequests")]
        public Input<bool>? RequiredSignedRequests { get; set; }

        /// <summary>
        /// The XML signature canonicalization method used when digesting and signing the SAML response. Unfortunately, many service providers do not correctly implement the XML signature specifications and force a specific canonicalization method. This setting allows you to change the canonicalization method to match the service provider. Often, service providers don’t even document their required method. You might need to contact enterprise support at the service provider to figure out what method they use.
        /// </summary>
        [Input("xmlSignatureCanonicalizationMethod")]
        public Input<string>? XmlSignatureCanonicalizationMethod { get; set; }

        /// <summary>
        /// The location to place the XML signature when signing a successful SAML response.
        /// </summary>
        [Input("xmlSignatureLocation")]
        public Input<string>? XmlSignatureLocation { get; set; }

        public FusionAuthApplicationSamlv2ConfigurationArgs()
        {
        }
        public static new FusionAuthApplicationSamlv2ConfigurationArgs Empty => new FusionAuthApplicationSamlv2ConfigurationArgs();
    }
}
