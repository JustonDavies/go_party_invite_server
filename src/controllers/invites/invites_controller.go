//-- Package Declaration -----------------------------------------------------------------------------------------------
package invites

//-- Imports -----------------------------------------------------------------------------------------------------------
import (
	"errors"
	"fmt"
	"github.com/JustonDavies/fhtagn/secrets"
	"github.com/labstack/echo"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"net/http"
)

//-- Constants ---------------------------------------------------------------------------------------------------------

//-- Structs -----------------------------------------------------------------------------------------------------------
type RSVP struct {
	Code string `form:"invocation_code"`

	FirstName      string `form:"first_name"`
	LastName       string `form:"last_name"`
	PhoneNumber    string `form:"phone_number"`
	Email          string `form:"email"`
	MealPreference string `form:"sentience_preference"`

	Guest               string `form:"plus_one"`
	GuestMealPreference string `form:"plus_one_sentience_preference"`

	Notes string `form:"plea"`
}

//-- Exported Functions ------------------------------------------------------------------------------------------------
func Index(context echo.Context) error {
	//-- Authenticate ----------
	{
		//Taken care of via middleware
	}

	//-- Authorize ----------
	{
		//No Action required
	}

	//-- Shared variables ----------

	//-- Parse event ----------
	{

	}

	//-- Connect to database ----------

	//-- Assert Existence ----------
	{

	}

	//-- Action ---------
	{

	}

	//-- Response ----------
	{
		return context.Render(http.StatusOK, `invite_index`, nil)
	}

	return nil
}

func Read(context echo.Context) error {
	//-- Authenticate ----------
	{
		//Taken care of via middleware
	}

	//-- Authorize ----------
	{
		//No Action required
	}

	//-- Shared variables ----------
	var formData = &struct {
		InvocationCode string `form:"invocation_code"`
	}{}
	var guest *secrets.Guest

	//-- Parse event ----------
	{
		if err := context.Bind(formData); err != nil {
			return err
		} else if subject, found := guestMap()[formData.InvocationCode]; !found {
			return errors.New(fmt.Sprintf(`Failed rite... '%s' was not a valid invocation`, formData.InvocationCode))
		} else {
			guest = subject
		}
	}

	//-- Connect to database ----------

	//-- Assert Existence ----------
	{

	}

	//-- Action ---------
	{

	}

	//-- Response ----------
	{
		return context.Render(http.StatusOK, `invite_edit`, map[string]interface{}{
			"invocation_code": formData.InvocationCode,
			"guest":           guest,
		})
	}

	return nil
}

func Update(context echo.Context) error {
	//-- Authenticate ----------
	{
		//Taken care of via middleware
	}

	//-- Authorize ----------
	{
		//No Action required
	}

	//-- Shared variables ----------
	var formData = &struct {
		InvocationCode string `form:"invocation_code"`
	}{}
	var rsvp = &RSVP{}

	//-- Parse event ----------
	{
		if err := context.Bind(formData); err != nil {
			return err
		} else if _, found := guestMap()[formData.InvocationCode]; !found {
			return errors.New(fmt.Sprintf(`Invocation %s was not found`, formData.InvocationCode))
		} else if err := context.Bind(rsvp); err != nil {
			return err
		}
	}

	//-- Connect to database ----------

	//-- Assert Existence ----------
	{

	}

	//-- Action ---------
	{

		var sender = mail.NewEmail(`The Void`, `Juston.Davies@gmail.com`)
		var subject = `The pact is sealed...`
		var message = fmt.Sprintf("<p>%s %s,</p>"+
			"<p>You have completed your dark rite and your commitment will be rewarded.</p>"+
			"<p>If you have any questions or problems please feel free to reply to this email or reach out to either of the hosts and we will make...arrangements for you.</p>"+
			"<p>Your RSVP is as follows:</p>"+
			"<p>"+
			"&nbsp;Name: %s %s<br/>"+
			"&nbsp;Phone Number: %s<br/>"+
			"&nbsp;Email: %s<br/>"+
			"&nbsp;Food Preference: %s<br/>"+
			"&nbsp;Guest: %t<br/>"+
			"&nbsp;Guest Food Preference: %s<br/>"+
			"&nbsp;Appeals & Requests: %s<br/>"+
			"</p>"+
			"<p> We will see you at: 2850 W Long Ave, Littleton, CO 80120 @ 5:30 PM</p>"+
			"<p>It won't be long now...</p>",
			rsvp.FirstName,
			rsvp.LastName,
			rsvp.FirstName,
			rsvp.LastName,
			rsvp.PhoneNumber,
			rsvp.Email,
			rsvp.MealPreference,
			rsvp.Guest == `on`,
			rsvp.GuestMealPreference,
			rsvp.Notes)

		var confirmation = mail.NewSingleEmail(
			sender,
			subject,
			mail.NewEmail(`The Void`, rsvp.Email),
			message,
			message,
		)

		var notification = mail.NewSingleEmail(
			sender,
			subject,
			sender,
			message,
			message,
		)

		var client = sendgrid.NewSendClient(secrets.SendgridAPIKey)

		if result, err := client.Send(confirmation); err != nil || !(result.StatusCode >= 200 && result.StatusCode <= 299) {
			log.Println(`Action: Unable to send confirmation email -> `, result.StatusCode, err)
		}

		if result, err := client.Send(notification); err != nil || !(result.StatusCode >= 200 && result.StatusCode <= 299) {
			log.Println(`Action: Unable to send notification email -> `, result.StatusCode, err)
		}
	}

	//-- Response ----------
	{
		return context.Render(http.StatusOK, `invite_confirmed`, nil)
	}

	return nil
}

//-- Internal Functions ------------------------------------------------------------------------------------------------
func guestMap() map[string]*secrets.Guest {
	var output = map[string]*secrets.Guest{}

	for _, guest := range secrets.GuestList {
		var code = guest.Code
		if len(code) > 0 {
			output[code] = guest
		}
	}

	return output
}
