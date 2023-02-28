package main

import (
    "fmt"
    "main/internal"
)

func main() {
    contacts.SetSupport("Служба поддержки")
    fmt.Println(contacts.GetContact())
    fmt.Println("Email:", contacts.Email)
} 