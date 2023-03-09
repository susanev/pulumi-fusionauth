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

    public sealed class FusionAuthApplicationRegistrationConfigurationMobilePhoneGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not SAML Single Logout for this SAML IdP is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("required")]
        public Input<bool>? Required { get; set; }

        public FusionAuthApplicationRegistrationConfigurationMobilePhoneGetArgs()
        {
        }
        public static new FusionAuthApplicationRegistrationConfigurationMobilePhoneGetArgs Empty => new FusionAuthApplicationRegistrationConfigurationMobilePhoneGetArgs();
    }
}
