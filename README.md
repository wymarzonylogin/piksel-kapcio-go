# Piksel Kapcio Go
This package is a simple captcha implementation for **Go**. It is a port of original **PHP** implementation of **Piksel Kapcio** library (https://github.com/wymarzonylogin/piksel-kapcio). For general description  please go to original [PHP implementation Piksel Kapcio docs](https://github.com/wymarzonylogin/piksel-kapcio/blob/master/README.md). The major difference between PHP and Go implementations is lack of session handling in Go port. Besides that, they are supposed to behave in an exactly same manner, having identical configuration options. This document covers details specific to Go implementation.

![piksel-kapcio-title](https://wymarzonylog.in/img/github/piksel-kapcio/piksel-kapcio-title.png)

## Session handling
I decided not to introduce dependancy on any session implementation for this package. There is a strong chance you are already using a different implementation than the one I would choose. This makes usage of this package a bit more cumbersome than it's PHP counterpart, but i believe it's still better than introducing dependancy that could be redundant. This document, however, covers fully operating example with use of great and widely used https://github.com/gorilla/sessions.

## General flow
Typical flow for captcha consists of 2 steps: captcha generation and captcha validation. 

Typically web application endpoint generating captcha has its own URL and does the following:
- generate text code and it's corresponding graphical representation (image)
- save generated text code in users session, so further validation is possible
- serve image to user

Usually you want to use captcha to protect forms from automated submission, so you show this image in/next to some form that user fills in along with additional field, in which user is supposed to retype text he can read from the served captcha image.

You will need to add captcha validation in endpoint that is processing form submission. This validation is simply based on comparing text for captcha that user submitted through the form with the code that was generated and stored in users session in generation step. If code stored in session and one submitted by user are the same, captcha was solved properly and we can assume form was submitted by a human, not a bot.

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
First, you need to import this package.

Then you can generate captcha with `GenerateCode(config)` function (config struct has to be passed as an argument, but it can be empty, as generator has a set of defaults).
This call returns two values: text (string) representation of code (here `code`) and corresponding image data that you can serve to user (here `imageData`). Text representation should be stored in session for further validation. Image should be served to user.

### Captcha validation
To check if user provided captcha solution is correct, simply compare uppercased user submitted value of code with one stored in session. Please pay attention to **uppercased** part - this package is case insensitive and operates on uppercase letters! Even if you provided custom list of lowercase words, codes returned for them by `GenerateCode` call will be using capital letters only.

## Text generation
There are two ways of generating code - random string and randomly selected word from defined list. 

Random strings contain only uppercase letters [A-Z] and digits [0-9]. You can define length of random code with any value from range [1,36]. In case provided value is outside of this range, generation of code will fallback to default length of 4.

With default configuration, image representing generated random string code looks like that:

![serve-image-default](https://wymarzonylog.in/img/github/piksel-kapcio/serve-image-default.png)

Codes generated from provided custom list may be modified - first of all the letters will be uppercase. If custom word contains unsupported characters, they will be replaced with asterisks (`*`). In that case, also image presented to user will show asterisks. This lets captcha be solvable by user even if image generator would not be able to provide image with unsupported characters.

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

## Image generation


## Configuration
## Full example