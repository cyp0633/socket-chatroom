<script setup lang="ts">
import {defineComponent} from 'vue';
import {NH1} from 'naive-ui';
defineComponent({
    name:"TopBar",
    props:{}
})
</script>

<template>
    <div class="w-full h-auto">
        <div class="container m-4">
            <n-h1 prefix="bar">Socket Chatroom</n-h1>
        </div>
    </div>
</template>