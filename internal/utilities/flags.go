/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package utilities

import (
	"strings"

	"github.com/kubefirst/kubefirst/internal/progress"
	"github.com/kubefirst/kubefirst/internal/types"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetFlags(cmd *cobra.Command, cloudProvider string) (types.CliFlags, error) {
	cliFlags := types.CliFlags{}
	alertsEmailFlag, err := cmd.Flags().GetString("alerts-email")
	if err != nil {
		progress.Error(err.Error())
		return cliFlags, err
	}

	cloudRegionFlag, err := cmd.Flags().GetString("cloud-region")
	if err != nil {
		progress.Error(err.Error())
		return cliFlags, err
	}

	clusterNameFlag, err := cmd.Flags().GetString("cluster-name")
	if err != nil {
		progress.Error(err.Error())
		return cliFlags, err
	}

	dnsProviderFlag, err := cmd.Flags().GetString("dns-provider")
	if err != nil {
		progress.Error(err.Error())
		return cliFlags, err
	}

	domainNameFlag, err := cmd.Flags().GetString("domain-name")
	if err != nil {
		progress.Error(err.Error())
		return cliFlags, err
	}

	githubOrgFlag, err := cmd.Flags().GetString("github-org")
	if err != nil {
		progress.Error(err.Error())
		return cliFlags, err
	}
	githubOrgFlag = strings.ToLower(githubOrgFlag)

	gitlabGroupFlag, err := cmd.Flags().GetString("gitlab-group")
	if err != nil {
		progress.Error(err.Error())
		return cliFlags, err
	}
	gitlabGroupFlag = strings.ToLower(gitlabGroupFlag)

	gitProviderFlag, err := cmd.Flags().GetString("git-provider")
	if err != nil {
		progress.Error(err.Error())
		return cliFlags, err
	}

	gitProtocolFlag, err := cmd.Flags().GetString("git-protocol")
	if err != nil {
		progress.Error(err.Error())
		return cliFlags, err
	}

	gitopsTemplateURLFlag, err := cmd.Flags().GetString("gitops-template-url")
	if err != nil {
		progress.Error(err.Error())
		return cliFlags, err
	}

	gitopsTemplateBranchFlag, err := cmd.Flags().GetString("gitops-template-branch")
	if err != nil {
		progress.Error(err.Error())
		return cliFlags, err
	}

	useTelemetryFlag, err := cmd.Flags().GetBool("use-telemetry")
	if err != nil {
		progress.Error(err.Error())
		return cliFlags, err
	}

	if cloudProvider == "aws" {
		ecrFlag, err := cmd.Flags().GetBool("ecr")
		if err != nil {
			progress.Error(err.Error())
			return cliFlags, err
		}

		cliFlags.Ecr = ecrFlag
	}

	if cloudProvider == "google" {
		googleProject, err := cmd.Flags().GetString("google-project")
		if err != nil {
			progress.Error(err.Error())
			return cliFlags, err
		}

		cliFlags.GoogleProject = googleProject
	}

	cliFlags.AlertsEmail = alertsEmailFlag
	cliFlags.CloudRegion = cloudRegionFlag
	cliFlags.ClusterName = clusterNameFlag
	cliFlags.DnsProvider = dnsProviderFlag
	cliFlags.DomainName = domainNameFlag
	cliFlags.GitProtocol = gitProtocolFlag
	cliFlags.GitProvider = gitProviderFlag
	cliFlags.GithubOrg = githubOrgFlag
	cliFlags.GitlabGroup = gitlabGroupFlag
	cliFlags.GitopsTemplateBranch = gitopsTemplateBranchFlag
	cliFlags.GitopsTemplateURL = gitopsTemplateURLFlag
	cliFlags.UseTelemetry = useTelemetryFlag
	cliFlags.CloudProvider = cloudProvider

	viper.Set("flags.alerts-email", cliFlags.AlertsEmail)
	viper.Set("flags.cluster-name", cliFlags.ClusterName)
	viper.Set("flags.dns-provider", cliFlags.DnsProvider)
	viper.Set("flags.domain-name", cliFlags.DomainName)
	viper.Set("flags.git-provider", cliFlags.GitProvider)
	viper.Set("flags.git-protocol", cliFlags.GitProtocol)
	viper.Set("flags.cloud-region", cliFlags.CloudRegion)
	viper.Set("kubefirst.cloud-provider", cloudProvider)
	viper.WriteConfig()

	return cliFlags, nil
}
