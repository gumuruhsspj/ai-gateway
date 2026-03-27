const { Client, LocalAuth } = require('whatsapp-web.js');
const qrcode = require('qrcode-terminal');
const axios = require('axios');

const client = new Client({
    authStrategy: new LocalAuth() // Biar nggak usah scan QR terus-terusan
});

client.on('qr', (qr) => {
    qrcode.generate(qr, {small: true}); // QR bakal muncul di terminal kamu
});

client.on('ready', () => {
    console.log('WhatsApp Bot Siap!');
});

client.on('message', async msg => {
    // Kirim pesan WA ke Laravel kamu
    try {
        const response = await axios.post('http://127.0.0.1:8000/api/chat', {
            message: msg.body,
            div: 'sales' // default divisi
        });

        // Balas langsung ke WhatsApp user
        msg.reply(response.data.reply);
    } catch (error) {
        console.error("Gagal konek ke Laravel");
    }
});

client.initialize();