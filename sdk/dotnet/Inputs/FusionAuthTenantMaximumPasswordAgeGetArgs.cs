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

    public sealed class FusionAuthTenantMaximumPasswordAgeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password maximum age in days. The number of days after which FusionAuth will require a user to change their password. Required when systemConfiguration.maximumPasswordAge.enabled is set to true.
        /// </summary>
        [Input("days")]
        public Input<int>? Days { get; set; }

        /// <summary>
        /// When true, FusionAuth will handle username collisions by generating a random suffix.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public FusionAuthTenantMaximumPasswordAgeGetArgs()
        {
        }
        public static new FusionAuthTenantMaximumPasswordAgeGetArgs Empty => new FusionAuthTenantMaximumPasswordAgeGetArgs();
    }
}
