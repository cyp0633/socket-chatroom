<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { NH1, NLayout, NLayoutContent, NLayoutHeader, NLayoutSider, NP, NConfigProvider, useOsTheme, darkTheme } from 'naive-ui'
import ConversationItemVue from "./components/ConversationItem.vue";
const { t, availableLocales: languages, locale } = useI18n();

var users = [
  "Alice",
  "Bob",
  "Charlie",
  "David",
]

const osTheme = useOsTheme()
const theme = osTheme.value === 'dark' ? darkTheme : null

document.body.addEventListener("click", function (event) {
  event.preventDefault();
});
</script>

<template>

  <body>
    <n-config-provider :theme="theme">
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
          <n-layout-content class="h-screen">
            <router-view class="h-full" />
          </n-layout-content>
        </n-layout>
      </n-layout>
    </n-config-provider>
  </body>
</template>

