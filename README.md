# Piksel Kapcio Go
This package is a simple captcha implementation for **Go**. It is a port of original **PHP** implementation of **Piksel Kapcio** library (https://github.com/wymarzonylogin/piksel-kapcio). For general description  please check out original [PHP implementation Piksel Kapcio docs](https://github.com/wymarzonylogin/piksel-kapcio/blob/master/README.md). The major difference between PHP and Go implementations is lack of session handling in Go port. Besides that, they are supposed to behave in an exactly same manner, having identical configuration options. This document covers details specific to Go implementation.

![piksel-kapcio-title](https://wymarzonylog.in/img/github/piksel-kapcio/piksel-kapcio-title.png)

## Index
1. [Session Handling](#session-handling)
2. [General flow](#general-flow)
3. [Text generation](#test-generation)
4. [Image generation](#image-generation)
5. [Configuration](#configuration)
6. [Full example](#full-example)

## Session handling
I decided not to introduce dependancy on any session implementation for this package. There is a strong chance you are already using a different implementation than the one I would choose. This makes usage of this package a bit more cumbersome than it's PHP counterpart, but i believe it's still better than introducing dependancy that could be redundant. This document, however, covers fully operating example with use of great and widely used https://github.com/gorilla/sessions.

## General flow
Typical flow for captcha consists of 2 steps: captcha generation and captcha validation. 

Web application endpoint generating captcha does the following:
- generate text code and it's corresponding graphical representation (image)
- save generated text code in users session, so later it can be retreived for validation
- serve image to user (user's browser)

Usually you want to use captcha to protect forms from automated submission, so you show this image in/next to some form that user fills in along with additional field, in which user is supposed to retype text he can read from the served captcha image.

You will need to add captcha validation in endpoint that is processing form submission. This validation is simply based on comparing text for captcha that user submitted through the form with the code that was generated and stored in users session in generation step. If code stored in session and code submitted by user are equal, captcha has been solved properly and we can assume form was submitted by a human, not a bot.

### Captcha generation
Below you can see code for the simplest endpoint that generates and serves captcha image:

```go
package main

import (
	"image/png"
	"log"
	"net/http"

	pikselkapcio "github.com/wymarzonylogin/piksel-kapcio-go"
)

func captchaHandler(w http.ResponseWriter, r *http.Request) {
	config := pikselkapcio.Config{}
	code, imageData := pikselkapcio.GenerateCode(config)

	// custom session handling - save generated code to session for further comparison with  user input
	// ...
	// session.Values["captcha_session_key"] = code
	// ...
	// end of custom session handling

	if err := png.Encode(w, imageData); err != nil {
		log.Println("Error while encoding image.")
	}

	w.Header().Set("Content-Type", "image/png")
}
```
Generate captcha with `GenerateCode(config)` function (config struct has to be passed as an argument, but it can be empty, as generator has a set of defaults for options).
This call returns two values: text (string) representation of code (here `code`) and corresponding image data that you can serve to user (here `imageData`). Text representation should be stored in session for further validation. Image should be served to user.

### Captcha validation
To check if user provided captcha solution is correct, simply compare uppercased code value submitted by user with the code stored in session. Please pay attention to **uppercased** part - this package is case insensitive and operates on uppercase letters. Even if you provide custom list of lowercase words, codes returned for them by `GenerateCode` call will be using capital letters instead of lowercase letters.

## Text generation
There are two ways of generating code - random string and randomly selected word from defined list. 

Random strings contain only uppercase letters [A-Z] and digits [0-9]. You can define length of random code with any integer value within [1,36] range. In case provided value is outside of this range, generation of code will fallback to default length of 4.

With default configuration, image representing generated random string code looks like that:

![serve-image-default](https://wymarzonylog.in/img/github/piksel-kapcio/serve-image-default.png)

Codes generated based on provided custom words list may be modified - first of all the lowercase letters will be replaced with uppercase counterparts. If custom word contains unsupported characters, they will be replaced with asterisks (`*`). In that case, also image presented to user will show asterisks. This lets captcha be solvable by user even if image generator would not be able to draw all the characters contained in custom word.

Here is the list of supported characters:
- digits [0-9]
- latin uppercase letters [A-Z]
- latin lowercase letters [a-z] (those are being converted to uppercase on code generation)
- empty space (` `)
- dot (`.`)
- basic arithmetical operator signs (`+-*/=`)
- question mark (`?`)
- round brackets (`()`)

Or if you prefer it as a full list:
`"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz.+-*/=?() "`

It's recommended to use only supported characters in custom words.

## Image generation
Generated image consists of scaled 7 x 7 pixel blocks representing each character of generated code. Each character occupies 5 x 5 pixels in the middle, but there is 1 pixel wide outline (filled with background color) around each character to make characters not touch each other. Each block has it's background color and foreground color.

You can define scale parameter (by default it's 5). With default scale, each character is represented by 35 x 35 pixel block on the result image (7 x scale). With scale 1, each character would be represented by literally 7 x 7 pixel area of the image.

You can define custom color pairs for generating image. Each color pair contains definition of background color and foreground color. If only one color pair is defined, all the characters on the image will have the same background and foreground color. If more than one color pair is defined, color pairs are assigned to characters in one of two ways:
- randomly (default)
- in a sequence cycle (1st character gets 1st defined color pair, 2nd character gets 2nd pair etc.)

There is default set of few color pairs defined.

Be careful while defining color pairs, not to set the same or very similar value for both background and foreground color of a given pair. This can make the image not readable because of not enough or no contrast at all.

## Configuration
Config structure that is passed to `GenerateCode` function is defined as:
```go
type Config struct {
	Scale               int
	TextGenerationMode  int8
	RandomTextLength    int8
	CustomWords         []string
	ColorHexStringPairs []ColorHexStringPair
	ColorPairsRotation  int8
}
```
### Scale
`Scale` affects the size of generated image. Default value is 5. Each pixel of the result image is scaled, which means that by default dot (`.`) character on the image would be actually a 5 x 5 pixels square.
### TextGenerationMode
For `TextGenerationMode` you should use one of available constants:
- `TextGenerationRandom` for generating random code that consists of letters and digits only (default behavior).
- `TextGenerationCustomWords` for generating code as a random selection  from provided custom words list.
### RandomTextLength
You can provide number from range [1,36]. If passed value ist out of this range, default length is used. Default length is 4. This parameter doesn't affect code generation behavior if `TextGenerationCustomWords` is used.
### CustomWords
Provide custom words list as:

`[]string{"word1", "word2"...},`

Even if custom words are provided, `TextGenerationMode` must be set to `TextGenerationCustomWords`. Otherwise, random code would be generated.
### ColorHexStringPairs
Provide custom color pairs as:
```go
[]pikselkapcio.ColorHexStringPair{
    {
        BackgroundColor: "FFFFFF",
        ForegroundColor: "000000",
    },
    {
        BackgroundColor: "000000",
        ForegroundColor: "FFFFFF",
    },
    ...
}
```
You can provide just one color pair if you want. If you provide empty set of color pairs, default color pairs will be used.
### ColorPairsRotation
For `ColorPairsRotation` you should use one of available constants:
- `ColorPairsRotationRandom` for picking random color pair out of defined color pairs for each character (default behavior). Notice: this is not random color generation.
- `ColorPairsRotationSequence` for cycling through color pairs in a sequence to select color pair for a character.

Check full example below to see how configuration options are set. `RandomTextLength` that you can see there is obsolete, as custom words list is used anyway.

## Full example
Here is full example of the most simple app using Piksel Kapcio with [Gorilla sessions](https://github.com/gorilla/sessions.) as sessions implementation:
```go
package main

import (
	"flag"
	"image/png"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	pikselkapcio "github.com/wymarzonylogin/piksel-kapcio-go"
)

var store = sessions.NewCookieStore([]byte("abcdef0123456789"))

func main() {
	flag.Parse()
	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/validate", validateHandler)
	http.HandleFunc("/captcha", captchaHandler)
	http.ListenAndServe(":8080", nil)
}

//Handler showing form with captcha field to fill in by user
func homepageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<html>
			<head>
				<meta charset="utf-8" />
				<title>PikselKapcio Test</title>
			</head>
			<body>
				<h1>PikselKapcio Test</h1>
				<form action="/validate" method="POST">
					<div>
						<img src="/captcha?" />
					</div>
					<div>
						<label for="verification-code-input">Type in code from picture above</label><br />
						<input id="verification-code-input" type="text" name="user_solution" />
					</div>
					<div>
						<input type="submit" value="Submit" />
					</div>
				</form>
			</body>
		</html>
	`))
}

//Handler generating image and storing associated code in session
func captchaHandler(w http.ResponseWriter, r *http.Request) {
	config := pikselkapcio.Config{
		Scale:              4,
		RandomTextLength:   10,
		TextGenerationMode: pikselkapcio.TextGenerationCustomWords,
		CustomWords:        []string{"Gopher", "Elephant", "Squirrel", "Turtle", "Octopus", "Hamster"},
		ColorHexStringPairs: []pikselkapcio.ColorHexStringPair{
			{
				BackgroundColor: "FFFF88",
				ForegroundColor: "44BBAA",
			},
			{
				BackgroundColor: "770066",
				ForegroundColor: "CCFF00",
			},
			{
				BackgroundColor: "446688",
				ForegroundColor: "FF4477",
			},
		},
		ColorPairsRotation: pikselkapcio.ColorPairsRotationSequence,
	}

	code, imageData := pikselkapcio.GenerateCode(config)

	//custom session handling - save generated code to session for further comparison with user input
	session, _ := store.Get(r, "session-name")
	session.Values["kapcio-code"] = code
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//end of custom session handling

	if err := png.Encode(w, imageData); err != nil {
		log.Println("Error while encoding image.")
	}

	w.Header().Set("Content-Type", "image/png")
}

//Handler receiving submitted form including users solution for captcha
func validateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	submittedCode := strings.ToUpper(r.Form.Get("user_solution"))
	var expectedCode string

	//custom session code retrieval
	session, _ := store.Get(r, "session-name")
	if session.Values["kapcio-code"] != nil {
		expectedCode = session.Values["kapcio-code"].(string)

		//clearing value in session to prevent resubmission via page refresh
		session.Values["kapcio-code"] = nil
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	//end of custom session code retrieval

	var resultMessage string
	if expectedCode != "" && expectedCode == submittedCode {
		resultMessage = "OK! Correct code!"
	} else {
		resultMessage = "Invalid code"
	}

	w.Write([]byte(resultMessage))
}

```
If you run this code, after accessing `http://127.0.0.1:8080/`, you should see form with captcha image:

![piksel-kapcio-title](https://wymarzonylog.in/img/github/piksel-kapcio/example-form.png)

Homepage (url `/`) displays form with one field, named `user_solution`. Above the field, the captcha image generated by PikselKapcio is displayed (it comes from `/captcha` url). Generator is triggered in `captchaHandler` with fully custom configuration. After generation, text representation of the code is stored in session under `kapcio-code` key.

Submitted form is processed by `validateHandler`. Value of code submitted by user is retrieved there from request (`user_solution` field) and uppercased. Then it's compared with expected code retreived from user session and appropriate message informing about comparison result is displayed.

To make example more robust, additional check on code validation is made:
`expectedCode != ""`

This is supposed to prevent case, when user enters the validation page directly, without having any code set in session and without submitting any.

Please note that depending on what you want to achieve and what is your use case, that after reading code value from session while validating, session value may need to be cleared in order to prevent resubmission of the same form without having new image shown - by simply refreshing validate page (or by pressing "back" button in browser - which would show previously submitted form but wouldn't reload captcha image, so would not affect code stored in session - and pressing submit again)
