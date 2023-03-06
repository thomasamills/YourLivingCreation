package conversation_code

import (
	"errors"
	"time"
)

type GameLoopManagerRegistry interface {
	AddGameLoop(sessionId string, manager GameLoopManager) error
	GetGameLoop(sessionId string) (GameLoopManager, error)
	StopGameLoop(sessionId string) error
	PurgeGameLoops() error
}

type GameLoopManagerRegistryImpl struct {
	gameLoopManagers map[string]GameLoopManager
}

// TODO find a way to persist the game loop stuff

func NewGameLoopManagerRegistry() GameLoopManagerRegistry {
	return &GameLoopManagerRegistryImpl{
		gameLoopManagers: make(map[string]GameLoopManager, 0),
	}
}

func (g GameLoopManagerRegistryImpl) AddGameLoop(sessionId string, manager GameLoopManager) error {
	g.gameLoopManagers[sessionId] = manager
	return nil
}

func (g GameLoopManagerRegistryImpl) GetGameLoop(sessionId string) (GameLoopManager, error) {
	if gameLoop, ok := g.gameLoopManagers[sessionId]; ok {
		return gameLoop, nil
	}
	return nil, errors.New("game loop not found")
}

func (g GameLoopManagerRegistryImpl) StopGameLoop(sessionId string) error {
	if gameLoop, ok := g.gameLoopManagers[sessionId]; ok {
		gameLoop.StopGameLoop()
	}
	return errors.New("no game loop to stop")
}

func (g GameLoopManagerRegistryImpl) PurgeGameLoops() error {
	for {
		time.Sleep(time.Second * 10)
		for _, gameLoopManager := range g.gameLoopManagers {
			if !gameLoopManager.IsAlive(time.Second * 60) {
				gameLoopManager.StopGameLoop()
			}
		}
	}
}
