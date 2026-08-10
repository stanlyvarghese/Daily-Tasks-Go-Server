package main

import (
	"fmt"
	"net/http" 
	// "os"
)

//db, err := sql.Open("sqlite3", "app.db")

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Kanban Board</title>
	<link rel="stylesheet" href="/static/css/style.css">
</head>

<body>

	<header class="navbar">
		<div class="navbar-brand">
			<h1>Kanban Board</h1>
		</div>

		<nav class="navbar-links">
			<a href="/" class="nav-link active">Board</a>
			<a href="/about" class="nav-link">About</a>
			<a href="/contact" class="nav-link">Contact</a>
		</nav>
	</header>

	<main class="container">

		<div class="page-header">
			<div>
				<h2>My Tasks</h2>
				<p>Organize your work and keep track of your progress.</p>
			</div>

			<button class="add-task-button">+ Add Task</button>
		</div>

		<section class="kanban-board">

			<div class="kanban-column">
				<div class="column-header">
					<h3>To Do</h3>
					<span class="task-count">7</span>
				</div>

				<div class="task-list">

					<div class="task-card">
						<h4>Learn Go + Templ + HTMX</h4>
						<p>Complete the fundamentals of the Go web stack.</p>
						<span class="task-label development">Development</span>
					</div>

					<div class="task-card">
						<h4>Fill out Mom and Dad Visa applications</h4>
						<span class="task-label personal">Personal</span>
					</div>

					<div class="task-card">
						<h4>Talk to City Sound about speakers</h4>
						<p>Discuss speaker options for Shastri Nagar.</p>
						<span class="task-label work">Work</span>
					</div>

					<div class="task-card">
						<h4>Visit Casio store</h4>
						<span class="task-label personal">Personal</span>
					</div>

					<div class="task-card">
						<h4>Fill out physical Voter ID application</h4>
						<span class="task-label documents">Documents</span>
					</div>

					<div class="task-card">
						<h4>Talk to videographer</h4>
						<span class="task-label work">Work</span>
					</div>

					<div class="task-card">
						<h4>Create DSLR course for Divyanshu &amp; Ashish</h4>
						<span class="task-label development">Development</span>
					</div>

				</div>
			</div>

			<div class="kanban-column">
				<div class="column-header">
					<h3>In Progress</h3>
					<span class="task-count">0</span>
				</div>

				<div class="empty-state">
					<span>○</span>
					<p>No tasks in progress</p>
				</div>
			</div>

			<div class="kanban-column">
				<div class="column-header">
					<h3>Done</h3>
					<span class="task-count">0</span>
				</div>

				<div class="empty-state">
					<span>✓</span>
					<p>Completed tasks will appear here</p>
				</div>
			</div>

		</section>

	</main>

</body>
</html>`)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>About - Kanban Board</title>
	<link rel="stylesheet" href="/static/css/style.css">
</head>

<body>

	<header class="navbar">
		<div class="navbar-brand">
			<h1>Kanban Board</h1>
		</div>

		<nav class="navbar-links">
			<a href="/" class="nav-link">Board</a>
			<a href="/about" class="nav-link active">About</a>
			<a href="/contact" class="nav-link">Contact</a>
		</nav>
	</header>

	<main class="simple-page">

		<section class="info-card">
			<span class="page-icon">📋</span>

			<h2>About Kanban Board</h2>

			<p>
				This is a simple Kanban board built while learning
				Go web development.
			</p>

			<p>
				The project will eventually use Go, Templ, HTMX,
				and SQLite to create a complete task management application.
			</p>

			<a href="/" class="primary-button">Back to Board</a>
		</section>

	</main>

</body>
</html>`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Contact - Kanban Board</title>
	<link rel="stylesheet" href="/static/css/style.css">
</head>

<body>

	<header class="navbar">
		<div class="navbar-brand">
			<h1>Kanban Board</h1>
		</div>

		<nav class="navbar-links">
			<a href="/" class="nav-link">Board</a>
			<a href="/about" class="nav-link">About</a>
			<a href="/contact" class="nav-link active">Contact</a>
		</nav>
	</header>

	<main class="simple-page">

		<section class="info-card">
			<span class="page-icon">✉</span>

			<h2>Contact Me</h2>

			<p>
				If you have questions, suggestions, or feedback
				about this project, feel free to get in touch.
			</p>

			<a href="/" class="primary-button">Back to Board</a>
		</section>

	</main>

</body>
</html>`)
}

func main() {

	// dir, _ := os.Getwd()
	// fmt.Println("Working directory:", dir)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)

}
