import { ref, computed } from "vue";
import { defineStore } from "pinia";

export const useCounterStore = defineStore("counter", () => {
  const count = ref(0);
  const doubleCount = computed(() => count.value * 2);
  function increment() {
    count.value++;
  }

  return { count, doubleCount, increment };
});

export const useIpStore = defineStore("ip", () => {
  const ip = ref("");
  function setIp(newIp: string) {
    ip.value = newIp;
  }
  return { ip, setIp };
});