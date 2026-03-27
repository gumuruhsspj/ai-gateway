<template>
  <div id="app" class="container">
    <nav>
      <button @click="view = 'admin'">Admin Panel</button>
      <button @click="view = 'demo'">Demo Chat</button>
    </nav>

    <hr />

    <div v-if="view === 'admin'">
      <h2>Config Division</h2>
      <div class="form">
        <input v-model="form.name" placeholder="Name (e.g. Sales)" />
        <input v-model="form.slug" placeholder="Slug (e.g. sales)" />
        <textarea v-model="form.context" placeholder="Context..."></textarea>
        <textarea v-model="form.rules" placeholder="Rules..."></textarea>
        <select v-model="form.tone">
          <option value="formal">Formal</option>
          <option value="casual">Casual</option>
        </select>
        <button @click="saveDiv">Save Division</button>
      </div>
    </div>

    <div v-if="view === 'demo'">
      <h2>Chat AI (Div: {{ currentDiv }})</h2>
      <div class="chat-box" ref="chatWindow">
        <div v-for="(m, i) in messages" :key="i" :class="m.role">
          <strong>{{ m.role }}:</strong> {{ m.text }}
        </div>
      </div>
      <div class="input-area">
        <input v-model="userInput" @keyup.enter="sendChat" placeholder="Type..." />
        <button @click="sendChat">Send</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, nextTick } from 'vue';
import axios from 'axios';

// Konfigurasi API (Arahkan ke Laravel)
const api = axios.create({ baseURL: 'http://localhost:8000/api' });

const view = ref('admin');
const form = ref({ name: '', slug: '', context: '', rules: '', tone: 'formal' });
const messages = ref([]);
const userInput = ref('');
const currentDiv = new URLSearchParams(window.location.search).get('div') || 'cs';
const chatWindow = ref(null);

// SAVE DIVISION (Admin)
const saveDiv = async () => {
  await api.post('/divisions', form.value);
  alert('Saved!');
};

// SEND CHAT & SPEECH (Demo)
const sendChat = async () => {
  const msg = userInput.value;
  messages.value.push({ role: 'user', text: msg });
  userInput.value = '';

  const { data } = await api.post('/chat', { message: msg, div: currentDiv });
  
  messages.value.push({ role: 'assistant', text: data.reply });

  // 🔊 Bonus: Speech Synthesis
  const talk = new SpeechSynthesisUtterance(data.reply);
  window.speechSynthesis.speak(talk);

  // Auto Scroll
  await nextTick();
  chatWindow.value.scrollTop = chatWindow.value.scrollHeight;
};
</script>

<style>
.chat-box { height: 300px; overflow-y: auto; border: 1px solid #ccc; padding: 10px; margin-bottom: 10px; }
.user { color: blue; text-align: right; }
.assistant { color: green; }
.form { display: flex; flex-direction: column; gap: 10px; max-width: 400px; }
</style>