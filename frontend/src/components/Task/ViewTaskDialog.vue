<script setup lang="ts">
import { useTaskStore } from "@/stores/task";
import { Tasks } from "@/types/Tasks";
import { defineEmits, defineProps } from "vue";
const props = defineProps({
  isDialogVisible: {
    type: Boolean,
    required: true,
  },
  title: {
    type: String,
    required: true,
  },
});

const emit = defineEmits(["update:isDialogVisible", "title", "form"]);

const taskStore = useTaskStore();
const form = reactive<Tasks>(taskStore.task);

const isDialogVisible = computed({
  get: () => props.isDialogVisible,
  set: (val: boolean) => emit("update:isDialogVisible", val),
});

const titleComfirmDialog = computed({
  get: () => props.title,
  set: (val: boolean) => emit("title", val),
});

const handleDialog = (status: boolean) => {
  isDialogVisible.value = status;
};

watch(
  () => isDialogVisible.value,
  (newValue: boolean) => {
    console.log("isDialogVisible", newValue);
    Object.assign(form, taskStore.task);
  }
);
</script>

<template>
  <v-dialog v-model="isDialogVisible" max-width="400" persistent>
    <v-card>
      <template #title>
        <h1 class="text-xl font-semibold text-primary">
          {{ titleComfirmDialog }}
        </h1>
      </template>
      <v-card-text>
        <!-- <h1 class="text-base">Are you sure you want to do this?</h1> -->
        <v-form>
          <v-text-field
            readonly
            variant="outlined"
            v-model="form.title"
            label="หัวข้องาน"
          />
          <v-select
            class="mt-3"
            variant="outlined"
            v-model="form.status"
            readonly
            label="สถานะ"
            :items="taskStore.statusList"
            item-title="label"
            item-value="value"
          />
          <v-textarea
            class="mt-3"
            readonly
            variant="outlined"
            v-model="form.description"
            label="รายละเอียด"
          />
        </v-form>
      </v-card-text>
      <template v-slot:actions>
        <v-spacer></v-spacer>
        <v-btn @click="handleDialog(false)" color="secondary">ปิด</v-btn>
      </template>
    </v-card>
  </v-dialog>
</template>
