# FizzBuzz API 
Une API permettant de générer des séquences fizzbuzz.
L'API est protégée par des clés présente dans le header de requete "X-API-key".

## Installation 
1. Récupérer les sources
2. Lancer la compilation : `./build`
3. Lancer le programme : `./api-fizzbuzz`

## Utilisation
L'api dispose d'une route: 

`GET /fizzbuzz`

paramètres : 

 - nombreMax : Le dernier nombre évalué. Par défaut 15.

## Exemple

`curl -H "X-API-key: cletest" "http://localhost:8080/fizzbuzz`

Renvoie : 

`{"reponse" : "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz"}`
