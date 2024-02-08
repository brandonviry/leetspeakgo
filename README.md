# leetspeakgo converter

Ce programme Go permet de convertir un texte en leetspeak, une forme d'écriture alternative utilisée sur Internet.Ceci est un petit exercice d'entrainement et c'est mon premier programme en go .

## Fonctionnalités

Conversion des mots en leetspeak en utilisant une table de conversion prédéfinie.
Prise en charge de la conversion de texte saisi par l'utilisateur ou passé en ligne de commande.
Utilisation de la bibliothèque JSON pour charger la table de conversion depuis un fichier JSON.



## Fichier 


### leetspeak.json

Ce fichier JSON contient la table de conversion utilisée pour convertir les caractères en leetspeak.

### build.bat 

Ce fichier permet de compiler le source go pour plusieur OS 

### go.mod

### leetspeakgo.go 

Ce fichier contient la fonction de lecture de fichier json ,  pour la conversion d'un text  en leetspeak. Il charge les règles de conversion à partir du fichier leetspeak.json e, le code principal du programme et permet à l'utilisateur de saisir du texte à convertir en leetspeak.

### run_for_linux , run_for_mac , run_for_windows 

Ces fichier sont des lanceur de programme pour 3 difeerent: os linux , mac,windows .

### leetspeakgo.exe , leetspeakgo_macos , leetspeakgo_linux

Ces fichier sont les programme pour 3 difeerent os: linux , mac,windows .

## Utilisation

Pour utiliser ce programme :

Pour linux:

* Exécutez run_for_linux .
* Saisissez le texte que vous souhaitez convertir en leetspeak lorsque vous y êtes invité.
* Le texte converti sera affiché à l'écran.

Pour macos:

* Exécutez run_for_mac .
* Saisissez le texte que vous souhaitez convertir en leetspeak lorsque vous y êtes invité.
* Le texte converti sera affiché à l'écran.

Pour windows:

* Exécutez run_for_windows .
* Saisissez le texte que vous souhaitez convertir en leetspeak lorsque vous y êtes invité.
* Le texte converti sera affiché à l'écran.

Si vous souhaitez passer le texte en ligne de commande, vous pouvez également exécuteravec le texte en argument.

Exemple d'utilisation en ligne de commande :

* pour linux:
```bash
./leetspeakgo_linux hello
```

* pour macos:
```bash
./leetspeakgo_macos hello
```

* pour windows:

```bat
leetspeakgo.exe hello
```


## Table de conversion

La table de conversion utilise des substitutions pour convertir chaque lettre en leetspeak. Voici un aperçu de la table de conversion :

```json
{
  "a": "4",
  "b": "8",
  "c": "(",
  "d": ")",
  ...
}
```

## Auteur

Ce jeu a été créé par VIRY brandon alias Chikara974.

## Licence

Ce projet est sous AUCUNE licence .
