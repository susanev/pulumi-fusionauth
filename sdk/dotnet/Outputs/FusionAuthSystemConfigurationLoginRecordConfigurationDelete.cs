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
    public sealed class FusionAuthSystemConfigurationLoginRecordConfigurationDelete
    {
        /// <summary>
        /// Whether or not FusionAuth should delete the login records based upon this configuration. When true the loginRecordConfiguration.delete.numberOfDaysToRetain will be used to identify login records that are eligible for deletion. When this value is set to false login records will be preserved forever.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The number of days to retain login records.
        /// </summary>
        public readonly int? NumberOfDaysToRetain;

        [OutputConstructor]
        private FusionAuthSystemConfigurationLoginRecordConfigurationDelete(
            bool? enabled,

            int? numberOfDaysToRetain)
        {
            Enabled = enabled;
            NumberOfDaysToRetain = numberOfDaysToRetain;
        }
    }
}
