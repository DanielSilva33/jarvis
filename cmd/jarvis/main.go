package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/DanielSilva33/jarvis/internal/speech"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println("Starting Jarvis...")

	// Inicializar o cliente de Speech-to-Text
	client, err := speech.NewSpeechClient()
	if err != nil {
		log.Fatalf("Error creating Speech-to-Text client: %v", err)
	}
	defer client.Close()

	// Escutar áudio por 5 segundos e transcrever
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	text, err := speech.ListenAndTranscribe(ctx, client)
	if err != nil {
		log.Fatalf("Error in Speech-to-Text: %v", err)
	}

	fmt.Printf("Transcribed text: %s\n", text)
}
