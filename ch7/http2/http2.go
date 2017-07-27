func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    switch req.URL.Path {
    case "/list":
        for item, price := range db {
            fmt.Fprintf(w, "%s: %s\n", item, price)
        }
    case "/price":
        item := req.URL.Query().Get("item")
        price, ok := db(item)
        if !ok {
            w.WriteHeader(http.StatusNotFound)
            fmt.Fprintf(w, "no such item: %q\n", item)
            return
        }
    default:
        w.WriterHeader(http.StatusNotFound)
        fmt.Fprintf(w, "no such page: %s\n", req.URL)
    }
}
