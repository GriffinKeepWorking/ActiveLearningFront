package main

import(
  "fmt"
  "encoding/json"
  "log"
  "net/http"
  // "strconv"
  "bufio"
  "os"
  // "path"
  // "math/rand"
  // "strconv"
  "github.com/gorilla/mux"
  // "github.com/gorilla/websocket"
  )

// type Book struct {
//   ID string `json:"id"`
//   Author *Author `json:"author"`
// }

// type Author struct {
//   FirstName string `json:"firstname"`
//   LastName string `json:"lastname"`
// }

type imageData struct {
  URL string `json:"url"`
  Type string `json:"type"`
  Time int `json:"time"`
  ID int `json:"id"` //Task ID
}

type allURLS struct {
  // URL string `json:"url"`
  URL []string `json:"url"`
}

type realTimeData struct {
  Data []imageData `json:"data"`
}

type msg struct {
  Num int
}

// var book imageData
var sortedData imageData

var dashboardData realTimeData

// var imageArray []string

// var books []Book

func setupHandler(w *http.ResponseWriter) {
  (*w).Header().Set("Access-Control-Allow-Origin", "*")
  (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func initialiseImageData(s *imageData) {
  s.URL = "";
  s.Type = "";
  s.Time = 0;
  s.ID = 0;
}

func initialiseDashboardData(s *realTimeData) {
  var tempSlice []imageData
  s.Data = tempSlice
}

func addRealTimeData(s imageData) {
  dashboardData.Data = append(dashboardData.Data, s)
  fmt.Println("Real time data so far = ", dashboardData.Data)
}

func main() {

  // var tempUrlArray []string;
  //Read data from file
  // tempUrlArray = append(tempUrlArray, 
  //             "http://r.ddmcdn.com/s_f/o_1/cx_462/cy_245/cw_1349/ch_1349/w_720/APL/uploads/2015/06/caturday-shutterstock_149320799.jpg", 
  //             "https://amp.businessinsider.com/images/592f4169b74af41b008b5977-750-563.jpg",
  //             "https://d17fnq9dkz9hgj.cloudfront.net/uploads/2012/11/150838553-tips-for-first-30-days-dog-632x475.jpg",
  //             "https://thenypost.files.wordpress.com/2017/12/nyc-streets.jpg?quality=90&strip=all",
  //             "https://upload.wikimedia.org/wikipedia/commons/a/a6/Congress_Street%2C_Portland_ME.jpg",
  //             "https://static1.squarespace.com/static/5744c9228a65e2b59a618fd0/594d67db78d171a5cb2baa2d/594d67f7b8a79b161a9a35a9/1498245278890/OpenStreets-103.jpg",
  //             "https://static1.squarespace.com/static/556b8c0ee4b024eba5bd8f0b/t/561533d4e4b0c82ae6a5e060/1444230102463/Complete+Streets+NYC?format=750w")

  // fmt.Println("tempUrlArray: ", tempUrlArray)
  // allImageUrls = allURLS{URL: tempUrlArray}

  initialiseDashboardData(&dashboardData)

  // imageArray = image

  // images = append(images, allURLS{URL: "http://r.ddmcdn.com/s_f/o_1/cx_462/cy_245/cw_1349/ch_1349/w_720/APL/uploads/2015/06/caturday-shutterstock_149320799.jpg"});
  // images = append(images, allURLS{URL: "https://amp.businessinsider.com/images/592f4169b74af41b008b5977-750-563.jpg"});
  // images = append(images, allURLS{URL: "https://d17fnq9dkz9hgj.cloudfront.net/uploads/2012/11/150838553-tips-for-first-30-days-dog-632x475.jpg"});
  // images = append(images, allURLS{URL: "https://thenypost.files.wordpress.com/2017/12/nyc-streets.jpg?quality=90&strip=all"});
  // images = append(images, allURLS{URL: "https://upload.wikimedia.org/wikipedia/commons/a/a6/Congress_Street%2C_Portland_ME.jpg"});
  // images = append(images, allURLS{URL: "https://static1.squarespace.com/static/5744c9228a65e2b59a618fd0/594d67db78d171a5cb2baa2d/594d67f7b8a79b161a9a35a9/1498245278890/OpenStreets-103.jpg"});
  // images = append(images, allURLS{URL: "https://static1.squarespace.com/static/556b8c0ee4b024eba5bd8f0b/t/561533d4e4b0c82ae6a5e060/1444230102463/Complete+Streets+NYC?format=750w"});

  r := mux.NewRouter() // Assumes type if := is used

  r. HandleFunc("/fetchImages/{key}", fetchImages).Methods("GET")

  r.Path("/").Handler(http.FileServer(http.Dir("./dist/")))
  // r.HandleFunc("/", home).Methods("GET")
  r.Path("/index_bundle.js").Handler(http.FileServer(http.Dir("./dist/")))


  // r.Path("/").Queries("id", "{id}").HandlerFunc(fetchImages).Methods("GET")
  r.Path("/fetchImages").Queries("id", "{id}").HandlerFunc(fetchImages).Methods("GET")
  r.Path("/fetchImages").HandlerFunc(fetchImages).Methods("GET")
  // r.Path("/fetchImages").HandlerFunc(fetchImages).Methods("GET")

  r. HandleFunc("/postImages", postImages).Methods("POST")
  r. HandleFunc("/realTimeImageData", realTimeImageData).Methods("GET")

  // r. HandleFunc("/hello", getHelloGet).Methods("GET")
  // r. HandleFunc("/hello", getHelloPost).Methods("POST")

  fmt.Println("Running...")
  log.Fatal(http.ListenAndServe(":8080", r)) //Port and router

  // fmt.Println("Hello World!")
}

// func home(w http.ResponseWriter, r *http.Request) {
//   fp := path.Join("./images", "index.xml")
//   fmt.Println("Serving ", fp)
//   http.ServeFile(w, r, fp)
// }

func realTimeImageData(w http.ResponseWriter, r *http.Request) {
  setupHandler(&w);
  w.Header().Set("Content-Type", "application/json")

  // fmt.Println("Sending ", dashboardData)

  json.NewEncoder(w).Encode(dashboardData)

  initialiseDashboardData(&dashboardData)

  // fmt.Println("tempUrlArray: ", tempUrlArray)
  // allImageUrls = allURLS{URL: tempUrlArray}

}

func postImages(w http.ResponseWriter, r *http.Request) {
  // w.Header().Set("Content-Type", "application/json")
  // w.Header().Set("Access-Control-Allow-Origin", "*")
  setupHandler(&w)

  json.NewDecoder(r.Body).Decode(&sortedData)

  fmt.Println("Image: ", sortedData.URL)
  fmt.Println("Type: ", sortedData.Type)
  fmt.Println("Time taken: ", sortedData.Time)

  addRealTimeData(sortedData)

}

// func enableCors(w *http.ResponseWriter) {
//   (*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

func fetchImages(w http.ResponseWriter, r *http.Request) {
  setupHandler(&w)
  vars := mux.Vars(r)

  for key, value := range vars {
    fmt.Println("Key:", key, "Value:", value)
  }

  file, err := os.Open("tasks/task" + vars["id"] + ".txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  var tempUrlArray []string;
  for scanner.Scan() {
    //Read data from file
    tempUrlArray = append(tempUrlArray, 
                scanner.Text())
      // fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  var allImageUrls allURLS

  allImageUrls = allURLS{URL: tempUrlArray}
  // Return Images according to ID
  // Read URLs from files
  
  // fmt.Println("ID = ", vars)

  fmt.Println("Sending: ", allImageUrls)
  
  w.Header().Set("Content-Type", "application/json")
  // json.NewEncoder(w).Encode(imageArray)
  // fmt.Println("Sending ", allImageUrls)
  json.NewEncoder(w).Encode(allImageUrls)
  // fmt.Println("After encoding")

}

// go get -u github.com/gorilla/mux

// package main

// import (
//     "fmt"
//     "net"
//     "os"
// )

// const (
//     CONN_HOST = "localhost"
//     CONN_PORT = "3333"
//     CONN_TYPE = "tcp"
// )

// func main() {
//     // Listen for incoming connections.
//     l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
//     if err != nil {
//         fmt.Println("Error listening:", err.Error())
//         os.Exit(1)
//     }
//     // Close the listener when the application closes.
//     defer l.Close()
//     fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
//     for {
//         // Listen for an incoming connection.
//         conn, err := l.Accept()
//         if err != nil {
//             fmt.Println("Error accepting: ", err.Error())
//             os.Exit(1)
//         }
//         // Handle connections in a new goroutine.
//         go handleRequest(conn)
//     }
// }

// // Handles incoming requests.
// func handleRequest(conn net.Conn) {
//   // Make a buffer to hold incoming data.
//   buf := make([]byte, 1024)
//   // Read the incoming connection into the buffer.
//   _, err := conn.Read(buf)
//   if err != nil {
//     fmt.Println("Error reading:", err.Error())
//   }
//   // Send a response back to person contacting us.
//   conn.Write([]byte("Message received."))
//   // Close the connection when you're done with it.
//   conn.Close()
// }



// func realTimeHandler(w http.ResponseWriter, r *http.Request) {

//   fmt.Println("Getting data");

//   if r.Header.Get("Origin") != "http://"+r.Host {
//     http.Error(w, "Origin not allowed", 403)
//     return
//   }

//   setupHandler(&w)

//   conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
//   if err != nil {
//     http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
//   }

//   go echo(conn)
// }

// func echo(conn *websocket.Conn) {
//   for {
//     m := msg{}

//     err := conn.ReadJSON(&m)
//     if err != nil {
//       fmt.Println("Error reading json.", err)
//     }

//     fmt.Printf("Sending message: %#v\n", m)

//     if err = conn.WriteJSON("Hello"); err != nil {
//       fmt.Println(err)
//     }
//   }
// }
