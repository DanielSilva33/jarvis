package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"

	speech "cloud.google.com/go/speech/apiv1"
	"github.com/gordonklaus/portaudio"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println("Starting Jarvis...")

	// Inicializar o cliente de Speech-to-Text
	client, err := speech.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Error creating Speech-to-Text client: %v", err)
	}
	defer client.Close()

	// Inicializar PortAudio
	portaudio.Initialize()
	defer portaudio.Terminate()

	// Configuração do áudio
	buffer := &bytes.Buffer{}
	stream, err := portaudio.OpenDefaultStream(1, 0, 16000, 1024, func(in []int16) {
		for _, sample := range in {
			buffer.WriteByte(byte(sample & 0xFF))
			buffer.WriteByte(byte(sample >> 8))
		}
	})
	if err != nil {
		log.Fatalf("Error opening audio stream: %v", err)
	}
	defer stream.Close()

	fmt.Println("Listening for 5 seconds... Say something!")

	// Configurar um timeout de 5 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Garante que o recurso seja liberado após o uso

	if err := stream.Start(); err != nil {
		log.Fatalf("Error starting stream: %v", err)
	}

	// Espera o timeout
	select {
	case <-ctx.Done():
		fmt.Println("Recording time has ended.")
	}

	if err := stream.Stop(); err != nil {
		log.Fatalf("Error stopping stream: %v", err)
	}

	// Processar o áudio capturado
	text, err := transcribeAudio(client, buffer.Bytes())
	if err != nil {
		log.Fatalf("Error transcribing audio: %v", err)
	}

	fmt.Printf("Transcribed text: %s\n", text)
}

// Função para transcrever áudio usando o Google Speech-to-Text
func transcribeAudio(client *speech.Client, audioData []byte) (string, error) {
	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 16000,
			LanguageCode:    "pt-BR",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: audioData},
		},
	}

	resp, err := client.Recognize(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("error in recognition API: %v", err)
	}

	var transcript string
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			transcript += alt.Transcript + " "
		}
	}
	return transcript, nil
}
