package main

/*
cat config.json
{
   "nome": "Test",
   "email": "test@test.com",
   "sexo": "F",
   "idade": 22,
   "outros": "",
}
*/

import (
    "fmt"
    "os"
    "json"
    "io/ioutil"
)

type jsonobject struct {
    nome string
    email string
    sexo string
    idade int
    outros string
}

// Main function
// I realize this function is much too simple I am simply at a loss to

func main() {
    file, e := ioutil.ReadFile("./config.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))

    //m := new(Dispatch)
    //var m interface{}
    var jsontype jsonobject
    json.Unmarshal(file, &jsontype)
    fmt.Printf("Results: %v\n", jsontype)
}
