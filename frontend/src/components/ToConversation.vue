<script setup lang="ts">
import {NAvatar, NCard} from 'naive-ui'
defineProps<{
    name: string;
    msg: string;
}>()
</script>

<template>
    <div class="grid grid-cols-2 place-content-start gap-4">
        <div class="shrink">
            <n-avatar round size="medium">{{ name }}</n-avatar>
        </div>
        <div class="shrink">
            <n-card>
                {{ msg }}
            </n-card>
        </div>
    </div>
</template>