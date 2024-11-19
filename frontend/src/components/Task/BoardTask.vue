<script lang="ts" setup>
import TaskService from "@/services/TaskService";
import { ref, computed } from "vue";
import { Tasks } from "@/types/Tasks";
import { useTaskStore } from "@/stores/task";
import { toast } from "vue3-toastify";
import { AxiosResponse } from "axios";
import { Response } from "@/types/Response";

interface TouchPosition {
  x: number;
  y: number;
}
const initState: Tasks = {
  id: "0",
  title: "",
  description: "",
  status: "in_progress",
  created_at: new Date().toISOString(),
  updated_at: new Date().toISOString(),
};
const taskStore = useTaskStore();
const draggedTask = ref<any>(initState);
const touchDragging = ref(false);
const isDragging = ref(false);
const isDragOver = ref<string | null>(null);
const touchPosition = ref<TouchPosition>({ x: 0, y: 0 });
const initialTouchPosition = ref<TouchPosition>({ x: 0, y: 0 });
const taskCount = computed(() => {
  const draft = taskStore.tasks.result.filter(
    (item: Tasks) => item.status === "draft"
  ).length;
  const in_progress = taskStore.tasks.result.filter(
    (item: Tasks) => item.status === "in_progress"
  ).length;
  const completed = taskStore.tasks.result.filter(
    (item: Tasks) => item.status === "completed"
  ).length;
  const payload = {
    draft,
    in_progress,
    completed,
  };

  return payload;
});
const ghostCardStyle = computed(() => ({
  transform: `translate(${touchPosition.value.x}px, ${touchPosition.value.y}px)`,
  opacity: "0.8",
  width: "200px",
  zIndex: 1000,
  cursor: "grabbing",
}));
const updateItemDialog = reactive({
  visible: false,
  title: "เเก้ไขรายการ",
  description: "",
  status: "",
});
const handleUpdateTask = async (id: string, payload: Tasks) => {
  try {
    const { data, status }: AxiosResponse = await TaskService.update(
      id,
      payload
    );
    const { message, error }: Response = data;
    if (status !== 200 || error) {
      throw new Error(message);
    }
    taskStore.tasks.result = taskStore.tasks.result.map((item: Tasks) => {
      const newItem = { ...item, ...payload };
      if (item.id === id) return newItem;
      return item;
    });
  } catch (error) {
    toast.error("เกิดข้อผิดพลาด");
    console.error("handleUpdateTask error", error);
  }
};

const filteredTasks = (status: string) => {
  return taskStore.tasks.result.filter((task: Tasks) => task.status === status);
};

const onDragStart = (task: Tasks) => {
  draggedTask.value = task;
  isDragging.value = true;
  document.body.style.cursor = "grabbing";
};

const onDragEnd = () => {
  isDragging.value = false;
  document.body.style.cursor = "";
};

const onDragOver = (status: string) => {
  isDragOver.value = status;
};

const onDragLeave = () => {
  isDragOver.value = null;
};

const onTaskTouchStart = (event: TouchEvent, task: Tasks) => {
  event.preventDefault();
  draggedTask.value = task;
  touchDragging.value = true;
  isDragging.value = true;

  const touch = event.touches[0];
  initialTouchPosition.value = {
    x: touch.clientX,
    y: touch.clientY,
  };
  touchPosition.value = {
    x: touch.clientX,
    y: touch.clientY,
  };
};

const onTouchMove = (event: TouchEvent) => {
  if (touchDragging.value) {
    event.preventDefault();
    const touch = event.touches[0];
    touchPosition.value = {
      x: touch.clientX - 100,
      y: touch.clientY - 25,
    };
  }
};

const onTouchEnd = (newStatus: string) => {
  if (touchDragging.value && draggedTask.value) {
    draggedTask.value.status = newStatus;
    touchDragging.value = false;
    isDragging.value = false;

    const taskElement = document.querySelector(
      `[data-task-id="${draggedTask.value.id}"]`
    );
    if (taskElement) {
      taskElement.classList.add("dropped");
      setTimeout(() => taskElement.classList.remove("dropped"), 500);
    }
    draggedTask.value = null;
  }
};

const onDrop = async (newStatus: string) => {
  try {
    if (draggedTask.value) {
      draggedTask.value.status = newStatus;
      isDragOver.value = null;

      const taskElement = document.querySelector(
        `[data-task-id="${draggedTask.value.id}"]`
      );
      if (taskElement) {
        taskElement.classList.add("dropped");
        setTimeout(() => taskElement.classList.remove("dropped"), 500);

        const payload: Tasks = {
          ...draggedTask.value,
          status: newStatus,
        };
        await handleUpdateTask(draggedTask.value.id, payload);
      }
      // end if (taskElement)
      draggedTask.value = null;
    }
  } catch (error) {
    console.error("onDrop error", error);
  }
};

const viewItemDialog = reactive({
  visible: false,
  title: "รายละเอียดข้อมูล",
  description: "",
  status: "",
});

const addItemDialog = reactive({
  visible: false,
  title: "เพิ่มรายการงาน (Task List)",
});

const deleteItemDialog = reactive({
  id: "",
  visible: false,
  title: "ยืนยันการทำรายการ?",
  text: "ถ้าหากยืนยันการทำรายการนี้ จะไม่สามารถกู้คืนได้",
});

const handleViewItem = (item: Tasks) => {
  if (item.id) {
    taskStore.task = item;
  }
  viewItemDialog.visible = true;
};

const handleUpdateItem = (item: Tasks) => {
  if (item.id) {
    taskStore.task = item;
  }
  updateItemDialog.visible = true;
};

const handleDeleteItem = (id: string) => {
  deleteItemDialog.title = `ยืนยันการลบรายการ?`;
  deleteItemDialog.visible = true;
  deleteItemDialog.id = id;
};

const deleteItem = () => {
  taskStore.deleteTaskById(deleteItemDialog.id);
};
</script>

<template>
  <div>
    <section class="w-full flex align-center mb-4 gap-3">
      <v-text-field
        v-model="taskStore.search.search"
        class="w-full none"
        variant="outlined"
        single-line
        hide-details
        label="ค้นหา"
        clearable
      />
    </section>
    <div class="flex flex-col md:flex-row md:space-x-4 space-y-4 md:space-y-0">
      <div
        v-for="status in taskStore.statusList"
        :key="status.value"
        class="w-full md:w-1/3 p-4 rounded-lg border-2 border-gray-200 overflow-y-scroll scrollbar-hide h-[450px]"
      >
        <div class="flex flex-align-center gap-3">
          <h2 class="text-xl text-on-surface-variant">
            {{ status.label }}
          </h2>
          <h2 class="text-xl text-on-surface-variant">
            {{
              status.value === "draft"
                ? taskCount.draft
                : status.value === "in_progress"
                ? taskCount.in_progress
                : taskCount.completed
            }}
          </h2>
        </div>

        <div
          class="w-full rounded-md h-1 mt-4 mb-4"
          :class="{
            'bg-[LightSteelBlue]': status.value === 'draft',
            'bg-yellow-200': status.value === 'in_progress',
            'bg-green-400': status.value === 'completed',
          }"
        ></div>
        <div
          class="min-h-[300px] p-2 bg-white rounded-lg"
          :class="{
            'border-dashed border-2 border-gray-300 ':
              isDragOver === status.value,
          }"
          @drop="onDrop(status.value)"
          @dragover.prevent="onDragOver(status.value)"
          @dragleave.prevent="onDragLeave"
          @touchmove="onTouchMove"
          @touchend="onTouchEnd(status.value)"
        >
          <transition-group name="card-animation" tag="div">
            <TaskCard
              v-for="task in filteredTasks(status.value)"
              :key="task.id"
              :task="task"
              :dragging="draggedTask?.id === task.id"
              draggable="true"
              @handleViewItem="handleViewItem"
              @handleUpdateItem="handleUpdateItem"
              @handleDeleteItem="handleDeleteItem"
              @dragstart="onDragStart"
              @dragend="onDragEnd"
              @touchstart.stop="onTaskTouchStart($event, task)"
              class="task-card"
              :class="{
                'cursor-grab': !isDragging,
                'cursor-grabbing': isDragging,
                'being-dragged': draggedTask?.id === task.id,
              }"
            />
          </transition-group>
        </div>
      </div>
    </div>

    <!-- Ghost card for touch drag -->
    <div
      v-if="touchDragging"
      class="fixed bg-white shadow-lg rounded-md p-2 pointer-events-none cursor-grabbing"
      :style="ghostCardStyle"
    >
      {{ draggedTask?.title }}
    </div>

    <ConfirmDeleteDialog
      v-model:isDialogVisible="deleteItemDialog.visible"
      v-model:title="deleteItemDialog.title"
      v-model:text="deleteItemDialog.text"
      @confirm="deleteItem"
    />

    <ViewTaskDialog
      v-model:isDialogVisible="viewItemDialog.visible"
      v-model:title="viewItemDialog.title"
    />

    <AddTaskDialog
      v-model:isDialogVisible="addItemDialog.visible"
      v-model:title="addItemDialog.title"
    />

    <UpdateTaskDialog
      v-model:isDialogVisible="updateItemDialog.visible"
      v-model:title="updateItemDialog.title"
    />
  </div>
</template>

<style scoped>
.task-card {
  touch-action: none;
  user-select: none;
  transition: transform 0.2s, box-shadow 0.2s;
}

.task-card:hover {
  cursor: grab;
}

.task-card.being-dragged {
  cursor: grabbing;
  transform: scale(1.02);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.bg-primary border-2 {
  background-color: #f3f4f6;
  border: 2px dashed #60a5fa;
}

.card-animation-enter-active,
.card-animation-leave-active {
  transition: all 0.3s ease;
}

.card-animation-enter-from {
  transform: translateY(20px);
  opacity: 0;
}

.card-animation-leave-to {
  transform: translateY(-20px);
  opacity: 0;
}

.drag-placeholder {
  border: 2px dashed #bbb;
  height: 60px;
  background-color: #f3f4f6;
  margin-bottom: 8px;
}

.dragging {
  opacity: 0.5;
}

.dropped {
  animation: dropped-effect 0.5s ease;
}

@keyframes dropped-effect {
  0% {
    transform: scale(1.2);
    background-color: #34d399;
  }
  100% {
    transform: scale(1);
    background-color: #ffffff;
  }
}

/* Cursor styles */
.cursor-grab {
  cursor: grab;
}

.cursor-grabbing {
  cursor: grabbing;
}

@media (max-width: 768px) {
  .drag-placeholder {
    height: 50px;
  }
}
</style>
