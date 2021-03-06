package main

import (
	"github.com/covergates/covergates/config"
	"github.com/covergates/covergates/core"
	"github.com/covergates/covergates/modules/charts"
	"github.com/covergates/covergates/modules/git"
	"github.com/covergates/covergates/modules/hook"
	"github.com/covergates/covergates/modules/report"
	"github.com/covergates/covergates/modules/scm"
	"github.com/covergates/covergates/modules/session"
	"github.com/covergates/covergates/service/coverage"
	"github.com/google/wire"
)

var serviceSet = wire.NewSet(
	provideSCMService,
	provideSession,
	provideCoverageService,
	provideChartService,
	provideGit,
	provideReportService,
	provideHookService,
)

func provideSCMService(
	config *config.Config,
	userStore core.UserStore,
	git core.Git,
) core.SCMService {
	return &scm.Service{
		Config:    config,
		UserStore: userStore,
		Git:       git,
	}
}

func provideSession() core.Session {
	return &session.Session{}
}

func provideCoverageService() core.CoverageService {
	return &coverage.Service{}
}

func provideChartService() core.ChartService {
	return &charts.ChartService{}
}

func provideGit() core.Git {
	return &git.Service{}
}

func provideReportService() core.ReportService {
	return &report.Service{}
}

func provideHookService(
	SCM core.SCMService,
	RepoStore core.RepoStore,
	ReportStore core.ReportStore,
	ReportService core.ReportService,
) core.HookService {
	return &hook.Service{
		SCM:           SCM,
		RepoStore:     RepoStore,
		ReportService: ReportService,
		ReportStore:   ReportStore,
	}
}
