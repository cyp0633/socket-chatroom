<template>
    <div
        class="mt-5 mb-5 p-2 border-solid border-gray-300 border-l border-t border-r border-b border-light-blue-500 rounded-md relative">
        <div :class="{
            'w-2.5 h-2.5 border-gray-300 transform absolute': true,
            'border-l border-t rotate-45 -top-1.5 left-4': placement === 'top-start',
            'border-l border-t rotate-45 -top-1.5 inset-x-2/4 -translate-x-1.5': placement === 'top',
            'border-l border-t rotate-45 -top-1.5 right-4': placement === 'top-end',
            'border-l border-t -rotate-45 top-4 -left-1.5': placement === 'left-start',
            'border-l border-t -rotate-45 inset-y-2/4 -left-1.5 -translate-y-1.5': placement === 'left',
            'border-l border-t -rotate-45 bottom-4 -left-1.5': placement === 'left-end',
            'border-r border-b rotate-45 -bottom-1.5 left-4': placement === 'bottom-start',
            'border-r border-b rotate-45 -bottom-1.5 inset-x-2/4 -translate-x-1.5': placement === 'bottom',
            'border-r border-b rotate-45 -bottom-1.5 right-4': placement === 'bottom-end',
            'border-r border-b -rotate-45 top-4 -right-1.5': placement === 'right-start',
            'border-r border-b -rotate-45 inset-y-2/4 -right-1.5 -translate-y-1.5': placement === 'right',
            'border-r border-b -rotate-45 bottom-4 -right-1.5': placement === 'right-end',
        }"></div>
        <slot></slot>
    </div>
</template>
  
<script lang="ts">
import {
    defineComponent,
    PropType
} from 'vue';
import { useThemeVars } from 'naive-ui';

enum EnumPlacement {
    TopStart = "top-start",
    Top = "top",
    TopEnd = "top-end",
    LeftStart = "left-start",
    Left = "left",
    LeftEnd = "left-end",
    BottomStart = "bottom-start",
    Bottom = "bottom",
    BottomEnd = "bottom-end",
    RightStart = "right-start",
    Right = "right",
    RightEnd = "right-end"
}

export default defineComponent({
    name: 'ConversationBubble',
    props: {
        placement: {
            type: String,
            default: 'top-start'
        }
    }
});
</script>
  