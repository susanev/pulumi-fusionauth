// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth.Outputs
{

    [OutputType]
    public sealed class FusionAuthTenantMinimumPasswordAge
    {
        /// <summary>
        /// When true, FusionAuth will handle username collisions by generating a random suffix.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The password minimum age in seconds. When enabled FusionAuth will not allow a password to be changed until it reaches this minimum age. Required when systemConfiguration.minimumPasswordAge.enabled is set to true.
        /// </summary>
        public readonly int? Seconds;

        [OutputConstructor]
        private FusionAuthTenantMinimumPasswordAge(
            bool? enabled,

            int? seconds)
        {
            Enabled = enabled;
            Seconds = seconds;
        }
    }
}
