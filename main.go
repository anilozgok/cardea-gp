package main

import (
	"fmt"
	"github.com/anilozgok/cardea-gp/internal/config"
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/handler"
	"github.com/anilozgok/cardea-gp/pkg/middleware"
	"github.com/anilozgok/cardea-gp/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var isLocalMode = os.Getenv("CARDEA_GP_LOCAL_MODE") == "true"

func init() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	l, err := cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("fail to build log. err: %s", err))
	}

	zap.ReplaceGlobals(l)
}

func main() {
	configs, err := config.Get()
	if err != nil {
		zap.L().Fatal("failed to read configs", zap.Error(err))
	}

	db := database.
		New(configs).
		Initialize()

	repo := database.NewRepository(db)

	register := handler.NewRegisterHandler(repo)
	login := handler.NewLoginHandler(repo)
	logout := handler.NewLogoutHandler()
	getUsers := handler.NewGetUsersHandler(repo)
	me := handler.NewMeHandler()
	createWorkout := handler.NewCreateWorkoutHandler(repo)
	listCoachWorkouts := handler.NewListCoachWorkoutHandler(repo)
	listUserWorkouts := handler.NewListUserWorkoutHandler(repo)

	fpCtx := new(utils.ForgotPasswordCtx)
	checkUser := handler.NewCheckUserHandler(repo, configs, fpCtx)
	verifyPasscode := handler.NewVerifyPasscodeHandler(fpCtx)
	updatePassword := handler.NewUpdatePasswordHandler(repo, fpCtx)

	profileHandler := handler.NewProfileHandler(repo)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	// Ensure the images directory exists
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		err := os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	app.Static("/uploads", "./uploads")

	if isLocalMode {
		app.Use(logger.New())
	}

	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("OK")
	})

	r := app.Group("/api/v1")
	auth := r.Group("/auth")
	auth.Post("/register", register.Handle)
	auth.Post("/login", login.Handle)
	auth.Post("/logout", logout.Handle)
	auth.Get("/check-user", checkUser.Handle)
	auth.Put("/verify-passcode", verifyPasscode.Handle)
	auth.Put("/update-password", updatePassword.Handle)

	user := r.Group("/user")
	user.Get("/", middleware.AuthMiddleware, middleware.RoleAdmin, getUsers.Handle)
	user.Get("/me", middleware.AuthMiddleware, me.Handle)
	user.Get("/workouts", middleware.AuthMiddleware, middleware.RoleUser, listUserWorkouts.Handle)

	workout := r.Group("/workout")
	workout.Post("/", middleware.AuthMiddleware, middleware.RoleCoach, createWorkout.Handle)
	workout.Get("/", middleware.AuthMiddleware, middleware.RoleCoach, listCoachWorkouts.Handle)

	// Profile routes
	profile := r.Group("/profile")
	profile.Post("/", middleware.AuthMiddleware, middleware.RoleUser, profileHandler.CreateProfile)
	profile.Get("/", middleware.AuthMiddleware, middleware.RoleUser, profileHandler.GetProfile)
	profile.Put("/", middleware.AuthMiddleware, middleware.RoleUser, profileHandler.UpdateProfile)
	profile.Post("/upload-photo", middleware.AuthMiddleware, middleware.RoleUser, profileHandler.UploadPhoto)

	go func() {
		err = app.Listen(":8080")
		if err != nil {
			zap.L().Fatal("error while starting server", zap.Error(err))
		}
	}()
	zap.L().Info("server started successfully on port :8080")

	gracefulShutdown(app)
}

func gracefulShutdown(app *fiber.App) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-signalCh

	zap.L().Info("shutting down server")
	if err := app.Shutdown(); err != nil {
		zap.L().Error("error occurred while shutting down server", zap.Error(err))
	}
}
