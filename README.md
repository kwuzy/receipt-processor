# receipt-processor

# Dependencies 
1. Go Version: 1.23.3

# Running
1. Clone the project
2. Install dependencies ```go mod init receipt-processor```
3. Run server ```go run main.go```
4. Hit the endpoints however you like (e.g. postman, cURL, etc.). Here's a sample cURL
```
curl -X POST http://localhost:8080/receipt -d '{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}' -H "Content-Type: application/json"
```