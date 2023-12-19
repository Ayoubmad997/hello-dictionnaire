package main

import (
    "fmt"
)


func Remove(entries map[string]string,word string)map[string]string{
    delete(entries, word)
	return entries
}

func Get(entries map[string]string,word string)string{
	return entries[word]
}

func List(entries map[string]string)map[string]string{
	return entries
}

func main() {

    entries := map[string]string{}

    entries["MAD"] = "étudiant"
    entries["aziz"] = "prof"
    entries["estiam"] = "école"


    Remove(entries,"estiam")
	Get(entries,"MAD")
	List(entries)

    for word, definition := range entries {
        fmt.Printf("%s: %s\n", word, definition)
    }
}
