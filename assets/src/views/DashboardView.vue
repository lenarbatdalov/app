<script setup>
import { ref, onMounted } from 'vue'

const users = ref([])

const name = ref("")
const location = ref("")
const title = ref("")

function getAllUsers() {
  fetch("./users", {
    method: "GET"
  })
  .then(response => {
    if (response.ok) {
      return response.json()
    } else {
      throw new Error(null)
    }
  })
  .then(result => {
    if (result.data !== null) {
      console.log(result.data)
    }
  })
  .catch(_ => {
    console.log('Ошибка соединения с сервером')
  })
}

function createUser() {
  fetch("./user", {
    method: "POST",
    redirect: 'follow',
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      name: name.value,
      location: location.value,
      title: title.value
    })
  })
  .then(response => {
    console.log(response)
    if (response.ok) {
      getAllUsers()
    } else {
      console.log(response.json()?.message || "провалилось создание пользователя")
      throw new Error(null)
    }
  })
}

onMounted(_ => getAllUsers())
</script>

<template>
  <div class="dashboard">
    <h1>This is an dashboard page</h1>

    <div></div>

    <div class="dashboard-create_user">
      <input type="text" v-model="name">
      <input type="text" v-model="location">
      <input type="text" v-model="title">

      <button @click="createUser">Save</button>
    </div>
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .dashboard {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
