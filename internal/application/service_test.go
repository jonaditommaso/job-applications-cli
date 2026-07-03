package application

import "testing"

func TestListApplicationsUsesCurrentApplications(t *testing.T) {
	InitializeApplications([]JobApplication{{
		Id:      1,
		Company: "Example Corp",
		Role:    "Software Engineer",
		Status:  Applied,
	}})

	ListApplications()
}

func TestSharedStatePersistsAcrossOperations(t *testing.T) {
	InitializeApplications([]JobApplication{{
		Id:      1,
		Company: "Example Corp",
		Role:    "Software Engineer",
		Status:  Applied,
	}})

	AddApplication(JobApplication{
		Id:      2,
		Company: "Acme",
		Role:    "Platform Engineer",
		Status:  Interview,
	})

	UpdateApplication(2, "Role", "Senior Platform Engineer")
	UpdateApplication(2, "Status", Offer)

	if len(currentApplications) != 2 {
		t.Fatalf("expected 2 applications, got %d", len(currentApplications))
	}

	if currentApplications[1].Role != "Senior Platform Engineer" {
		t.Fatalf("expected role to be updated, got %q", currentApplications[1].Role)
	}

	if currentApplications[1].Status != Offer {
		t.Fatalf("expected status to be updated, got %q", currentApplications[1].Status)
	}
}
