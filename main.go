package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
)

func main() {
	// Function to fetch the outbound IP address
	getOutboundIP := func() net.IP {
		conn, err := net.Dial("udp", "8.8.8.8:80")
		if err != nil {
			fmt.Println("Error determining IP:", err)
			return nil
		}
		defer conn.Close()

		localAddr := conn.LocalAddr().(*net.UDPAddr)
		return localAddr.IP
	}

	// Handler for the homepage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Fetch the IP address
		ipAddress := getOutboundIP()

		// Prepare data for the template
		data := struct {
			IPAddress string
		}{
			IPAddress: ipAddress.String(),
		}

		// Parse and execute the HTML template
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, data)
	})

	// Serve static files (like images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Specify the port and start the server
	port := ":5000"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
