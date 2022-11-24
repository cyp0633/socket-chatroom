<script setup lang="ts">
import { NAvatar } from 'naive-ui';
defineProps<{
    name: string;
}>()
</script>

<template>
    <div class="flex flex-row m-4">
        <div class="flex mx-2">
            <n-avatar round size="medium">{{ name }}</n-avatar>
        </div>
        <div class="flex mx-2">
            <slot>{{ name }}</slot>
        </div>
    </div>
</template>