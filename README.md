# FizzBuzz API

Une API REST performante et sécurisée développée en Go avec le framework Gin. Elle permet de générer des séquences FizzBuzz personnalisables via des paramètres d'URL.

## Sécurité
L'API est protégée. Chaque requête doit inclure une clé valide dans le header :
X-API-Key: <votre_clé>

---

## Installation et Lancement

1. Cloner le projet :
   git clone git@github.com:LEVACHER-Louis/api-fizzbuzz.git
   cd api-fizzbuzz

2. Gérer les dépendances :
   go mod download

3. Compiler et Lancer :
   go build -o fizzbuzz-api
   ./fizzbuzz-api [flags]

---

## Configuration
Vous pouvez configurer l'accès et le réseau via des drapeaux (flags) au démarrage :

| Flag | Raccourci | Description |
| :--- | :--- | :--- |
| --key | -k | Ajoute une clé d'API (cumulable). |
| --keyfile | | Charge un fichier contenant une clé par ligne. |
| --port | -p | Port du serveur (par défaut : 8080). |

Exemple : ./fizzbuzz-api -k "admin123" --keyfile "./secrets/keys.txt" -p 9000

---

## Utilisation de l'API

### Route unique : GET /fizzbuzz

#### Paramètres de requête (Query Params) :
* nombreMax (int) : Le nombre jusqu'auquel générer la séquence. Par défaut : 15.
* mots[<diviseur>] (map) : Définit un remplacement personnalisé. 
    > Note : L'utilisation de ce paramètre désactive les comportements par défaut (3 pour Fizz et 5 pour Buzz).

---

## Exemples de requêtes

### 1. Requête standard (Valeurs par défaut)
curl -H "X-API-Key: cletest" "http://localhost:8080/fizzbuzz"

Réponse : {"reponse": "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz"}

### 2. Requête personnalisée (FizzBuzz complexe)
Générer jusqu'à 20, avec "Foo" pour les multiples de 2 et "Bar" pour les multiples de 7 :
curl -H "X-API-Key: cletest" "http://localhost:8080/fizzbuzz?nombreMax=20&mots[2]=Foo&mots[7]=Bar"

Réponse : {"reponse": "1 Foo 3 Foo 5 Foo Bar Foo 9 Foo 11 Foo 13 FooBar 15 Foo 17 Foo 19 Foo"}

---

## Architecture et Dépendances
Le projet est structuré de manière modulaire :
* internal/ : Logique métier de l'algorithme FizzBuzz.
* api/ : Handlers Gin et définition des routes.

Stack technique :
* Go 1.24.0
* Gin Gonic v1.11.0
* pflag v1.0.10
