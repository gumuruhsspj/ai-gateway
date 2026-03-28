
-----

# 🚀 AI Gateway - Multi-Engine WhatsApp Bridge

[](https://go.dev/)
[](https://laravel.com/)
[](https://vuejs.org/)
[](https://nodejs.org/)
[](https://opensource.org/licenses/MIT)

Sistem Gateway AI modular yang menghubungkan WhatsApp dengan Engine Kecerdasan Buatan (LLM) melalui integrasi Laravel API. Dirancang untuk efisiensi tinggi, kompatibel dengan hardware i3 RAM 2GB.

-----

## 🏗️ Project Structure

```text
/ai-gateway
├── _backend         # Laravel 10 API (Central Logic & Database)
├── _frontend        # VueJS 3 Modular (Dashboard & QR Monitor)
├── _wa_bot_node     # WhatsApp Client (Legacy Node.js version)
└── _wa_bot_go       # WhatsApp Client (High-Performance Go version) - RECOMMENDED
```

## 🛠️ Tech Stack & Requirements

### Core Components

  * **Backend:** PHP 8.2, Laravel 10, MySQL/PostgreSQL.
  * **Frontend:** Vue.js 3, Vite, TailwindCSS.
  * **WhatsApp Engine:**
      * **Go:** `whatsmeow` library (Low Memory Consumption).
      * **Node:** `whatsapp-web.js` / `baileys` (Optional).

### 🧠 AI Engine Options

Sistem ini mendukung integrasi dua jalur pemrosesan AI:

1.  **Groq Cloud API (Fast Access):** Menggunakan LPU™ untuk akses super cepat. Tidak membebani server lokal. (Sangat disarankan untuk RAM 2GB).
2.  **Local Ollama (Self-Hosted):** Privasi data total.
      * **Konsekuensi:** Butuh spesifikasi tinggi (**RAM 16GB+ & GPU VRAM 8GB+**) untuk menjalankan model seperti Llama3 atau Mistral secara lancar.

### Hardware Minimum 

  * **CPU:** Intel i3 Gen 4 (Haswell) or better.
  * **RAM:** 2GB Minimum (Cloud AI Mode) / 16GB+ (Local AI Mode).
  * **OS:** Windows 10 / Linux antiX (CLI Mode Optimized).

-----

## 🚀 Quick Start

### 1\. Setup Backend (Laravel)

```bash
cd _backend
composer install
cp .env.example .env
php artisan key:generate
php artisan migrate
php artisan serve
```

### 2\. Setup Bot Engine (Go) - *Result Oriented Choice*

```bash
cd _wa_bot_go

# Build for Windows
go build -o wa-bot.exe

# Build for Linux (antiX)
set GOOS=linux
set GOARCH=amd64
go build -o wa-bot-linux
./wa-bot.exe
```

### 3\. Setup Frontend (VueJS)

```bash
cd _frontend
npm install
npm run dev
```

-----

## 🛡️ Features

  - [x] **Modular Architecture:** Ganti bot engine (Node/Go) tanpa merusak API utama.
  - [x] **Hybrid AI Path:** Pilih antara Groq (Online) atau Ollama (Local).
  - [x] **Real-time QR:** QR Code otomatis tersimpan di `_backend/public/qrcodes`.
  - [x] **Memory Optimized:** Versi Go hanya mengonsumsi RAM \< 50MB.
  - [x] **Cross-Platform:** Siap dideploy di Windows Server maupun Linux antiX.

-----

## 📉 Scalability & Memory Benchmarks

Estimasi penggunaan RAM total (OS + Laravel + Database + WhatsApp Engine) berdasarkan jumlah akun/session yang aktif secara bersamaan.

| Jumlah Akun (QR Login) | Engine Node.js (Puppeteer) | Engine Go (Whatsmeow) | Selisih Efisiensi |
| :--- | :--- | :--- | :--- |
| **1 Akun** | \~1.2 GB RAM | **\~0.7 GB RAM** | 41% Lebih Hemat |
| **5 Akun** | \~2.5 GB RAM | **\~0.9 GB RAM** | 64% Lebih Hemat |
| **10 Akun** | \~4.5 GB RAM | **\~1.1 GB RAM** | 75% Lebih Hemat |
| **20 Akun** | \~8.0 GB RAM | **\~1.5 GB RAM** | 81% Lebih Hemat |
| **50 Akun** | \~18.0 GB RAM | **\~2.8 GB RAM** | 84% Lebih Hemat |

### 🔍 Technical Insight

  * **Engine Go (`_wa_bot_go`):** Berjalan pada level *Raw Socket*. Satu server RAM 2GB mampu melayani puluhan akun personil sekaligus secara stabil menggunakan **Groq API**.
  * **Engine Node.js (`_wa_bot_node`):** Setiap akun membuka satu instance browser virtual yang memakan resource eksponensial. Tidak disarankan untuk hardware terbatas.
  * **Ollama Warning:** Jika menjalankan Ollama di server yang sama dengan RAM 2GB, sistem dipastikan akan **Hang/Freeze**. Gunakan Ollama hanya pada dedicated AI server.

-----

## 👨‍💻 Developed By

**Gumuruh Samudra Sabar**
*Senior Software Developer & IT Consultant*
*Team of FGroupIndonesia*

-----

> **Note:** Gunakan versi Go (`_wa_bot_go`) dan Groq API untuk stabilitas maksimal pada server spek ekonomis.

-----
