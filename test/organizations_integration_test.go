package test

import (
	"context"
	"log/slog"
	"os"
	"sort"
	"testing"

	"github.com/esnet/gdg/pkg/test_tooling/common"

	"github.com/esnet/gdg/internal/config"

	"github.com/esnet/gdg/internal/service"
	"github.com/esnet/gdg/pkg/test_tooling"
	"github.com/gosimple/slug"
	"github.com/grafana/grafana-openapi-client-go/models"
	"golang.org/x/exp/slices"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationCrud(t *testing.T) {
	if os.Getenv(test_tooling.EnableTokenTestsEnv) == test_tooling.FeatureEnabled {
		t.Skip("Skipping Token configuration, Organization CRUD requires Basic SecureData")
	}
	config.InitGdgConfig(common.DefaultTestConfig)
	var r *test_tooling.InitContainerResult
	err := Retry(context.Background(), DefaultRetryAttempts, func() error {
		r = test_tooling.InitTest(t, service.DefaultConfigProvider, nil)
		return r.Err
	})
	assert.NotNil(t, r)
	assert.NoError(t, err)
	defer func() {
		err := r.CleanUp()
		if err != nil {
			slog.Warn("Unable to clean up after test", "test", t.Name())
		}
	}()
	apiClient := r.ApiClient
	orgs := apiClient.ListOrganizations(service.NewOrganizationFilter(), true)
	assert.Equal(t, len(orgs), 1)
	mainOrg := orgs[0]
	assert.Equal(t, mainOrg.Organization.ID, int64(1))
	assert.Equal(t, mainOrg.Organization.Name, "Main Org.")
	newOrgs := apiClient.UploadOrganizations(service.NewOrganizationFilter())
	assert.Equal(t, len(newOrgs), 4)
	assert.True(t, slices.Contains(newOrgs, "DumbDumb"))
	assert.True(t, slices.Contains(newOrgs, "Moo"))
	assert.True(t, slices.Contains(newOrgs, "testing"))
	// Filter Org
	orgs = apiClient.ListOrganizations(service.NewOrganizationFilter("DumbDumb"), true)
	assert.Equal(t, len(orgs), 1)
	assert.Equal(t, orgs[0].Organization.Name, "DumbDumb")
}

func TestOrganizationUserMembership(t *testing.T) {
	if os.Getenv(test_tooling.EnableTokenTestsEnv) == "1" {
		t.Skip("Skipping Token configuration, Organization CRUD requires Basic SecureData")
	}
	config.InitGdgConfig(common.DefaultTestConfig)
	var r *test_tooling.InitContainerResult
	err := Retry(context.Background(), DefaultRetryAttempts, func() error {
		r = test_tooling.InitTest(t, service.DefaultConfigProvider, nil)
		return r.Err
	})
	assert.NotNil(t, r)
	assert.NoError(t, err)
	defer func() {
		err := r.CleanUp()
		if err != nil {
			slog.Warn("Unable to clean up after test", "test", t.Name())
		}
	}()
	apiClient := r.ApiClient
	// Create Orgs in case they aren't already present.
	apiClient.UploadOrganizations(service.NewOrganizationFilter())
	orgs := apiClient.ListOrganizations(service.NewOrganizationFilter(), true)
	sort.Slice(orgs, func(a, b int) bool {
		return orgs[a].Organization.ID < orgs[b].Organization.ID
	})
	newOrg := orgs[2]
	// Create Users in case they aren't already present.
	apiClient.UploadUsers(service.NewUserFilter(""))
	// get users
	users := apiClient.ListUsers(service.NewUserFilter(""))
	assert.Equal(t, len(users), 3)
	var orgUser *models.UserSearchHitDTO
	for _, u := range users {
		if u.Login == "tux" {
			orgUser = u
			break
		}
	}
	assert.NotNil(t, orgUser)
	// Reset if any state exists.
	err = apiClient.DeleteUserFromOrg(slug.Make(newOrg.Organization.Name), orgUser.ID)
	assert.Nil(t, err)
	// Start CRUD test
	orgUsers := apiClient.ListOrgUsers(newOrg.Organization.ID)
	assert.Equal(t, len(orgUsers), 1)
	assert.Equal(t, orgUsers[0].Login, "admin")
	assert.Equal(t, orgUsers[0].Role, "Admin")

	err = apiClient.AddUserToOrg("Admin", slug.Make(newOrg.Organization.Name), orgUser.ID)
	assert.Nil(t, err)
	orgUsers = apiClient.ListOrgUsers(newOrg.Organization.ID)
	assert.Equal(t, len(orgUsers), 2)
	assert.Equal(t, orgUsers[1].Role, "Admin")
	err = apiClient.UpdateUserInOrg("Viewer", slug.Make(newOrg.Organization.Name), orgUser.ID)
	orgUsers = apiClient.ListOrgUsers(newOrg.Organization.ID)
	assert.Nil(t, err)
	assert.Equal(t, orgUsers[1].Role, "Viewer")
	err = apiClient.DeleteUserFromOrg(slug.Make(newOrg.Organization.Name), orgUser.ID)
	orgUsers = apiClient.ListOrgUsers(newOrg.Organization.ID)
	assert.Equal(t, len(orgUsers), 1)
	assert.Nil(t, err)
}

func TestOrganizationProperties(t *testing.T) {
	if os.Getenv(test_tooling.EnableTokenTestsEnv) == test_tooling.FeatureEnabled {
		t.Skip("Skipping Token configuration, Organization CRUD requires Basic SecureData")
	}
	config.InitGdgConfig(common.DefaultTestConfig)
	var r *test_tooling.InitContainerResult
	err := Retry(context.Background(), DefaultRetryAttempts, func() error {
		r = test_tooling.InitTest(t, service.DefaultConfigProvider, nil)
		return r.Err
	})
	assert.NotNil(t, r)
	assert.NoError(t, err)
	defer func() {
		err := r.CleanUp()
		if err != nil {
			slog.Warn("Unable to clean up after test", "test", t.Name())
		}
	}()
	apiClient := r.ApiClient
	_, err = apiClient.UploadDashboards(service.NewDashboardFilter("", "", ""))
	assert.NoError(t, err)
	prefs, err := apiClient.GetOrgPreferences("Main Org.")
	assert.Nil(t, err)
	prefs.HomeDashboardUID = "000000003"
	prefs.Theme = "dark"
	prefs.WeekStart = "Saturday"
	err = apiClient.UploadOrgPreferences("Main Org.", prefs)
	assert.Nil(t, err)
	prefs, err = apiClient.GetOrgPreferences("Main Org.")
	assert.Nil(t, err)
	assert.NotNil(t, prefs)
	assert.Equal(t, prefs.Theme, "dark")
	assert.Equal(t, prefs.WeekStart, "Saturday")
	assert.Equal(t, prefs.HomeDashboardUID, "000000003")
	apiClient.DeleteAllDashboards(service.NewDashboardFilter("", "", ""))
}
