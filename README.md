# Groupie Trackers
Groupie Trackers is a Go-based backend application that interacts with a RESTful API to fetch and manipulate data about musical artists, their concert locations, dates, and relationships. This project aims to create a user-friendly website that visualizes this data effectively.

## 📌 Project Overview
Groupie Tracker is a web application that fetches and displays data about music artists, their concerts, and related information from an API. Users can browse through artist details and concert locations in a structured and interactive manner.

## Objectives
The application connects to an API with four main parts:

- Artists: Information about bands and artists, including their names, images, formation year, first album date, and members.
- Locations: Locations of their past and/or upcoming concerts.
- Dates: Dates of their past and/or upcoming concerts.
- Relation: Links between artists, dates, and locations.
- The goal is to build a website that displays this information using various data visualizations such as blocks, cards, tables, lists, and graphics.

## Features
Fetch and display artist information, concert dates, and locations.
Visualize data through different UI components.
Handle client-server communication effectively.
Implement features based on client-triggered actions.
## Technology Stack
Backend: Go (Golang)
Frontend: HTML/CSS for the user interface
API: RESTful API for data retrieval
# Groupie Tracker



## 💁️ Project Structure
The project follows a modular structure to ensure clean and maintainable code:

```
GROUPIE-TRACKER/
│── cmd/
│   └── main.go              # Entry point of the application
│
├── handler/                 # Handles HTTP requests
│   ├── Detail_Func.go       # Handles artist details page
│   ├── Groupie_Func.go      # Handles main page request
│   ├── Style_Func.go        # Manages styles for the website
│
├── helpers/
│   ├── fetchingById.go      # Fetches data by artist ID
│   ├── renderTemplates.go   # Utility for rendering templates
│
├── routes/
│   ├── routes.go            # Handles routing for the application
│
├── static/
│   ├── images/              # Stores static images
│   ├── card_Detail.css      # Styling for artist details page
│   ├── index.css            # Styling for homepage
│   ├── status_Page.css      # Styling for status/error pages
│
├── template/                # HTML templates
│   ├── detailsCard.html     # Template for artist details
│   ├── index.html           # Homepage template
│   ├── statusPage.html      # Error/status page template
│
├── tools/
│   ├── Tools.go             # Contains data structures and utility functions
│
├── go.mod                   # Go module file
└── README.md                # Project documentation
```

## 🛠 Features
- Fetch and display artist information dynamically.
- View artist details, including name, image, members, and concerts.
- Handle 404 and error pages gracefully.
- Clean and structured project organization.

## 🚀 Installation & Setup
### Prerequisites
- Install [Go](https://go.dev/)

### Steps
1. Clone this repository:
   ```sh
   git clone https://github.com/yourusername/groupie-tracker.git
   cd groupie-tracker
   ```
2. Initialize Go modules:
   ```sh
   go mod tidy
   ```
3. Run the project:
   ```sh
   go run cmd/main.go
   ```
4. Open your browser and visit:
   ```
   http://localhost:8080
   ```

## 📺 API Source
This project fetches data from the [Groupie Tracker API](https://groupietrackers.herokuapp.com/api/).

## 🛠 Technologies Used
- **Go**: Backend development
- **HTML & CSS**: Frontend
- **net/http**: Server and routing
- **encoding/json**: Handling API responses
- **text/template**: HTML rendering

## 📝 License
This project is open-source and available under the [MIT License](LICENSE).

