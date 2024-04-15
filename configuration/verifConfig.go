package configuration

import (
	"errors"
	"log"
)

func verifConfig(global Configuration) {
	if global.RandomGeneration && global.StopCamera {
		log.Fatal(errors.New("le fichier de config n'est pas valide, RandomGeneration n'est pas compatible avec StopCamera"))
	}
	if global.FloorKind == 1 && global.RandomGeneration {
		log.Fatal(errors.New("le fichier de config n'est pas valide, FloorKind 1 n'est pas compatible avec RandomGeneration"))
	}
	if global.StopCamera && global.RoundEarth {
		log.Fatal(errors.New("le fichier de config n'est pas valide, StopCamera n'est pas compatible avec RoundEarth"))
	}
	if global.NumTileForDebug < 1 {
		log.Fatal(errors.New("le fichier de config n'est pas valide, NumTileForDebug doit être positif"))
	}
	if global.FloorKind != 0 && global.FloorKind != 1 && global.FloorKind != 2 {
		log.Fatal(errors.New("le fichier de config n'est pas valide, FloorKind doit être de 0, 1 ou 2"))
	}
	if global.CameraMode != 0 && global.CameraMode != 1 && global.CameraMode != 2 {
		log.Fatal(errors.New("le fichier de config n'est pas valide, CameraMode doit être de 0, 1 ou 2"))
	}
	if global.NumTileY < 2 || global.NumTileX < 2 {
		log.Fatal(errors.New("le fichier de config n'est pas valide, NumTile doivent être supérieur à 1"))
	}
	if global.RandomGeneration && global.ChunkSize < 2 {
		log.Fatal(errors.New("le fichier de config n'est pas valide, ChunkSize doivent être supérieur à 1"))
	}
	if global.TileSize < 1 {
		log.Fatal(errors.New("le fichier de config n'est pas valide, TileSize doivent être positif"))
	}
	if global.NumFramePerCharacterAnimImage < 1 || global.NumFramePerCharacterAnimImage > 5 {
		log.Fatal(errors.New("le fichier de config n'est pas valide, NumFramePerCharacterAnimImage doivent être compris entre 1 et 5 inclus"))
	}
	if global.NumCharacterAnimImages < 1 || global.NumCharacterAnimImages > 5 {
		log.Fatal(errors.New("le fichier de config n'est pas valide, NumCharacterAnimImages doivent être compris entre 1 et 5 inclus"))
	}
	if global.BigStep < 10 {
		log.Fatal(errors.New("le fichier de config n'est pas valide, BigStep doit être supérieur ou égal à 10"))
	}
	if global.FloorKind == 0 && global.RoundEarth {
		log.Fatal(errors.New("le fichier de config n'est pas valide, FloorKind = 0 n'est pas compatible avec RoundEarth"))
	}
	if global.FloorKind == 0 && global.StopCamera {
		log.Fatal(errors.New("le fichier de config n'est pas valide, FloorKind = 0 n'est pas compatible avec StopCamera"))
	}
	if global.FloorKind != 2 && global.RandomGeneration {
		log.Fatal(errors.New("le fichier de config n'est pas valide, FloorKind doit être à 2 avec la génération aléatoire"))
	}
}
