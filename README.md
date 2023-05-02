# Simple Web Server Built With Go and Gin
Hi everyone! I have been learning Go for the past week and decided to work on my first web server just to practice the basics.

## To test the server 

```sh
$ git clone https://github.com/Davidjustice28/first-go-server.git
$ go run server.go 
```

### Routes

\- GET /colors/ (Retrieves all colors)
<br>
\- POST /color/:color (Adds a new color)
<br>
\- PUT /colors/:index/:color (Change a color at a given index)

\**Currently is very basic, but as I continue learning Go, I will be adding more complexity to the API*

### Coming Soon
1. CRUD operations with a NoSQL or SQL Database
2. Third party API integration
3. Create/Utilize middleware
4. Organize file structure as it gets bigger
5. Provide helpful documentation