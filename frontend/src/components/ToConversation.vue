<script setup lang="ts">
import { NAvatar, NCard, NP } from 'naive-ui'
import ConversationBubbleVue from "./ConversationBubble.vue";
defineProps<{
    name: string;
    msg: string;
}>()
</script>

<template>
    <div class="flex flex-row m-4">
        <div class="flex">
            <n-avatar round size="medium">{{ name }}</n-avatar>
        </div>
        <div class="flex">
            <conversation-bubble-vue class="max-w-3xl mx-4" placement="left-end">
                {{ msg }}
            </conversation-bubble-vue>
        </div>
    </div>
</template>