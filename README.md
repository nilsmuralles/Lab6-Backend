# Series Tracker API 🖥️ 
The Series Tracker API is designed to help users manage and track their backlog of series. It provides a set of RESTful endpoints that allow the application to communicate with a database, retrieve series-related data, and send responses in JSON format.

With this API, users can add, update, and remove series from their watchlist, track their progress, and organize their series efficiently.

## Features 📦
- ✅ CRUD Operations – Create, read, update, and delete series in the database.
- ✅ Track Series Progress – Update the status, episode, and ranking of a series.
- ✅ Retrieve Series Data – Fetch all series or get details of a specific one.
- ✅ RESTful API – Uses standard HTTP methods for easy integration.
- ✅ JSON Responses – Structured and easy-to-use data format.
- ✅ Database Integration – Seamless communication with a relational database.
- ✅ Swagger Documentation – API is documented with Swagger for easy exploration.
- ✅ Fully Dockerized – The API and database run in containers for easy deployment.
- ✅ One-Command Setup – Run the entire stack with:

## Getting started 🚀
### Prerequisits 🛠️
- Docker installed for utilizing the app 🐳
- Git to be able to clone this repository 🐱

### Installation 💾
1️⃣ Clone the repository and cd into the project
```sh
git clone https://github.com/nilsmuralles/Lab6-Backend.git
cd Lab6-Backend
```

2️⃣ Build and run the Docker compose file
```bash
docker compose up --build
```

3️⃣ Explore the API
- Make requests to [http://localhost:8080/api](http://localhost:8080/api).
- Check the Swagger documentation through [Swagger UI](http://localhost:8080/swagger/index.html).

4️⃣ Stop the API
```bash
docker compose down
```

## Usage 📌
| Method  | Endpoint               | Description |
|---------|------------------------|-------------|
| `GET`   | `/series`                | Get all the series from the tracker. |
| `GET`   | `/series/:id`            | Fetch details of a specific series by its ID. |
| `POST`  | `/series`                | Add a series to the tracker. |
| `PUT`   | `/series/:id`            | Update the details of an existing series. |
| `DELETE`| `/series/:id`            | Remove a series from the database. |
| `PATCH` | `/series/:id/:direction` | Use **upvote** or **downvote** to upgrade or downgrede the ranking |
| `PATCH` | `/series/:id/status`     | Update the watch status of a series (e.g., Watching, Completed, Dropped). |
| `PATCH` | `/series/:id/episode`    | Update the last watched episode of a series. |

## Running frontend 🖼️
![Image](https://github.com/user-attachments/assets/539d1060-991f-4966-9a32-ec35c7e03984)
