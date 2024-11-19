## Challenge Overview

The goal of this challenge was to build a web service that processes a receipt, calculates points based on predefined rules, and returns the total points for that receipt. I have implemented two API endpoints:

1. /receipts/process (POST): Receives a receipt, processes it, and returns an ID that uniquely identifies that receipt.
2. /receipts/{id}/points (GET): Accepts the receipt ID and returns the points calculated based on various rules.

## Endpoints

1. POST /receipts/process
   - Purpose: To submit a receipt for processing.
   - Input: JSON object containing receipt details (retailer name, purchase date, items, total, etc.).
   - Output: JSON object with a unique receipt ID.

2. GET /receipts/{id}/points
   - Purpose: To retrieve the points awarded for a processed receipt.
   - Input: The receipt ID (from the POST request).
   - Output: JSON object with the total points.


## How to Run

### Using Docker

1. Build the Docker Image:
   
   -> build the Docker image using the following command:
   docker build -t fetch_exercise .
   
   ->Run the Docker Container: Once the image is built, run the container:
   docker run -p 8080:8080 fetch_exercise

This will start the server and it will be accessible at http://localhost:8080.


2. Using Go (without Docker)

Install Dependencies: If you are not using Docker, run the following to install required dependencies:

-> go mod download

-> go run main.go

This will also make the server accessible at http://localhost:8080.

## Testing the API

1. Submit a Receipt (POST /receipts/process)
URL: http://localhost:8080/receipts/process
Method: POST
Body: JSON containing receipt details
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "total": "35.35",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },
    {
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },
    {
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },
    {
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },
    {
      "shortDescription": "Klarbrunn 12-PK 12 FL OZ",
      "price": "12.00"
    }
  ]
}
Example Response:

{
  "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
}


2. Get Points for a Receipt (GET /receipts/{id}/points)
URL: http://localhost:8080/receipts/{id}/points (Replace {id} with the generated receipt ID)
Method: GET
Example Request:
http://localhost:8080/receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points
Example Response:
{
  "points": 28
}

## Conclusion
This service allows you to process a receipt, calculate points based on predefined rules, and retrieve the points for each receipt using its unique ID. You can test the service using tools like Thunder Client or Postman by interacting with the provided endpoints.

If you encounter any issues or have questions, feel free to reach pratheeksha.rajashekar@gmail.com for further assistance.
