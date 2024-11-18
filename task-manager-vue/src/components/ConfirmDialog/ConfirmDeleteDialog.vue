<script setup lang="ts">
import { defineEmits, defineProps } from "vue";

const props = defineProps({
  isDialogVisible: {
    type: Boolean,
    required: true,
  },
  title: {
    type: String,
    required: false,
  },
  text: {
    type: String,
    required: false,
  },
});

const emit = defineEmits([
  "update:isDialogVisible",
  "confirm",
  "title",
  "text",
]);

const isDialogVisible = computed({
  get: () => props.isDialogVisible,
  set: (val: boolean) => emit("update:isDialogVisible", val),
});

const titleComfirmDialog = computed({
  get: () => props.title,
  set: (val: boolean) => emit("title", val),
});

const textComfirmDialog = computed({
  get: () => props.text,
  set: (val: boolean) => emit("text", val),
});

const handleConfirm = () => {
  isDialogVisible.value = false;
  emit("confirm");
};
const handleCancel = () => {
  isDialogVisible.value = false;
};
</script>

<template>
  <div class="text-center pa-4">
    <v-dialog v-model="isDialogVisible" max-width="400" persistent>
      <v-card>
        <template #text>
          <section class="w-full pt-1 flex gap-2">
            <v-icon color="warning">mdi-alert-circle</v-icon>
            <h1 class="text-lg font-semibold">{{ titleComfirmDialog }}</h1>
          </section>
          <section>
            <h1 class="text-base">{{ textComfirmDialog }}</h1>
          </section>
        </template>
        <section class="flex justify-end pb-2 pr-2">
          <div>
            <v-btn @click="handleCancel" color="secondary" variant="text"
              >ยกเลิก</v-btn
            >
            <v-btn @click="handleConfirm" variant="text">ตกลง</v-btn>
          </div>
        </section>
      </v-card>
    </v-dialog>
  </div>
</template>
