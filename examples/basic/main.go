package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	zjucourses "github.com/eWloYW8/zju-courses-go-sdk"
	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

func main() {
	// Create client with cookie-based authentication.
	// You need to obtain cookies by logging in via browser or SSO.
	client := zjucourses.NewClient()

	// Set cookies from browser (copy from browser DevTools > Application > Cookies)
	client.SetCookieString("SESSION=your-session-cookie-here")

	// Or set cookies programmatically
	client.SetCookies([]*http.Cookie{
		{Name: "SESSION", Value: "your-session-cookie-value"},
	})

	ctx := context.Background()

	// --- List My Courses ---
	courses, err := client.Courses.ListMyCourses(ctx, &model.ListOptions{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		log.Fatalf("Failed to list courses: %v", err)
	}
	fmt.Printf("Found %d courses (total: %d)\n", len(courses.Courses), courses.Total)
	for _, c := range courses.Courses {
		fmt.Printf("  [%d] %s (%s)\n", c.ID, c.DisplayName, c.CourseCode)
	}

	if len(courses.Courses) == 0 {
		return
	}

	// --- Get Course Details ---
	courseID := courses.Courses[0].ID
	course, err := client.Courses.GetCourse(ctx, courseID)
	if err != nil {
		log.Fatalf("Failed to get course: %v", err)
	}
	fmt.Printf("\nCourse: %s\n", course.DisplayName)
	fmt.Printf("  Department: %s\n", course.Department.Name)
	fmt.Printf("  Students: %d\n", course.StudentsCount)

	// --- List Modules ---
	modules, err := client.Courses.ListModules(ctx, courseID)
	if err != nil {
		log.Fatalf("Failed to list modules: %v", err)
	}
	fmt.Printf("\nModules (%d):\n", len(modules.Modules))
	for _, m := range modules.Modules {
		fmt.Printf("  [%d] %s\n", m.ID, m.Name)
	}

	// --- Get Todos ---
	todos, err := client.Notifications.ListTodos(ctx)
	if err != nil {
		log.Fatalf("Failed to list todos: %v", err)
	}
	fmt.Printf("\nTodos (%d):\n", len(todos.TodoList))
	for _, t := range todos.TodoList {
		fmt.Printf("  [%s] %s - %s (due: %s)\n", t.Type, t.CourseName, t.Title, t.EndTime)
	}

	// --- Get Homework Submission Status ---
	status, err := client.Courses.GetHomeworkSubmissionStatus(ctx, courseID)
	if err != nil {
		log.Fatalf("Failed to get homework status: %v", err)
	}
	fmt.Printf("\nHomework Status (%d):\n", len(status.HomeworkActivities))
	for _, h := range status.HomeworkActivities {
		fmt.Printf("  [%d] %s (%s)\n", h.ID, h.Status, h.StatusCode)
	}

	// --- Get Academic Years ---
	years, err := client.User.ListMyAcademicYears(ctx)
	if err != nil {
		log.Fatalf("Failed to list academic years: %v", err)
	}
	for _, y := range years.AcademicYears {
		fmt.Printf("  Academic Year: %s (active: %v)\n", y.Name, y.IsActive)
	}

	// --- Get Global Config ---
	config, err := client.Admin.GetGlobalConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to get global config: %v", err)
	}
	fmt.Printf("\nPlatform: %s\n", config.APM.Environment)

	// --- Check AI Ability ---
	aiAbility, err := client.AirCredit.HasAIAbility(ctx)
	if err != nil {
		log.Fatalf("Failed to check AI ability: %v", err)
	}
	fmt.Printf("Has AI Ability: %v\n", aiAbility.HasAnyCourseAIAbility)
}
