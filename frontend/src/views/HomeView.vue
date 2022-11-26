<script setup lang="ts">
import { ref } from "vue";
import { NInput, NButton, useMessage,NH1,NP,NA } from "naive-ui";
import { SetIP } from "../../wailsjs/go/main/App";

const serverIP = ref("")
const message = useMessage()

function setIP() {
  SetIP(serverIP.value)
  console.log("Trying to set server IP", serverIP.value);
  message.success("设置成功")
}
</script>

<template>
    <div class="w-1/3 absolute left-1/2 top-1/2 mt-20 translate-x-[-50%] translate-y-[-50%] justify-center">
      <n-h1>请设定服务器 IP 地址</n-h1>
      <n-input v-model:value="serverIP" type="text" placeholder="服务器地址" />
      <div class="h-3" />
      <n-button secondary type="primary" @click="setIP()">设定</n-button>
      <n-p>by @cyp0633, <n-a href="https://github.com/cyp0633/socket-chatroom">GitHub</n-a></n-p>
    </div>
</template>
