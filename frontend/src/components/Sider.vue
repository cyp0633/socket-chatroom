<script setup lang="ts">
import { useRouter } from 'vue-router';
import ConversationItemVue from './ConversationItem.vue';
import { reactive, ref, onMounted } from 'vue';
import { FetchClients } from '../../wailsjs/go/main/App';
const users = reactive([] as string[]);
const router = useRouter();
const currUser = ref("");

onMounted(async () => {
    // fetch clients every 5 seconds
    setInterval(async () => {
        console.log("Fetching clients");
        const clients = await FetchClients();
        users.splice(0, users.length, ...clients);
    }, 5000);
});

function goToConversation(ip: string) {
    router.push('/conversation/' + ip);
    currUser.value = ip;
    console.log("Current user: ", currUser.value);
}

</script>

<template>
    <div class="h-full">
        <div v-for="user in users" :key="user">
            <conversation-item-vue :name="user" @nav-to-conversation="(dest) => goToConversation(dest)" :highlight="currUser==user"/>
        </div>
    </div>
</template>