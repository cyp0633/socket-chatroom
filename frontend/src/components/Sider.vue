<script setup lang="ts">
import { useRouter } from 'vue-router';
import ConversationItemVue from './ConversationItem.vue';
import { reactive, ref, onMounted } from 'vue';
import { FetchClients } from '../../wailsjs/go/main/App';
const users = reactive([
    "10.0.0.1",
    "10.0.0.2",
    "10.0.0.3",
    "10.0.0.4",
    "10.0.0.5",
    "10.0.0.6",
    "10.0.0.7",
    "10.0.0.8",
    "10.0.0.9",
    "10.0.0.10",
    "10.0.0.11",
    "10.0.0.12",

])
const router = useRouter();

onMounted(async () => {
    // fetch clients every 5 seconds
    setInterval(async () => {
        console.log("Fetching clients");
        const clients = await FetchClients();
        users.splice(0, users.length, ...clients);
    }, 5000);

});

</script>

<template>
    <div class="h-full">
        <div v-for="user in users" :key="user">
            <conversation-item-vue :name="user" @nav-to-conversation="(dest) => router.push('/conversation/' + dest)" />
        </div>
    </div>
</template>