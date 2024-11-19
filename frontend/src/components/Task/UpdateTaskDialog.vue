<script setup lang="ts">
import TaskService from "@/services/TaskService";
import { useTaskStore } from "@/stores/task";
import { CreateTasks, Tasks } from "@/types/Tasks";
import { defineEmits, defineProps } from "vue";
import { toast } from "vue3-toastify";

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

const initForm: CreateTasks = {
  id: "",
  title: "",
  description: "",
  status: "in_progress",
};

let form = reactive<CreateTasks>(initForm);
const formRef = ref<any>(null);
const rules = {
  title: [(v: string) => !!v || "กรุณากรอกหัวข้องาน"],
  description: [(v: string) => !!v || "กรุณากรอกรายละเอียด"],
  status: [(v: string) => !!v || "กรุณาเลือกสถานะ"],
};

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

const handleSubmit = async () => {
  taskStore.tasks.loading = true;
  try {
    const { valid } = await formRef.value.validate();
    // เช็กว่า form ถูกกรอกครบหรือไม่
    if (!valid) {
      toast.error("กรุณากรอกข้อมูลให้ครบ");
      return false;
    }

    const { data, status } = await TaskService.update(taskStore.task.id, form);
    const { error, message, result } = data;
    const res = data;
    delete res.result;
    Object.assign(taskStore.tasks, taskStore.tasks, res);

    if (status !== 200 || error) {
      toast.error("เเก้ไขข้อมูลไม่สําเร็จ");
      console.error("handleSubmit error", message);
      return false;
    }

    toast.success("เเก้ไขข้อมูลเรียบร้อย");
    isDialogVisible.value = false;
    taskStore.tasks.result = taskStore.tasks.result.map((item: Tasks) => {
      if (item.id === taskStore.task.id) return result;
      return item;
    });
  } catch (error) {
    console.error("handleSubmit error", error);
  } finally {
    taskStore.tasks.loading = false;
  }
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
        <v-form ref="formRef">
          <v-text-field
            clearable
            variant="outlined"
            v-model="form.title"
            label="หัวข้องาน"
            :rules="rules.title"
          />
          <v-select
            class="mt-3"
            variant="outlined"
            v-model="form.status"
            clearable
            :rules="rules.status"
            label="สถานะ"
            :items="taskStore.statusList"
            item-title="label"
            item-value="value"
          />
          <v-textarea
            class="mt-3"
            clearable
            variant="outlined"
            v-model="form.description"
            label="รายละเอียด"
            :rules="rules.description"
          />
        </v-form>
      </v-card-text>
      <template v-slot:actions>
        <v-spacer></v-spacer>
        <v-btn @click="handleDialog(false)" color="secondary">ยกเลิก</v-btn>
        <v-btn @click="handleSubmit">ยืนยัน</v-btn>
      </template>
    </v-card>
  </v-dialog>
</template>
