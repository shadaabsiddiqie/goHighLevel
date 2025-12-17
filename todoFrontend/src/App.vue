<script setup>
import { ref } from "vue";
import { todoClient } from "./client.js";
import { urlClient } from "./client.js";

const shorten_url = ref("");
async function shortenURL(url) {
  const res = await urlClient.shortenURL({ url });
  shorten_url.value = res.shortenedUrl;
  return res.shortenedUrl;
}

const longUrl = ref("");



const title = ref("");
const todos = ref([]);

async function loadTodos() {
  const res = await todoClient.listTodos({});
  todos.value = res.todos;
}

async function addTodo() {
  if (!title.value) return;

  await todoClient.createTodo({ title: title.value });
  title.value = "";
  await loadTodos();
}

loadTodos();
</script>

<template>
  <div style="padding: 20px">
    <h2>Todo App</h2>

    <input v-model="title" placeholder="New todo" />
    <button @click="addTodo">Add</button>

    <ul>
      <li v-for="t in todos" :key="t.id">
        {{ t.title }}
      </li>
    </ul>
  </div>
  <div>
    <h2>URL Shortener</h2>
    <input v-model="longUrl" placeholder="Enter URL to shorten" />
    <button @click="shortenURL(longUrl)">Shorten URL</button>
    <p>Shortened URL: {{ shorten_url }}</p>
  </div>

</template>
