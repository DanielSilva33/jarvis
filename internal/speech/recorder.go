package speech

import (
	"bytes"
	"context"
	"log"

	"github.com/gordonklaus/portaudio"
)

// RecordAudio grava áudio do microfone por 5 segundos e retorna os dados em bytes.
func RecordAudio(ctx context.Context, buffer *bytes.Buffer) error {
	// Inicializar PortAudio
	portaudio.Initialize()
	defer portaudio.Terminate()

	// Configurar o stream de áudio
	stream, err := portaudio.OpenDefaultStream(1, 0, 16000, 1024, func(in []int16) {
		for _, sample := range in {
			buffer.WriteByte(byte(sample & 0xFF))
			buffer.WriteByte(byte(sample >> 8))
		}
	})
	if err != nil {
		return err
	}
	defer stream.Close()

	// Iniciar gravação
	if err := stream.Start(); err != nil {
		return err
	}
	log.Println("Recording audio...")

	// Manter o stream aberto até o contexto expirar
	<-ctx.Done()
	log.Println("Recording finished.")

	if err := stream.Stop(); err != nil {
		return err
	}

	return nil
}
