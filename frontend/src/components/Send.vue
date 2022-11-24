<script setup lang="ts">
import { defineProps, ref } from 'vue'
import { NInput, NButton } from 'naive-ui'

var msg = ref("")

function sendMsg() {
    console.log(msg.value)
}
</script>

<template>
    <div class="w-full m-4">
        <n-input v-model:value="msg" type="textarea" placeholder="发送消息..." />
    </div>
    <div class="flex flex-row-reverse">
        <n-button class="m-4" type="primary" @click="$emit('sendMsg',msg)">发送</n-button>
    </div>
</template>