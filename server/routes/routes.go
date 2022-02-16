package routes

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4/middleware"
	"kry-go/server"
	"kry-go/server/handlers"
	"os"
)

func InitRouter(server *server.Server) {
	contestHandler := handlers.MakeContestHandler(server)
	problemHandler := handlers.MakeProblemHandler(server)
	submissionHandler := handlers.MakeSubmissionHandler(server)
	tagHandler := handlers.MakeTagHandler(server)
	testcaseHandler := handlers.MakeTestcaseHandler(server)
	userHandler := handlers.MakeUserHandler(server)
	loginHandler := handlers.MakeLoginHandler(server)

	// TODO: Consider using redis-session
	sessionSecret := os.Getenv("SESSION_SECRET")
	server.Echo.Use(session.Middleware(sessions.NewCookieStore([]byte(sessionSecret))))
	server.Echo.Use(middleware.Logger())

	g := server.Echo.Group("/api")

	g.GET("/login/:loginType", loginHandler.LoginUser)
	g.POST("/login/:loginType/callback", loginHandler.LoginNewUser)

	// TODO: Define resources in URLs
	g.GET("/contest/:contestId", contestHandler.GetContest)
	g.GET("/problem/:problemId", problemHandler.GetProblem)
	g.GET("/submission/:problemId", submissionHandler.GetSubmission)
	g.GET("/tag/:problemId", tagHandler.GetTag)
	g.GET("/testcase/:problemId", testcaseHandler.GetTestcase)
	g.GET("/user/:userId", userHandler.GetUser)
}
