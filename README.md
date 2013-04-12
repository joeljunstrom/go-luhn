Generating and validating Luhn numbers in GO
=======

    import (
      "github.com/joeljunstrom/go-luhn"
      "fmt"
    )

    // Checking if a string is a valid luhn
    luhn.Valid("1234") //= false
    luhn.Valid("562246784655") //= true

    // Generating a valid luhn string of a specified size
    randomLuhn := luhn.Generate(12)
    fmt.Println(randomLuhn) //= "802252051072"

    // Generating a valid luhn string of a specified size
    // with a given prefix
    randomLuhnWithPrefix := luhn.GenerateWithPrefix(10, "12345")
    fmt.Println(randomLuhnWithPrefix) //= "1234533220"

