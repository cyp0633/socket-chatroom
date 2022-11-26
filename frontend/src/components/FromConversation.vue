<script setup lang="ts">
import { NAvatar, NCard,NP } from 'naive-ui'
import ConversationBubbleVue from "./ConversationBubble.vue";
interface props {
    name: string;
    msg: string;
}
const props = defineProps<props>()
</script>

<template>
    <div class="flex flex-row-reverse m-4">
        <div class="flex">
            <n-avatar round size="medium">{{ name }}</n-avatar>
        </div>
        <div class="flex">
            <conversation-bubble-vue class="max-w-3xl mx-4" placement="right-end">
                {{ msg }}
            </conversation-bubble-vue>
        </div>
    </div>
</template>