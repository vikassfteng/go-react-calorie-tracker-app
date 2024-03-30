# Calorie Tracker

Calorie Tracker is a web application that allows users to track their daily calorie intake. The backend is developed using Go, and the frontend is built with React and Bootstrap. MongoDB is used for the database.

## Features

- Track daily calorie intake
- Add new entries for each meal
- Update the details of an entry
- Update the ingredients of an entry
- Delete an entry

## System Architecture

1. User Interface (React + Bootstrap): This is where the user interacts with the application. It communicates with the Backend to fetch, create, update, and delete entries.

2. Backend (Go): This is the server-side of your application. It processes requests from the User Interface, performs operations on the database, and returns responses.

3. Database (MongoDB): This is where your application's data is stored. The Backend communicates with the Database to perform CRUD (Create, Read, Update, Delete) operations.

## Installation

### Backend

1. Navigate to the backend directory: `cd backend`
2. Install the dependencies: `go get`
3. Start the server: `go run main.go`

### Frontend

1. Navigate to the frontend directory: `cd frontend`
2. Install the dependencies: `npm install`
3. Start the server: `npm start`

## Usage

### Track daily calorie intake

The main page displays the total calorie intake for the day.

### Add a new entry

To add a new entry, click on the "Add Entry" button and fill out the form.

### Update an entry

To update an entry, click on the "Update Entry" button next to the entry you want to update. This will open a modal with a form. Fill out the form and click "Update".

### Update the ingredients of an entry

To update the ingredients of an entry, click on the "Update Ingredients" button next to the entry. This will open a modal with a form. Fill out the form and click "Update".

### Delete an entry

To delete an entry, click on the "Delete" button next to the entry you want to delete.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.




Here's the picture how app will look like<img width="1664" alt="Screenshot 2024-03-30 at 4 03 28 PM" src="https://github.com/vikassfteng/go-react-calorie-tracker-app/assets/70806739/fc5bb1b8-fe48-4e03-977d-4b3eb47def30">


