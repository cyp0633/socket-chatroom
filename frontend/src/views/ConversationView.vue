<script setup lang="ts">
import FromConversationVue from "@/components/FromConversation.vue";
import ToConversationVue from "@/components/ToConversation.vue";
import SendVue from "@/components/Send.vue";
import { FetchMessages, SendMessage } from "../../wailsjs/go/main/App";
import { onMounted, onUpdated, ref } from "vue";
import { useIpStore } from '@/stores/counter';
import { useRoute } from "vue-router";
interface Message {
    type: string;
    content: string;
    name: string;
}
const route = useRoute();
const ip = route.params.ip as string;
const name = ip.match(/[0-9]{1,3}$/)![0];
const ipStore = useIpStore();
const msg = ref<Message[]>([
    {
        type: "from",
        content: "Hello, World!",
        name: "Alice",
    },
    {
        type: "to",
        content: "Hello, Alice!",
        name: "Bob",
    },
    {
        type: "from",
        content: "Hello, World!",
        name: "Alice",
    },
    {
        type: "to",
        content: "Hello, Alice!",
        name: "Bob",
    }, {
        type: "from",
        content: "Hello, World!",
        name: "Alice",
    },
    {
        type: "to",
        content: "Hello, Alice!",
        name: "Bob",
    }, {
        type: "from",
        content: "Hello, World!",
        name: "Alice",
    },
    {
        type: "to",
        content: "Hello, Alice!",
        name: "Bob",
    }, {
        type: "from",
        content: "Hello, World!",
        name: "Alice",
    },
    {
        type: "to",
        content: "Hello, Alice!",
        name: "Bob",
    }, {
        type: "from",
        content: "Hello, World!",
        name: "Alice",
    },
    {
        type: "to",
        content: "Hello, Alice!",
        name: "Bob",
    }, {
        type: "from",
        content: "Hello, World!",
        name: "Alice",
    },
    {
        type: "to",
        content: "Hello, Alice!",
        name: "Bob",
    },
]);

function sendMsg(message: string) {
    msg.value.push({
        type: "from",
        content: message,
        name: "Alice",
    });
    SendMessage(ip, message)
    console.log("Sent message: " + message);
}

// 每五秒钟拉取消息
onUpdated(async () => {
    ipStore.ip=ip;
    // execute every 5 seconds
    setInterval(async () => {
        const messages = await FetchMessages(ip);
        if (messages == null) {
            msg.value = [] as Message[];
        }
        else {
            msg.value = messages as Message[];
        }
        console.log("New messages with ", ip, ": ", messages);
    }, 5000);
});
</script>

<template>
    <div class="h-2/3 overflow-y-scroll">
        <div v-for="message in msg" :key="message.content">
            <template v-if="message.type === 'from'">
                <from-conversation-vue :msg="message.content" name="我" />
            </template>
            <template v-else>
                <to-conversation-vue :msg="message.content" :name="name" />
            </template>
        </div>
    </div>
    <div class="h-1/3 mx-8">
        <send-vue @send-msg="(n) => sendMsg(n)" />
    </div>
</template>