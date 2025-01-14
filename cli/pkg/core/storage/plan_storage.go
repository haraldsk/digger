package storage

type PlanStorage interface {
	StorePlan(localPlanFilePath string, storedPlanFilePath string) error
	StorePlanFile(fileContents []byte, artifactName string, fileName string) error
	RetrievePlan(localPlanFilePath string, storedPlanFilePath string) (*string, error)
	DeleteStoredPlan(storedPlanFilePath string) error
	PlanExists(storedPlanFilePath string) (bool, error)
}
