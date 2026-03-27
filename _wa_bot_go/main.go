package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"os/signal"
	"syscall"

	_ "modernc.org/sqlite"
	"github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

const qrPath = "../ai-gateway/public/qrcodes/wa-qr.png"
const laravelWebhook = "http://127.0.0.1:8000/api/chat"

var client *whatsmeow.Client

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		// Ambil teks pesan
		userMsg := v.Message.GetConversation()
		if v.Info.IsGroup || userMsg == "" {
			return
		}

		fmt.Printf("Pesan dari %s: %s\n", v.Info.Sender.String(), userMsg)

		// 1. Tembak ke Laravel
		go func() {
			payload, _ := json.Marshal(map[string]string{
				"message": userMsg,
				"div":     "sales",
			})

			resp, err := http.Post(laravelWebhook, "application/json", bytes.NewBuffer(payload))
			if err != nil {
				fmt.Println("❌ Gagal konek ke Laravel:", err)
				return
			}
			defer resp.Body.Close()

			var result map[string]string
			json.NewDecoder(resp.Body).Decode(&result)

			// 2. Balas ke WhatsApp (Struktur baru pakai proto)
			if reply, ok := result["reply"]; ok {
				responseMsg := &waE2E.Message{
					Conversation: proto.String(reply),
				}
				client.SendMessage(context.Background(), v.Info.Sender, responseMsg)
			}
		}()
	}
}

func main() {

	dir := filepath.Dir(qrPath)

	err := os.MkdirAll(dir, os.ModePerm)
    if err != nil {
        fmt.Println("❌ Gagal membuat folder QR:", err)
    } else {
        fmt.Println("✅ Folder QR siap/sudah ada.")
    }

	// WhatsMeow versi baru butuh Context di awal
	ctx := context.Background()

	// New sekarang minta Context & Logger (kita kasih nil saja biar enteng)
	container, err := sqlstore.New(ctx, "sqlite", "file:session.db?_pragma=foreign_keys(1)", nil)
	if err != nil { panic(err) }

	// GetFirstDevice sekarang butuh Context
	deviceRes, err := container.GetFirstDevice(ctx)
	if err != nil { panic(err) }

	client = whatsmeow.NewClient(deviceRes, nil)
	client.AddEventHandler(eventHandler)

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(ctx)
		err = client.Connect()
		if err != nil { panic(err) }

		for evt := range qrChan {
			if evt.Event == "code" {
				err := qrcode.WriteFile(evt.Code, qrcode.Medium, 256, qrPath)
				if err != nil {
					fmt.Println("❌ Gagal simpan QR:", err)
				} else {
					fmt.Println("✅ QR PNG Terupdate! Silahkan cek di Vue.")
				}
			}
		}
	} else {
		err = client.Connect()
		if err != nil { panic(err) }
		fmt.Println("✅ WhatsApp Terhubung Otomatis!")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	client.Disconnect()
}