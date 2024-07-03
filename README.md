# Simple go web server example

A basic Go application that demonstrates serving static files and providing simple API endpoints.

## Description

This Go application sets up a web server that serves static files from a `static` folder and provides two API endpoints:

- `/host`: Returns the hostname of the server and any error encountered during hostname retrieval.
- `/hi`: Returns a simple greeting message.
- `/index`: It demonstrates a simple home page with a background color and an image served from the `static/images` directory.

## Prerequisites

- Go installed locally. You can download it from [golang.org](https://golang.org/dl/).
- Basic understanding of Go programming and HTTP servers.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/BNeerajKrishna/my-go-app.git
  
2. Move to directory
   ```bash
   cd my-go-app
  
## Usage

1. Run the code
   ```bash
   go run main.go

2. Open a web browser and navigate to`http://localhost:5000` to view the static files served from the `static` folder.

## File Structure

- `main.go`: Entry point of the Go application.
- `static/`: Folder containing static files served by the web server.
- `static/images/`: Directory for images served by the /index endpoint.

## API Endpoints

- `/host`: Returns a response with the hostname and any error encountered during retrieval.

    response:
    ```
    Hi, hostname is <hostname> and error is <error>

- `/hi`: Returns a simple greeting message.

   response:
   ```
   Hello
   
- `/index`: Returns an HTML page with an embedded image,when accessing `/index`, the server                   returns an HTML page that includes an image served from the `static/images` directory.

   HTML response:
   ```html
   <!DOCTYPE html>
   <html>
   <head>
     <title>Index Page</title>
   </head>
   <body>
     <h2>Index page</h2>
     <p>Here is the image:</p>
     <img src="/images/one_piece.jpg" alt="One Piece Image">
   </body>
   </html>
 
## Configuration

- The server runs on port 5000 by default. You can modify the port variable in the main function to change the port number.

