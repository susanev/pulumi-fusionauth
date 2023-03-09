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

    public sealed class FusionAuthIdpAppleApplicationConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the Application to apply this configuration to.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// The top-level button text to use on the FusionAuth login page for this Identity Provider.
        /// </summary>
        [Input("buttonText")]
        public Input<string>? ButtonText { get; set; }

        /// <summary>
        /// Determines if a UserRegistration is created for the User automatically or not. If a user doesn’t exist in FusionAuth and logs in through an identity provider, this boolean controls whether or not FusionAuth creates a registration for the User in the Application they are logging into.
        /// </summary>
        [Input("createRegistration")]
        public Input<bool>? CreateRegistration { get; set; }

        /// <summary>
        /// Determines if this provider is enabled. If it is false then it will be disabled globally.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The top-level space separated scope that you are requesting from Apple.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
        /// </summary>
        [Input("servicesId")]
        public Input<string>? ServicesId { get; set; }

        /// <summary>
        /// The Apple App ID Prefix, or Team ID found in your Apple Developer Account which has been configured for Sign in with Apple.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public FusionAuthIdpAppleApplicationConfigurationGetArgs()
        {
        }
        public static new FusionAuthIdpAppleApplicationConfigurationGetArgs Empty => new FusionAuthIdpAppleApplicationConfigurationGetArgs();
    }
}
