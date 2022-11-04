# shopbridge-go
Shopbridge API problem solution in Go. Just a simple product inventory API made with Gin, GORM, Postgresql and ofc GO.


## Endpoints:

| Endpoint | HTTP Method | Usage |
| :---         |     :---:      |  :---: |
| /item   | POST | Enter new item into Inventory   |
| /item | GET | Return all items in the inventory |
| /item/<item_id> | GET | Fetch data for a single item based on ID |
| /item/<item_id> | PUT | Update fields for item based on ID |
| /item/<item_id> | DELETE | Delete item from Inventory based on ID |
