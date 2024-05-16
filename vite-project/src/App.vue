<script setup lang = "ts" >  
  import { ref } from "vue"
  import {v4 as uuidv4} from "uuid"

  type TodoItem = {
    task: string
    complete: boolean
    id: string
  }
  type Todos = TodoItem[]

  const todos = ref<Todos>([])
  const input = ref("")

  function handleSubmitForm(e: KeyboardEvent) {
    if (e.key === "Enter") {
      addTodo()
    }
  }
  
  const addTodo = () => {
    if (input.value !== "") {
      todos.value = [...todos.value, { task: input.value, complete: false, id: uuidv4() }]
      input.value = ""
    } else {
      alert("Введите что-нибудь")
    }
  }  

  const deleteTodo = (todoId: string) => {
    todos.value = todos.value.filter((todo) => todo.id !==  todoId)
  }

  const toggleTodo = (todoId: string) => {
    const todoIndex = todos.value.findIndex((todo) => todo.id === todoId);
    if (todoIndex !== -1) {
      todos.value[todoIndex].complete = !todos.value[todoIndex].complete;
    }
  }
</script>

<template>
  <section>
    <h1>TODO</h1>
    <input type="text" v-model="input" placeholder="Add something" @keyup="handleSubmitForm"/>
    <button @click="addTodo">ADD</button>
    <div class="pole">
      <ul>
      <li v-for="(todo) in todos" :key="todo.id" :class="{'completed' : todo.complete}">
      {{ todo.task }}
      <button @click="toggleTodo(todo.id)">TOGGLE</button>
      <button @click="deleteTodo(todo.id)">DELETE</button>
      </li>
    </ul>
    </div>
  </section>
</template>

<style>
template {
  display: flex;
  flex-direction: column;
}

svg {
  width: 20px;
  height: 20px;
}

.pole {
  width: 500px;
  height: 500px;
  border: 1px solid white;
}

ul {
  list-style-type: none;
}

.completed {
 text-decoration: line-through;
 color: gray;
}

.not-completed {
 text-decoration: none;
 color: black;
}


  h1 {
    font-size: 100px;
  }

  input {
    width: 300px;
    height: 30px;
    font-size: 15px;
  }
</style>
