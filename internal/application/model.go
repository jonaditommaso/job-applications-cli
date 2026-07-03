package application

type ApplicationStatus string

const (
	Applied       ApplicationStatus = "applied"
	Interview     ApplicationStatus = "interview"
	TechnicalTest ApplicationStatus = "technical_test"
	Offer         ApplicationStatus = "offer"
	Rejected      ApplicationStatus = "rejected"
	Ghosted       ApplicationStatus = "ghosted"
)

type WorkMode string

const (
	Remote    WorkMode = "remote"
	Hybrid    WorkMode = "hybrid"
	Worldwide WorkMode = "worldwide"
	EUonly    WorkMode = "eu_only"
)

type JobApplication struct {
	Id                int
	Company           string
	Role              string
	Status            ApplicationStatus
	ApplicationDate   *string
	Notes             *string
	CreatedAt         string
	ApplicationViewed bool
	Contacted         bool
	Channel           string
	OtherCandidates   *int
	WorkMode          *WorkMode
	CvInEnglish       *bool
	EnglishRequired   *bool
	SalaryMinCompany  *float64
	SalaryMaxCompany  *float64
	SalaryExpectation *float64
	Technologies      []string
	TechnologiesNice  []string
	Country           *string
	City              *string
}
