package services

type LoggerService interface {
	LogActivity(uuid string, action string, statusCode string, message string)
	AnalyzeActivities()
}
