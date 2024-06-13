# Party Invitation Platform(RSVP Now)

* Welcome to the Party Invitation Platform! 
* This web application allows users to RSVP for a party, view the list of attendees, and see personalized messages based on their RSVP status.

## Features

- **Welcome Page:** A friendly invitation to RSVP for the party.
- **RSVP Form:** Collects attendee details including name, email, phone number, and attendance status.
- **Thank You Page:** A personalized thank you message for those who accepted party invite.
- **Sorry Page:** A personalized message for those who declined party invite.
- **Attendee List:** A list of attendees who have confirmed their attendance.

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- A web browser

### Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/party-invitation-platform.git
```

```bash
cd party-invitation-platform
```
2. Create HTML templates:

Create the following HTML files in the same directory as your Go file.

* layout.html
* welcome.html
* form.html
* list.html
* thanks.html
* sorry.html

Create the following go files:

* main.go
* go.mod

3. Run the project on the terminal 

```bash
go run . 
```

## Author

[Benard Opiyo](https://github.com/benardopiyo)
