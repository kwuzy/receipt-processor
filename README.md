# receipt-processor
 
## Go Version: 1.23.3

# How to Run
1. Clone the project ```git clone https://github.com/kwuzy/receipt-processor.git```
2. Change directory ```cd receipt-processor```
3. Install dependencies ```go mod tidy```
4. Run server ```go run main.go```
5. Hit the endpoints however you like (e.g. postman, cURL, etc.). There are sample cURLs further below

# Additional Notes
1. Decided to include a regular GET because it just felt right - and helped me test. If we were to go into production we could just not expose this endpoint but I imagine other services could find it useful
2. Consider addition validating on POST `/process`. I added comments in the POST endpoint for these but I chose not to do these because I decided this is just a system that stores and evaluates the points, we could have another system that actually checks validity more thoroughly.
    1. Check if total matches item total
    2. Check if the exact same receipt's already been uploaded
    3. Check if the date/time are valid
3. Consider not erroring when determining points - if we want to give the user the benefit of the doubt. I decided to stop at errors right now because I figured this doesn't apply full validations but it does apply some.
4. Two things I think are very out of scope for now but should be noted
    1. Rate limiting - I assume this will get hit a lot
    2. Tests - Would definitely need tests if we were to deploy this

## POST
### Process Receipt
```
curl -X POST http://localhost:8080/receipts/process -d '{
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
## GET
### Get Receipt by ID
```
curl http://localhost:8080/receipts/{id}
```
### Get Receipt Points
```
curl http://localhost:8080/receipts/{id}/points
```