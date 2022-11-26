<script setup lang="ts">
import { defineComponent, onUpdated } from 'vue';
import { useIpStore } from '@/stores/counter';
import { NH1 } from 'naive-ui';
const ipStore = useIpStore();
defineComponent({
    name: "TopBar",
    props: {}
})
</script>

<template>
    <div class="w-full h-auto">
        <div class="container m-4">
            <n-h1 prefix="bar">Socket Chatroom - 与 {{ ipStore.ip }} 对话中</n-h1>
        </div>
    </div>
</template>