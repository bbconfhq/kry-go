package routes

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"kry-go/base"
	_ "kry-go/docs"
	"kry-go/routes/handler"
	"os"
	"strings"
)

func InitRouter(server *base.Server) {
	contestHandler := handler.MakeContestHandler(server)
	problemHandler := handler.MakeProblemHandler(server)
	submissionHandler := handler.MakeSubmissionHandler(server)
	tagHandler := handler.MakeTagHandler(server)
	testcaseHandler := handler.MakeTestcaseHandler(server)
	userHandler := handler.MakeUserHandler(server)
	loginHandler := handler.MakeLoginHandler(server)

	// TODO: Consider using redis-session
	sessionSecret := os.Getenv("SESSION_SECRET")
	server.Echo.Use(session.Middleware(sessions.NewCookieStore([]byte(sessionSecret))))
	server.Echo.Use(middleware.Logger())

	if strings.ToLower(os.Getenv("DEBUG")) == "true" {
		server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	g := server.Echo

	g.GET("/login/:loginType", loginHandler.LoginUser)
	g.POST("/login/:loginType/callback", loginHandler.LoginOrRegisterUser)

	// ContestHandler
	g.GET("/contest", contestHandler.GetContests)
	g.POST("/contest", contestHandler.PostContest)
	g.GET("/contest/:contestId", contestHandler.GetContest)
	g.PUT("/contest/:contestId", contestHandler.PutContest)
	g.DELETE("/contest/:contestId", contestHandler.DeleteContest)

	// ProblemHandler
	g.GET("/problem", problemHandler.GetProblems)
	g.POST("/problem/:problemId", problemHandler.PostProblem)
	g.GET("/problem/:problemId", problemHandler.GetProblem)
	g.PUT("/problem/:problemId", problemHandler.PutProblem)
	g.DELETE("/problem/:problemId", problemHandler.DeleteProblem)

	// TODO: Make decision how to handle submissions
	// SubmissionHandler
	g.GET("/submission", submissionHandler.GetSubmissions)
	g.POST("/submission/:submissionId", submissionHandler.PostSubmission)
	g.GET("/submission/:submissionId", submissionHandler.GetSubmission)
	//g.PUT("/submission/:submissionId", submissionHandler.PutSubmission)
	//g.DELETE("/submission/:submissionId", submissionHandler.DeleteSubmission)

	// TagHandler
	g.GET("/tag", tagHandler.GetTags)
	g.POST("/tag", tagHandler.PostTag)
	g.GET("/tag/:tagId", tagHandler.GetTag)
	g.PUT("/tag/:tagId", tagHandler.PutTag)
	g.DELETE("/tag/:tagId", tagHandler.DeleteTag)

	// TODO: Make decision how to manage testcases
	// TestcaseHandler
	g.GET("/testcase", testcaseHandler.GetTestcases)
	g.POST("/testcase", testcaseHandler.PostTestcase)
	g.GET("/testcase/:testcaseId", testcaseHandler.GetTestcase)
	g.PUT("/testcase/:testcaseId", testcaseHandler.PutTestcase)
	g.DELETE("/testcase/:testcaseId", testcaseHandler.DeleteTestcase)

	// UserHandler
	g.GET("/user", userHandler.GetUsers)
	g.POST("/user", userHandler.PostUser)
	g.GET("/user/:userId", userHandler.GetUser)
	g.PUT("/user/:userId", userHandler.PutUser)
	g.DELETE("/user/:userId", userHandler.DeleteUser)
}
