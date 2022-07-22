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

    public sealed class FusionAuthIdpExternalJwtApplicationConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the Application to apply this configuration to.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

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

        public FusionAuthIdpExternalJwtApplicationConfigurationArgs()
        {
        }
    }
}
