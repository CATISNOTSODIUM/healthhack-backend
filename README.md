# Backend for healthhack 
## Todo list
- [x] Set up database schema
- [ ] User queries, mutation, Google authentication
- [ ] API adding history
- [ ] API Connect to open ai
### Issues
- **Cannot connect to database:** Make sure not to use School Wifi to connect to database. You can use VPN to bypass though.
## Set up
Make sure to configure your `.env` file before starting the server.
```
PORT=YOUR_PORT_NUMBER (DEFAULT = 8080)
DSN=[YOUR_POSTGRESQL_URL]
OPENAI_TOKEN=[YOUR_TOKEN] # for text extraction
```
Then, you can download relevant libraries and start the server.
``` bash
go mod download
go run cmd/server/main.go
```
## Database Design (Subjected to change)
You guys can edit the design from 
![alt text](figure/database.png)