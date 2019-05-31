package route

import (
    "github.com/gorilla/mux"
    asana "github.com/heaptracetechnology/microservice-asana/asana"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "CreateProject",
        "POST",
        "/createproject",
        asana.CreateProject,
    },
    Route{
        "CreateTask",
        "POST",
        "/createtask",
        asana.CreateTask,
    },
    Route{
        "DeleteProject",
        "POST",
        "/deleteproject",
        asana.DeleteProject,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
