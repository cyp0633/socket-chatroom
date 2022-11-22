<script setup lang="ts">
import { NAvatar } from 'naive-ui';
defineProps<{
    name: string;
}>()
</script>

<template>
    <div class="grid grid-cols-2 place-content-start gap-4">
        <div class="shrink">
            <n-avatar round size="medium">{{ name }}</n-avatar>
        </div>
        <div class="shrink">
            <slot>{{ name }}</slot>
        </div>
    </div>
</template>