
---

```markdown
# 🚀 AI Gateway - Multi-Engine WhatsApp Bridge

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![Laravel Version](https://img.shields.io/badge/Laravel-10.x-FF2D20?style=for-the-badge&logo=laravel)](https://laravel.com/)
[![Vue Version](https://img.shields.io/badge/Vue.js-3.x-4FC08D?style=for-the-badge&logo=vuedotjs)](https://vuejs.org/)
[![Node.js Version](https://img.shields.io/badge/Node.js-18.x-339933?style=for-the-badge&logo=nodedotjs)](https://nodejs.org/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

Sistem Gateway AI modular yang menghubungkan WhatsApp dengan Engine Kecerdasan Buatan (LLM) melalui integrasi Laravel API. Dirancang untuk efisiensi tinggi, kompatibel dengan hardware i3 RAM 2GB.

---

## 🏗️ Project Structure

```text
/ai-gateway
├── _backend         # Laravel 10 API (Central Logic & Database)
├── _frontend        # VueJS 3 Modular (Dashboard & QR Monitor)
├── _wa_bot_node     # WhatsApp Client (Legacy Node.js version)
└── _wa_bot_go       # WhatsApp Client (High-Performance Go version) - RECOMMENDED
```

---

## 🛠️ Tech Stack & Requirements

### Core Components
* **Backend:** PHP 8.2, Laravel 10, MySQL/PostgreSQL.
* **Frontend:** Vue.js 3, Vite, TailwindCSS.
* **WhatsApp Engine:** * **Go:** `whatsmeow` library (Low Memory Consumption).
    * **Node:** `whatsapp-web.js` / `baileys` (Optional).

### Hardware Recommendation
* **CPU:** Intel i3 Gen 4 (Haswell) or better.
* **RAM:** 2GB Minimum.
* **OS:** Windows 10 / Linux antiX (CLI Mode Optimized).

---

## 🚀 Quick Start

### 1. Setup Backend (Laravel)
```bash
cd _backend
composer install
cp .env.example .env
php artisan key:generate
php artisan migrate
php artisan serve
```

### 2. Setup Bot Engine (Go) - *Result Oriented Choice*
```bash
cd _wa_bot_go
# Build for Windows
go build -o wa-bot.exe
# Build for Linux (antiX)
set GOOS=linux && set GOARCH=amd64 && go build -o wa-bot-linux
./wa-bot.exe
```

### 3. Setup Frontend (VueJS)
```bash
cd _frontend
npm install
npm run dev
```

---

## 🛡️ Features
- [x] **Modular Architecture:** Ganti bot engine (Node/Go) tanpa merusak API utama.
- [x] **Real-time QR:** QR Code otomatis tersimpan di `_backend/public/qrcodes`.
- [x] **Memory Optimized:** Versi Go hanya mengonsumsi RAM < 50MB.
- [x] **Cross-Platform:** Siap dideploy di Windows Server maupun Linux antiX.

---

## 👨‍💻 Developed By
**Gumuruh Samudra Sabar**
*Senior Software Developer & IT Consultant*
*Team of FGroupIndonesia*

---
> **Note:** Gunakan versi Go (`_wa_bot_go`) untuk stabilitas jangka panjang dan penggunaan resource yang lebih hemat pada server spek lebih ekonomis.
```
