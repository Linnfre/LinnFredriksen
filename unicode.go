package unicode

import (
  "fmt"
)

var oversetter string

// Kode for Oppgave 4a
func Translate(expression string, language string) string {
  if language == "is" && expression == "nord og sør" {
    oversetter :=  "\xE2\x80\x9C" + expression + "\xE2\x80\x9D" + " oversatt til islandsk blir: " + 
    "\xE2\x80\x9C\x6E\x6F\x72\xF0\x75\x72\x20\x6F\x67\x20\x73\x75\xF0\x75\x72\xE2\x80\x9D\n"
    return oversetter
  }
  if language == "jp" && expression == "nord og sør" {
    oversetter :=  "\xE2\x80\x9C" + expression + "\xE2\x80\x9D" + " oversatt til japansk blir: " + "\xE2\x80\x9C\x5317\x3068\x5357\xE2\x80\x9D\n"
    return oversetter
  }

  return oversetter
}

// Kode for Oppgave 4b
func UnicodeCodePointDemo() {
	// Hva er dette?
	// Er det likt på MS Windows og macOS?
  fmt.Println("\xf0\x9F\x98\x80")
  fmt.Println("\xf0\x9F\x98\x97")
  // Demonstrerer at deler av et tegn representert med flere bytes
  // kan ikke tolkes innenfor koden (unicode)
  fmt.Println("\xf0\x9F\x98")
  fmt.Println("\xf0\x9F")
  fmt.Println("\xf0")
}

