package emailx

import (
	"github.com/matcornic/hermes"
	"io/ioutil"
)

//type Product struct {
//	Name string
//	Link string
//	Logo string
//}

func GenerateEmail() error {
	// Configure hermes by setting a theme and you product info
	h := hermes.Hermes{
		// Optional Theme
		Product: hermes.Product{
			Name: "Hermes",
			Link: "http://www.baidu.com",
			// Optional product logo
			Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
	}

	// Generate an email
	email := hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Confirm your account",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	// Generate an HTML email with the provided contents
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return nil
	}

	// Generate the plaintext version of the email
	//emailText, err := h.GenerateHTML(email)
	//if err != nil {
	//	return nil
	//}

	// Optionally, preveiw the generated the HTML
	err = ioutil.WriteFile("preview.html", []byte(emailBody), 0644)
	return err
}
