package handlers

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/page"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/util"
	"html/template"
	"log"
	"net/http"
)

// AdminHandler handles GET-request at /admin
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	//check that user is logged in and is admin/teacher

	//find classes admin/teacher own

	// Data for displaying on screen
	data := struct {
		PageTitle   string
		Menu        page.Menu
		Navbar      page.Menu
		Courses     page.Courses
		Assignments []page.Assignment
	}{
		PageTitle: "Homepage",
		Menu:      util.LoadMenuConfig("configs/menu/dashboard.json"),
		Navbar:    util.LoadMenuConfig("configs/menu/site.json"),
		Courses:   util.LoadCoursesConfig("configs/dd.json"), // dd = dummy data
	}

	w.WriteHeader(http.StatusOK)

	//parse templates
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/navbar.html", "web/dashboard/sidebar.html", "web/dashboard/index.html")

	if err != nil {
		log.Println(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", data); err != nil {
		log.Println(err)
	}
}

// AdminCourseHandler handles GET-request at /admin/course
func AdminCourseHandler(w http.ResponseWriter, r *http.Request) {
	// Data for displaying on screen
	data := struct {
		PageTitle   string
		Menu        page.Menu
		Navbar      page.Menu
		Courses     page.Courses
		Assignments []page.Assignment
	}{
		PageTitle: "Dashboard - Courses",
		Menu:      util.LoadMenuConfig("configs/menu/dashboard.json"),
		Navbar:    util.LoadMenuConfig("configs/menu/site.json"),
		Courses:   util.LoadCoursesConfig("configs/dd.json"), // dd = dummy data
	}

	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/sidebar.html", "web/dashboard/course/index.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", data); err != nil {
		log.Fatal(err)
	}
}

// AdminCreateCourseHandler handles GET-request at /admin/course/create
func AdminCreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/sidebar.html", "web/dashboard/course/create.html")

	if err != nil {
		ErrorHandler(w, r, 404)
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle string
		Menu      page.Menu
		Navbar    page.Menu
	}{
		PageTitle: "Dashboard - Create course",
		Menu:      util.LoadMenuConfig("configs/menu/dashboard.json"),
		Navbar:    util.LoadMenuConfig("configs/menu/site.json"),
	}); err != nil {
		log.Fatal(err)
	}
}

// AdminCreateCourseRequest handles POST-request at /admin/course/create
// Inserts a new course to the database
func AdminCreateCourseRequest(w http.ResponseWriter, r *http.Request) {
	session, err := db.CookieStore.Get(r, "login-session") //get session
	if err != nil {
		log.Fatal(err)
		ErrorHandler(w, r, http.StatusInternalServerError) //error getting session 500
		return
	}

	//check if user is already logged in
	user := getUser(session)

	//todo check that user is a teacher!

	// TODO: talk to database and stuff
	course := page.Course{
		Code:        r.FormValue("code"),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Year:        r.FormValue("year"),
		Semester:    r.FormValue("semester"),
	}

	//todo add year, semester and link names to the mix
	rows, err := db.DB.Query("INSERT INTO course(coursecode, coursename, year, semester, description, teacher) VALUES(?, ?, ?, ?, ?, ?)",
		course.Code, course.Name, course.Year, course.Semester, course.Description, user.ID, )

	if err != nil {
		//todo log error
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	http.Redirect(w, r, "/", http.StatusFound) //success redirect to homepage
}

// AdminUpdateCourseHandler handles GET-request at /admin/course/update/{id}
func AdminUpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/sidebar.html", "web/dashboard/course/update.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle string
		Menu      page.Menu
		Navbar    page.Menu
	}{
		PageTitle: "Dashboard - Update course",
		Menu:      util.LoadMenuConfig("configs/menu/dashboard.json"),
		Navbar:    util.LoadMenuConfig("configs/menu/site.json"),
	}); err != nil {
		log.Fatal(err)
	}
}

// AdminUpdateCourseRequest handles POST-request at /admin/course/update/{id}
func AdminUpdateCourseRequest(w http.ResponseWriter, r *http.Request) {

}

// AdminAssignmentHandler handles GET-request at /admin/assignment
func AdminAssignmentHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	//parse template
	temp, err := template.ParseFiles("web/dashboard/layout.html", "web/dashboard/sidebar.html", "web/dashboard/assignment/index.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle string
		Menu      page.Menu
		Navbar    page.Menu
	}{
		PageTitle: "Dashboard - Assignments",
		Menu:      util.LoadMenuConfig("configs/menu/dashboard.json"),
		Navbar:    util.LoadMenuConfig("configs/menu/site.json"),
	}); err != nil {
		log.Fatal(err)
	}
}
