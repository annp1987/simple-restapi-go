package main
import (
        "encoding/json"
        "log"
        "fmt"
        // "net/http" has method to implement HTTP clients and servers
        "net/http"
        // route http request
        "github.com/gorilla/mux"
)

// staff struct
type Staff struct {
        Id string `json:"id`
        Name string `json:"name`
}

// list stafffs
var list_staffs []Staff

func get_staffs(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(list_staffs)
}

func get_staff(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        params := mux.Vars(r)
        for _, item := range list_staffs {
                if item.Id == params["id"]{
                        json.NewEncoder(w).Encode(item)
                        return
                }
        }
        // if not found return {}
        fmt.Fprintf(w, "Resource not found!\n")
}

func create_staff(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        var staff Staff
        _ = json.NewDecoder(r.Body).Decode(&staff)
        list_staffs = append(list_staffs, staff)
        json.NewEncoder(w).Encode(staff)
}

func delete_staff(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        params := mux.Vars(r)
        for index, item := range list_staffs {
                if item.Id == params["id"] {
                        list_staffs = append(list_staffs[:index], list_staffs[index+1:]...)
                        return
                }
        }
        fmt.Fprintf(w, "Resource not found!")
}

func main() {
        router := mux.NewRouter()
        list_staffs = append(list_staffs, Staff{Id: "1", Name: "abc"})
        list_staffs = append(list_staffs, Staff{Id: "2",Name: "xyz"})
        router.HandleFunc("/api",get_staffs).Methods("GET")
        router.HandleFunc("/api/{id}", get_staff).Methods("GET")
        router.HandleFunc("/api", create_staff).Methods("POST")
        router.HandleFunc("/api/{id}", delete_staff).Methods("DELETE")
        log.Fatal(http.ListenAndServe(":8000", router))
}
