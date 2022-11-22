<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { NH1, NLayout, NLayoutContent, NLayoutHeader, NLayoutSider } from 'naive-ui'
const { t, availableLocales: languages, locale } = useI18n();
import ConversationItemVue from "./components/ConversationItem.vue";
import FromConversationVue from "./components/FromConversation.vue";
import ToConversationVue from "./components/ToConversation.vue";

var users = [
  "Alice",
  "Bob",
  "Charlie",
  "David",
]

document.body.addEventListener("click", function (event) {
  event.preventDefault();
});
</script>

<template>

  <body>
    <n-layout>
      <n-layout-header class="h-16">
        <n-h1 prefix="bar" class="m-4">
          Socket Chatroom
        </n-h1>
      </n-layout-header>
      <n-layout has-sider>
        <n-layout-sider>
          <div v-for="user in users" :key="user">
            <conversation-item-vue :name="user" />
          </div>
        </n-layout-sider>
        <n-layout-content>
          <from-conversation-vue name="Alice" msg="Hello" />
          <to-conversation-vue name="Bob" msg="Hi" />
        </n-layout-content>
      </n-layout>
    </n-layout>
  </body>
</template>

