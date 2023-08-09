# Firebase Auth example

This example shows how to use Firebase Auth for a golang server.

## Setup

1. Create a Firebase project using the [Firebase Developer Console](https://console.firebase.google.com).
2. Enable the **Email/Password** sign-in provider in the **Authentication > SIGN-IN METHOD** tab.
3. Download a service account from the Firebase Console ([Settings > Service Accounts](https://console.firebase.google.com/project/_/settings/serviceaccounts/adminsdk)) and set the environment variable `GOOGLE_APPLICATION_CREDENTIALS` to the path to the downloaded service account credentials file.
4. Add your `FIREBASE_API_KEY` key in .env file
5. Run `go run main.go` to start the server.
6. Visit `http://localhost:8080` to view the sample.
