<script setup lang="ts">
import { NAvatar, NP } from 'naive-ui';
const props = defineProps<{
    name: string;
}>()

function logNav() {
    console.log(props.name)
}
</script>

<template>
    <div class="h-20">
        <div class="flex flex-row z-10 absolute m-4">
            <div class="flex mx-2">
                <n-avatar round size="medium">{{ name }}</n-avatar>
            </div>
            <div class="flex mx-2">
                <n-p>{{ name }}</n-p>
            </div>
        </div>
        <div class="w-full h-20 z-20 absolute" @click="$emit('navToConversation', name)" />
    </div>
</template>