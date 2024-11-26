package speech

import (
	"bytes"
	"context"
	"fmt"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "cloud.google.com/go/speech/apiv1/speechpb"
)

// NewSpeechClient inicializa o cliente de Speech-to-Text.
func NewSpeechClient() (*speech.Client, error) {
	return speech.NewClient(context.Background())
}

// TranscribeAudio transcreve os dados de áudio em texto.
func TranscribeAudio(client *speech.Client, audioData []byte) (string, error) {
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

// ListenAndTranscribe combina gravação de áudio e transcrição.
func ListenAndTranscribe(ctx context.Context, client *speech.Client) (string, error) {
	var buffer bytes.Buffer

	// Grava o áudio do microfone
	err := RecordAudio(ctx, &buffer)
	if err != nil {
		return "", fmt.Errorf("error recording audio: %v", err)
	}

	// Transcrever o áudio
	return TranscribeAudio(client, buffer.Bytes())
}
