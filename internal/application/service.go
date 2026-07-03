package application

import "fmt"

var currentApplications []JobApplication

func InitializeApplications(applications []JobApplication) {
	currentApplications = append([]JobApplication(nil), applications...)
}

func ListApplications() {
	for _, app := range currentApplications {
		fmt.Printf(
			"Company: %s, Role: %s, Status: %s\n",
			app.Company,
			app.Role,
			app.Status,
		)
	}
}

func GetApplication(appID int) (*JobApplication, error) {
	for i := range currentApplications {
		if currentApplications[i].Id == appID {
			return &currentApplications[i], nil
		}
	}
	return nil, fmt.Errorf("application not found")
}

func AddApplication(newApp JobApplication) {
	currentApplications = append(currentApplications, newApp)
}

func UpdateApplication(appID int, field string, value any) {
	for i := range currentApplications {
		if currentApplications[i].Id == appID {
			switch field {
			case "Company":
				currentApplications[i].Company = value.(string)
			case "Role":
				currentApplications[i].Role = value.(string)
			case "Status":
				currentApplications[i].Status = value.(ApplicationStatus)
			}
			return
		}
	}
}
