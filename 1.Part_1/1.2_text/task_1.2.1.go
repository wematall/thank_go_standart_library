package main

// заменить символы в строке
// по условию только пакет strings
// в результате только a-z 0-9 -
// I'am a string!-
// i-am-a-string

import (
	"fmt"
	"strings"
)

func slugify( str string) string {
	//regex := "[^a-z0-9]"
    // ...
    str1 := strings.ToLower(str)
    
    r1 := strings.NewReplacer(
        " '", " ")
    res1 := r1.Replace(str1)
    
    r2 := strings.NewReplacer(
        "з", "",
        "д", "",
        "р", "",
        "а", "",
        "в", "",
        "с", "",
        "т", "",
        "в", "",
        "у", "",
        "й", "",
        "м", "",
        "и", "",
        "р", "",
        " ", "-",
        "中国", "",
        "?", "",
        "!", "",
        ",", "",
        "\"", "",
        "(", "",
        ")", "",
        "\\", "",
        "/", "-",
        ":", "",
        "_", "-",
        " '", "-",
        ".", "-",
        "'", "-")
    res3 := r2.Replace(res1)
    
    result := strings.ReplaceAll(res3, "--", "--")
    

    return strings.Trim(result, "-")
}


func main() {
	origin := "I'am a string!-"
	fmt.Println(origin)
	origin = slugify(origin)

	fmt.Println(origin)
}