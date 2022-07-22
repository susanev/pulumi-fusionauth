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

    public sealed class FusionAuthTenantMultiFactorConfigurationEmailArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, FusionAuth will handle username collisions by generating a random suffix.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Id of the SMS template that is used when notifying a user to complete a multi-factor authentication request.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        public FusionAuthTenantMultiFactorConfigurationEmailArgs()
        {
        }
    }
}
