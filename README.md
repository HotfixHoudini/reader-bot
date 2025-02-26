
### WEEK-WIN
#### Project descriptiom

## Setup Instructions
To run this project locally, follow these steps:

#### Prerequisites
- Go installed on your machine.

- Steps:
Clone the repository:

```bash
git clone https://github.com/your-username/your-project-name.git
cd your-project-name
```

Install dependencies:
- Install Gin using Go’s package manager:

```bash
go get github.com/gin-gonic/gin
```

#### API Documentation
#### Base URL
- All API endpoints are accessible via:

```
http://localhost:8080/api/v1
```

### Endpoints

Get Current Date and Time
- URL:  `/datetime`
- Method: GET
  ##### Response Format:

```json
{
  "email": "ilesanmibry@gmail.com",
  "current_datetime": "2025-01-30T06:14:05Z",
  "github_url": "https://github.com/TheCodeGhinux/HNG12-zero-gin"
}
```

##### Description: This endpoint returns the current date and time in ISO 8601 format (UTC).
- Example Request:

```
GET http://localhost:3000/api/datetime
```

- Example Response:

```json
{
  "email": "ilesanmibry@gmail.com",
  "current_datetime": "2025-01-30T06:14:05Z",
  "github_url": "https://github.com/TheCodeGhinux/HNG12-zero-gin"
}
```

- Example Usage
Fetching the Current Date and Time using cURL:
```bash
curl http://localhost:8080/api/
```

- Using JavaScript (fetch API):
```javascript
fetch('http://localhost:8080/api/datetime')
  .then(response => response.json())
  .then(data => console.log(data));
```

### License
This project is licensed under the MIT License - see the LICENSE file for details.

