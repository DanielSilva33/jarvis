package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/DanielSilva33/jarvis/internal/speech"
	"github.com/DanielSilva33/jarvis/utils"
)

func main() {
	// Load .env file
	utils.LoadEnv()

	fmt.Println("Starting Jarvis...")

	// Inicializar o cliente de Speech-to-Text
	speechClient, err := speech.NewSpeechClient()
	if err != nil {
		log.Fatalf("Error creating Speech-to-Text client: %v", err)
	}
	defer speechClient.Close()

	// Escutar áudio por 5 segundos e transcrever
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	text, err := speech.ListenAndTranscribe(ctx, speechClient)
	if err != nil {
		log.Fatalf("Error in Speech-to-Text: %v", err)
	}

	fmt.Printf("Transcribed text: %s\n", text)

	// Inicializar o cliente ChatGPT
	// apiKey := os.Getenv("OPENAI_API_KEY")
	// if apiKey == "" {
	// 	log.Fatal("Missing OpenAI API key")
	// }
	// chatClient := chatgpt.NewChatGPTClient(apiKey)

	// // Obter resposta do ChatGPT
	// response, err := chatClient.GetJarvisResponse(text)
	// if err != nil {
	// 	log.Fatalf("Error getting response from ChatGPT: %v", err)
	// }

	// fmt.Printf("Jarvis: %s\n", response)
}
