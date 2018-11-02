### Purpose:
  This project exists as a simple web application for displaying information about a planned event and allowing guests to RSVP using unique codes and email confirmations to both the guest and the host.

  This project was also designed to have minimal software and hardware requirements, be platform agnostic, and self contained.

  In this case, the planned event is a Halloween party.

### Screenshot:
![Screen Shot](https://raw.githubusercontent.com/JustonDavies/go_party_invite_server/master/assets/images/screenshot.png "Screen Shot")

### Dependencies:
  This project depends on Go, vgo and some small selection of supporting tools.

  - Go: It is recommended you install Go `1.11.0` via `goenv` Instructions for installing can be found [here](https://github.com/syndbg/goenv) or by running the following commands:
    ```
    git clone https://github.com/syndbg/goenv.git ~/.goenv

    echo `export GOENV_ROOT="$HOME/.goenv"
          export PATH="$GOENV_ROOT/bin:$PATH"
          export PATH="$GOENV_ROOT/shims:$PATH"
          eval "$(goenv init -)"
          export GOPATH="$HOME/workspace/go"` > ~/.bash_profile
    source ~/.bash_profile

    goenv install 1.10.3
    goenv global  1.10.3
    ```

  - vgo: Dependencies are managed via the Go 1.11 built-in dependency management tool vgo to download dependencies needed for building:
    ```
    go mod init
    go mod tidy
    go mod download
    ```

### Configuration:
  Minimal configuration is required for this project, just download the dependencies (as stated above) and follow the instructions for the `secrets` below to get started.

### Secrets / Infrastructure:
  This project assumes that the predicate supporting services have already been deployed.
  The project expects the supporting information to be described in an appropriate `secrets/secrets.go` file that follows / requires the following format:

  ```
package secrets

const SendgridAPIKey = `...`

var GuestList = []*Guest{
	{ `code`, `first_name`, `last_name`, `phone_number`, `email_address`, false},
}

type Guest struct {
	Code string

	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string

	GuestPermitted bool
}
  ```

### Deploying:
  To compile (and run) this code and deploy you need only build the binary via the following commands:

  ```
  make
  bin/fhtagn
  ```

### Output:
  After deploying a web application will be available for use at the stated URL provided by the log output.

#### Authentication
  This application requires no authentication to use, connections over SSL is strongly encouraged for security purposes

#### Endpoints
This application has just one HTTPS endpoint

`GET /`
  - Parameters:
    - URL: This endpoint will not acknowledge URL encoded parameters
    - Body: This endpoint will not acknowledge and parameters included in the body
  - Exceptions:
    - Internal Error: If the endpoint hits a critical error while encoding the results, or reading user attributes it will return a user friendly error page and a status 500 RFC 7231, 6.6.1
  - Return:
    - If no errors are encountered the endpoint will return a web form with event details and a status 200 RFC 7231, 6.3.1

`POST /attune`
  - Parameters:
    - URL: This endpoint will not acknowledge URL encoded parameters
    - Body: This endpoint expects a form-data encoded invocation included in the BODY of the request
  - Exceptions:
    - Not Found: If the endpoint is unable to find a guest with the provided code it will return a user friendly error page and a status 404 RFC 7231, 6.6.1
    - Internal Error: If the endpoint hits a critical error while encoding the results, or reading user attributes it will return a user friendly error page and a status 500 RFC 7231, 6.6.1
  - Return:
    - If no errors are encountered the endpoint will return a web form with the guest details and an ability to RSVP with a status 200 RFC 7231, 6.3.1

`POST /invoke`
  - Parameters:
    - URL: This endpoint will not acknowledge URL encoded parameters
    - Body: This endpoint expects a form-data encoded RSVP with invocation code included in the BODY of the request
  - Exceptions:
    - Not Found: If the endpoint is unable to find a guest with the provided code it will return a user friendly error page and a status 404 RFC 7231, 6.6.1
    - Internal Error: If the endpoint hits a critical error while encoding the results, or reading user attributes it will return a user friendly error page and a status 500 RFC 7231, 6.6.1
  - Return:
    - If no errors are encountered the endpoint will return a web form with a confirmation of the RSVP with a status 200 RFC 7231, 6.3.1

### Current Deployment
This application is currently deployed and supported by infrastructure at: https://void.justondavies.net