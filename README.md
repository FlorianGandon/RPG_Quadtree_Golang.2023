# project-quadtree
Code source initial pour le projet d'introduction au développement (R1.01) et de SAÉ implémentation d'un besoin client (SAE1.01), année 2023-2024.


**Paquet à installer !!**
sudo apt install libgtk-3-dev

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

Projet de :
 - Armel CLOAREC
 - Florian GANDON
