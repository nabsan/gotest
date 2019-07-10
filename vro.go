package main

import (
    "encoding/json"
    "fmt"
    "strings"
    "io/ioutil"
    "log"
)

/** JSONデコード用に構造体定義 */
type Person struct {
    Id       int    `json:"id"`
    Name     string `json:"name"`
    Birthday string `json:"birthday"`
}

func main() {
    // JSONファイル読み込み
    bytes, err := ioutil.ReadFile("vro.json")
    if err != nil {
        log.Fatal(err)
    }
    // JSONデコード
    var persons []Person
    if err := json.Unmarshal(bytes, &persons); err != nil {
        log.Fatal(err)
    }
    // デコードしたデータを表示
    for _, p := range persons {
        fmt.Printf("%d : %s\n", p.Id, p.Name)
    }

    fmt.Printf("------------mid ------------- \n")
    var decode_data interface{}
    if err := json.Unmarshal(bytes, &decode_data); err != nil {
        log.Fatal(err)
    }
    // 表示
    for _, data := range decode_data.([]interface{}) {
        var d = data.(map[string]interface{})
        e := strings.Replace(d["name"].(string),"waka", "saka",1)
        f := strings.Replace(e,"ka", "ta",1)
        fmt.Printf("%d : %s (%s) (%s) \n", int(d["id"].(float64)), d["name"], e, f)

    }
    fmt.Printf("------------e n d2 ------------- \n")
}
