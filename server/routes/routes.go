package routes

import (
	"github.com/labstack/echo/v4/middleware"
	"kry-go/server"
	"kry-go/server/handlers"
)

func InitRouter(server *server.Server) {
	contestHandler := handlers.MakeContestHandler(server)
	problemHandler := handlers.MakeProblemHandler(server)
	submissionHandler := handlers.MakeSubmissionHandler(server)
	tagHandler := handlers.MakeTagHandler(server)
	testcaseHandler := handlers.MakeTestcaseHandler(server)
	userHandler := handlers.MakeUserHandler(server)

	server.Echo.Use(middleware.Logger())

	g := server.Echo.Group("/api")

	// TODO: 어떤 URL에 어떤 Resource를 제공할 것인가
	g.GET("/contest/:contestId", contestHandler.GetContest)
	g.GET("/problem/:problemId", problemHandler.GetProblem)
	g.GET("/submission/:problemId", submissionHandler.GetSubmission)
	g.GET("/tag/:problemId", tagHandler.GetTag)
	g.GET("/testcase/:problemId", testcaseHandler.GetTestcase)
	g.GET("/user/:userId", userHandler.GetUser)
}
