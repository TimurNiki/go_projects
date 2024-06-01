package main

import (
	"crypto/rand"
	"errors"
	"strings"
	"unicode/utf8"

	// "html/template"
	// "log"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TimurNiki/go_api_tutorial/books/snippetbox/internal/models"
	"github.com/TimurNiki/go_api_tutorial/books/snippetbox/internal/validator"
	"github.com/go-playground/form/v4"
	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// if r.URL.Path != "/" {
	// 	app.notFound(w)
	// 	return
	// }

	// Because httprouter matches the "/" path exactly, we can now remove the
	// manual check of r.URL.Path != "/" from this handler.

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Call the newTemplateData() helper to get a templateData struct containing
	// the 'default' data (which for now is just the current year), and add the
	// snippets slice to it.
	data := app.newTemplateData(r)
	data.Snippets = snippets

	// Use the new render helper.
	app.render(w, http.StatusOK, "home.tmpl.html", data)

	// files := []string{
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/partials/nav.tmpl",
	// 	"./ui/html/pages/home.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	// // Create an instance of a templateData struct holding the slice of
	// // snippets.
	// data := &templateData{
	// 	Snippets: snippets,
	// }

	// // Pass in the templateData struct when executing the template.
	// err = ts.ExecuteTemplate(w, "base", data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// for _, snippet := range snippets {
	// 	fmt.Fprintf(w, "%+v\n", snippet)
	// }

	// files := []string{
	// 	"./ui/html/base.tmpl.html",
	// 	"./ui/html/pages/home.tmpl.html",
	// 	"./ui/html/partials/nav.tmpl.html",
	// }

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.errorLog.Print(err.Error())
	// 	http.Error(w, "Internal Server Error", 500)
	// 	return
	// }
	// We then use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in, which for now we'll
	// leave as nil.
	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.errorLog.Print(err.Error())
	// 	http.Error(w, "Internal Server Error", 500)
	// }
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	// When httprouter is parsing a request, the values of any named parameters
	// will be stored in the request context. We'll talk about request context
	// in detail later in the book, but for now it's enough to know that you can
	// use the ParamsFromContext() function to retrieve a slice containing these
	// parameter names and values like so:
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return

	}

	// id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// if err != nil || id < 1 {
	// 	app.notFound(w)
	// 	return
	// }

	// Use the SnippetModel object's Get method to retrieve the data for a
	// specific record based on its ID. If no matching record is found,
	// return a 404 Not Found response.
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Use the PopString() method to retrieve the value for the "flash" key.
	// PopString() also deletes the key and value from the session data, so it
	// acts like a one-time fetch. If there is no matching key in the session
	// data this will return the empty string
	// flash:= app.sessionManager.PopString(r.Context(), "flash")

	data := app.newTemplateData(r)
	data.Snippet = snippet

	// Use the new render helper.
	app.render(w, http.StatusOK, "view.tmpl.html", data)

	// Initialize a slice containing the paths to the view.tmpl file,
	// plus the base layout and navigation partial that we made earlier.
	// files := []string{
	// 	"./ui/html/base.tmpl.html",
	// 	"./ui/html/partials/nav.tmpl.html",
	// 	"./ui/html/pages/view.tmpl.html",
	// }
	// // Parse the template files...
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	// // Create an instance of a templateData struct holding the snippet data.
	// data := &templateData{
	// 	Snippet: snippet,
	// }

	// // Pass in the templateData struct when executing the template.
	// err = ts.ExecuteTemplate(w, "base", data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// // And then execute them. Notice how we are passing in the snippet
	// // data (a models.Snippet struct) as the final parameter?
	// err = ts.ExecuteTemplate(w, "base", snippet)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// // Write the snippet data as a plain-text HTTP response body
	// fmt.Fprintf(w, "%+v", snippet)

	// fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	data := app.newTemplateData(r)
	// Initialize a new createSnippetForm instance and pass it to the template.
	// Notice how this is also a great opportunity to set any default or
	// 'initial' values for the form --- here we set the initial value for the
	// snippet expiry to 365 days.
	data.Form = snippetCreateForm{
		Expires: 365,
	}
	app.render(w, http.StatusOK, "create.tmpl.html", data)

	// Add a new snippetCreate handler, which for now returns a placeholder
	// response. We'll update this shortly to show a HTML form.
	// w.Write([]byte("Display the form for creating a new snippet..."))

	// if r.Method != http.MethodPost {
	// 	w.Header().Set("Allow", http.MethodPost)
	// 	app.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }
	// w.Write([]byte("Create a new snippet..."))

}

// * Define a snippetCreateForm struct to represent the form data and validation
// errors for the form fields. Note that all the struct fields are deliberately
// exported (i.e. start with a capital letter). This is because struct fields
// must be exported in order to be read by the html/template package when
// rendering the template.
// * Update our snippetCreateForm struct to include struct tags which tell the
// decoder how to map HTML form values into the different struct fields. So, for
// example, here we're telling the decoder to store the value from the HTML form
// input with the name "title" in the Title field. The struct tag `form:"-"`
// tells the decoder to completely ignore a field during decoding.
type snippetCreateForm struct {
	Title   string `form:"title"`
	Content string `form:"content"`
	Expires int    `form:"expires"`
	// FieldErrors map[string]string
	// Remove the explicit FieldErrors struct field and instead embed the Validator
	// type. Embedding this means that our snippetCreateForm "inherits" all the
	// fields and methods of our Validator type (including the FieldErrors field).
	validator.Validator `form:"-"`
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	// First we call r.ParseForm() which adds any data in POST request bodies
	// to the r.PostForm map. This also works in the same way for PUT and PATCH
	// requests. If there are any errors, we use our app.ClientError() helper to
	// send a 400 Bad Request response to the user.
	// err := r.ParseForm()
	// if err != nil {
	// 	app.clientError(w, http.StatusBadRequest)
	// 	return
	// }

	var form snippetCreateForm

	err = app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	//* Use the r.PostForm.Get() method to retrieve the title and content
	// from the r.PostForm map.
	// title := r.PostForm.Get("title")
	// content := r.PostForm.Get("content")

	//* The r.PostForm.Get() method always returns the form data as a *string*.
	// However, we're expecting our expires value to be a number, and want to
	// represent it in our Go code as an integer. So we need to manually covert
	// the form data to an integer using strconv.Atoi(), and we send a 400 Bad
	// Request response if the conversion fails.
	// expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	// if err != nil {
	// 	app.clientError(w, http.StatusBadRequest)
	// 	return
	// }

	// form := snippetCreateForm{
	// 	Title:   r.PostForm.Get("title"),
	// 	Content: r.PostForm.Get("content"),
	// 	Expires: expires,
	// 	// FieldErrors: make(map[string]string),
	// 	// Remove the above code and instead use the validator.
	// }

	// Because the Validator type is embedded by the snippetCreateForm struct,
	// we can call CheckField() directly on it to execute our validation checks.
	// CheckField() will add the provided key and error message to the
	// FieldErrors map if the check does not evaluate to true. For example, in
	// the first line here we "check that the form.Title field is not blank". In
	// the second, we "check that the form.Title field has a maximum character
	// length of 100" and so on.
	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	// Use the Valid() method to see if any of the checks failed. If they did,
	// then re-render the template passing in the form in the same way as
	// before.
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.tmpl.html", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Use the Put() method to add a string value ("Snippet successfully
	// created!") and the corresponding key ("flash") to the session data
	app.sessionManager.Put(r.Content(), "flash", "Snippet successfully created!")

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)

	// Initialize a map to hold any validation errors for the form fields.
	// fieldErrors := make(map[string]string)

	// Check that the title value is not blank and is not more than 100
	// characters long. If it fails either of those checks, add a message to the
	// errors map using the field name as the key.
	// if strings.TrimSpace(form.Title) == "" {
	// 	form.FieldErrors["title"] = "This field cannot be blank"
	// } else if utf8.RuneCountInString(form.Title) > 100 {
	// 	form.FieldErrors["title"] = "This field cannot be more than 100 characters long"
	// }

	// Check that the Content value isn't blank.
	// if strings.TrimSpace(form.Content) == "" {
	// 	form.FieldErrors["content"] = "This field cannot be blank"
	// }

	// Check the expires value matches one of the permitted values (1, 7 or
	// 365).
	// if form.Expires != 1 && form.Expires != 7 && form.Expires != 365 {
	// 	form.FieldErrors["expires"] = "This field must equal 1, 7 or 365"
	// }

	// If there are any errors, dump them in a plain text HTTP response and
	// return from the handler
	// if len(form.FieldErrors) > 0 {
	// fmt.Fprint(w, form.FieldErrors)
	// return

	//* If there are any validation errors re-display the create.tmpl template,
	// passing in the snippetCreateForm instance as dynamic data in the Form
	// field. Note that we use the HTTP status code 422 Unprocessable Entity
	// when sending the response to indicate that there was a validation error.
	// data := app.newTemplateData(r)
	// data.Form = form
	// app.render(w, http.StatusUnprocessableEntity, "create.tmpl.html", data)
	// return
	// }

	// 	 // Checking if the request method is a POST is now superfluous and can be
	//  // removed, because this is done automatically by httprouter.
	//  title := "O snail"
	//  content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	//  expires := 7

	//  id,err:=app.snippets.Insert(title,content,expires)
	//  if err!=nil{
	// 	app.serverError(w,err)
	// 	return
	//  }
	// Update the redirect path to use the new clean URL format.
	// http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

type userSignupForm struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) userSignUp(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userSignupForm{}
	app.render(w, http.StatusOK, "signup.tmpl", data)
}

func (app *application) userSignUpPost(w http.ResponseWriter, r *http.Request) {
	// Declare an zero-valued instance of our userSignupForm struct.
	var form userSignupForm

	// Parse the form data into the userSignupForm struct.
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Validate the form contents using our helper functions.
	form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must be at least 8 characters long")

	// If there are any errors, redisplay the signup form along with a 422
	// status code.
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "signup.tmpl", data)
		return
	}

	// Try to create a new user record in the database. If the email already
	// exists then add an error message to the form and re-display it.
	err = app.Users.Insert(form.Name, form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.AddFieldError("email", "Email address is already in use")
			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, http.StatusUnprocessableEntity, "signup.tmpl.html", data)
		} else {
			app.serverError(w, err)
		}
	}
	
	// Otherwise add a confirmation flash message to the session confirming that
	// their signup worked.
	app.sessionManager.Put(r.Context(), "flash", "Your signup was successful. Please log in.")
	// And redirect the user to the login page.
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)

}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {

}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {

}

func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {

}
