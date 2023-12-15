package main

import (
	"fmt"
)

const (
	french  = "French"
	spanish = "Spanish"

	frenchHelloPrefix  = "Bonjour, "
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

// Les fonctions publiques commencent par une majuscule et pour les privées par une minuscule

// On veut séparer le code avec le domaine métier (ici retourner Hello, world)
// et le code qui vient de l'extérieur (fmt.Println)
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// Dans la signature de la fonction, nous avons une valeur de retour nommée qui est prefix
// Cela va créer cette variable dans la fonction avec une valeur par défaut ""
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
