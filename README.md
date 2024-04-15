# Création d'une base de rpg composée de tuile en Golang 2023
C'est un projet pour l'université créer par Loig JEZEQUEL. Il y avait deux objectifs principaux. Le premier enregistrer la carte dans un quadtree et non pas dans un tableau de tableau. Le deuxième et d'ajouter des modulesau programme de base en pouvant facilement les activer et les désactiver grâce à un fichier config.json.

Le projet : 
<br>
Code source initial pour le projet d'introduction au développement (R1.01) et de SAÉ implémentation d'un besoin client (SAE1.01), année 2023-2024.

## Installation

Install RPG_Quadtree_Golang.2023 with git

```bash
  git clone https://github.com/FlorianGandon/RPG_Quadtree_Golang.2023.git
```
### ⚠️ If you are on Linux, you need to install :

```bash
  sudo apt install libgtk-3-dev
```


## Run Locally

Install the project

### With the bash

Go to the folder 'cmd' in the project directory

```bash
  cd RPG_Quadtree_Golang.2023/cmd/
```

Build the project

```bash
  go build .\main.go
```

Run the executable file

### With GoLand (JetBrains)

<img src="portfolio/goland configuration.png">

## Screenshots

### Generation thanks floor file

<img src="portfolio/fenêtre RPG.png">

### Random Generation

<img src="portfolio/fenêtre random.png">

### Debug Mode (key D)

<img src="portfolio/fenêtre debug.png">

### The config file in JSON

<img src="portfolio/fenêtre config.png">

## Documentation

Première partie, explication de config.json :

NumTileX et NumTileY représentent la taille en tuile de l'affichage.

FloorKind représente la structure de la carte :
- 0 le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
- 1 la carte est dans un tableau de tableau
- 2 la carte est dans un quadtree

CameraMode modifie le comportement de la caméra :
- 0 la caméra ne bouge pas
- 1 la caméra suit partiellement le joueur quand il a fini d'avancer
- 2 la caméra est fluide

StopCamera empêche la caméra à afficher des zones en dehors de la carte sauf dans le cas où le carte est plus petite que l'affichage alors la carte sera au centre.

RoundEarth répète la carte à l'infini et le joueur ne peut quitter la carte.

RandomGeneration créé un terrain aléatoirement infini en quadtree avec les options ChunkSize et Seed.
ChunkSize définit la taille d'un chunk (modifie donc le temps de chargement et la mémoire).
Seed initialise l'aléatoire
RandomGenerationSave Permet de sauvegarder et charger la carte générée (ctrl+s save et ctrl+o load)

ConnectedTile modifie les tuiles en fonctions des tuiles qui sont autour.

ActiveBigStep marche si le terrain est infini (FloorKind = 0 ou RandomGeneration = true) permet de se déplacer plus vite en restant appuyé sur Ctrl + les flèches directionnelles.
BigStep est la taille des pas.

BetterCharacter ajoute une meilleure ombre (une particule transparente) et fait voler le joueur en débug mode.
BetterBlocking empêche le joueur de marcher sur l'eau et sur les murs.

ActivePortal permet d'avoir deux portails (I pour les placer / supprimer, T pour se téléporter).
ActiveParticlesPortal ajoute 10 particules aux portails

La touche espace met le jeu en fullScreen et inversement

ActiveScrolling permet de changer la taille d'affichage avec la molette de la souris.


## Authors

- [@FlorianGandon](https://github.com/FlorianGandon)
- [@ArmelCloarec](https://github.com/Zolkennn)
