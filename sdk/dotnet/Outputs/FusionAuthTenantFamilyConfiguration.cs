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
    public sealed class FusionAuthTenantFamilyConfiguration
    {
        /// <summary>
        /// Whether to allow child registrations.
        /// </summary>
        public readonly bool? AllowChildRegistrations;
        /// <summary>
        /// The unique Id of the email template to use when confirming a child.
        /// </summary>
        public readonly string? ConfirmChildEmailTemplateId;
        /// <summary>
        /// Indicates that child users without parental verification will be permanently deleted after tenant.familyConfiguration.deleteOrphanedAccountsDays days.
        /// </summary>
        public readonly bool? DeleteOrphanedAccounts;
        /// <summary>
        /// The number of days from creation child users will be retained before being deleted for not completing parental verification. Value must be greater than 0.
        /// </summary>
        public readonly int? DeleteOrphanedAccountsDays;
        /// <summary>
        /// When true, FusionAuth will handle username collisions by generating a random suffix.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The unique Id of the email template to use when a family request is made.
        /// </summary>
        public readonly string? FamilyRequestEmailTemplateId;
        /// <summary>
        /// The maximum age of a child. Value must be greater than 0.
        /// </summary>
        public readonly int? MaximumChildAge;
        /// <summary>
        /// The minimum age to be an owner. Value must be greater than 0.
        /// </summary>
        public readonly int? MinimumOwnerAge;
        /// <summary>
        /// Whether a parent email is required.
        /// </summary>
        public readonly bool? ParentEmailRequired;
        /// <summary>
        /// The unique Id of the email template to use for parent registration.
        /// </summary>
        public readonly string? ParentRegistrationEmailTemplateId;

        [OutputConstructor]
        private FusionAuthTenantFamilyConfiguration(
            bool? allowChildRegistrations,

            string? confirmChildEmailTemplateId,

            bool? deleteOrphanedAccounts,

            int? deleteOrphanedAccountsDays,

            bool? enabled,

            string? familyRequestEmailTemplateId,

            int? maximumChildAge,

            int? minimumOwnerAge,

            bool? parentEmailRequired,

            string? parentRegistrationEmailTemplateId)
        {
            AllowChildRegistrations = allowChildRegistrations;
            ConfirmChildEmailTemplateId = confirmChildEmailTemplateId;
            DeleteOrphanedAccounts = deleteOrphanedAccounts;
            DeleteOrphanedAccountsDays = deleteOrphanedAccountsDays;
            Enabled = enabled;
            FamilyRequestEmailTemplateId = familyRequestEmailTemplateId;
            MaximumChildAge = maximumChildAge;
            MinimumOwnerAge = minimumOwnerAge;
            ParentEmailRequired = parentEmailRequired;
            ParentRegistrationEmailTemplateId = parentRegistrationEmailTemplateId;
        }
    }
}
