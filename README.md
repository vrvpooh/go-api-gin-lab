# Student API with Gin

## How to Run
1. Clone this repository to your local machine or download a ZIP file and extract.
2. Ensure you have Go installed.
3. Open your terminal in the project root folder (go-api-gin-lab).
4. Install dependencies:
   ```bash
   go mod tidy
   ```
5. Run the server:
   ```bash
   go run main.go
   ```
The server will start on http://localhost:8080/students

## Available API endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/students` | Create a new student |
| GET | `/students` | Get all students |
| GET | `/students/:id` | Get student by ID |
| PUT | `/students/:id` | Update an existing student |
| DELETE | `/students/:id` | Delete a student |

## Test the API
### Testing with Postman / Thunder Client (Extensions on VScode)
#### 1. Create Student (POST)

Endpoint: `POST /students`

URL: http://localhost:8080/students

Method: POST

Request Body (JSON):
```json
{
  "id": "66090001",
  "name": "John Doe",
  "major": "Computer Science",
  "gpa": 3.80
}
```

Success: Status 201 Created

Validation Fail: Status 400 Bad Request (e.g., GPA > 4.0 or empty name)

#### 2. Get All Students (GET)

Endpoint: `GET /students`

URL: http://localhost:8080/students

Method: GET

Success: Status 200 OK

#### 3. Get student by ID

Endpoint: `GET /students/:id`

URL: http://localhost:8080/students/:id (/students/66090001)

Method: GET

Success: Status 200 OK

#### 4. Update Student (PUT) (Challenge 1 in Lab4)

Endpoint: `PUT /students/:id`

URL: http://localhost:8080/students/:id (/students/66090001)

Method: PUT

Request Body (JSON):
```json
{
  "id": "66090001",
  "name": "John Meow",
  "major": "CS",
  "gpa": 3.50
}
```

Success: Status 200 OK

Not Found: Status 404 Not Found (If ID does not exist)

#### 5. Delete Student (DELETE) (Challenge 2 in Lab4)

Endpoint: `DELETE /students/:id`

URL: http://localhost:8080/students/:id (/students/66090001)

Method: DELETE

Success: Status 204 No Content

Not Found: Status 404 Not Found




