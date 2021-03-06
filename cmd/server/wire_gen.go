// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/covergates/covergates/config"
	"github.com/jinzhu/gorm"
)

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Injectors from wire.go:

func InitializeApplication(config2 *config.Config, db *gorm.DB) (application, error) {
	session := provideSession()
	loginMiddleware := provideLogin(config2)
	databaseService := provideDatabaseService(db)
	userStore := provideUserStore(databaseService)
	git := provideGit()
	scmService := provideSCMService(config2, userStore, git)
	coverageService := provideCoverageService()
	chartService := provideChartService()
	reportService := provideReportService()
	repoStore := provideRepoStore(databaseService)
	reportStore := provideReportStore(databaseService)
	hookService := provideHookService(scmService, repoStore, reportStore, reportService)
	routers := provideRouter(session, config2, loginMiddleware, scmService, coverageService, chartService, reportService, hookService, reportStore, repoStore)
	mainApplication := newApplication(routers, databaseService)
	return mainApplication, nil
}
