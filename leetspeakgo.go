package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// lectureJson lit le fichier JSON spécifié et renvoie son contenu sous forme de map[string]string.
func lectureJson(filename string) map[string]string {
	// Lecture du contenu du fichier JSON
	jsonBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Création d'une map pour stocker les données JSON
	leetMap := make(map[string]string)
	// Décodage du contenu JSON dans la map
	err = json.Unmarshal(jsonBytes, &leetMap)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return leetMap
}

// conv convertit une chaîne de caractères en leet speak en utilisant les données du fichier JSON.
func conv(mot string) string {
	// Lecture du fichier JSON contenant les correspondances leet
	ref := lectureJson("./leetspeak.json")
	var lettre []string

	// Parcours de chaque caractère de la chaîne d'entrée
	for _, char := range mot {
		// Si le caractère existe dans la map, le remplacer par sa version leet
		if val, ok := ref[string(char)]; ok {
			lettre = append(lettre, val)
		} else {
			// Si le caractère n'a pas de version leet, le laisser tel quel
			lettre = append(lettre, string(char))
		}
	}

	// Concaténation des caractères convertis en une seule chaîne
	result := ""
	for _, l := range lettre {
		result += l
	}

	return result
}

func main() {
	var mot string

	// Vérifie si un argument a été passé en ligne de commande
	if len(os.Args) > 1 {
		// Si oui, utilise le premier argument comme mot à convertir en leet speak
		mot = os.Args[1]
	} else {
		// Si aucun argument n'a été passé, demande à l'utilisateur de saisir un mot
		fmt.Print("Veuillez saisir quelque chose : ")
		_, err := fmt.Scanln(&mot)
		if err != nil {
			fmt.Println("Erreur de saisie :", err)
			return
		}
	}

	// Conversion de la chaîne en leet speak
	resultat := conv(mot)
	// Affichage du résultat
	fmt.Println("Résultat de la conversion:", resultat)
}
