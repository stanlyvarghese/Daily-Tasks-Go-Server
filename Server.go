package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Home</title>
</head>
<body>
    <div style="display: flex; background-color: black; width: 100vw;">
        
        <div style="display: grid; place-content: center;  width: 80vw;"> <h1 style="color: white;">Home</h1> </div>
        
        <div style="display: grid; place-content: center; width: 10vw; color: white;">
            <a href="http://localhost:8080/about"> <button type="button">About</button></a>
        </div>
        <div style="display: grid; place-content: center; width: 10vw; color: white;"><a href="http://localhost:8080/contact"> <button type="button"> Contact Us</button></a></div>
    
</div>
<ol style="font-size: 20px;">
            <li>Learn Go+Templ+HTMX</li>
			<li>Fill out Mom and Dad Visa applications</li>
            <li>Talk to City Sound for speakers in Shastri Nagar</li>
            <li>Visit Casio store</li>
            <li>Fill out application for physical Voter ID</li>
            <li>Talk to videographer</li>
            <li>Create a DSLR course for Divyanshu & Ashish</li>
        </ol>
</body>
</html>`)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Home</title>
</head>
<body>
    <div style="display: flex; background-color: black; width: 100vw;">
        
        <div style="display: grid; place-content: center;  width: 80vw;"> <h1 style="color: white;">About</h1> </div>
        
        <div style="display: grid; place-content: center; width: 10vw; color: white;">
            <a href="http://localhost:8080/"> <button type="button">Home</button></a>
        </div>
        <div style="display: grid; place-content: center; width: 10vw; color: white;"><a href="http://localhost:8080/contact"> <button type="button"> Contact Us</button></a></div>
    
</div>
</body>
</html>`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Home</title>
</head>
<body>
    <div style="display: flex; background-color: black; width: 100vw;">
        
        <div style="display: grid; place-content: center;  width: 80vw;"> <h1 style="color: white;">Contact Me</h1> </div>
        
        <div style="display: grid; place-content: center; width: 10vw; color: white;">
            <a href="http://localhost:8080/"> <button type="button">Home</button></a>
        </div>
        <div style="display: grid; place-content: center; width: 10vw; color: white;"><a href="http://localhost:8080/about"> <button type="button"> About </button></a></div>
    
</div>
</body>
</html>`)
	fmt.Fprintln(w, "Contact me at : 7007787263")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}
