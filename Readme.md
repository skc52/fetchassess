git clone <repository_url>
cd FETCH_TAKE_HOME

Build the docker image

- docker build -t fetch-assess .

Run the docker container

- docker run -p 8081:8081 fetch-assess

Open Postman
Select Method as POST with the link set up as http://localhost:8081/receipts/process
Add Json Body
{
"retailer": "M&M Corner Market",
"purchaseDate": "2022-03-20",
"purchaseTime": "14:33",
"items": [
{
"shortDescription": "Gatorade",
"price": "2.25"
},{
"shortDescription": "Gatorade",
"price": "2.25"
},{
"shortDescription": "Gatorade",
"price": "2.25"
},{
"shortDescription": "Gatorade",
"price": "2.25"
}
],
"total": "9.00"
}

And Hit Send
This will create a receipt and respond back with the newly created receipt's ID

With Method set as GET with link http://localhost:8081/receipts/{id}/points where id is the receiptId, Hit send and it will respond back with points for the receipt.
{
"points": 109
}
