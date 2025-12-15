<script setup>
import { ref } from "vue";
import { todoClient } from "./client.js";

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
</template>
