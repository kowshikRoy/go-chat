package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`
		<html>
		<head>
			<title>Chat</title>
		</head>
		<body>
			Let's chat
		</body>
		</html>
		`))
	})

	//start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
