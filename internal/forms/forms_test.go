package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	request := httptest.NewRequest("POST", "/goobar", nil)
	form := New(request.PostForm)
	if !form.Valid() {
		t.Error("Empty form should be Valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData

	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {

	values := url.Values{}
	form := New(values)

	if form.Has("whatever") {
		t.Error("(form:51) form shows has field when it does not")
	}

	postedData := url.Values{}

	postedData.Add("a", "a")
	form = New(postedData)

	has := form.Has("a")
	if !has {
		t.Error("shows form dows not have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {

	values := url.Values{}
	form := New(values)

	// We have not added any items to the form data, so minlegth should fail.
	if form.MinLength("afield", 3) {
		t.Error("MinLegth detected and passed a field when it should not")
	}

	postedValues := url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minlength of 100 met when data is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")
	form = New(postedValues)

	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows minlength of 1 is not met when it is")
	}

	isError := form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {

	request := httptest.NewRequest("POST", "/goobar", nil)
	form := New(request.PostForm)

	// We have not added any items to the form data, so minlegth should fail.
	form.IsEmail("email")
	if form.Valid() {
		t.Error("email to isEmail detected and passed a field when it should not")
	}

	postedValues := url.Values{}
	postedValues.Add("email", "me@here.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "me@here")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid for invalid email address")
	}
}
